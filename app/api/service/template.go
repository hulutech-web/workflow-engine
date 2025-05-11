package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
)

/*
用来操作数据库查表
*/
type TemplateService interface {
	Index(ctx *gin.Context) (*PageResult, error)
	List(ctx *gin.Context) ([]models.Template, error)
	Store(ctx *gin.Context, part models.Template) error
	Update(ctx *gin.Context, part models.Template) error
	Show(ctx *gin.Context, id int) *models.Template
	Destroy(ctx *gin.Context, id uint) error
	TemplateForm(ctx *gin.Context, id int) ([]models.TemplateForm, error)
}

type templateService struct {
	db *gorm.DB
}

func (d templateService) Index(ctx *gin.Context) (*PageResult, error) {
	var tmpls []models.Template
	paginatorService := NewPaginatorServiceImpl(d.db, ctx)
	err, result := paginatorService.SearchByParams(nil, nil).ResultPagination(&tmpls)
	return result, err
}

func (d templateService) List(ctx *gin.Context) ([]models.Template, error) {
	depts := []models.Template{}
	d.db.Model(&models.Template{}).Find(&depts)
	return depts, nil
}

func (d templateService) Store(ctx *gin.Context, dept models.Template) error {

	tx := d.db.Model(&models.Template{}).Create(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d templateService) Update(ctx *gin.Context, dept models.Template) error {
	tx := d.db.Model(&models.Template{}).Where("id=?", dept.ID).Updates(&dept)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d templateService) Show(ctx *gin.Context, id int) *models.Template {
	dept := models.Template{}
	tx := d.db.Model(&models.Template{}).Where("id=?", id).First(&dept)
	if tx.Error == nil {
		return &dept
	}
	if tx.Error != nil {
		return nil
	}
	return nil

}

func (d templateService) Destroy(ctx *gin.Context, id uint) error {
	tx := d.db.Model(&models.Template{}).Where("id=?", id).Delete(&models.Template{
		Model: models.Model{
			ID: id,
		},
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (d templateService) TemplateForm(ctx *gin.Context, id int) ([]models.TemplateForm, error) {
	forms := []models.TemplateForm{}
	d.db.Model(&models.TemplateForm{}).Where("template_id=?", id).Find(&forms)
	return forms, nil
}
func NewTemplateService(db *gorm.DB) TemplateService {
	return &templateService{db: db}
}
