package route

import (
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/spf13/cast"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"slices"
	"strings"
)

type Engine struct {
	fx.In
	Srv service.EngineService
	db  *gorm.DB
}

func workflowRoutes(w Engine, router *types.ApiRouter) {
	//	部门
	deptCtrl := controllers.NewDeptController()
	router.Resource("dept", deptCtrl)
	router.Get("dept/list", deptCtrl.List)
	router.Post("dept/bindmanager", deptCtrl.BindManager)
	router.Post("dept/binddirector", deptCtrl.BindDirector)

	//	员工
	empCtrl := controllers.NewEmpController()
	router.Resource("emp", empCtrl)
	router.Post("emp/search", empCtrl.Search)
	router.Get("emp/options", empCtrl.Options)
	router.Post("emp/bind", empCtrl.BindUser)
	//流程
	flowCtrl := controllers.NewFlowController()
	router.Resource("flow", flowCtrl)
	router.Get("flow/list", flowCtrl.List)
	router.Get("flow/create", flowCtrl.Create)
	//流程设计
	router.Get("flow/flowchart/{id}", flowCtrl.FlowDesign)
	router.Post("flow/publish", flowCtrl.Publish)

	//entry节点
	entryCtrl := controllers.NewEntryController()
	router.Get("flow/{id}/entry", entryCtrl.Create)
	router.Post("entry", entryCtrl.Store)
	router.Get("entry/{id}", entryCtrl.Show)
	router.Get("entry/{id}/entrydata", entryCtrl.EntryData)
	//流程重发
	router.Post("entry/resend", entryCtrl.Resend)
	//流程轨迹
	flowlinkCtrl := controllers.NewFlowlinkController()
	router.Post("flowlink", flowlinkCtrl.Update)

	//模板
	templateCtrl := controllers.NewTemplateController()
	router.Resource("template", templateCtrl)

	//模板控件
	templateformCtrl := controllers.NewTemplateformController()
	router.Get("template/{id}/templateform", templateformCtrl.Index)
	router.Post("templateform", templateformCtrl.Store)
	router.Put("templateform/{id}", templateformCtrl.Update)
	router.Delete("templateform/{id}", templateformCtrl.Destroy)
	router.Get("templateform/{id}", templateformCtrl.Show)
	router.Post("flow/templateform", templateformCtrl.FlowTemplateForm)

	//	流程
	processCtrl := controllers.NewProcessController()
	router.Resource("process", processCtrl)
	router.Get("process/attribute", processCtrl.Attribute)
	router.Post("process/con", processCtrl.Condition)
	router.Post("process/list", processCtrl.List)

	//	审批流转
	procCtrl := controllers.NewProcController()
	router.Get("proc/{entry_id}", procCtrl.Index)
	//同意
	router.Post("pass", procCtrl.Pass)
	//驳回
	router.Post("unpass", procCtrl.UnPass)
}

func (r *EntryController) Create(ctx http.Context) http.Response {
	flow_id := ctx.Request().RouteInt("id")
	var flow models.Flow
	facades.Orm().Query().Model(&models.Flow{}).Where("id", flow_id).
		With("Template.TemplateForms").Find(&flow)
	return httpfacades.NewResult(ctx).Success("", flow)
}

func (r *EntryController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *EntryController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	var entry models.Entry
	facades.Orm().Query().Model(&models.Entry{}).With("EntryDatas").With("Flow.Template.TemplateForms").Where("id", id).Find(&entry)
	return httpfacades.NewResult(ctx).Success("", entry)
}

func (r *EntryController) EntryData(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	var entrydata []models.EntryData
	var entry models.Entry
	query := facades.Orm().Query()
	query.Model(&models.Entry{}).Where("id=?", id).Find(&entry)
	//当时子流程时，需要查找当前流程的父流程
	query.Model(&models.EntryData{}).Where("entry_id=?", id).OrWhere("entry_id=?", entry.Pid).Find(&entrydata)

	last_flowlink := models.Flowlink{}
	query.Model(&models.Flowlink{}).Where("next_process_id=?", entry.ProcessID).
		Where("type=?", "Condition").Find(&last_flowlink)
	plugin_configs := official_plugins.PluginConfig{}
	//找上一个process
	query.Model(&official_plugins.PluginConfig{}).Where("process_id=?", last_flowlink.ProcessID).Find(&plugin_configs)
	return httpfacades.NewResult(ctx).Success("", http.Json{
		"entry":          entry,
		"entrydata":      entrydata,
		"plugin_configs": plugin_configs,
	})
}

func (r *EntryController) Store(ctx http.Context) http.Response {
	//添加发起节点
	flow_id := ctx.Request().InputInt("flow_id")
	var user models.Emp
	facades.Auth(ctx).User(&user)

	flowlink := models.Flowlink{}
	facades.Orm().Query().Table("flowlinks").Where("flowlinks.flow_id=?", cast.ToUint(flow_id)).Where("flowlinks.type=?", "Condition").Join("left join processes on flowlinks.id=processes.id").
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
	facades.Orm().Query().Raw(dbSql).Scan(&flowlink)
	var withFlowlink models.Flowlink
	facades.Orm().Query().Model(&models.Flowlink{}).Where("id=?", flowlink.ID).
		With("Process").With("NextProcess").Find(&withFlowlink)
	//校验提交的数据
	validRule, validMsg := r.dynamicValidator.DynamicValidate(flow_id)
	validator, err := facades.Validation().Make(r.dynamicValidator.DynamicValidateField(ctx), validRule, validation.Messages(validMsg))
	if err != nil {
		return httpfacades.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), "")
	}
	if validator.Fails() {
		return httpfacades.NewResult(ctx).ValidError("", validator.Errors().All())
	}
	query := facades.Orm().Query()
	var entry models.Entry
	entry.Title = ctx.Request().Input("title")
	entry.FlowID = cast.ToUint(flow_id)
	entry.EmpID = user.ID
	entry.Circle = 1
	entry.Status = 0
	err = query.Model(&models.Entry{}).Create(&entry)

	var withEntry models.Entry
	query.Model(&models.Entry{}).Where("id=?", entry.ID).With("Flow").With("Emp.Dept").With("Procs").With("EnterProcess").
		Find(&withEntry)
	//进程初始化
	//第一步看是否指定审核人

	err = r.workflow.SetFirstProcessAuditor(withEntry, withFlowlink)

	//向entrydata中插入数据
	for key, val := range ctx.Request().All() {
		if key == "title" || key == "flow_id" {
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
				entryData.FieldName = key
				entryData.FieldValue = newVal
				query.Model(&models.EntryData{}).Create(&entryData)
			} else {
				var entryData models.EntryData
				entryData.FlowID = cast.ToInt(flow_id)
				entryData.EntryID = cast.ToInt(entry.ID)
				entryData.FieldName = key
				entryData.FieldValue = cast.ToString(val)
				query.Model(&models.EntryData{}).Create(&entryData)
			}
		}
	}
	if err != nil {
		return httpfacades.NewResult(ctx).Error(http.StatusInternalServerError, err.Error(), "")
	}
	//流程表单数据插入，需要goravel的验证规则
	return httpfacades.NewResult(ctx).Success("发起成功", entry)
}

func (r *EntryController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *EntryController) Destroy(ctx http.Context) http.Response {
	return nil
}

// 重发
func (r *EntryController) Resend(ctx http.Context) http.Response {
	entry_id := ctx.Request().Input("entry_id")
	entry := models.Entry{}
	query := facades.Orm().Query()
	query.Model(&models.Entry{}).Where("id=?", entry_id).Where("status=?", -1).With("Flow").With("Emp.Dept").With("Procs").With("EnterProcess").
		Find(&entry)

	flow := models.Flow{}

	query.Model(&models.Flow{}).Where("id=?", entry.FlowID).Where("is_publish=?", true).Find(&flow)
	if flow.ID == 0 {
		return httpfacades.NewResult(ctx).Error(http.StatusInternalServerError, "流程未发布，请检查", "")
	}
	var flowlink models.Flowlink

	sql := fmt.Sprintf("SELECT * FROM `flowlinks` WHERE `flow_id` = %d "+
		"AND EXISTS (SELECT 1 FROM `processes` WHERE `processes`.`id` = `flowlinks`.`process_id` AND `processes`.`position` = 0) ORDER BY `sort` ASC LIMIT 1;", entry.FlowID)
	query.Raw(sql).Scan(&flowlink)
	if flowlink.ID == 0 {
		return httpfacades.NewResult(ctx).Error(http.StatusInternalServerError, "节点关系错误，请检查", "")
	}
	var withFlowlink models.Flowlink
	facades.Orm().Query().Model(&models.Flowlink{}).Where("id=?", flowlink.ID).
		With("Process").With("NextProcess").Find(&withFlowlink)
	//零值更新
	var map_entry = make(map[string]interface{})
	map_entry["circle"] = entry.Circle + 1
	map_entry["child"] = 0
	map_entry["status"] = 0
	query.Model(&models.Entry{}).Where("id=?", entry.ID).Update(map_entry)
	newEntry := models.Entry{}
	query.Model(&models.Entry{}).Where("id=?", entry.ID).With("Flow").With("Emp.Dept").With("Procs").With("EnterProcess").Find(&newEntry)

	err := r.workflow.SetFirstProcessAuditor(newEntry, withFlowlink)
	if err != nil {
		return httpfacades.NewResult(ctx).Error(http.StatusInternalServerError, "系统错误，请检查", "")
	}
	return httpfacades.NewResult(ctx).Success("重发成功", entry)
}

type DynamicValidator struct {
}

func NewDynamicValidator() *DynamicValidator {
	return &DynamicValidator{}
}
func (r *DynamicValidator) DynamicValidate(flow_id int) (map[string]string, map[string]string) {
	var flow models.Flow
	facades.Orm().Query().Model(&models.Flow{}).Where("id", flow_id).Find(&flow)
	template := models.Template{}
	facades.Orm().Query().Model(&models.Template{}).Where("id=?", flow.TemplateID).Find(&template)
	if template.ID == 0 {
		return make(map[string]string), nil
	}
	template_forms := []models.TemplateForm{}

	facades.Orm().Query().Model(&models.TemplateForm{}).Where("template_id", template.ID).Find(&template_forms)
	var validateMap = make(map[string]string)
	var messageMap = make(map[string]string)
	ruleSlice := []string{"required", "string", "uint", "min_len", "max_len", "max", "min", "ne", "date", "file", "image", "number", "email", "slice"}
	for _, template_form := range template_forms {
		if template_form.FieldRules != nil {
			for _, rule := range template_form.FieldRules {
				//	如果ruleSlice中存在rule，则添加到validateMap中
				//最终形成的结构类似于required|uint|
				var truthRule string
				if slices.Contains(ruleSlice, rule.RuleName) {
					if rule.RuleName == "min_len" || rule.RuleName == "max_len" || rule.RuleName == "max" || rule.RuleName == "min" || rule.RuleName == "ne" {
						truthRule += fmt.Sprintf("%s:%s|", rule.RuleName, rule.RuleValue)
					} else {

						truthRule += fmt.Sprintf("%s|", rule.RuleName)
					}
					if rule.RuleName == "file" {
						//实际上是一个文本类型
						truthRule += fmt.Sprintf("%s|", "string")
					}
					if rule.RuleName == "required" {
						messageMap[fmt.Sprintf("%s.%s", template_form.Field, rule.RuleName)] = fmt.Sprintf("%s%s%s", "错误", rule.RuleTitle, rule.RuleValue)
					} else {
						messageMap[fmt.Sprintf("%s.%s", template_form.Field, rule.RuleName)] = fmt.Sprintf("%s[%s]%s", "错误", rule.RuleTitle, rule.RuleValue)
					}
					//将truthRule最后一个|去掉
					validateMap[template_form.Field] += truthRule
				}
			}
		}
	}
	//去掉validateMap中每个value的最后一个|
	for key, val := range validateMap {
		validateMap[key] = strings.TrimRight(val, "|")
	}
	return validateMap, messageMap
}

// 如果提交数据为int64类型，将其转换为int类型
func (r *DynamicValidator) DynamicValidateField(ctx http.Context) map[string]any {
	result := map[string]any{}
	requests := ctx.Request().All()
	for key, val := range requests {
		atype := reflect.TypeOf(val)
		if atype.Name() == "float64" {
			result[key] = int(val.(float64))
		} else {
			result[key] = val
		}
	}
	return result
}
