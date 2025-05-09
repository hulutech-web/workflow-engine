package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/dromara/carbon/v2"
	"time"
)

type RuleItem struct {
	RuleName  string `json:"rule_name" form:"rule_name"`
	RuleTitle string `json:"rule_title" form:"rule_title"`
	RuleValue string `json:"rule_value" form:"rule_value"`
}
type Rule []RuleItem

func (t *Rule) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Rule) Value() (driver.Value, error) {
	//如果t为nil,返回nil
	return json.Marshal(t)
}

type FieldValue []string

func (t *FieldValue) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t FieldValue) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// CarbonDateTime 自定义 carbon.DateTime 类型
type CarbonDateTime struct {
	*carbon.Carbon
}

// NewCarbonDateTime 创建一个新的 CarbonDateTime
func NewCarbonDateTime(c *carbon.Carbon) CarbonDateTime {
	if c == nil {
		return CarbonDateTime{carbon.Now()}
	}
	return CarbonDateTime{c}
}

// Scan 实现 sql.Scanner 接口
func (c *CarbonDateTime) Scan(value interface{}) error {
	if value == nil {
		c.Carbon = carbon.Now()
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		c.Carbon = carbon.CreateFromStdTime(v)
	case []byte:
		// 尝试解析为时间字符串
		str := string(v)
		c.Carbon = carbon.Parse(str)
	case string:
		c.Carbon = carbon.Parse(v)
	default:
		return fmt.Errorf("unsupported type for CarbonDateTime.Scan: %T", value)
	}

	return nil
}

// Value 实现 driver.Valuer 接口
func (c CarbonDateTime) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.ToDateTimeString(), nil
}

// IsZero 判断是否为零值
func (c CarbonDateTime) IsZero() bool {
	return c.Carbon == nil || c.Carbon.IsZero()
}

// GormDataType 定义 GORM 数据类型
func (CarbonDateTime) GormDataType() string {
	return "datetime"
}

// MarshalJSON 实现 json.Marshaler 接口
func (c CarbonDateTime) MarshalJSON() ([]byte, error) {
	if c.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(c.ToDateTimeString())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (c *CarbonDateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		c.Carbon = carbon.Now()
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	c.Carbon = carbon.Parse(str)

	return nil
}

// String 实现 Stringer 接口
func (c CarbonDateTime) String() string {
	if c.IsZero() {
		return ""
	}
	return c.ToDateTimeString()
}

// Now 获取当前时间
func Now() CarbonDateTime {
	return NewCarbonDateTime(carbon.Now())
}

// Parse 从字符串解析时间
func Parse(layout, value string) (CarbonDateTime, error) {
	dt := carbon.ParseByFormat(value, layout)
	return NewCarbonDateTime(dt), nil
}

// FromTime 从 time.Time 创建
func FromTime(t time.Time) CarbonDateTime {
	return NewCarbonDateTime(carbon.CreateFromStdTime(t))
}

// Unix 从 Unix 时间戳创建
func Unix(sec int64, nsec int64) CarbonDateTime {
	return NewCarbonDateTime(carbon.CreateFromTimestamp(sec))
}
