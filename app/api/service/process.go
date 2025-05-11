package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
)

type ProcessService interface {
	Index(ctx *gin.Context) (*PageResult, error)
	List(ctx *gin.Context, req req.ProcessReq) ([]models.Process, error)
	Store(ctx *gin.Context, part models.Process) error
	Update(ctx *gin.Context, part models.Process) error
	Show(ctx *gin.Context, id int) *models.Process
	Destroy(ctx *gin.Context, id int) error
	Attribute(ctx *gin.Context) error
	Condition(ctx *gin.Context) error
}

type processService struct {
	db *gorm.DB
}

func (f *processService) Index(ctx *gin.Context) (*PageResult, error) {
	var depts []models.Process
	paginatorService := NewPaginatorServiceImpl(f.db, ctx)
	err, result := paginatorService.SearchByParams(nil, nil).ResultPagination(&depts)
	return result, err
}

func (f *processService) List(ctx *gin.Context, req req.ProcessReq) ([]models.Process, error) {
	flow_id := req.FlowID
	processes := []models.Process{}
	f.db.Model(&models.Process{}).
		Where("flow_id=?", flow_id).Preload("Flow").Find(&processes)
	return processes, nil
}

func (f *processService) Store(ctx *gin.Context, part models.Process) error {
	tx := f.db.Model(&models.Process{}).Create(&part)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (f *processService) Update(ctx *gin.Context, part models.Process) error {
	return nil
}
func (f *processService) Show(ctx *gin.Context, id int) *models.Process {
	return nil
}

func (f *processService) Destroy(ctx *gin.Context, id int) error {
	return nil
}
func (f *processService) Attribute(ctx *gin.Context) error {
	return nil
}
func (f *processService) Condition(ctx *gin.Context) error {
	return nil
}

func NewProcessService(db *gorm.DB) ProcessService {
	return &processService{db: db}
}
