package core

import (
	"errors"
	"reflect"
	"sync"
)

type Container interface {
	Bind(string, interface{})
	Singleton(string, func(Container) interface{})
	Make(string) (interface{}, error)
	Has(string) bool
}

type service struct {
	constructor func(Container) interface{}
	instance    interface{}
	singleton   bool
}

type container struct {
	services map[string]service
	mu       sync.RWMutex
}

func NewContainer() Container {
	return &container{
		services: make(map[string]service),
	}
}

func (c *container) Bind(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Func {
		c.services[key] = service{
			constructor: func(c Container) interface{} {
				return val.Call([]reflect.Value{reflect.ValueOf(c)})[0].Interface()
			},
			singleton: false,
		}
	} else {
		c.services[key] = service{
			instance:  value,
			singleton: true,
		}
	}
}

func (c *container) Singleton(key string, constructor func(Container) interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.services[key] = service{
		constructor: constructor,
		singleton:   true,
	}
}

func (c *container) Make(key string) (interface{}, error) {
	c.mu.RLock()
	srv, exists := c.services[key]
	c.mu.RUnlock()

	if !exists {
		return nil, errors.New("service not found")
	}

	if srv.singleton && srv.instance != nil {
		return srv.instance, nil
	}

	if srv.constructor != nil {
		instance := srv.constructor(c)
		if srv.singleton {
			c.mu.Lock()
			srv.instance = instance
			c.services[key] = srv
			c.mu.Unlock()
		}
		return instance, nil
	}

	return srv.instance, nil
}

func (c *container) Has(key string) bool {
	c.mu.RLock()
	_, exists := c.services[key]
	c.mu.RUnlock()
	return exists
}
