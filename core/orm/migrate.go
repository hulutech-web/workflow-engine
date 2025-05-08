package orm

import (
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	m := dst()
	err := db.AutoMigrate(m...)
	if err != nil {
		return err
	}
	return nil
}

func dst() []interface{} {
	return []interface{}{
		&models.User{},
	}
}
