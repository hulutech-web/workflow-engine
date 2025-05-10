package models

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

func init() {
	// 设置默认时区为上海（中国标准时间）
	carbon.SetTimezone(carbon.PRC) // PRC 代表中华人民共和国时区
	carbon.SetLocale("zh-CN")
}

type Model struct {
	ID uint `gorm:"primaryKey" json:"id" form:"id"`
	Timestamps
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type Timestamps struct {
	CreatedAt carbon.DateTime `gorm:"autoCreateTime;column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" form:"updated_at" json:"updated_at"`
}
