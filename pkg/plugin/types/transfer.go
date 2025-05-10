package types

import (
	"database/sql/driver"
	"encoding/json"
)

type FieldValue []string
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

func (t *FieldValue) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t FieldValue) Value() (driver.Value, error) {
	return json.Marshal(t)
}
