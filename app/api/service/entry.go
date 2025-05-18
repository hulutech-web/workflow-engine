package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service/workflow"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"strings"
)

/*
用来操作数据库查表
*/
type EntryService interface {
	Create(ctx *gin.Context) (models.Flow, error)
	Store(ctx *gin.Context, r req.FlowIDReq) error
	Show(ctx *gin.Context, id int) *models.Entry
	EntryData(ctx *gin.Context, id int) (error, map[string]interface{})
	Resend(ctx *gin.Context, entry req.EntryIDReq) error
}

type entryService struct {
	db    *gorm.DB
	cache *cache.Redis
	wf    workflow.EngineImpl
}

func (d entryService) Create(ctx *gin.Context) (models.Flow, error) {
	flow_id := ctx.Param("id")
	var flow models.Flow
	d.db.Model(&models.Flow{}).Where("id", flow_id).
		Preload("Template.TemplateForms").Find(&flow)
	return flow, nil
}

func (d entryService) Store(ctx *gin.Context, r req.FlowIDReq) error {
	flow_id := r.FlowID
	res := req.GetAuth(ctx)
	user_id := res.UserId
	var user models.Emp

	query := d.db
	query.Model(&models.Emp{}).Where("id=?", user_id).First(&user)
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
	//validRule, validMsg := r.dynamicValidator.DynamicValidate(flow_id)
	//validator, err := facades.Validation().Make(r.dynamicValidator.DynamicValidateField(ctx), validRule, validation.Messages(validMsg))
	//if err != nil {
	//	return nil, err
	//}
	//if validator.Fails() {
	//	return nil, err
	//}
	var entry models.Entry
	entry.Title = ctx.Query("title")
	entry.FlowID = cast.ToUint(flow_id)
	entry.EmpID = user.ID
	entry.Circle = 1
	entry.Status = 0
	query.Model(&models.Entry{}).Create(&entry)

	var withEntry models.Entry
	query.Model(&models.Entry{}).Where("id=?", entry.ID).Preload("Flow").Preload("Emp.Dept").Preload("Procs").Preload("EnterProcess").
		Find(&withEntry)
	//进程初始化
	//第一步看是否指定审核人

	err := d.wf.SetFirstProcessAuditor(withEntry, withFlowlink)
	queryParams := ctx.Request.URL.Query()
	for key, values := range queryParams {
		if key == "title" || key == "flow_id" {
			continue
		}

		// 取第一个值或拼接多个值
		var newVal string
		if len(values) == 1 {
			newVal = values[0]
		} else {
			newVal = strings.Join(values, ",")
		}

		entryData := models.EntryData{
			FlowID:     cast.ToInt(flow_id),
			EntryID:    cast.ToInt(entry.ID),
			FieldName:  key,
			FieldValue: newVal,
		}
		query.Model(&models.EntryData{}).Create(&entryData)
	}

	/*for key, values := range queryParams {
		// 只取第一个值（如果同一个 key 有多个值）
		p[key] = values[0]
	}
	//向entrydata中插入数据
	for key, val := range queryParams {
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
	}*/
	if err != nil {
		return err
	}
	return nil
}

func (d entryService) Show(ctx *gin.Context, id int) *models.Entry {
	var entry models.Entry
	d.db.Model(&models.Entry{}).Where("id=?", id).Preload("EntryDatas").Preload("Flow.Template.TemplateForms").Find(&entry)

	return &entry
}

func (d entryService) EntryData(ctx *gin.Context, id int) (error, map[string]interface{}) {
	var entrydata []models.EntryData
	var entry models.Entry
	query := d.db
	query.Model(&models.Entry{}).Where("id=?", id).Find(&entry)
	//当时子流程时，需要查找当前流程的父流程
	query.Model(&models.EntryData{}).Where("entry_id=?", id).Or("entry_id=?", entry.Pid).Find(&entrydata)

	last_flowlink := models.Flowlink{}
	query.Model(&models.Flowlink{}).Where("next_process_id=?", entry.ProcessID).
		Where("type=?", "Condition").Find(&last_flowlink)
	return nil, map[string]interface{}{
		"entry":     entry,
		"entrydata": entrydata,
	}
}

func (d *entryService) Resend(ctx *gin.Context, et req.EntryIDReq) error {
	entry := models.Entry{}
	query := d.db
	query.Model(&models.Entry{}).Where("id=?", et.EntryID).Where("status=?", -1).
		Preload("Flow").Preload("Emp.Dept").Preload("Procs").Preload("EnterProcess").
		Find(&entry)

	flow := models.Flow{}

	query.Model(&models.Flow{}).Where("id=?", entry.FlowID).Where("is_publish=?", true).Find(&flow)
	if flow.ID == 0 {
		return errors.New("流程未发布，请检查")
	}
	var flowlink models.Flowlink

	sql := fmt.Sprintf("SELECT * FROM `flowlinks` WHERE `flow_id` = %d "+
		"AND EXISTS (SELECT 1 FROM `processes` WHERE `processes`.`id` = `flowlinks`.`process_id` AND `processes`.`position` = 0) ORDER BY `sort` ASC LIMIT 1;", entry.FlowID)
	query.Raw(sql).Scan(&flowlink)
	if flowlink.ID == 0 {
		return errors.New("节点关系错误，请检查")
	}
	var withFlowlink models.Flowlink
	query.Model(&models.Flowlink{}).Where("id=?", flowlink.ID).
		Preload("Process").Preload("NextProcess").Find(&withFlowlink)
	//零值更新
	var map_entry = make(map[string]interface{})
	map_entry["circle"] = entry.Circle + 1
	map_entry["child"] = 0
	map_entry["status"] = 0
	query.Model(&models.Entry{}).Where("id=?", entry.ID).Updates(map_entry)
	newEntry := models.Entry{}
	query.Model(&models.Entry{}).Where("id=?", entry.ID).
		Preload("Flow").Preload("Emp.Dept").Preload("Procs").Preload("EnterProcess").Find(&newEntry)

	err := d.wf.SetFirstProcessAuditor(newEntry, withFlowlink)
	if err != nil {
		return err
	}
	return nil
}
func NewEntryService(db *gorm.DB, cache *cache.Redis, wf workflow.EngineImpl) EntryService {
	//wf := workflow.NewEngin()
	return &entryService{
		db:    db,
		cache: cache,
		wf:    wf,
	}
}
