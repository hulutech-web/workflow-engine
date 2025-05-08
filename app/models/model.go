package models

import "gorm.io/gorm"

type Model struct {
	ID        uint  `gorm:"primary_key" json:"id"`
	CreatedAt int64 `json:"created_at" gorm:"type:bigint;column:created_at;autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"type:bigint;column:updated_at;autoUpdateTime"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"type:datetime;column:deleted_at;index"`
}
