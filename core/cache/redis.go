package cache

import (
	"context"
	"fmt"
	"github.com/hulutech-web/workflow-engine/core/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"time"
)

type Redis struct {
	Instance *redis.Client
	Ctx      context.Context
	config   *config.Config
}

func NewRedis(config *config.Config) (*Redis, error) {
	dsn := fmt.Sprintf("redis://:%s@%s:%d/%d",
		config.Redis.Password,
		config.Redis.Host,
		config.Redis.Port,
		config.Redis.Db)
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("redis 解析dsn失败: %v", err)
	}
	poolSize := config.Redis.PoolSize
	if opt.PoolSize == 0 {
		opt.PoolSize = 10 // 默认连接池大小
	}
	opt.MinIdleConns = 5                   // 最小空闲连接数
	opt.MaxActiveConns = poolSize * 2      // 最大连接数
	opt.PoolTimeout = 1 * time.Second      // 获取连接超时时间
	opt.MaxIdleConns = poolSize - 5        // 最大空闲连接数
	opt.ConnMaxIdleTime = 10 * time.Minute // 连接最大空闲时间
	opt.DialTimeout = 10 * time.Second     // 连接建立超时时间
	opt.ReadTimeout = 5 * time.Second      // 读操作超时
	opt.WriteTimeout = 5 * time.Second     // 写操作超时

	ctx := context.Background()
	client := redis.NewClient(opt)
	return &Redis{
		config:   config,
		Instance: client,
		Ctx:      ctx,
	}, nil
}

var Module = fx.Provide(NewRedis)

func (r *Redis) Close() error {
	return r.Instance.Close()
}

// Ping 测试redis连接
func (r *Redis) Ping() error {
	pCtx, cancel := context.WithTimeout(r.Ctx, 3*time.Second)
	defer cancel()
	_, err := r.Instance.Ping(pCtx).Result()
	if err != nil {
		return fmt.Errorf("redis 连接失败: %v", err)
	}
	return nil
}
