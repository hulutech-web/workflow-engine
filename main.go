package main

import (
	"fmt"
	"github.com/hulutech-web/workflow-engine/auth"
	"github.com/hulutech-web/workflow-engine/database"
	"github.com/hulutech-web/workflow-engine/logging"
	"github.com/redis/go-redis"
	"workflow-engine/config"
)

func main() {
	// 加载配置
	cfg, err := config.DefaultLoader().Load()
	if err != nil {
		panic(err)
	}

	// 解析数据库配置
	var dbConfig config.DatabaseConfig
	if err := cfg.Unmarshal(&dbConfig); err != nil {
		panic(err)
	}

	// 初始化数据库
	db, err := database.NewGormRepository(
		dbConfig.Driver,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
			dbConfig.Params,
		),
	)
	if err != nil {
		panic(err)
	}

	// 解析Redis配置
	var redisConfig config.RedisConfig
	if err := cfg.Unmarshal(&redisConfig); err != nil {
		panic(err)
	}

	// 初始化Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	// 解析JWT配置
	var jwtConfig config.JWTConfig
	if err := cfg.Unmarshal(&jwtConfig); err != nil {
		panic(err)
	}

	// 初始化认证
	authManager := auth.NewJwtAuth(
		redisClient,
		auth.JwtConfig{
			Secret:        jwtConfig.Secret,
			AccessExpiry:  jwtConfig.AccessExpiry,
			RefreshExpiry: jwtConfig.RefreshExpiry,
			TokenRotation: true,
		},
		NewUserRepository(db),
	)

	// 解析日志配置
	var logConfig config.LoggingConfig
	if err := cfg.Unmarshal(&logConfig); err != nil {
		panic(err)
	}

	// 初始化日志
	logger := logging.NewZapLogger(logConfig.Level == "debug")
	if logConfig.FilePath != "" {
		logger.AddFileOutput(logConfig.FilePath)
	}

	// 监听配置变化
	go func() {
		for range cfg.Watch() {
			logger.Info("configuration changed, reloading...")
			// 重新加载配置并更新相关组件
		}
	}()
}
