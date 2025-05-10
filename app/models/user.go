package models

import "github.com/hulutech-web/workflow-engine/pkg/plugin/types"

type User struct {
	Model
	Username     string      `json:"username" gorm:"type:varchar(255);not null;comment:'用户名'"`
	Phone        types.Phone `json:"phone" gorm:"type:varchar(11);unique;comment:'手机号'"`
	Password     string      `json:"password" gorm:"type:varchar(32);not null"`
	Salt         string      `gorm:"not null;default:'';comment:'加密盐巴'"`
	Email        string      `json:"email" gorm:"type:varchar(125);"`
	Avatar       string      `json:"avatar" gorm:"type:varchar(255);"`
	Nickname     string      `json:"nickname" gorm:"type:varchar(255);"`
	RoleId       uint        `json:"role_id" gorm:"type:int;default:0"`
	IsMultipoint uint8       `gorm:"not null;default:0;comment:'多端登录: 0=否, 1=是''"`
	IsDisable    uint8       `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	TenantId     uint        `gorm:"not null;default:0;comment:'租户ID'"`
}
