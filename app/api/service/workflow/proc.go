package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service/workflow"
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
)

/*
用来操作数据库查表
*/
type ProcService interface {
	Index(ctx *gin.Context, entry_id int) ([]models.Proc, error)
	Pass(ctx *gin.Context, r req.ProcPass) error
	UnPass(ctx *gin.Context, r req.ProcUnPass) error
}

type procService struct {
	db *gorm.DB
	wf *workflow.Engine
}

func (d procService) Index(ctx *gin.Context, entry_id int) ([]models.Proc, error) {
	var procs []models.Proc
	query := d.db
	query.Model(&models.Proc{}).Where("entry_id=?", entry_id).Preload("Entry.Emp").Find(&procs)
	return procs, nil
}

func (d procService) Pass(ctx *gin.Context, r req.ProcPass) error {
	var user models.Emp
	res := req.GetAuth(ctx)
	user_id := res.UserId
	d.db.Model(&models.Emp{}).Where("user_id=?", user_id).Find(&user)
	process_id := r.ProcessID
	content := r.Content
	err := d.wf.Pass(process_id, user, content)
	if err != nil {
		return errors.New(fmt.Sprintf("审批失败:%s", err.Error()))
	}
	return nil
}

func (d procService) UnPass(ctx *gin.Context, r req.ProcUnPass) error {
	var user models.Emp
	res := req.GetAuth(ctx)
	user_id := res.UserId
	d.db.Model(&models.Emp{}).Where("user_id=?", user_id).Find(&user)
	withUser := models.Emp{}
	d.db.Model(&models.Emp{}).Where("id=?", user.ID).Preload("Dept").Find(&withUser)
	proc_id := r.ProcID
	content := r.Content

	d.wf.UnPass(proc_id, withUser, content)
	return nil
}

func NewProcService(db *gorm.DB, wf *workflow.Engine) ProcService {
	return &procService{
		db: db,
		wf: wf,
	}
}
