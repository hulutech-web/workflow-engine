package models

type User struct {
	Model
	Username string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(32);not null"`
	Email    string `json:"email" gorm:"type:varchar(125);"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);"`
	Nickname string `json:"nickname" gorm:"type:varchar(255);"`
	Status   uint8  `json:"status" gorm:"type:int;default:1"`
	RoleId   uint   `json:"role_id" gorm:"type:int;default:0"`
}
