package models

import "time"

// LicenseKey 秘钥表结构体
type LicenseKey struct {
	Model
	KeyID       string           `json:"key_id" gorm:"column:key_id;unique;" comment:"秘钥KeyID"`
	KeyValue    string           `json:"key_value" gorm:"column:key_value;not null" comment:"秘钥内容（加密存储）"`
	TenantID    uint             `json:"tenant_id" gorm:"column:tenant_id;index" comment:"租户ID（未售出时为NULL）"`
	PackageID   uint             `json:"package_id" gorm:"null" comment:"关联套餐ID"`
	Status      LicenseKeyStatus `json:"status" gorm:"index;default:'inactive'" comment:"状态（inactive/active/expired/disabled）"`
	PurchasedAt *time.Time       `json:"purchased_at" comment:"购买时间（未售出为NULL）"`
	ActivatedAt *time.Time       `json:"activated_at" comment:"激活时间"`
	ExpiresAt   *time.Time       `json:"expires_at" gorm:"index" comment:"过期时间"`
	UsageLimit  *int             `json:"usage_limit" comment:"使用次数限制（NULL表示不限制）"`
	UsageCount  int              `json:"usage_count" gorm:"default:0" comment:"已使用次数"`
	Notes       string           `json:"notes" gorm:"type:text" comment:"备注信息"`
}

type LicenseKeyStatus string

const (
	StatusInactive LicenseKeyStatus = "inactive" // 未激活
	StatusActive   LicenseKeyStatus = "active"   // 已激活
	StatusExpired  LicenseKeyStatus = "expired"  // 已过期
	StatusDisabled LicenseKeyStatus = "disabled" // 已禁用
)
