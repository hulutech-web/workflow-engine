package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

/*
用来操作数据库查表
*/
type EntryService interface {
	Create(ctx *gin.Context) (models.Flow, error)
	Store(ctx *gin.Context) (*models.Entry, error)
	Show(ctx *gin.Context, id int) *models.Entry
	EntryData(ctx *gin.Context, id int) error
	Resend(ctx *gin.Context, id int, user_id int) error
}

type entryService struct {
	db *gorm.DB
}

func (d entryService) Create(ctx *gin.Context) (models.Flow, error) {
	flow_id := ctx.Param("id")
	var flow models.Flow
	d.db.Model(&models.Flow{}).Where("id", flow_id).
		Preload("Template.TemplateForms").Find(&flow)
	return flow, nil
}

func (d entryService) Store(ctx *gin.Context) (*models.Entry, error) {
	flow_id := ctx.Query("flow_id")
	var user models.Emp
	facades.Auth(ctx).User(&user)
	query := d.db
	flowlink := models.Flowlink{}
	query.Table("flowlinks").Where("flowlinks.flow_id=?", cast.ToUint(flow_id)).Where("flowlinks.type=?", "Condition").
		Joins("left join processes on flowlinks.id=processes.id").
		Where("processes.position=?", 0).Order("sort  ASC").Find(&flowlink)
	dbSql := fmt.Sprintf("SELECT * "+
		"FROM `flowlinks` "+
		"WHERE `flow_id` = %d "+
		"  AND `type` = 'Condition' "+
		"  AND EXISTS ("+
		"    SELECT 1 "+
		"    FROM `processes` "+
		"   WHERE `flowlinks`.`process_id` = `processes`.`id` "+
		"      AND `processes`.`position` = 0"+
		"  ) "+
		"ORDER BY `sort` ASC "+
		"LIMIT 1;", flow_id)
	query.Raw(dbSql).Scan(&flowlink)
	var withFlowlink models.Flowlink
	query.Model(&models.Flowlink{}).Where("id=?", flowlink.ID).
		Preload("Process").Preload("NextProcess").Find(&withFlowlink)
	//校验提交的数据
	validRule, validMsg := r.dynamicValidator.DynamicValidate(flow_id)
	validator, err := facades.Validation().Make(r.dynamicValidator.DynamicValidateField(ctx), validRule, validation.Messages(validMsg))
	if err != nil {
		return nil, err
	}
	if validator.Fails() {
		return nil, err
	}
	var entry models.Entry
	entry.Title = ctx.Query("title")
	entry.FlowID = cast.ToUint(flow_id)
	entry.EmpID = user.ID
	entry.Circle = 1
	entry.Status = 0
	err = query.Model(&models.Entry{}).Create(&entry)

	var withEntry models.Entry
	query.Model(&models.Entry{}).Where("id=?", entry.ID).Preload("Flow").Preload("Emp.Dept").Preload("Procs").Preload("EnterProcess").
		Find(&withEntry)
	//进程初始化
	//第一步看是否指定审核人

	err = r.workflow.SetFirstProcessAuditor(withEntry, withFlowlink)

	//向entrydata中插入数据
	for key, val := range ctx.Params {
		if string(key) == "title" || string(key) == "flow_id" {
			continue
		} else {
			//判断val的类型，如果是[]string,则转换为解析为字符串

			if reflect.TypeOf(val).Kind() == reflect.Slice {
				var sliceStr []string
				//将val解析为sliceStr
				for _, v := range val.([]interface{}) {
					sliceStr = append(sliceStr, cast.ToString(v))
				}
				var newVal string
				newVal = strings.Join(sliceStr, ",")
				var entryData models.EntryData
				entryData.FlowID = cast.ToInt(flow_id)
				entryData.EntryID = cast.ToInt(entry.ID)
				entryData.FieldName = string(key)
				entryData.FieldValue = newVal
				query.Model(&models.EntryData{}).Create(&entryData)
			} else {
				var entryData models.EntryData
				entryData.FlowID = cast.ToInt(flow_id)
				entryData.EntryID = cast.ToInt(entry.ID)
				entryData.FieldName = string(key)
				entryData.FieldValue = cast.ToString(val)
				query.Model(&models.EntryData{}).Create(&entryData)
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (d entryService) Show(ctx *gin.Context, id int) *models.Entry {
	var entry models.Entry
	d.db.Model(&models.Entry{}).Where("id=?", id).Find(&entry)

	return &entry
}

func (d entryService) EntryData(ctx *gin.Context, id int) error {
	tx := d.db.Model(&models.Entry{}).Where("id=?", id).Delete(&models.Entry{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (d entryService) Resend(ctx *gin.Context, id int, user_id int) error {
	tx := d.db.Model(&models.Entry{}).Where("id=?", id).Update("user_id", user_id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func NewEntryService(db *gorm.DB) EntryService {
	return &entryService{db: db}
}
