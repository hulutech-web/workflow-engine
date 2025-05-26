package workflow

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/service/common"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"net/url"
)

type FlowTypeService interface {
	Index(ctx *gin.Context, query url.Values) (*common.PageResult, error)
	List(ctx *gin.Context) ([]models.Flowtype, error)
	Store(ctx *gin.Context, part models.Flowtype) error
	Update(ctx *gin.Context, part models.Flowtype) error
	Show(ctx *gin.Context, id int) *models.Flowtype
	Destroy(ctx *gin.Context, id int) error
}

type flowTypeService struct {
	db *gorm.DB
}

func (d flowTypeService) Index(ctx *gin.Context, query url.Values) (*common.PageResult, error) {
	var tmpls []models.Flowtype
	paginatorService := common.NewPaginatorServiceImpl(d.db, ctx)

	err, result := paginatorService.SearchByParams(query, nil).ResultPagination(&tmpls)
	return result, err
}
func (f *flowTypeService) List(ctx *gin.Context) ([]models.Flowtype, error) {
	return nil, nil
}

func (f *flowTypeService) Store(ctx *gin.Context, part models.Flowtype) error {
	tx := f.db.Model(&models.Flow{}).Create(&part)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (f *flowTypeService) Create(ctx *gin.Context) (error, []models.Template, []models.Flowtype) {
	var templates []models.Template
	var flowtypes []models.Flowtype
	f.db.Model(&models.Template{}).Find(&templates)
	f.db.Model(&models.Flowtype{}).Find(&flowtypes)
	return nil, templates, flowtypes
}
func (f *flowTypeService) Update(ctx *gin.Context, part models.Flowtype) error {
	f.db.Model(&models.Flowtype{}).Where("id=?", part.ID).Save(&part)
	return nil
}
func (f *flowTypeService) Show(ctx *gin.Context, id int) *models.Flowtype {
	flow := models.Flowtype{}
	f.db.Model(&models.Flowtype{}).Where("id=?", id).Find(&flow)
	return &flow
}

func (f *flowTypeService) Destroy(ctx *gin.Context, id int) error {

	f.db.Model(&models.Flowtype{}).Where("id=?", id).Delete(id, models.Flowtype{
		Model: models.Model{
			ID: cast.ToUint(id),
		},
	})
	return nil
}

func NewFlowTypeService(db *gorm.DB) FlowTypeService {
	return &flowTypeService{db: db}
}
