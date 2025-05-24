package orm

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	if err := initWorkflowData(db); err != nil {
		return err
	}
	if err := modifyNecessaryColumns(db); err != nil {
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
		&models.Config{},
		&models.User{},
		&models.AuthTenant{},
		&models.AuthMenu{},
		&models.AuthRole{},
		&models.AuthPerm{},
		&models.Dept{},
		&models.Emp{},
		&models.Entry{},
		&models.EntryData{},
		&models.Flow{},
		&models.Flowlink{},
		&models.Flowtype{},
		&models.Template{},
		&models.Proc{},
		&models.Process{},
		&models.ProcessVar{},
		&models.TemplateForm{},
		&models.Log{},
		&models.LicenseKey{},
		&models.LicensePackage{},
		&models.File{},
		&models.FileCate{},
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
			statements := strings.Split(sqlContent, ";")
			for _, stmt := range statements {
				stmt = strings.TrimSpace(stmt)
				if stmt == "" {
					continue
				}
				// 处理表名, 如果语句中{table_name}存在，则替换成实际表名
				stmt = strings.ReplaceAll(stmt, "{table_name}", tableName)
				if err = db.Exec(stmt).Error; err != nil {
					zap.S().Errorf("执行SQL语句失败：%s\n语句内容：%s", err.Error(), stmt)
				}
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

// 修改必要的模型中的字段，迁移
func modifyNecessaryColumns(query *gorm.DB) error {
	query.Exec("alter table flowlinks modify COLUMN  process_id int")
	query.Exec("alter table flowlinks modify COLUMN  next_process_id int")
	return nil
}

// 初始化流程引擎数据，便于测试
func initWorkflowData(query *gorm.DB) error {

	var isExistDept int64
	query.Model(&models.Dept{}).Count(&isExistDept)
	var isExistEmp int64
	query.Model(&models.Emp{}).Count(&isExistEmp)
	var isExistFlowtype int64
	query.Model(&models.Flowtype{}).Count(&isExistFlowtype)
	if isExistFlowtype == 0 && isExistDept == 0 && isExistEmp == 0 {
		//#region 部门部分
		/*部门部分*/
		query.Model(&models.Dept{}).Create(&models.Dept{
			DeptName:   "总部",
			PID:        0,
			ManagerID:  1,
			DirectorID: 1,
		})
		//2-技术部
		query.Model(&models.Dept{}).Create(&models.Dept{
			DeptName:   "技术部",
			PID:        1,
			Html:       "|-",
			ManagerID:  2,
			DirectorID: 3,
		})
		//3-财务部
		query.Model(&models.Dept{}).Create(&models.Dept{
			DeptName:   "财务部",
			PID:        1,
			Html:       "|-",
			ManagerID:  5,
			DirectorID: 6,
		})
		// 4-市场部
		query.Model(&models.Dept{}).Create(&models.Dept{
			DeptName:   "市场部",
			PID:        1,
			Html:       "|-",
			ManagerID:  8,
			DirectorID: 9,
		})
		//4-1市场部-销售部
		query.Model(&models.Dept{}).Create(&models.Dept{
			DeptName:   "市场部-销售部",
			PID:        4,
			Html:       "|-",
			ManagerID:  11,
			DirectorID: 12,
		})
		//4-2市场部-市场拓展部
		query.Model(&models.Dept{}).Create(&models.Dept{
			DeptName:   "市场部-市场拓展部",
			PID:        4,
			Html:       "|-",
			ManagerID:  14,
			DirectorID: 15,
		})
		//#endregtion
		//#region 员工部分

		password := util.ToolsUtil.MakeMd5("admin888")
		/*员工部分*/
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "董事长",
			WorkNo:   "0001",
			DeptID:   1,
			UserID:   1,
			Password: password,
		})
		//2-技术部
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "技术部-技术经理",
			WorkNo:   "10001",
			DeptID:   2,
			UserID:   2,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "技术部-技术主管",
			WorkNo:   "10002",
			DeptID:   2,
			UserID:   3,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "技术部-技术员",
			WorkNo:   "10003",
			DeptID:   2,
			UserID:   4,
			Password: password,
		})
		//3-财务部
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "财务部-经理",
			WorkNo:   "20001",
			DeptID:   3,
			UserID:   5,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "财务部-主管",
			WorkNo:   "20002",
			DeptID:   3,
			UserID:   6,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "财务部-财务员",
			WorkNo:   "20003",
			DeptID:   3,
			UserID:   7,
			Password: password,
		})
		// 4-市场部
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-经理",
			WorkNo:   "30001",
			Password: password,
			DeptID:   4,
			UserID:   8,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-主管",
			WorkNo:   "30002",
			DeptID:   4,
			UserID:   9,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-总部员工1",
			WorkNo:   "30003",
			DeptID:   4,
			UserID:   10,
			Password: password,
		})
		//4-1市场部-销售部
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-销售部-经理",
			WorkNo:   "30011",
			DeptID:   5,
			UserID:   11,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-销售部-主管",
			WorkNo:   "30012",
			DeptID:   5,
			UserID:   12,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-销售部-员工1",
			WorkNo:   "30013",
			DeptID:   5,
			UserID:   13,
			Password: password,
		})
		//4-1市场部-扩展部
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-扩展部-经理",
			WorkNo:   "30021",
			DeptID:   6,
			UserID:   14,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-扩展部-主管",
			WorkNo:   "30022",
			DeptID:   6,
			UserID:   15,
			Password: password,
		})
		query.Model(&models.Emp{}).Create(&models.Emp{
			Name:     "市场部-扩展部-员工1",
			WorkNo:   "30023",
			DeptID:   6,
			UserID:   16,
			Password: password,
		})
		//#endregtion

		/*流程类型*/
		var flowType models.Flowtype
		flowType.TypeName = "资金"
		query.Model(&flowType).Create(&flowType)
	}

	return nil
}
