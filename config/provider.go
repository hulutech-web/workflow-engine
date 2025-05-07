package config

import "time"

// Provider 配置提供者接口
type Provider interface {
	Config() *Config
}

// Configurable 可配置接口
type Configurable interface {
	Configure(cfg *Config) error
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `config:"driver"`
	Host     string `config:"host"`
	Port     int    `config:"port"`
	Username string `config:"username"`
	Password string `config:"password"`
	Database string `config:"database"`
	Params   string `config:"params"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `config:"host"`
	Port     int    `config:"port"`
	Password string `config:"password"`
	DB       int    `config:"db"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret        string        `config:"secret"`
	AccessExpiry  time.Duration `config:"access_expiry"`
	RefreshExpiry time.Duration `config:"refresh_expiry"`
}

// LoggingConfig 日志配置
type LoggingConfig struct {
	Level    string `config:"level"`
	FilePath string `config:"file_path"`
}
