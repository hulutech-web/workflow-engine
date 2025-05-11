package orm

import (
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	m := dst()
	err := db.AutoMigrate(m...)
	if err != nil {
		return err
	}
	InitTestUser(db)
	return nil
}

func dst() []interface{} {
	return []interface{}{
		&models.User{},
		&models.AuthTenant{},
		&models.AuthMenu{},
		&models.AuthRole{},
		&models.AuthPerm{},
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

func InitTestUser(db *gorm.DB) {
	//先将users表清空
	db.Exec("truncate table users")
	salt := util.ToolsUtil.RandomString(5)
	pwd := util.ToolsUtil.MakeMd5("admin888" + salt)
	users := []models.User{
		models.User{
			Username: "admin",
			Password: pwd,
			Nickname: "管理员",
			Phone:    "18888888888",
			Email:    "admin@admin.com",
			Salt:     salt,
			RoleId:   1,
			TenantId: 1,
		},
		models.User{
			Username: "user",
			Password: pwd,
			Nickname: "普通用户",
			Phone:    "19999999999",
			Email:    "user@user.com",
			Salt:     salt,
			RoleId:   2,
			TenantId: 1,
		},
	}
	db.Create(&users)
}
