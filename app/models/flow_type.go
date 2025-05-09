package models

type Flowtype struct {
	Model
	TypeName string `gorm:"column:type_name;not null;default:''" json:"type_name"`
	Flows    []Flow `gorm:"-"`
}
