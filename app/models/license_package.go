package models

type LicensePackage struct {
	Model
	Name         string        `json:"name" gorm:"size:50;not null" comment:"套餐名称"`
	DurationDays int           `json:"duration_days" gorm:"not null" comment:"有效期天数"`
	MaxUsers     int           `json:"max_users" gorm:"not null" comment:"最大用户数"`
	Features     string        `json:"features" gorm:"not null" comment:"权限配置"`
	Price        int           `json:"price" gorm:"comment:'价格（单位：分）'"`
	Status       PackageStatus `json:"status" gorm:"index;default:'active'" comment:"状态（active/inactive）"`
}

type PackageStatus string

const (
	PackageActive   PackageStatus = "active"   // 启用
	PackageInactive PackageStatus = "inactive" // 停用
)
