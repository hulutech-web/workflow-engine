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
	if poolSize == 0 {
		poolSize = 10
	}
	ctx := context.Background()
	client := redis.NewClient(opt)
	pCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err = client.Ping(pCtx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis 连接失败: %v", err)
	}
	return &Redis{
		config:   config,
		Instance: client,
		Ctx:      ctx,
	}, nil
}

var Module = fx.Provide(NewRedis)
