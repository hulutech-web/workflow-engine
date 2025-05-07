package source

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// FileSource 文件配置源
type FileSource struct {
	path     string
	interval time.Duration
	lastMod  time.Time
}

// NewFileSource 创建文件配置源
func NewFileSource(path string, interval time.Duration) *FileSource {
	return &FileSource{
		path:     path,
		interval: interval,
	}
}

func (f *FileSource) Name() string {
	return "file"
}

func (f *FileSource) Load() (map[string]interface{}, error) {
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	f.lastMod = info.ModTime()

	var values map[string]interface{}

	switch ext := filepath.Ext(f.path); ext {
	case ".json":
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&values); err != nil {
			return nil, err
		}
	case ".yaml", ".yml":
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&values); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	return normalizeKeys(values), nil
}

func (f *FileSource) Watch(ch chan<- struct{}) {
	if f.interval <= 0 {
		return
	}

	ticker := time.NewTicker(f.interval)
	defer ticker.Stop()

	for range ticker.C {
		info, err := os.Stat(f.path)
		if err != nil {
			continue
		}

		if info.ModTime().After(f.lastMod) {
			f.lastMod = info.ModTime()
			ch <- struct{}{}
		}
	}
}

func normalizeKeys(values map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range values {
		key := strings.ToLower(k)

		if nested, ok := v.(map[string]interface{}); ok {
			v = normalizeKeys(nested)
		}

		result[key] = v
	}

	return result
}
