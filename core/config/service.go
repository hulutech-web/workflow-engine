package config

import (
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
	"path/filepath"
)

const (
	configFile = "config.yaml"
	configType = "yaml"
)

func NewConfig() *Config {
	defaultConfig := &Config{
		Server: Server{
			Port:      8080,
			Host:      "0.0.0.0",
			Mode:      "debug",
			PublicUrl: "http://localhost:8080",
		},
		Database: Database{
			Driver:       "mysql",
			Host:         "127.0.0.1",
			Port:         3306,
			Username:     "workflow",
			Password:     "123456",
			Database:     "workflow",
			Params:       "charset=utf8mb4&parseTime=True&loc=Local",
			MaxOpenConns: 10,
			MaxIdleConns: 5,
			MaxLifeTime:  10,
			AutoMigrate:  true,
		},
		Redis: Redis{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "redis_QCHytt",
			Db:       0,
		},
		Jwt: Jwt{
			Secret:        "workflow",
			AccessExpiry:  "1h",
			RefreshExpiry: "12h",
		},
		Logging: Logging{
			Level:      "info",
			FilePath:   "public/logs/app.log",
			MaxSize:    10,
			MaxBackups: 10,
			MaxAge:     5,
		},
		Storage: Storage{
			PublicPrefix: "/uploads",
			LocalPath:    "public/uploads",
		},
	}
	conf := &Config{}
	viper.SetConfigType(configType)
	dir, _ := os.Getwd()
	viper.SetConfigFile(filepath.Join(dir, configFile))
	if err := viper.ReadInConfig(); err != nil {
		return defaultConfig
	}
	// 写入默认配置
	err := copier.Copy(defaultConfig, conf)
	if err != nil {
		return defaultConfig
	}
	// 可补充动态配置

	if err := viper.Unmarshal(conf); err != nil {
		return defaultConfig
	}

	return conf
}

var Module = fx.Provide(
	NewConfig,
)
