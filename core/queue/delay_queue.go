package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type DelayQueue struct {
	QueueName string
	C         *Redis
	PollInterval time.Duration // 轮询间隔
}

// NewDelayQueue 创建延时队列
func NewDelayQueue(queueName string, r *Redis) *DelayQueue {
	return &DelayQueue{
		QueueName:    queueName,
		C:            r,
		PollInterval: time.Second, // 默认1秒轮询间隔
	}
}

// Add 添加延时消息
func (dq *DelayQueue) Add(message string, delay time.Duration) error {
	// 计算到期时间戳
	score := time.Now().Add(delay).UnixNano()
	_, err := dq.C.Instance.ZAdd(context.Background(), dq.QueueName, redis.Z{
		Score:  float64(score),
		Member: message,
	}).Result()
	return err
}

// Poll 消费到期消息
func (dq *DelayQueue) Poll(handle func(message string) error) error {
	// 获取当前时间戳
	now := time.Now().UnixNano()
	
	// 查询到期消息
	messages, err := dq.C.Instance.ZRangeByScore(context.Background(), dq.QueueName, &redis.ZRangeBy{
		Min:    "0",
		Max:    string(rune(now)),
		Offset: 0,
		Count:  1,
	}).Result()

	if errors.Is(err, redis.Nil) {
		time.Sleep(dq.PollInterval)
		return nil
	} else if err != nil {
		log.Printf("Poll error: %v", err)
		return err
	}

	if len(messages) == 0 {
		time.Sleep(dq.PollInterval)
		return nil
	}

	// 原子操作：获取并删除消息
	message := messages[0]
	removed, err := dq.C.Instance.ZRem(context.Background(), dq.QueueName, message).Result()
	if err != nil {
		log.Printf("Remove message error: %v", err)
		return err
	}
	if removed == 0 {
		return nil // 消息已被其他消费者处理
	}

	return handle(message)
}

// Remove 移除指定消息
func (dq *DelayQueue) Remove(message string) error {
	_, err := dq.C.Instance.ZRem(context.Background(), dq.QueueName, message).Result()
	return err
}

// Size 获取队列长度
func (dq *DelayQueue) Size() int64 {
	count, err := dq.C.Instance.ZCard(context.Background(), dq.QueueName).Result()
	if err != nil {
		log.Printf("Size error: %v", err)
	}
	return count
}

// Clear 清空队列
func (dq *DelayQueue) Clear() {
	_, err := dq.C.Instance.Del(context.Background(), dq.QueueName).Result()
	if err != nil {
		log.Printf("Clear error: %v", err)
	}
}