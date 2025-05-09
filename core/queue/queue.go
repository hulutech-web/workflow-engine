package queue

import (
	"errors"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type Queue struct {
	C *cache.Redis
}

// NewQueue 创建队列
func NewQueue(r *cache.Redis) *Queue {
	return &Queue{C: r}
}

// Push 推送消息
func (q *Queue) Push(queueName string, message string) {
	_, err := q.C.Instance.RPush(q.C.Ctx, queueName, message).Result()
	if err != nil {
		zap.S().Warn("Push err: ", err)
	}
}

// RPop 消费消息
func (q *Queue) RPop(queueName string, handle func(message string) error) error {
	message, err := q.C.Instance.RPop(q.C.Ctx, queueName).Result()
	if errors.Is(err, redis.Nil) {
		time.Sleep(1 * time.Second)
		return nil
	} else if err != nil {
		zap.S().Warn("Pop err: ", err)
		return err
	}
	err = handle(message)
	if err != nil {
		return err
	}
	return nil
}

// LPop 消费消息
func (q *Queue) LPop(queueName string, handle func(message string) error) error {
	message, err := q.C.Instance.LPop(q.C.Ctx, queueName).Result()
	if errors.Is(err, redis.Nil) {
		time.Sleep(1 * time.Second)
		return nil
	} else if err != nil {
		zap.S().Warn("Pop err: ", err)
		return err
	}
	err = handle(message)
	if err != nil {
		return err
	}
	return nil

}

// Len 队列长度
func (q *Queue) Len(queueName string) int64 {
	l, err := q.C.Instance.LLen(q.C.Ctx, queueName).Result()
	if err != nil {
		zap.S().Warn("Len err: ", err)
	}
	return l
}

// Clear 清空队列
func (q *Queue) Clear(queueName string) {
	_, err := q.C.Instance.Del(q.C.Ctx, queueName).Result()
	if err != nil {
		zap.S().Warn("Del err: ", err)
	}
}
