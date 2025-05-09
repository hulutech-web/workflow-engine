package cache

import "C"
import (
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Queue struct {
	QueueName string
	C         *Redis
}

// NewQueue 创建队列
func NewQueue(queueName string, r *Redis) *Queue {
	return &Queue{QueueName: queueName, C: r}
}

// Push 推送消息
func (q *Queue) Push(message string) {
	_, err := q.C.Instance.RPush(C.Ctx, q.QueueName, message).Result()
	if err != nil {
		log.Println("Push err: ", err)
	}
}

// RPop 消费消息
func (q *Queue) RPop(handle func(message string) error) error {
	message, err := q.C.Instance.RPop(C.Ctx, q.QueueName).Result()
	if errors.Is(err, redis.Nil) {
		time.Sleep(1 * time.Second)
		return nil
	} else if err != nil {
		log.Println("Pop err: ", err)
		return err
	}
	err = handle(message)
	if err != nil {
		return err
	}
	return nil
}

// LPop 消费消息
func (q *Queue) LPop(handle func(message string) error) error {
	message, err := q.C.Instance.LPop(C.Ctx, q.QueueName).Result()
	if errors.Is(err, redis.Nil) {
		time.Sleep(1 * time.Second)
		return nil
	} else if err != nil {
		log.Println("Pop err: ", err)
		return err
	}
	err = handle(message)
	if err != nil {
		return err
	}
	return nil

}

// Len 队列长度
func (q *Queue) Len() int64 {
	l, err := q.C.Instance.LLen(C.Ctx, q.QueueName).Result()
	if err != nil {
		log.Println("Len err: ", err)
	}
	return l
}

// Clear 清空队列
func (q *Queue) Clear() {
	_, err := q.C.Instance.Del(C.Ctx, q.QueueName).Result()
	if err != nil {
		log.Println("Clear err: ", err)
	}
}
