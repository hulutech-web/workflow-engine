package orm

import (
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	models := dst()
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}

func dst() []interface{} {
	return []interface{}{}
}
