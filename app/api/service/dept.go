package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"gorm.io/gorm"
)

/*
用来操作数据库查表
*/
type DeptService interface {
	Index(ctx *gin.Context) (*PageResult, error)
	List(ctx *gin.Context) ([]models.Dept, error)
	Store(ctx *gin.Context, part models.Dept) (*models.Dept, error)
	Update(ctx *gin.Context, part models.Dept) (*models.Dept, error)
	Show(ctx *gin.Context, id int) *models.Dept
	Destroy(ctx *gin.Context, id int) error
	BindManager(ctx *gin.Context, manager_id int, dept_id int) (response.Response, error)
	BindDirector(ctx *gin.Context, director_id int, dept_id int) (response.Response, error)
}

type deptService struct {
	db *gorm.DB
}

func (d deptService) Index(ctx *gin.Context) (*PageResult, error) {
	var depts []models.Dept
	paginatorService := NewPaginatorServiceImpl(d.db, ctx)
	err, result := paginatorService.SearchByParams(nil, nil).ResultPagination(&depts)
	return result, err
}

func (d deptService) List(ctx *gin.Context) ([]models.Dept, error) {
	depts := []models.Dept{}
	d.db.Model(&models.Dept{}).Find(&depts)
	return depts, nil
}

func (d deptService) Store(ctx *gin.Context, dept models.Dept) (*models.Dept, error) {
	tx := d.db.Model(&models.Dept{}).Create(&dept)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dept, nil
}

func (d deptService) Update(ctx *gin.Context, dept models.Dept) (*models.Dept, error) {
	tx := d.db.Model(&models.Dept{}).Where("id=?", dept.ID).Save(&dept)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dept, nil

}

func (d deptService) Show(ctx *gin.Context, id int) *models.Dept {
	return nil

}

func (d deptService) Destroy(ctx *gin.Context, id int) error {
	return nil

}
func (d deptService) BindManager(ctx *gin.Context, manager_id int, dept_id int) (response.Response, error) {
	return response.Response{}, nil

}
func (d deptService) BindDirector(ctx *gin.Context, director_id int, dept_id int) (response.Response, error) {
	return response.Response{}, nil
}

func NewDeptService(db *gorm.DB) DeptService {
	return &deptService{db: db}
}
