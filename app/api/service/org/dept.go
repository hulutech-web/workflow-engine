package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"net/url"
)

/*
用来操作数据库查表
*/
type DeptService interface {
	Index(ctx *gin.Context, query url.Values) (*PageResult, error)
	List(ctx *gin.Context) ([]models.Dept, error)
	Store(ctx *gin.Context, part models.Dept) error
	Update(ctx *gin.Context, part models.Dept) error
	Show(ctx *gin.Context, id int) *models.Dept
	Destroy(ctx *gin.Context, id int) error
	BindManager(ctx *gin.Context, manager_id int, dept_id int) error
	BindDirector(ctx *gin.Context, director_id int, dept_id int) error
	DisplayTree(ctx *gin.Context, id int) ([]*models.Dept, error)
}

type deptService struct {
	db *gorm.DB
}

func (d deptService) Index(ctx *gin.Context, query url.Values) (*PageResult, error) {
	var tmpls []models.Dept
	paginatorService := NewPaginatorServiceImpl(d.db, ctx)

	err, result := paginatorService.SearchByParams(query, nil).ResultPagination(&tmpls, "Manager", "Director")
	return result, err
}
func (d deptService) List(ctx *gin.Context) ([]models.Dept, error) {
	depts := []models.Dept{}
	d.db.Model(&models.Dept{}).Find(&depts)
	return depts, nil
}

func (d deptService) Store(ctx *gin.Context, dept models.Dept) error {
	tx := d.db.Model(&models.Dept{}).Create(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d deptService) Update(ctx *gin.Context, dept models.Dept) error {
	tx := d.db.Model(&models.Dept{}).Where("id=?", dept.ID).Updates(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d deptService) Show(ctx *gin.Context, id int) *models.Dept {
	dept := models.Dept{}
	tx := d.db.Model(&models.Dept{}).Where("id=?", id).First(&dept)
	if tx.Error == nil {
		return &dept
	}
	if tx.Error != nil {
		return nil
	}
	return nil

}

func (d deptService) Destroy(ctx *gin.Context, id int) error {
	tx := d.db.Model(&models.Dept{}).Where("id=?", id).Delete(&models.Dept{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}
func (d deptService) BindManager(ctx *gin.Context, manager_id int, dept_id int) error {
	logrus.WithFields(logrus.Fields{
		"manager_id": manager_id,
		"dept_id":    dept_id,
	}).Info("返回成功")
	tx := d.db.Model(&models.Dept{}).Where("id = ?", dept_id).Update("manager_id", manager_id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (d deptService) BindDirector(ctx *gin.Context, director_id int, dept_id int) error {
	logrus.WithFields(logrus.Fields{
		"director_id": director_id,
		"dept_id":     dept_id,
	}).Info("返回成功")
	tx := d.db.Model(&models.Dept{}).Where("id = ?", dept_id).Update("director_id", director_id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (d deptService) DisplayTree(ctx *gin.Context, id int) ([]*models.Dept, error) {
	var depts []models.Dept
	d.db.Model(&models.Dept{}).Find(&depts)

	var dept models.Dept
	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Info("返回成功")
	res := dept.BuildDeptTree(depts, cast.ToUint(id))
	return res, nil
}

func NewDeptService(db *gorm.DB) DeptService {
	return &deptService{db: db}
}
