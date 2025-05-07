package database

import (
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	Find(out interface{}, where ...interface{}) error
	First(out interface{}, where ...interface{}) error
	Create(value interface{}) error
	Update(value interface{}) error
	Delete(value interface{}) error
	Where(query interface{}, args ...interface{}) Repository
	Preload(column string, conditions ...interface{}) Repository
	Transaction(fc func(Repository) error) error
	Model(value interface{}) Repository
	Select(query interface{}, args ...interface{}) Repository
}

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
