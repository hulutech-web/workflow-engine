package workflow

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/service/common"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"net/url"
)

type FlowService interface {
	Index(ctx *gin.Context, query url.Values) (*common.PageResult, error)
	List(ctx *gin.Context) ([]models.Flow, error)
	Store(ctx *gin.Context, part models.Flow) error
	Create(ctx *gin.Context) (error, []models.Template, []models.Flowtype)
	Update(ctx *gin.Context, part models.Flow) error
	Show(ctx *gin.Context, id int) *models.Flow
	Destroy(ctx *gin.Context, id int) error
	FlowDesign(ctx *gin.Context, id int) (error, models.Flow)
	Publish(ctx *gin.Context, flow_id int) error
}

type flowService struct {
	db *gorm.DB
}

func (d flowService) Index(ctx *gin.Context, query url.Values) (*common.PageResult, error) {
	var tmpls []models.Flow
	paginatorService := common.NewPaginatorServiceImpl(d.db, ctx)

	err, result := paginatorService.SearchByParams(query, nil).ResultPagination(&tmpls)
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
	f.db.Model(&models.Flow{}).Where("id=?", part.ID).Save(&part)
	return nil
}
func (f *flowService) Show(ctx *gin.Context, id int) *models.Flow {
	flow := models.Flow{}
	f.db.Model(&models.Flow{}).Where("id=?", id).Find(&flow)
	return &flow
}
func (f *flowService) FlowDesign(ctx *gin.Context, id int) (error, models.Flow) {
	flow := models.Flow{}
	tx := f.db.Model(&models.Flow{}).Where("id=?", id).Preload("Processes").Find(&flow)
	if tx.Error != nil {
		return tx.Error, flow
	}
	return nil, flow
}
func (f *flowService) Publish(ctx *gin.Context, flow_id int) error {
	flow := models.Flow{}
	f.db.Model(&models.Flow{}).Where("id=?", flow_id).Find(&flow)

	//如果设置了多个个开始步骤
	process_starts := []models.Process{}
	f.db.Model(&models.Process{}).Where("flow_id=?", flow_id).Where("position=?", 0).Find(&process_starts)
	if len(process_starts) > 1 {
		return errors.New("发布失败，只能设置一个开始步骤")
	}
	var fkCount1 int64
	f.db.Model(&models.Flowlink{}).Where("flow_id=?", flow_id).Where("type=?", "Condition").
		Count(&fkCount1)
	if fkCount1 <= 1 {
		return errors.New("发布失败，至少需要两个步骤")
	}

	var fkCount2 int64
	f.db.Model(&models.Flowlink{}).Where("flow_id=?", flow_id).Where("type=?", "Condition").
		Where("next_process_id=?", -1).Count(&fkCount2)
	if fkCount2 > 1 {
		return errors.New("发布失败，有步骤没有创建连接")
	}
	type Countf struct {
		Fid uint `json:"fid"`
		Pid uint `json:"pid"`
	}
	var flowlinkExists bool

	err := f.db.Table("flowlinks").
		Select("count(*) > 0").
		Joins("left join processes on flowlinks.process_id = processes.id").
		Where("flowlinks.flow_id = ?", flow_id).
		Where("processes.position = ?", 0).
		Scan(&flowlinkExists).
		Error

	if err != nil {
		return errors.New(fmt.Sprintf("关联查询错误:%s", err.Error()))
	}
	if !flowlinkExists {
		return errors.New("发布失败，请设置结束步骤")
	}
	flowlinks := []models.Flowlink{}
	err2 := f.db.Table("flowlinks").
		Select("flowlinks.*").
		Joins("JOIN processes ON flowlinks.process_id = processes.id").
		Where("flowlinks.flow_id = ?", flow_id).
		Where("flowlinks.type != ?", "Condition").
		Where("processes.position != ?", 0).
		Find(&flowlinks).
		Error

	if err2 != nil {
		return errors.New(fmt.Sprintf("关联查询错误:%s", err2.Error()))
	}
	for _, flowlink := range flowlinks {
		var cConditionMet bool
		var exists bool
		err3 := f.db.Table("flowlinks").
			Select("1").
			Joins("JOIN processes ON flowlinks.process_id = processes.id").
			Where("flowlinks.flow_id = ?", flow_id).
			Where("flowlinks.process_id = ?", flowlink.ProcessID).
			Where("flowlinks.type != ?", "Condition").
			Where("processes.position != ?", 0).
			Limit(1).
			Scan(&exists).
			Error

		if err3 != nil {
			return errors.New(fmt.Sprintf("flowlink关联查询错误:%s", err3.Error()))
		}
		cConditionMet = exists
		if !cConditionMet {

			return errors.New(fmt.Sprintf("发布失败，请给设置步骤审批权限:%s", flowlink.Auditor))
		}
	}

	flow.IsPublish = true
	f.db.Model(&models.Flow{}).Where("id=?", flow.ID).Save(&flow)
	return nil
}

func (f *flowService) Destroy(ctx *gin.Context, id int) error {

	f.db.Model(&models.Flow{}).Where("id=?", id).Delete(id, models.Flow{
		Model: models.Model{
			ID: cast.ToUint(id),
		},
	})
	return nil
}

func NewFlowService(db *gorm.DB) FlowService {
	return &flowService{db: db}
}
