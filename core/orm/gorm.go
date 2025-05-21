package orm

import (
	"fmt"
	"github.com/dromara/carbon/v2"
	"github.com/glebarez/sqlite"
	"github.com/hulutech-web/workflow-engine/core/config"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func NewOrm(config *config.Config) (*gorm.DB, error) {
	var err error
	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
		})
	conf := &gorm.Config{
		Logger:                 l,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Database.TablePrefix, // 表名前缀
			SingularTable: false,                       // 使用单一表名, eg. `User` => `user`
		},

		DisableForeignKeyConstraintWhenMigrating: true,
		// 关键新增配置
		NowFunc: func() time.Time {
			return carbon.Now().StdTime() // 使用 Carbon 的当前时间
		},
	}

	var db *gorm.DB
	switch config.Database.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Database,
			config.Database.Params)
		db, err = gorm.Open(mysql.Open(dsn), conf)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Database.Database), conf)
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s",
			config.Database.Host,
			config.Database.Port,
			config.Database.Username,
			config.Database.Password,
			config.Database.Database,
			config.Database.Params)
		db, err = gorm.Open(postgres.Open(dsn), conf)
	default:
		return nil, fmt.Errorf("不支持的数据库驱动: %s", config.Database.Driver)
	}
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %s", err.Error())
	}
	dist, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取sql.DB对象失败: %v", err)
	}
	dist.SetMaxOpenConns(config.Database.MaxOpenConns)
	dist.SetMaxIdleConns(config.Database.MaxIdleConns)
	dist.SetConnMaxLifetime(time.Duration(config.Database.MaxLifeTime) * time.Minute)
	if err = dist.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接失败: %v", err)
	}
	if config.Database.AutoMigrate {
		err = autoMigrate(db)
		if err != nil {
			return nil, fmt.Errorf("自动迁移失败: %v", err)
		}
	}
	return db, nil
}

var Module = fx.Provide(NewOrm)
