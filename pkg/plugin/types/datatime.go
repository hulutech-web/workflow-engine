package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// DateTime 自定义时间类型（解决 JSON 格式化和 GORM 自动更新问题）
type DateTime struct {
	time.Time
}

// Value 实现 driver.Valuer 接口（写入数据库）
func (t DateTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil // 转换为标准 time.Time
}

// Scan 实现 sql.Scanner 接口（从数据库读取）
func (t *DateTime) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		t.Time = v
	case string:
		if v == "" {
			t.Time = time.Time{}
			return nil
		}
		// 尝试解析常见时间格式（带/不带时区）
		parsedTime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			parsedTime, err = time.Parse(time.RFC3339, v)
			if err != nil {
				return fmt.Errorf("无法解析时间字符串 '%s'，支持的格式: YYYY-MM-DD HH:mm:ss 或 RFC3339", v)
			}
		}
		t.Time = parsedTime
		return nil
	case []byte: // 处理某些驱动以 []byte 形式返回空值的情况（如 MySQL 的空字符串）
		if len(v) == 0 {
			t.Time = time.Time{}
			return nil
		}
		return fmt.Errorf("不支持的字节值: %s", string(v))
	default:
		return fmt.Errorf("不支持的数据库类型: %T", value)
	}
	return nil
}

// MarshalJSON 自定义 JSON 格式
func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.UTC().Format(time.RFC3339)) // 强制 UTC
}
func (t *DateTime) UnmarshalJSON(data []byte) error {
	var strTime string
	if err := json.Unmarshal(data, &strTime); err != nil {
		return err
	}
	if strTime == "null" || strTime == "" {
		t.Time = time.Time{}
		return nil
	}
	// 支持解析 RFC3339 和常规时间格式
	parsedTime, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02 15:04:05", strTime)
		if err != nil {
			return fmt.Errorf("无效时间格式 '%s'，要求格式: RFC3339 或 YYYY-MM-DD HH:mm:ss", strTime)
		}
	}
	t.Time = parsedTime
	return nil
}
