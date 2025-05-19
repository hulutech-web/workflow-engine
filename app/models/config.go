package models

type Config struct {
	Model
	Type     string `gorm:"default:'';comment:'类型''"`
	Name     string `gorm:"not null;default:'';comment:'键'"`
	Value    string `gorm:"type:text;comment:'值'"`
	TenantId uint   `gorm:"not null;default:0;comment:'租户ID'"`
}
