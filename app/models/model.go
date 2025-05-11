package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
	"time"
)

func init() {
	// 设置默认时区为上海（中国标准时间）
	carbon.SetTimezone(carbon.PRC) // PRC 代表中华人民共和国时区
	carbon.SetLocale("zh-CN")
}

type Model struct {
	ID        uint     `gorm:"primaryKey" json:"id" form:"id"`
	CreatedAt DateTime `gorm:"autoCreateTime;column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt DateTime `gorm:"autoUpdateTime;column:updated_at" form:"updated_at" json:"updated_at"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// DateTime 自定义时间类型（解决 JSON 格式化和 GORM 自动更新问题）
type DateTime struct {
	carbon.DateTime
}

// Value 实现 driver.Valuer 接口（写入数据库）
func (t DateTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil // 转换为标准 time.Time
}

// Scan 实现 sql.Scanner 接口（从数据库读取）
func (t *DateTime) Scan(value interface{}) error {
	if value == nil {
		t.DateTime = carbon.DateTime{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		t.DateTime = *carbon.NewDateTime(carbon.CreateFromStdTime(v))
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
	return json.Marshal(t.ToDateTimeString()) // 格式化为 "YYYY-MM-DD HH:mm:ss"
}
