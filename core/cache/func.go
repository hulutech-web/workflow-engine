package cache

import (
	"bufio"
	"context"
	"go.uber.org/zap"
	"strings"
	"time"
)

func (ru *Redis) Info(sections ...string) (res map[string]string) {
	infoStr, err := ru.Instance.Info(context.Background(), sections...).Result()
	res = map[string]string{}
	if err != nil {
		zap.S().Errorf("redisUtil.Info err: err=[%+v]", err)
		return res
	}
	// string拆分多行
	lines, err := stringToLines(infoStr)
	if err != nil {
		zap.S().Errorf("stringToLines err: err=[%+v]", err)
		return res
	}
	// 解析成Map
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || strings.HasPrefix(lines[i], "# ") {
			continue
		}
		k, v := stringToKV(lines[i])
		res[k] = v
	}
	return res
}

// DBSize 当前数据库key数量
func (ru *Redis) DBSize() int64 {
	size, err := ru.Instance.DBSize(context.Background()).Result()
	if err != nil {
		zap.S().Errorf("redisUtil.DBSize err: err=[%+v]", err)
		return 0
	}
	return size
}

// Set 设置键值对
func (ru *Redis) Set(key string, value interface{}, timeSec int) bool {
	err := ru.Instance.Set(context.Background(),
		ru.config.Redis.Prefix+key, value, time.Duration(timeSec)*time.Second).Err()
	if err != nil {
		zap.S().Errorf("redisUtil.Set err: err=[%+v]", err)
		return false
	}
	return true
}

// Get 获取key的值
func (ru *Redis) Get(key string) string {
	res, err := ru.Instance.Get(context.Background(), ru.config.Redis.Prefix+key).Result()
	if err != nil {
		zap.S().Errorf("redisUtil.Get err: err=[%+v]", err)
		return ""
	}
	return res
}

// SSet 将数据放入set缓存
func (ru *Redis) SSet(key string, values ...interface{}) bool {
	err := ru.Instance.SAdd(context.Background(), ru.config.Redis.Prefix+key, values...).Err()
	if err != nil {
		zap.S().Errorf("redisUtil.SSet err: err=[%+v]", err)
		return false
	}
	return true
}

// SGet 根据key获取Set中的所有值
func (ru *Redis) SGet(key string) []string {
	res, err := ru.Instance.SMembers(context.Background(), ru.config.Redis.Prefix+key).Result()
	if err != nil {
		zap.S().Errorf("redisUtil.SGet err: err=[%+v]", err)
		return []string{}
	}
	return res
}

// HMSet 设置key, 通过字典的方式设置多个field, value对
func (ru *Redis) HMSet(key string, mapping map[string]string, timeSec int) bool {
	err := ru.Instance.HSet(context.Background(), ru.config.Redis.Prefix+key, mapping).Err()
	if err != nil {
		zap.S().Errorf("redisUtil.HMSet err: err=[%+v]", err)
		return false
	}
	if timeSec > 0 {
		if !ru.Expire(key, timeSec) {
			return false
		}
	}
	return true
}

// HSet 向hash表中放入数据,如果不存在将创建
func (ru *Redis) HSet(key string, field string, value string, timeSec int) bool {
	return ru.HMSet(key, map[string]string{field: value}, timeSec)
}

// HGet 获取key中field域的值
func (ru *Redis) HGet(key string, field string) string {
	res, err := ru.Instance.HGet(context.Background(), ru.config.Redis.Prefix+key, field).Result()
	if err != nil {
		zap.S().Errorf("redisUtil.HGet err: err=[%+v]", err)
		return ""
	}
	return res
}

// HExists 判断key中有没有field域名
func (ru *Redis) HExists(key string, field string) bool {
	res, err := ru.Instance.HExists(context.Background(), ru.config.Redis.Prefix+key, field).Result()
	if err != nil {
		zap.S().Errorf("redisUtil.HExists err: err=[%+v]", err)
		return false
	}
	return res
}

// HDel 删除hash表中的值
func (ru *Redis) HDel(key string, fields ...string) bool {
	err := ru.Instance.HDel(context.Background(), ru.config.Redis.Prefix+key, fields...).Err()
	if err != nil {
		zap.S().Errorf("redisUtil.HDel err: err=[%+v]", err)
		return false
	}
	return true
}

// Exists 判断多项key是否存在
func (ru *Redis) Exists(keys ...string) int64 {
	fullKeys := ru.toFullKeys(keys)
	cnt, err := ru.Instance.Exists(context.Background(), fullKeys...).Result()
	if err != nil {
		zap.S().Errorf("redisUtil.Exists err: err=[%+v]", err)
		return -1
	}
	return cnt
}

// Expire 指定缓存失效时间
func (ru *Redis) Expire(key string, timeSec int) bool {
	err := ru.Instance.Expire(context.Background(), ru.config.Redis.Prefix+key, time.Duration(timeSec)*time.Second).Err()
	if err != nil {
		zap.S().Errorf("redisUtil.Expire err: err=[%+v]", err)
		return false
	}
	return true
}

// TTL 根据key获取过期时间
func (ru *Redis) TTL(key string) int {
	td, err := ru.Instance.TTL(context.Background(), ru.config.Redis.Prefix+key).Result()
	if err != nil {
		zap.S().Errorf("redisUtil.TTL err: err=[%+v]", err)
		return 0
	}
	return int(td / time.Second)
}

// Del 删除一个或多个键
func (ru *Redis) Del(keys ...string) bool {
	fullKeys := ru.toFullKeys(keys)
	err := ru.Instance.Del(context.Background(), fullKeys...).Err()
	if err != nil {
		zap.S().Errorf("redisUtil.Del err: err=[%+v]", err)
		return false
	}
	return true
}

// toFullKeys 为keys批量增加前缀
func (ru *Redis) toFullKeys(keys []string) (fullKeys []string) {
	for _, k := range keys {
		fullKeys = append(fullKeys, ru.config.Redis.Prefix+k)
	}
	return
}

// stringToLines string拆分多行
func stringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

// stringToKV string拆分key和val
func stringToKV(s string) (string, string) {
	ss := strings.Split(s, ":")
	if len(ss) < 2 {
		return s, ""
	}
	return ss[0], ss[1]
}
