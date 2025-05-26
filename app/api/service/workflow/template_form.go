package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
)

/*
用来操作数据库查表
*/
type TemplateFormService interface {
	Index(ctx *gin.Context) (*PageResult, error)
	List(ctx *gin.Context) ([]models.TemplateForm, error)
	Store(ctx *gin.Context, part models.TemplateForm) error
	Update(ctx *gin.Context, part models.TemplateForm) error
	Show(ctx *gin.Context, id int) *models.TemplateForm
	Destroy(ctx *gin.Context, id int) error
}

type templateFormService struct {
	db *gorm.DB
}

func (d templateFormService) Index(ctx *gin.Context) (*PageResult, error) {
	var tmpls []models.TemplateForm
	paginatorService := NewPaginatorServiceImpl(d.db, ctx)
	err, result := paginatorService.SearchByParams(nil, nil).ResultPagination(&tmpls)
	return result, err
}

func (d templateFormService) List(ctx *gin.Context) ([]models.TemplateForm, error) {
	depts := []models.TemplateForm{}
	d.db.Model(&models.TemplateForm{}).Find(&depts)
	return depts, nil
}

func (d templateFormService) Store(ctx *gin.Context, dept models.TemplateForm) error {

	tx := d.db.Model(&models.TemplateForm{}).Create(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d templateFormService) Update(ctx *gin.Context, dept models.TemplateForm) error {
	tx := d.db.Model(&models.TemplateForm{}).Where("id=?", dept.ID).Updates(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d templateFormService) Show(ctx *gin.Context, id int) *models.TemplateForm {
	dept := models.TemplateForm{}
	tx := d.db.Model(&models.TemplateForm{}).Where("id=?", id).First(&dept)
	if tx.Error == nil {
		return &dept
	}
	if tx.Error != nil {
		return nil
	}
	return nil

}

func (d templateFormService) Destroy(ctx *gin.Context, id int) error {
	tx := d.db.Model(&models.TemplateForm{}).Where("id=?", id).Delete(&models.TemplateForm{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func NewTemplateFormService(db *gorm.DB) TemplateFormService {
	return &templateFormService{db: db}
}
