package orm

import (
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path"
)

func autoMigrate(db *gorm.DB) error {
	m := dst()
	err := db.AutoMigrate(m...)
	if err != nil {
		return err
	}
	if err := fillData(db, m); err != nil {
		return err
	}
	if err := modifyFieldType(db); err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"modifyFieldType": "修改部分表字段格式",
	}).Infof("修改表")
	return nil
}

// 修改相关数据表字段类型
func modifyFieldType(db *gorm.DB) error {
	err := db.Exec("ALTER TABLE flowlinks MODIFY COLUMN next_process_id INT DEFAULT 2").Error
	err = db.Exec("ALTER TABLE flowlinks MODIFY COLUMN process_id INT DEFAULT 2").Error
	if err != nil {
		return err
	}
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
		models.Log{},
	}
}

func fillData(db *gorm.DB, m []interface{}) error {
	for _, obj := range m {
		var count int64
		if err := db.Model(obj).Count(&count).Error; err != nil {
			zap.S().Error(err.Error())
			continue
		}
		if count > 0 {
			continue
		}
		tableName := DBTableName(db, obj)
		if checkSQLFile(tableName) {
			sqlContent, err := readSQLFile(tableName)
			if err != nil {
				zap.S().Error("读取SQL文件失败：%s", err.Error())
				continue
			}
			if err = db.Exec(sqlContent).Error; err != nil {
				zap.S().Error("执行SQL文件失败：%s", err.Error())
				continue
			}
		}
	}
	return nil

}

func DBTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	return stmt.Schema.Table
}

// 判断sql文件是否存在，如果文件不存在则返回错误，避免不必要的处理
func checkSQLFile(tableName string) bool {
	filePath, err := getSQLFilePath(tableName)
	if err != nil {
		zap.S().Error(err.Error())
		return false
	}
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 提取获取 SQL 文件路径的逻辑到单独的函数中，避免代码重复
func getSQLFilePath(tableName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(wd, "public", "sql", tableName+".sql"), nil
}

func readSQLFile(tableName string) (string, error) {
	sqlFilePath, err := getSQLFilePath(tableName)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		return "", err
	}
	//replacedContent := strings.ReplaceAll(string(data), "{table_name}", tableName)
	return string(data), nil
}
