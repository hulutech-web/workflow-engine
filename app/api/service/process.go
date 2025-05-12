package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/workflow/common"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"strings"
)

type ProcessService interface {
	Index(ctx *gin.Context) (*PageResult, error)
	List(ctx *gin.Context, req req.ProcessReq) ([]models.Process, error)
	Store(ctx *gin.Context, req req.ProReq) error
	Update(ctx *gin.Context, id int, processRequest common.ProcessRequest) error
	Show(ctx *gin.Context, id int) *models.Process
	Destroy(ctx *gin.Context, id int) error
	Attribute(ctx *gin.Context, id int) (error, map[string]any)
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

func (f *processService) Store(ctx *gin.Context, req req.ProReq) error {

	flow_id := req.FlowID
	left := req.Left
	top := req.Top
	tx := f.db.Begin()
	var flow models.Flow
	tx.Model(&models.Flow{}).Where("id=?", flow_id).Find(&flow)

	//步骤一
	var process models.Process
	process.FlowID = cast.ToInt(flow_id)
	process.ProcessName = "新建流程"
	process.StyleWidth = 200
	process.StyleHeight = 48
	process.Style = fmt.Sprintf("width:200px;height:48px;line-height:30px;color:#66CDAA;left:%s;top:%s;", left, top)
	process.PositionLeft = left
	process.PositionTop = top
	create := tx.Model(&models.Process{}).Create(&process)
	if create.Error != nil {
		tx.Rollback()
		return create.Error
	}
	//步骤二
	jsMap := common.Plumb{}
	if flow.Jsplumb == "" {
		//添加属性
		jsMap.Total = 1
		jsMap.List = map[string]common.Node{}
		listMap := map[string]common.Node{}
		node := common.Node{
			ID:          cast.ToInt(process.ID),
			FlowId:      process.FlowID,
			ProcessName: process.ProcessName,
			ProcessTo:   "",
			Icon:        "",
			Style:       process.Style,
		}
		listMap[cast.ToString(process.ID)] = node
		jsMap.List = listMap
		strByte, _ := json.Marshal(jsMap)
		logrus.WithFields(logrus.Fields{
			"strByte": string(strByte),
		}).Info("strByte")
		flow.IsPublish = false
		flow.Jsplumb = string(strByte)
		tx.Model(&models.Flow{}).Where("id=?", flow_id).Save(&flow)
		tx.Commit()
		ctx.JSON(200, gin.H{
			"id":           process.ID,
			"flow_id":      process.FlowID,
			"process_name": process.ProcessName,
			"process_to":   "",
			"icon":         "",
			"style":        process.Style,
		})
		return nil

	} else {
		//jsMap的list属性为二维数组
		var jsMapTemp common.Plumb
		//将flow中的Jsplumb转换为jsMapTemp
		if err := json.Unmarshal([]byte(flow.Jsplumb), &jsMapTemp); err != nil {
			tx.Rollback()
		}

		node := common.Node{
			ID:          cast.ToInt(process.ID),
			FlowId:      process.FlowID,
			ProcessName: process.ProcessName,
			ProcessTo:   "",
			Icon:        "",
			Style:       process.Style,
		}
		jsMapTemp.List[cast.ToString(process.ID)] = node
		jsMap = jsMapTemp
		//转换jsMap为json
		strByte, _ := json.Marshal(jsMap)
		flow.Jsplumb = string(strByte)
		flow.IsPublish = false
		tx.Model(&models.Flow{}).Where("id=?", flow_id).Save(&flow)
		tx.Commit()
		ctx.JSON(200, gin.H{
			"id":           process.ID,
			"flow_id":      process.FlowID,
			"process_name": process.ProcessName,
			"process_to":   "",
			"icon":         "",
			"style":        process.Style,
		})
		return nil
	}
}

func (f *processService) Update(ctx *gin.Context, id int, processRequest common.ProcessRequest) error {
	tx := f.db.Begin()

	var process models.Process
	find1 := tx.Model(&models.Process{}).Where("id=?", id).Find(&process)
	if find1.Error != nil {
		tx.Rollback()
		return errors.New("数据错误1")
	}
	logrus.WithFields(logrus.Fields{
		"process": process,
	}).Info("更新流程Update")
	if processRequest.ProcessPosition == 9 {
		var count int64
		tx.Model(&models.Flowlink{}).Where("process_id=?", id).Count(&count)
		if count > 1 {
			return errors.New("该节点是分支节点，不能设置为结束或起始步骤")
		}
	}
	if processRequest.ProcessPosition == 0 {
		tx.Model(&models.Process{}).Where("flow_id=?", process.FlowID).Where("position", 0).Update("position", 1)
		tx.Model(&models.Process{}).Where("flow_id=?", process.FlowID).Update("position", 0)
	}
	process.ProcessName = processRequest.ProcessName
	process.StyleColor = processRequest.StyleColor
	process.StyleHeight = processRequest.StyleHeight
	process.StyleWidth = processRequest.StyleWidth
	process.Style = fmt.Sprintf("width:%dpx;height:%dpx;line-height:30px;color:%s;left:%s;top:%s;",
		process.StyleWidth, process.StyleHeight, process.StyleColor, process.PositionLeft, process.PositionTop)
	process.Icon = processRequest.StyleIcon
	process.Position = processRequest.ProcessPosition
	process.ChildFlowID = processRequest.ChildFlowId
	process.ChildAfter = processRequest.ChildAfter
	process.ChildBackProcess = processRequest.ChildBackProcess
	save := tx.Model(&models.Process{}).Where("id=?", id).Save(&process)
	if save.Error != nil {
		tx.Rollback()
		return errors.New("数据错误2")
	}
	// 同步更新jsplumb json数据
	var flow models.Flow
	find := tx.Model(&models.Flow{}).Where("id=?", process.FlowID).Preload("Template.TemplateForms").Find(&flow)
	if find.Error != nil {
		tx.Rollback()
		return errors.New("数据错误4")
	}

	jsMap := common.Plumb{}
	//flow.Jsplum解析为jsMap
	err4 := json.Unmarshal([]byte(flow.Jsplumb), &jsMap)
	logrus.WithFields(logrus.Fields{
		"jsMap": jsMap,
	}).Info("jsMap结构")
	if err4 != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("解析数据错误4:%s", err4.Error()))
	}

	//需要将jsMap读取出来，然后再写回去
	for key, val := range jsMap.List {
		if key == cast.ToString(process.ID) {
			jsMap.List[key] = common.Node{
				ID:          cast.ToInt(process.ID),
				FlowId:      process.FlowID,
				ProcessTo:   val.ProcessTo,
				ProcessName: processRequest.ProcessName,
				Icon:        processRequest.StyleIcon,
				Style: fmt.Sprintf("width:%dpx;height:%dpx;line-height:30px;color:%s;left:%s;top:%s;",
					processRequest.StyleWidth, processRequest.StyleHeight, processRequest.StyleColor, process.PositionLeft, process.PositionTop),
			}
		}
	}
	jsplumbByte, err5 := json.Marshal(jsMap)
	if err5 != nil {
		tx.Rollback()
		return errors.New("解析数据错误5")
	}
	//更新流程图
	flow.Jsplumb = string(jsplumbByte)
	tx.Model(&models.Flow{}).Where("id=?", flow.ID).Update("jsplumb", flow.Jsplumb)
	tx.Model(&models.Flow{}).Where("id=?", flow.ID).Update("IsPublish", false)

	//更新步骤 流转条件 process_condition
	//根据ProcessCondition中的每一项分组，然后将每一组id相同的数据找出，将表达式合并为一个fmt.Sprintf("%s%s%s", condition.Field, condition.Operator, condition.Value)
	var conditionsMap map[int][]common.ProcessCondition
	if len(processRequest.ProcessCondition) > 0 {
		conditionsMap = groupConditionsById(processRequest.ProcessCondition)
	}
	//根据提交的conditionsMap更新process_var中的数据，如果有新数据，则新增，前提是只针对类型为int字段的数据，
	for _, conditions := range conditionsMap {
		for _, condition := range conditions {
			if condition.Field != "" {
				var exists_count int64
				tx.Model(&models.ProcessVar{}).
					Where("flow_id=?", flow.ID).
					Where("process_id=?", id).
					Where("expression_field=?", condition.Field).Count(&exists_count)
				if exists_count == 0 {
					//新增一条
					var newProcessVar models.ProcessVar
					newProcessVar.FlowID = cast.ToInt(flow.ID)
					newProcessVar.ProcessID = id
					newProcessVar.ExpressionField = condition.Field
					tx.Model(&models.ProcessVar{}).Create(&newProcessVar)
				}
			}
		}
	}

	for key, conditions := range conditionsMap {
		jsonStr, _ := json.Marshal(conditions)
		tx.Model(&models.Flowlink{}).Where("id=?", key).Update("expression", jsonStr)
	}

	//@改，如果当前的processRequest.AutoPerson=="0",更新当前的步骤
	if processRequest.AutoPerson == "0" {
		tx.Model(&models.Flowlink{}).Where("flow_id=?", flow.ID).Where("process_id=?", id).
			Where("type=?", "Condition").Update("auditor", processRequest.AutoPerson)
	}
	//权限处理
	if processRequest.AutoPerson != "0" {

		var fk models.Flowlink
		tx.Model(&fk).Where("flow_id=?", flow.ID).Where("process_id=?", id).Where("type=?", "Sys").Find(&fk)
		if fk.ID != 0 {
			fk.Auditor = cast.ToString(processRequest.AutoPerson)
			update := tx.Model(&models.Flowlink{}).Where("id=?", fk.ID).Update("auditor", processRequest.AutoPerson)
			if update.Error != nil {
				tx.Rollback()
				return errors.New("数据错误7")
			}
		} else {
			tx.Model(&models.Flowlink{}).Create(&models.Flowlink{
				FlowID:        flow.ID,
				Type:          "Sys",
				ProcessID:     cast.ToUint(id),
				Auditor:       cast.ToString(processRequest.AutoPerson),
				NextProcessID: 0,
				Sort:          100,
			})
		}
		//更新当前flowlink的Audiitor

		//删除其他权限
		tx.Model(&models.Flowlink{}).Where("flow_id=?", flow.ID).Where("process_id=?", id).
			Where("type!=?", "Condition").Where("type!=?", "Sys").Delete(&models.Flowlink{})
	} else {
		//指定部门
		if len(processRequest.RangeDeptIds) > 0 {
			var fkdept models.Flowlink
			tx.Model(&fkdept).Where("flow_id=?", flow.ID).Where("process_id=?", id).Where("type=?", "Dept").Find(&fkdept)
			if fkdept.ID != 0 {
				//id组成的数组，然后转换为字符串
				auditor := ""
				for _, dept := range processRequest.RangeDeptIds {
					auditor += cast.ToString(dept) + ","
				}
				//取消最后一个,号
				auditor = strings.TrimSuffix(auditor, ",")
				fkdept.Auditor = auditor
				tx.Model(&models.Flowlink{}).Where("id=?", fkdept.ID).Update("auditor", fkdept.Auditor)
			} else {
				auditor := ""
				for _, dept := range processRequest.RangeDeptIds {
					auditor += cast.ToString(dept) + ","
				}
				//去掉最后一个,号
				auditor = strings.TrimSuffix(auditor, ",")
				tx.Model(&models.Flowlink{}).Create(&models.Flowlink{FlowID: flow.ID, Type: "Dept", ProcessID: cast.ToUint(id), Auditor: auditor, NextProcessID: 0, Sort: 100})
			}
		} else {
			//删除部门权限
			tx.Model(&models.Flowlink{}).Where("flow_id=?", flow.ID).Where("process_id=?", id).
				Where("type=?", "Dept").Delete(&models.Flowlink{})
		}
		//	指定员工
		if len(processRequest.RangeEmpIds) > 0 {
			var fkemp models.Flowlink
			tx.Model(&fkemp).Where("flow_id=?", flow.ID).Where("process_id=?", id).Where("type=?", "Emp").Find(&fkemp)
			if fkemp.ID != 0 {
				//id组成的数组，然后转换为字符串
				auditor := ""
				for _, emp := range processRequest.RangeEmpIds {
					auditor += cast.ToString(emp) + ","
				}
				auditor = strings.TrimSuffix(auditor, ",")
				fkemp.Auditor = auditor
				tx.Model(&models.Flowlink{}).Where("id=?", fkemp.ID).Update("auditor", fkemp.Auditor)
			} else {
				auditor := ""
				for _, emp := range processRequest.RangeEmpIds {
					auditor += cast.ToString(emp) + ","
				}
				auditor = strings.TrimSuffix(auditor, ",")
				tx.Model(&models.Flowlink{}).Create(&models.Flowlink{FlowID: flow.ID, Type: "Emp", ProcessID: cast.ToUint(id), Auditor: auditor, NextProcessID: 0, Sort: 100})
			}
		} else {
			//	删除
			tx.Model(&models.Flowlink{}).Where("flow_id=?", flow.ID).Where("process_id=?", id).Where("type=?", "Emp").Delete(&models.Flowlink{
				Model: models.Model{
					ID: cast.ToUint(id),
				},
			})
		}
	}
	tx.Commit()
	return nil
}
func (f *processService) Show(ctx *gin.Context, id int) *models.Process {
	return nil
}

func (f *processService) Destroy(ctx *gin.Context, id int) error {
	return nil
}
func (f *processService) Attribute(ctx *gin.Context, id int) (error, map[string]interface{}) {
	process := models.Process{}
	tx := f.db
	tx.Model(&models.Process{}).Where("id=?", id).Find(&process)

	//1- //当前步骤的下一步操作
	next_process := []models.Flowlink{}
	tx.Model(&models.Flowlink{}).Where("process_id=?", process.ID).
		Where("flow_id=?", process.FlowID).Where("type=?", "Condition").Preload("Process").
		Preload("NextProcess").Find(&next_process)
	next_process_ids := []int{}
	tx.Model(&models.Flowlink{}).Where("process_id=?", process.ID).
		Where("flow_id=?", process.FlowID).Where("type=?", "Condition").Preload("Process").
		Preload("NextProcess").Pluck("next_process_id", &next_process_ids)
	beixuan_process := []models.Flowlink{}
	tx.Model(&models.Flowlink{}).Where("flow_id=?", process.FlowID).
		Where("type=?", "Condition").Where("process_id !=?", process.ID).
		Where("process_id not in (?)", next_process_ids).Preload("Process").Preload("NextProcess").Find(&beixuan_process)

	//	2-流程模板 表单字段
	flow := models.Flow{}

	fields := []models.TemplateForm{}
	tx.Model(&models.Flow{}).Where("id=?", process.FlowID).Preload("Template").Find(&flow)
	if flow.Template.ID != 0 {
		tfId := flow.Template.ID
		tx.Model(&models.TemplateForm{}).Where("template_id=?", tfId).Find(&fields)
	}

	//3-当前选择员工
	select_emps := []models.Emp{}
	auditor_emp_flowlink := models.Flowlink{}
	tx.Model(&models.Flowlink{}).Where("process_id=?", process.ID).
		Where("type=?", "Emp").Select("auditor").Find(&auditor_emp_flowlink)
	//depts按照,拆分
	empsSlice := []string{}
	for _, emp := range strings.Split(auditor_emp_flowlink.Auditor, ",") {
		empsSlice = append(empsSlice, emp)
	}
	tx.Model(&models.Emp{}).Where("id in (?)", empsSlice).Find(&select_emps)
	//4 -flowlinks
	flowlink := models.Flowlink{}
	sys := "0"
	tx.Model(&models.Flowlink{}).Where("process_id = ?", process.ID).Where("flow_id=?", process.FlowID).
		Where("type=?", "Sys").Find(&flowlink)
	if flowlink.Auditor != "" {
		sys = flowlink.Auditor
	}

	// 5-部门
	select_depts := []models.Dept{}
	auditor_dept_flowlink := models.Flowlink{}
	tx.Model(&models.Flowlink{}).Where("type=?", "Dept").Where("process_id=?", process.ID).
		Select("auditor").Find(&auditor_dept_flowlink)
	//depts按照,拆分
	deptsSlice := []string{}
	for _, dept := range strings.Split(auditor_dept_flowlink.Auditor, ",") {
		deptsSlice = append(deptsSlice, dept)
	}
	tx.Model(&models.Dept{}).Where("id in (?)", deptsSlice).Find(&select_depts)

	// 6-flow
	flows := []models.Flow{}
	tx.Model(&models.Flow{}).Where("is_publish=?", 1).Where("id!=?", process.FlowID).Find(&flows)

	processes := []models.Process{}
	tx.Model(&models.Process{}).Where("flow_id=?", process.FlowID).Find(&processes)
	var count int64
	var can_child bool
	tx.Model(&models.Flowlink{}).Where("process_id=?", process.ID).Where("type=?", "Condition").
		Count(&count)
	if count == 1 {
		can_child = true
	}
	return nil, map[string]interface{}{
		"process":         process,
		"next_process":    next_process,
		"beixuan_process": beixuan_process,
		"fields":          fields,
		"select_emps":     select_emps,
		"sys":             sys,
		"select_depts":    select_depts,
		"flows":           flows,
		"processes":       processes,
		"can_child":       can_child,
	}
}
func (f *processService) Condition(ctx *gin.Context) error {
	return nil
}

func groupConditionsById(conditions []common.ProcessCondition) map[int][]common.ProcessCondition {
	grouped := make(map[int][]common.ProcessCondition)
	for _, condition := range conditions {
		grouped[condition.Id] = append(grouped[condition.Id], condition)
	}
	return grouped
}

func NewProcessService(db *gorm.DB) ProcessService {
	return &processService{db: db}
}
