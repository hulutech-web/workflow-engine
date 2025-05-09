package queue

import (
	"context"
	"errors"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/core/config"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"sync"
	"time"
)

type DelayQueue struct {
	C            *cache.Redis
	PollInterval time.Duration  // 轮询间隔
	mu           sync.Mutex     // 互斥锁保障并发安全
	config       *config.Config // 配置
}

// NewDelayQueue 创建延时队列
func NewDelayQueue(r *cache.Redis, config *config.Config) *DelayQueue {
	return &DelayQueue{
		C:            r,
		PollInterval: time.Second, // 默认1秒轮询间隔
		config:       config,
	}
}

// Add 添加延时消息
func (dq *DelayQueue) Add(queueName string, message string, delay time.Duration) error {
	// 计算到期时间戳
	score := time.Now().Add(delay).UnixNano()
	fullKey := dq.config.Redis.Prefix + queueName
	_, err := dq.C.Instance.ZAdd(context.Background(), fullKey, redis.Z{
		Score:  float64(score),
		Member: message,
	}).Result()
	return err
}

// Poll 消费到期消息
func (dq *DelayQueue) Poll(queueName string, handle func(message string) error) error {
	// 获取当前时间戳
	now := time.Now().UnixNano()

	// 查询到期消息
	dq.mu.Lock()
	defer dq.mu.Unlock()

	fullKey := dq.config.Redis.Prefix + queueName
	messages, err := dq.C.Instance.ZRangeByScore(context.Background(), fullKey, &redis.ZRangeBy{
		Min:    "0",
		Max:    strconv.FormatInt(now, 10),
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
	removed, err := dq.C.Instance.ZRem(context.Background(), fullKey, message).Result()
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
func (dq *DelayQueue) Remove(queueName string, message string) error {
	fullKey := dq.config.Redis.Prefix + queueName
	_, err := dq.C.Instance.ZRem(context.Background(), fullKey, message).Result()
	return err
}

// Size 获取队列长度
func (dq *DelayQueue) Size(queueName string) int64 {
	fullKey := dq.config.Redis.Prefix + queueName
	count, err := dq.C.Instance.ZCard(context.Background(), fullKey).Result()
	if err != nil {
		log.Printf("Size error: %v", err)
	}
	return count
}

// Clear 清空队列
func (dq *DelayQueue) Clear(queueName string) {
	fullKey := dq.config.Redis.Prefix + queueName
	_, err := dq.C.Instance.Del(context.Background(), fullKey).Result()
	if err != nil {
		log.Printf("Clear error: %v", err)
	}
}

/**
e.g.:
dq := NewDelayQueue(redisClient)
// 添加5秒后到期的消息
dq.Add("queue1","message1", 5*time.Second)

// 消费消息
go func() {
    for {
        dq.Poll("queue1",func(msg string) error {
            // 处理消息
            return nil
        })
    }
}()
*/
