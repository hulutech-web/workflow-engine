package models

import "github.com/hulutech-web/workflow-engine/app/models/common"

type Model struct {
	ID        uint                  `gorm:"primary_key" json:"id"`
	CreatedAt common.CarbonDateTime `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt common.CarbonDateTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type SoftDelete struct {
	DeletedAt common.CarbonDateTime `json:"deleted_at" gorm:"type:datetime;column:deleted_at;index"`
}
