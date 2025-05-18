package queue

import (
	"context"
	"github.com/hulutech-web/workflow-engine/core/queue"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

type RoleService struct {
	delayQueue *queue.DelayQueue
	queue      *queue.Queue
}

func NewRoleService(delayQueue *queue.DelayQueue, queue *queue.Queue) *RoleService {
	return &RoleService{
		delayQueue: delayQueue,
		queue:      queue,
	}
}

func (s *RoleService) ConsumeDelayedRoles() {
	go func() {
		err := s.delayQueue.Poll("role.delayed", func(msg string) error {
			zap.S().Info("处理延迟角色队列消息", "message", msg)
			// 实际业务逻辑处理
			return nil
		})
		if err != nil {
			zap.S().Error("延迟队列消费错误", zap.Error(err))
		}
	}()
}

func (s *RoleService) ProduceDelayedRoleMessage(message string, delay time.Duration) error {
	return s.delayQueue.Add("role.delayed", message, delay)
}

func (s *RoleService) ConsumeRoles() {
	go func() {
		err := s.queue.LPop("role.normal", func(msg string) error {
			zap.S().Info("处理普通角色队列消息", "message", msg)
			// 实际业务逻辑处理
			return nil
		})
		if err != nil {
			zap.S().Error("普通队列消费错误", zap.Error(err))
		}
	}()
}

func (s *RoleService) ProduceRoleMessage(message string) {
	s.queue.Push("role.normal", message)
}

func RegisterRoleConsumer(lc fx.Lifecycle, srv *RoleService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			srv.ConsumeDelayedRoles()
			srv.ConsumeRoles()
			return nil
		},
	})
}

// TestQueueUsage 示例用法
func TestQueueUsage(srv *RoleService) {
	// 生产延迟10秒的消息
	srv.ProduceDelayedRoleMessage("role_update_task", 10*time.Second)
	// 生产立即处理的消息
	srv.ProduceRoleMessage("role_immediate_task")
}
