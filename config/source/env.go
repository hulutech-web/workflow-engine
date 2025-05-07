package source

import (
	"os"
	"strings"
)

// EnvSource 环境变量配置源
type EnvSource struct {
	prefix string
}

// NewEnvSource 创建环境变量配置源
func NewEnvSource(prefix string) *EnvSource {
	return &EnvSource{prefix: prefix}
}

func (e *EnvSource) Name() string {
	return "env"
}

func (e *EnvSource) Load() (map[string]interface{}, error) {
	envs := os.Environ()
	values := make(map[string]interface{})

	for _, env := range envs {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) != 2 {
			continue
		}

		key := pair[0]
		value := pair[1]

		if e.prefix != "" && !strings.HasPrefix(key, e.prefix) {
			continue
		}

		// 去掉前缀并转为小写
		configKey := strings.ToLower(strings.TrimPrefix(key, e.prefix))
		values[configKey] = value
	}

	return values, nil
}

func (e *EnvSource) Watch(ch chan<- struct{}) {
	// 环境变量通常不监听变化
}
