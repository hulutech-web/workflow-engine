package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

const (
	configFile = "config.yaml"
	configType = "yaml"
)

func NewConfig() *Config {
	defaultConfig := &Config{
		Server: Server{
			Port: 8080,
			Host: "0.0.0.0",
			Mode: "debug",
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
			Password: "",
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
	}
	conf := &Config{}
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
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
