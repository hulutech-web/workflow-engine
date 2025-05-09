package workflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dromara/carbon/v2"
	"github.com/hulutech-web/workflow-engine/app/api/workflow/common"
	"github.com/hulutech-web/workflow-engine/app/api/workflow/official_plugin"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
	"sync"
)

type EngineImpl interface {
	RegisterHook(name string, method reflect.Value) error
	NotifySendOne(id uint) error
	NotifyNextAuditor(id uint) error
	invokeHooks(hookName string, id uint)
	SetFirstProcessAuditor(entry models.Entry, flowlink models.Flowlink) error
	ExecPluginMethod(plugin_name string, flowID uint, processID uint) error
	ClearHooks(name string)
	GoToProcess(entry models.Entry, processID int) error
}

type Engine struct {
	db    *gorm.DB
	hooks map[string][]reflect.Value // 修改为 存储多个钩子函数
	mutex sync.Mutex
}

// RegisterHook 注册钩子方法
func (w *Engine) RegisterHook(name string, hookFunc reflect.Value) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// 初始化 hooks map
	if w.hooks == nil {
		w.hooks = make(map[string][]reflect.Value)
	}

	// 验证钩子函数
	hook := reflect.ValueOf(hookFunc)
	if hook.Kind() != reflect.Func {
		return fmt.Errorf("hook must be a function")
	}

	// 检查参数签名
	hookType := hook.Type()
	if hookType.NumIn() != 1 || hookType.In(0).Kind() != reflect.Uint {
		return fmt.Errorf("hook must have signature func(uint)")
	}

	w.hooks[name] = append(w.hooks[name], hook)
	return nil
}

// NotifySendOne 调用 NotifySendOne 钩子
func (w *Engine) NotifySendOne(id uint) error {
	if w == nil {
		fmt.Println("Workflow instance is nil in NotifySendOne!")
		return fmt.Errorf("workflow instance is nil")
	}
	fmt.Printf("BaseWorkflow.NotifySendOne :%d\n", id)

	w.invokeHooks("NotifySendOneHook", id)

	return nil
}

func (w *Engine) ClearHooks(name string) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	delete(w.hooks, name)
}

// NotifyNextAuditor 调用 NotifyNextAuditor 钩子
func (w *Engine) NotifyNextAuditor(id uint) error {
	if w == nil {
		fmt.Println("Workflow instance is nil in NotifyNextAuditor!")
		return fmt.Errorf("workflow instance is nil")
	}
	fmt.Printf("BaseWorkflow.NotifyNextAuditor:%d\n", id)

	w.invokeHooks("NotifyNextAuditorHook", id)

	return nil
}

// invokeHooks 用于依次调用所有注册的钩子方法
func (w *Engine) invokeHooks(hookName string, id uint) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	hooks, ok := w.hooks[hookName]
	if !ok {
		log.Printf("[Hook] %s not registered", hookName)
		return
	}

	for i, hook := range hooks {
		// 更严格的方法签名检查
		methodType := hook.Type()
		if methodType.NumIn() != 1 || methodType.In(0).Kind() != reflect.Uint {
			log.Printf("[Hook] %s[%d] invalid signature: expected func(uint)", hookName, i)
			continue
		}

		// 安全调用
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("[Hook] %s[%d] panic: %v", hookName, i, r)
				}
			}()

			log.Printf("[Hook] Calling %s[%d]...", hookName, i)
			hook.Call([]reflect.Value{reflect.ValueOf(id)})
			log.Printf("[Hook] %s[%d] completed", hookName, i)
		}()
	}
}

func (w *Engine) SetFirstProcessAuditor(entry models.Entry, flowlink models.Flowlink) error {
	return w.db.Transaction(func(tx *gorm.DB) error {
		var myFlowlink models.Flowlink

		var auditor_ids []int
		tx.Model(&models.Flowlink{}).Where("type != ?", "Condition").
			Where("process_id=?", flowlink.ProcessID).Find(&myFlowlink)

		var process_id int
		var process_name string
		if myFlowlink.ID == 0 {

			//第一步未指定审核人 自动进入下一步操作
			var proc models.Proc
			proc.FlowID = cast.ToInt(entry.FlowID)
			proc.ProcessID = cast.ToInt(flowlink.ProcessID)
			proc.ProcessName = flowlink.Process.ProcessName
			proc.EmpID = cast.ToInt(entry.EmpID)
			proc.EmpName = entry.Emp.Name
			proc.DeptName = entry.Emp.Dept.DeptName
			proc.AuditorID = cast.ToInt(entry.EmpID)
			proc.AuditorName = entry.Emp.Name
			proc.AuditorDept = entry.Emp.Dept.DeptName
			proc.Status = 9
			proc.Circle = entry.Circle
			proc.Concurrence = carbon.NewDateTime(carbon.Now())
			proc.EntryID = entry.ID
			tx.Transaction(func(tx2 *gorm.DB) error {
				create := tx2.Model(&models.Proc{}).Create(&proc)
				if create.RowsAffected == 0 {
					tx.Rollback()
					return errors.New("create proc failed")
				}
				return nil
			})

			auditor_ids = w.GetProcessAuditorIds(entry, flowlink.NextProcessID)
			process_id = flowlink.NextProcessID
			process_name = flowlink.NextProcess.ProcessName
			entry.ProcessID = cast.ToUint(flowlink.NextProcessID)
		} else {

			auditor_ids = w.GetProcessAuditorIds(entry, cast.ToInt(flowlink.ProcessID))
			process_id = cast.ToInt(flowlink.ProcessID)
			process_name = flowlink.Process.ProcessName
			entry.ProcessID = cast.ToUint(flowlink.ProcessID)

		}
		//步骤流转
		//步骤审核人
		var auditors_emps []models.Emp
		w.db.Transaction(func(tx4 *gorm.DB) error {
			tx4.Model(&models.Emp{}).Where("id IN (?)", auditor_ids).Preload("Dept").Find(&auditors_emps)
			if len(auditors_emps) < 1 {
				return errors.New("下一步骤未找到审批人")
			}
			return nil
		})
		for _, emp := range auditors_emps {
			var proc2 models.Proc
			proc2.EntryID = entry.ID
			proc2.FlowID = cast.ToInt(entry.FlowID)
			proc2.ProcessID = process_id
			proc2.ProcessName = process_name
			proc2.EmpID = cast.ToInt(emp.ID)
			proc2.EmpName = emp.Name
			proc2.DeptName = emp.Dept.DeptName
			proc2.Status = 0
			proc2.Circle = entry.Circle
			proc2.Concurrence = carbon.NewDateTime(carbon.Now())
			w.db.Model(&models.Proc{}).Create(&proc2)
		}

		w.db.Model(models.Entry{}).Where("id=?", entry.ID).Save(&entry)
		return nil
	})
}

func (w *Engine) GetProcessAuditorIds(entry models.Entry, next_process_id int) []int {
	var auditor_ids []int
	var flowlink models.Flowlink
	query := w.db
	query.Model(&models.Flowlink{}).Where("type = ?", "Sys").Where("process_id=?", next_process_id).Find(&flowlink)
	if flowlink.ID > 0 {
		if flowlink.Auditor == "-1000" {
			//发起人
			auditor_ids = append(auditor_ids, cast.ToInt(entry.EmpID))
		}
		if flowlink.Auditor == "-1001" {
			//发起人部门主管
			if entry.Emp.Dept.ID == 0 {
				return auditor_ids
			}
			auditor_ids = append(auditor_ids, cast.ToInt(entry.Emp.Dept.DirectorID))
		}
		if flowlink.Auditor == "-1002" {
			//发起人部门经理
			if entry.Emp.Dept.ID == 0 {
				return auditor_ids
			}
			auditor_ids = append(auditor_ids, cast.ToInt(entry.Emp.Dept.ManagerID))
		}
	} else {
		//	concurrent 并行
		//	1、指定员工
		concurrent_emp_flowlink := models.Flowlink{}
		query.Model(&models.Flowlink{}).Where("type = ?", "Emp").Where("process_id=?", next_process_id).Find(&concurrent_emp_flowlink)
		if concurrent_emp_flowlink.ID > 0 {
			Auditor_ids := []string{}
			//按照,分割concurrent_flowlink.Auditor
			Auditor_ids = strings.Split(concurrent_emp_flowlink.Auditor, ",")
			for _, id := range Auditor_ids {
				auditor_ids = append(auditor_ids, cast.ToInt(id))
			}
		}
		//	2、指定部门（指定部门时，可能指定多个部门，分别找到部门的主管，并找到对应的emp_id）
		concurrent_dept_flowlink := models.Flowlink{}
		query.Model(&models.Flowlink{}).Where("type = ?", "Dept").Where("process_id=?", next_process_id).
			Find(&concurrent_dept_flowlink)

		if concurrent_dept_flowlink.ID > 0 {
			dept_id_strs := []string{}
			//按照,分割concurrent_flowlink.Auditor
			dept_id_strs = strings.Split(concurrent_dept_flowlink.Auditor, ",")
			dept_ids := []int{}
			for _, id := range dept_id_strs {
				dept_ids = append(dept_ids, cast.ToInt(id))
			}
			emp_ids := []int{}
			//默认查找部门主管director_id，它对应着员工的id
			query.Model(&models.Dept{}).Select("director_id").Where("id IN (?)", dept_ids).Pluck("director_id", &emp_ids)
			for _, id := range emp_ids {
				auditor_ids = append(auditor_ids, id)
			}
		}
		//	3、指定角色，暂时不需要
	}
	ret_auditor_ids := uniqueSlice(auditor_ids)
	//	对auditor_ids去重
	return ret_auditor_ids

}

// 辅助函数，从slice中去重
func uniqueSlice(slice []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, value := range slice {
		if _, ok := seen[value]; !ok {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}

func (w *Engine) GoToProcess(entry models.Entry, processID int) error {
	auditor_ids := w.GetProcessAuditorIds(entry, processID)
	auditors := []models.Emp{}
	w.db.Model(&models.Emp{}).Preload("Dept").Where("id in (?)", auditor_ids).Find(&auditors)

	if len(auditors) < 1 {
		return errors.New("未找到下一步步骤审批人")
	}
	current_time := carbon.NewDateTime(carbon.Now())
	processName := ""
	w.db.Model(&models.Process{}).Where("id=?", processID).Select("ProcessName").
		Scan(&processName)

	for _, auditor := range auditors {
		w.db.Model(&models.Proc{}).Create(&models.Proc{
			EntryID:     entry.ID,
			FlowID:      cast.ToInt(entry.FlowID),
			ProcessID:   cast.ToInt(processID),
			ProcessName: processName,
			EmpID:       cast.ToInt(auditor.ID),
			EmpName:     auditor.Name,
			DeptName:    auditor.Dept.DeptName,
			Circle:      entry.Circle,
			Status:      0,
			IsRead:      0,
			Concurrence: current_time,
		})
	}
	return nil
}

// 流转
func (w *Engine) Transfer(process_id int, user models.Emp, content string) error {
	tx := w.db
	var emp models.Emp
	tx.Model(&models.Emp{}).Preload("Dept").Where("user_id=?", user.ID).Find(&emp)
	var proc models.Proc
	tx.Model(&models.Proc{}).Preload("Entry.Emp.Dept").Preload("Entry.ParentEntry").Where("process_id=?", process_id).
		Where("emp_id=?", emp.ID).Where("status=?", 0).Find(&proc)
	if proc.ID == 0 {
		return errors.New("未绑定员工，请设置员工绑定")
	}
	var fkcount int64
	tx.Model(&models.Flowlink{}).Where("process_id=?", proc.ProcessID).Where("type=?", "Condition").Count(&fkcount)

	if fkcount > 1 {
		//	情况一：有条件
		pvar := models.ProcessVar{}
		tx.Model(&models.ProcessVar{}).Where("process_id=?", process_id).Find(&pvar)
		var field_value string
		tx.Model(&models.EntryData{}).Select("field_value").
			Where("entry_id=?", proc.EntryID).
			Where("field_name=?", pvar.ExpressionField).Pluck("field_value", &field_value)

		flowlinks := []models.Flowlink{}
		tx.Model(&models.Flowlink{}).Where("process_id=?", proc.ProcessID).
			Where("type=?", "Condition").Find(&flowlinks)
		var flowlink models.Flowlink //满足条件的flowlink
		field := pvar.ExpressionField
		for _, m := range flowlinks {
			if m.Expression == "" {
				return errors.New("未设置流转条件，无法流转，请联系流程设置人员")
			}

			if m.Expression == "1" {
				flowlink = m
				break
			} else {
				//m.Expression
				type ResultCount struct {
					Number int `json:"number"`
				}
				var resultCount ResultCount
				processConditions := []common.ProcessCondition{}
				json.Unmarshal([]byte(m.Expression), &processConditions)
				if len(processConditions) > 0 {
					//检查语法错误(使用mysql数条件表达式
					conditionSql := ""
					for _, condition := range processConditions {
						if condition.Field != field {
							return errors.New("没有该条件字段，请检查")
						} else {
							conditionSql += fmt.Sprintf(" `field_value` %s %s %s", condition.Operator, condition.Value, condition.Extra)
						}
					}
					conditionSql = fmt.Sprintf("SELECT count(*) as number FROM entrydatas WHERE entry_id=%d and flow_id=%d and (%s) and (`field_name`='%s')",
						proc.EntryID, proc.FlowID, conditionSql, field)
					//还需要条件entry_id和flow_id
					err := w.db.Raw(conditionSql).Scan(&resultCount)
					if err != nil {
						return errors.New("条件语法错误，请检查")
					}
					if resultCount.Number > 0 {
						flowlink = m
						break
					}
				}
			}
		}
		if flowlink.ID == 0 {
			return errors.New("未找到符合条件的流转条件，无法流转")
		}
		var PreloadFlowlink models.Flowlink
		w.db.Model(&models.Flowlink{}).Preload("NextProcess").Where("id=?", flowlink.ID).Find(&PreloadFlowlink)
		auditor_ids := w.GetProcessAuditorIds(*proc.Entry, PreloadFlowlink.NextProcessID)
		if len(auditor_ids) == 0 {
			return errors.New("未找到下一步骤审批人")
		}
		auditors := []models.Emp{}
		tx.Model(&models.Emp{}).Where("id IN (?)", auditor_ids).Preload("Dept").Find(&auditors)
		if len(auditors) == 0 {
			return errors.New("未找到下一步骤审批人")
		}
		curr_time := carbon.NewDateTime(carbon.Now())
		for _, auditor := range auditors {
			tx.Model(&models.Proc{}).Create(&models.Proc{
				EntryID:     proc.EntryID,
				FlowID:      cast.ToInt(proc.FlowID),
				ProcessID:   PreloadFlowlink.NextProcessID,
				ProcessName: PreloadFlowlink.NextProcess.ProcessName,
				EmpID:       cast.ToInt(auditor.ID),
				EmpName:     auditor.Name,
				DeptName:    auditor.Dept.DeptName,
				Circle:      proc.Entry.Circle,
				Status:      0,
				IsRead:      0,
				Concurrence: curr_time,
			})
			//通知下一个审批人
			//通知发起人，被驳回
			w.NotifyNextAuditor(auditor.ID)
		}
		procEntry := models.Entry{}
		tx.Model(&models.Entry{}).Where("id=?", proc.EntryID).Find(&procEntry)
		procEntry.ProcessID = cast.ToUint(flowlink.NextProcessID)
		tx.Model(&models.Entry{}).Where("id=?", procEntry.ID).Save(&procEntry)
		//判断是否存在父进程
		if proc.Entry.Pid > 0 {
			proc2Entry := models.Entry{}
			tx.Model(&models.Entry{}).Where("id=?", proc.EntryID).Find(&proc2Entry)
			partentEntry := models.Entry{}
			tx.Model(&models.Entry{}).Where("pid=?", proc.ID).Find(&partentEntry)
			partentEntry.Child = flowlink.NextProcessID
			tx.Model(&models.Entry{}).Where("id=?", partentEntry.ID).Save(&partentEntry)
		}
	} else {
		fklink := models.Flowlink{}
		tx.Model(&models.Flowlink{}).Preload("Process").Preload("NextProcess").Where("process_id=?", proc.ProcessID).
			Where("type=?", "Condition").Find(&fklink)
		if fklink.Process.ChildFlowID > 0 {
			// 创建子流程
			child_entry := models.Entry{}
			tx.Model(&models.Entry{}).Preload("Flow").
				Preload("Process").Preload("EnterProcess").
				Preload("Emp.Dept").
				Where("pid=?", proc.Entry.ID).
				Where("circle=?", proc.Entry.Circle).Find(&child_entry)
			if child_entry.ID == 0 {
				new_child_entry := models.Entry{
					Title:          proc.Entry.Title,
					FlowID:         cast.ToUint(fklink.Process.ChildFlowID),
					EmpID:          cast.ToUint(proc.Entry.EmpID),
					Status:         0,
					Pid:            cast.ToInt(proc.Entry.ID),
					Circle:         proc.Entry.Circle,
					EnterProcessID: cast.ToInt(fklink.ProcessID),
					EnterProcID:    cast.ToInt(proc.ID),
				}
				tx.Model(&models.Entry{}).Create(&new_child_entry)

				tx.Model(&models.Entry{}).Where("id=?", new_child_entry.ID).
					Preload("Flow").
					Preload("Process").Preload("EnterProcess").
					Preload("Emp.Dept").
					Find(&new_child_entry)
				child_entry = new_child_entry
			}

			child_flowlink := models.Flowlink{}
			exec_sql := "SELECT * FROM flowlinks AS f " +
				"WHERE f.flow_id = (SELECT child_flow_id FROM processes WHERE id = ? AND f.type = 'Condition' " +
				"AND EXISTS (SELECT * FROM processes AS p WHERE p.id = f.process_id AND p.position = 0) " +
				"ORDER BY f.sort ASC " +
				"LIMIT 1);"
			tx.Raw(exec_sql, fklink.ProcessID).Scan(&child_flowlink)
			tx.Model(&models.Flowlink{}).Where("id=?", child_flowlink.ID).Preload("Process").
				Preload("NextProcess").Find(&child_flowlink)
			err := w.SetFirstProcessAuditor(child_entry, child_flowlink)
			if err != nil {
				return err
			}
			tx.Model(&models.Entry{}).Where("id=?", child_entry.Pid).Update("child", child_entry.ProcessID)
		} else {
			if fklink.NextProcessID == -1 {
				//最后一步
				tx.Model(&models.Entry{}).Where("id=?", proc.EntryID).Save(models.Entry{
					Status:    9,
					ProcessID: fklink.ProcessID,
				})

				if proc.Entry.Pid > 0 {
					if proc.Entry.EnterProcess.ChildAfter == 1 {
						//同时结束父流程
						parentEntry := models.Entry{}
						tx.Model(&models.Entry{}).Where("id=?", proc.Entry.Pid).Find(&parentEntry)
						map_entry := make(map[string]interface{})
						map_entry["status"] = 9
						map_entry["child"] = 0
						tx.Model(&models.Entry{}).Where("id=?", parentEntry.ID).Save(&map_entry)
						//通知发起人，审批结束
						w.NotifySendOne(proc.Entry.ID)
					} else {
						//	进入设置的父流程步骤
						if proc.Entry.EnterProcess.ChildBackProcess > 0 {
							w.GoToProcess(*proc.Entry.ParentEntry, proc.Entry.EnterProcess.ChildBackProcess)
							proc.Entry.ParentEntry.ProcessID = cast.ToUint(proc.Entry.EnterProcess.ChildBackProcess)
							//	通知设置的父流程步骤中的审批人
							//ins := NewBaseWorkflow()
							//ins.NotifySendOne(cast.ToUint(proc.AuditorID))
						} else {
							//默认进入父流程步骤下一步
							parentFlowlink := models.Flowlink{}
							tx.Model(&models.Flowlink{}).Where("process_id=?", proc.Entry.EnterProcessID).
								Where("type=?", "Condition").Find(&parentFlowlink)
							if parentFlowlink.NextProcessID == -1 {
								parentEntry := models.Entry{}
								tx.Model(&models.Entry{}).Where("id=?", proc.Entry.Pid).Find(&parentEntry)
								map_entry := make(map[string]interface{})
								map_entry["process_id"] = cast.ToUint(proc.Entry.EnterProcess.ChildBackProcess)
								map_entry["status"] = 9
								map_entry["child"] = 0
								tx.Model(&models.Entry{}).Where("id=?", parentEntry.ID).Save(&map_entry)

								var notifyProc models.Proc
								tx.Model(&models.Proc{}).Where("id=?", proc.ID).Preload("Emp").Find(&notifyProc)
								//通知发起人，审批结束
								w.NotifySendOne(proc.Entry.EmpID)
							} else {
								w.GoToProcess(*proc.Entry.ParentEntry, parentFlowlink.NextProcessID)
								proc.Entry.ParentEntry.ProcessID = cast.ToUint(parentFlowlink.NextProcessID)
								parentEntry := models.Entry{}
								tx.Model(&models.Entry{}).Where("id=?", proc.Entry.Pid).Find(&parentEntry)
								map_entry := make(map[string]interface{})
								map_entry["process_id"] = parentFlowlink.NextProcessID
								map_entry["status"] = 0
								tx.Model(&models.Entry{}).Where("id=?", parentEntry.ID).Save(&map_entry)
								//通知到下一个审批人
								w.NotifySendOne(cast.ToUint(proc.AuditorID))
							}
						}
						pentry := models.Entry{}
						tx.Model(&models.Entry{}).Where("id=?", proc.Entry.ParentEntry.ID).Find(&pentry)
						map_entry := make(map[string]interface{})
						map_entry["child"] = 0
						tx.Model(&models.Entry{}).Where("id=?", pentry.ID).Save(&map_entry)

					}
				} else {
					var notifyProc models.Proc
					tx.Model(&models.Proc{}).Where("id=?", proc.ID).Find(&notifyProc)
				}
			} else {
				auditor_ids := w.GetProcessAuditorIds(*proc.Entry, fklink.NextProcessID)
				auditors := []models.Emp{}
				tx.Model(&models.Emp{}).Where("id in (?)", auditor_ids).Preload("Dept").Find(&auditors)
				if len(auditors) < 1 {
					return errors.New("未找到下一步步骤审批人")
				}
				for _, auditor := range auditors {
					tx.Model(&models.Proc{}).Create(&models.Proc{
						EntryID:     proc.Entry.ID,
						FlowID:      cast.ToInt(proc.FlowID),
						ProcessID:   cast.ToInt(fklink.NextProcessID),
						ProcessName: fklink.NextProcess.ProcessName,
						EmpID:       cast.ToInt(auditor.ID),
						EmpName:     auditor.Name,
						DeptName:    auditor.Dept.DeptName,
						Circle:      proc.Entry.Circle,
						Concurrence: carbon.NewDateTime(carbon.Now()),
						Status:      0,
						IsRead:      0,
					})
					//通知下一个审批人
					w.NotifyNextAuditor(auditor.ID)
				}
				tx.Model(&models.Entry{}).Where("id=?", proc.Entry.ID).Update("process_id", cast.ToUint(fklink.NextProcessID))
				//	判断是否存在父进程
				var parentEntry models.Entry
				tx.Model(&models.Entry{}).Where("id=?", proc.Entry.Pid).Find(&parentEntry)
				if parentEntry.Pid > 0 {
					parentEntry.Child = cast.ToInt(fklink.NextProcessID)
					tx.Model(&models.Entry{}).Where("id=?", parentEntry.ID).Save(&parentEntry)
				}
			}
		}
	}

	var plugin_configs []official_plugin.PluginConfig
	tx.Model(official_plugin.PluginConfig{}).Where("process_id=?", process_id).Find(&plugin_configs)

	plugin_configs_str, _ := json.Marshal(plugin_configs)

	tx.Model(&models.Proc{}).
		Where("entry_id=?", proc.EntryID).
		Where("process_id=?", proc.ProcessID).
		Where("circle=?", proc.Entry.Circle).
		Where("status=?", 0).Save(models.Proc{
		Status:      1,
		AuditorID:   cast.ToInt(emp.ID),
		AuditorName: emp.Name,
		DeptName:    emp.Dept.DeptName,
		Content:     content,
		Beizhu:      string(plugin_configs_str),
		Concurrence: carbon.NewDateTime(carbon.Now()),
	})
	FlowID := cast.ToUint(proc.FlowID)
	ProcessID := cast.ToUint(proc.ProcessID)

	w.ExecPluginMethod("DistributePlugin", FlowID, ProcessID)

	return nil
}

func (w *Engine) Pass(process_id int, user models.Emp, content string) error {
	return w.Transfer(process_id, user, content)
}

func (w *Engine) UnPass(proc_id int, user models.Emp, content string) {
	var proc models.Proc
	query := w.db
	var emp models.Emp
	query.Model(&models.Emp{}).Where("user_id=?", user.ID).Find(&emp)
	query.Model(&models.Proc{}).Where("id=?", proc_id).Preload("Entry").Find(&proc)
	todoProc := models.Proc{}
	query.Model(&models.Proc{}).
		Where("entry_id=?", proc.EntryID).
		Where("process_id=?", proc.ProcessID).
		Where("circle=?", proc.Entry.Circle).
		Where("status=?", 0).Find(&todoProc)
	todoProc.Status = 1
	todoProc.AuditorID = cast.ToInt(emp.ID)
	todoProc.AuditorName = user.Name
	todoProc.AuditorDept = user.Dept.DeptName
	todoProc.Concurrence = carbon.NewDateTime(carbon.Now())
	todoProc.Content = content
	todoProc.IsRead = 1
	todoProc.Status = -1
	query.Model(&models.Proc{}).Where("id=?", todoProc.ID).Save(&todoProc)
	query.Model(&models.Entry{}).Where("id=?", proc.EntryID).Update("status", -1)
	if proc.Entry.Pid > 0 {
		var parentEntry models.Entry
		query.Model(&models.Entry{}).Where("id=?", proc.Entry.Pid).Find(&parentEntry)
		parentEntry.Child = proc.ProcessID
		parentEntry.Status = -1
		query.Model(&models.Entry{}).Where("id=?", parentEntry.ID).Save(&parentEntry)
	}
	w.NotifySendOne(proc.Entry.EmpID)

}

// 执行插件方法
func (w *Engine) ExecPluginMethod(plugin_name string, flowID uint, processID uint) error {
	ctor := GetCollectorIns()
	return ctor.DoPluginsExec(plugin_name, flowID, processID)
}

func NewEngin(db *gorm.DB) EngineImpl {
	return &Engine{
		hooks: make(map[string][]reflect.Value),
		mutex: sync.Mutex{},
		db:    db,
	}
}
