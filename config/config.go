package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// Config 配置中心
type Config struct {
	values  map[string]interface{}
	mu      sync.RWMutex
	sources []Source
}

// Source 配置源接口
type Source interface {
	Name() string
	Load() (map[string]interface{}, error)
	Watch(chan<- struct{})
}

// New 创建配置中心
func New(sources ...Source) *Config {
	return &Config{
		values:  make(map[string]interface{}),
		sources: sources,
	}
}

// Load 加载所有配置
func (c *Config) Load() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, source := range c.sources {
		values, err := source.Load()
		if err != nil {
			return fmt.Errorf("source %s load error: %v", source.Name(), err)
		}

		for k, v := range values {
			c.values[strings.ToLower(k)] = v
		}
	}

	return nil
}

// Watch 监听配置变化
func (c *Config) Watch() <-chan struct{} {
	ch := make(chan struct{}, 1)
	for _, source := range c.sources {
		go source.Watch(ch)
	}
	return ch
}

// Get 获取配置值
func (c *Config) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.values[strings.ToLower(key)]
	return val, ok
}

// Set 设置配置值
func (c *Config) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[strings.ToLower(key)] = value
}

// Unmarshal 将配置解析到结构体
func (c *Config) Unmarshal(v interface{}) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New("target must be a pointer to struct")
	}

	return c.unmarshalStruct(rv)
}

func (c *Config) unmarshalStruct(rv reflect.Value) error {
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldValue := rv.Field(i)

		// 获取字段tag
		tag := field.Tag.Get("config")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}

		// 从配置中获取值
		configValue, ok := c.values[tag]
		if !ok {
			continue
		}

		// 设置字段值
		if err := setFieldValue(fieldValue, configValue); err != nil {
			return fmt.Errorf("field %s: %v", field.Name, err)
		}
	}

	return nil
}

func setFieldValue(field reflect.Value, value interface{}) error {
	if !field.CanSet() {
		return nil
	}

	val := reflect.ValueOf(value)
	if val.Type().ConvertibleTo(field.Type()) {
		field.Set(val.Convert(field.Type()))
		return nil
	}

	return fmt.Errorf("cannot convert %v to %v", val.Type(), field.Type())
}
