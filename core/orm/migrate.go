package orm

import (
	"github.com/hulutech-web/workflow-engine/app/models"
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
	return []interface{}{
		models.User{},
		models.Dept{},
		models.Emp{},
		models.Entry{},
		models.EntryData{},
		models.Flow{},
		models.Flowlink{},
		models.Flowtype{},
		models.Template{},
		models.Proc{},
		models.Process{},
		models.ProcessVar{},
		models.TemplateForm{},
	}
}
