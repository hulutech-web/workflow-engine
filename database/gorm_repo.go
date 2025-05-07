package database

import (
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Find(out interface{}, where ...interface{}) error {
	return r.db.Find(out, where...).Error
}

func (r *GormRepository) First(out interface{}, where ...interface{}) error {
	return r.db.First(out, where...).Error
}

func (r *GormRepository) Create(value interface{}) error {
	return r.db.Create(value).Error
}

func (r *GormRepository) Update(value interface{}) error {
	return r.db.Save(value).Error
}

func (r *GormRepository) Delete(value interface{}) error {
	return r.db.Delete(value).Error
}

func (r *GormRepository) Where(query interface{}, args ...interface{}) Repository {
	return &GormRepository{db: r.db.Where(query, args...)}
}

func (r *GormRepository) Preload(column string, conditions ...interface{}) Repository {
	return &GormRepository{db: r.db.Preload(column, conditions...)}
}

func (r *GormRepository) Transaction(fc func(Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fc(&GormRepository{db: tx})
	})
}

func (r *GormRepository) Model(value interface{}) Repository {
	return &GormRepository{db: r.db.Model(value)}
}

func (r *GormRepository) Select(query interface{}, args ...interface{}) Repository {
	return &GormRepository{db: r.db.Select(query, args...)}
}

// GormProvider 数据库服务提供者
type GormProvider struct {
	DSN string
}

func (p *GormProvider) Register(c Container) {
	c.Singleton("db", func(c Container) interface{} {
		db, err := gorm.Open(mysql.Open(p.DSN), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return db
	})

	c.Bind("repository", func(c Container) interface{} {
		db, _ := c.Make("db")
		return NewGormRepository(db.(*gorm.DB))
	})
}

func (p *GormProvider) Boot(c Container) {
	// 数据库迁移可以在这里执行
}
