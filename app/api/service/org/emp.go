package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
	"net/url"
)

/*
用来操作数据库查表
*/
type EmpService interface {
	Index(ctx *gin.Context, query url.Values) (*PageResult, error)
	List(ctx *gin.Context) ([]models.Emp, error)
	Store(ctx *gin.Context, part models.Emp) error
	Update(ctx *gin.Context, part models.Emp) error
	Show(ctx *gin.Context, id int) *models.Emp
	Destroy(ctx *gin.Context, id int) error
	BindUser(ctx *gin.Context, id int, user_id int) error
}

type empService struct {
	db *gorm.DB
}

func (d empService) Index(ctx *gin.Context, query url.Values) (*PageResult, error) {
	var tmpls []models.Emp
	paginatorService := NewPaginatorServiceImpl(d.db, ctx)

	err, result := paginatorService.SearchByParams(query, nil).ResultPagination(&tmpls, "Dept")
	return result, err
}

func (d empService) List(ctx *gin.Context) ([]models.Emp, error) {
	depts := []models.Emp{}
	d.db.Model(&models.Emp{}).Find(&depts)
	return depts, nil
}

func (d empService) Store(ctx *gin.Context, dept models.Emp) error {
	tx := d.db.Model(&models.Emp{}).Create(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d empService) Update(ctx *gin.Context, dept models.Emp) error {
	tx := d.db.Model(&models.Emp{}).Where("id=?", dept.ID).Updates(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d empService) Show(ctx *gin.Context, id int) *models.Emp {
	dept := models.Emp{}
	tx := d.db.Model(&models.Emp{}).Where("id=?", id).First(&dept)
	if tx.Error == nil {
		return &dept
	}
	if tx.Error != nil {
		return nil
	}
	return nil

}

func (d empService) Destroy(ctx *gin.Context, id int) error {
	tx := d.db.Model(&models.Emp{}).Where("id=?", id).Delete(&models.Emp{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (d empService) BindUser(ctx *gin.Context, id int, user_id int) error {
	tx := d.db.Model(&models.Emp{}).Where("id=?", id).Update("user_id", user_id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func NewEmpService(db *gorm.DB) EmpService {
	return &empService{db: db}
}
