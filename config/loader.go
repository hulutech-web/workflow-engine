package config

import (
	"fmt"
	"path/filepath"
	"time"
)

// Loader 配置加载器
type Loader struct {
	sources []Source
}

// NewLoader 创建配置加载器
func NewLoader() *Loader {
	return &Loader{}
}

// WithEnv 添加环境变量配置源
func (l *Loader) WithEnv(prefix string) *Loader {
	l.sources = append(l.sources, NewEnvSource(prefix))
	return l
}

// WithFile 添加文件配置源
func (l *Loader) WithFile(path string, watchInterval time.Duration) *Loader {
	absPath, err := filepath.Abs(path)
	if err != nil {
		absPath = path
	}
	l.sources = append(l.sources, NewFileSource(absPath, watchInterval))
	return l
}

// Load 加载配置
func (l *Loader) Load() (*Config, error) {
	cfg := New(l.sources...)
	if err := cfg.Load(); err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	return cfg, nil
}

// DefaultLoader 默认配置加载器
func DefaultLoader() *Loader {
	return NewLoader().
		WithEnv("APP_").
		WithFile("config.yaml", 0)
}
