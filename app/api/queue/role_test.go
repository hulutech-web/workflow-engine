package queue

import (
	"github.com/hulutech-web/workflow-engine/core/cache"
	"testing"
	"time"

	"github.com/hulutech-web/workflow-engine/core/config"
	"github.com/hulutech-web/workflow-engine/core/queue"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestMain(m *testing.M) {
	// 初始化测试配置
	config.NewConfig()
	m.Run()
}

func createTestModule(tb fxtest.TB) fx.Option {
	return fx.Options(
		fx.Provide(
			func() *config.Config {
				return config.NewConfig()
			},
			cache.NewRedis,
			queue.NewDelayQueue,
			queue.NewQueue,
			NewRoleService,
		),
		fx.Invoke(RegisterRoleConsumer),
		fx.NopLogger,
	)
}

func TestNormalQueue(t *testing.T) {
	// 初始化测试环境
	var srv *RoleService
	app := fxtest.New(t, createTestModule(t), fx.Populate(&srv))
	app.RequireStart()
	defer app.RequireStop()

	t.Run("测试普通队列完整流程", func(t *testing.T) {

		// 清空测试队列
		srv.queue.Clear("role.normal")

		// 生产测试消息
		msg := "test_normal_message"
		srv.ProduceRoleMessage(msg)

		// 创建结果通道
		resultChan := make(chan string, 1)

		// 替换原消费者逻辑
		err := srv.queue.LPop("role.normal", func(m string) error {
			resultChan <- m
			return nil
		})
		assert.NoError(t, err)

		// 验证消息消费
		select {
		case received := <-resultChan:
			assert.Equal(t, msg, received)
		case <-time.After(3 * time.Second):
			t.Fatal("普通队列消息消费超时")
		}
	})
}

func TestDelayedQueue(t *testing.T) {
	var srv *RoleService
	app := fxtest.New(t, createTestModule(t), fx.Populate(&srv))
	app.RequireStart()
	defer app.RequireStop()

	t.Run("测试延迟队列定时触发", func(t *testing.T) {

		srv.delayQueue.Clear("role.delayed")

		startTime := time.Now()
		msg := "test_delayed_message"

		// 生产延迟2秒的消息
		err := srv.ProduceDelayedRoleMessage(msg, 2*time.Second)
		assert.NoError(t, err)

		resultChan := make(chan string, 1)

		go func() {
			for {
				err = srv.delayQueue.Poll("role.delayed", func(m string) error {
					resultChan <- m
					return nil
				})
				if err != nil {
					t.Error(err)
					return
				}
			}
		}()

		select {
		case received := <-resultChan:
			elapsed := time.Since(startTime)
			assert.Equal(t, msg, received)
			assert.True(t, elapsed >= 2*time.Second, "延迟时间不足")
		case <-time.After(5 * time.Second):
			t.Fatal("延迟队列消息消费超时")
		}
	})
}
