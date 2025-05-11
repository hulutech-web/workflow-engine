package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
)

type FlowService interface {
	Index(ctx *gin.Context) (*PageResult, error)
	List(ctx *gin.Context) ([]models.Flow, error)
	Store(ctx *gin.Context, part models.Flow) error
	Create(ctx *gin.Context) (error, []models.Template, []models.Flowtype)
	Update(ctx *gin.Context, part models.Flow) error
	Show(ctx *gin.Context, id int) *models.Flow
	Destroy(ctx *gin.Context, id int) error
	FlowDesign(ctx *gin.Context, id int) (error, models.Flow)
}

type flowService struct {
	db *gorm.DB
}

func (f *flowService) Index(ctx *gin.Context) (*PageResult, error) {
	var depts []models.Flow
	paginatorService := NewPaginatorServiceImpl(f.db, ctx)
	err, result := paginatorService.SearchByParams(nil, nil).ResultPagination(&depts)
	return result, err
}

func (f *flowService) List(ctx *gin.Context) ([]models.Flow, error) {
	return nil, nil
}

func (f *flowService) Store(ctx *gin.Context, part models.Flow) error {
	tx := f.db.Model(&models.Flow{}).Create(&part)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (f *flowService) Create(ctx *gin.Context) (error, []models.Template, []models.Flowtype) {
	var templates []models.Template
	var flowtypes []models.Flowtype
	f.db.Model(&models.Template{}).Find(&templates)
	f.db.Model(&models.Flowtype{}).Find(&flowtypes)
	return nil, templates, flowtypes
}
func (f *flowService) Update(ctx *gin.Context, part models.Flow) error {
	return nil
}
func (f *flowService) Show(ctx *gin.Context, id int) *models.Flow {
	return nil
}
func (f *flowService) FlowDesign(ctx *gin.Context, id int) (error, models.Flow) {
	flow := models.Flow{}
	tx := f.db.Model(&models.Flow{}).Where("id=?", id).Preload("Processes").Find(&flow)
	if tx.Error != nil {
		return tx.Error, flow
	}
	return nil, flow
}
func (f *flowService) Destroy(ctx *gin.Context, id int) error {
	return nil
}

func NewFlowService(db *gorm.DB) FlowService {
	return &flowService{db: db}
}
