package models

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint             `gorm:"primary_key" json:"id"`
	CreatedAt carbon.Timestamp `json:"created_at" gorm:"type:bigint;column:created_at;autoCreateTime"`
	UpdatedAt carbon.Timestamp `json:"updated_at" gorm:"type:bigint;column:updated_at;autoUpdateTime"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:datetime;column:deleted_at;index"`
}
