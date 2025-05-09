package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	"go.uber.org/fx"
)

type dept struct {
	fx.In
	Srv service.DeptService
}

func deptRoutes(a dept, r *types.ApiRouter) {
	r.POST("/dept", a.Store)
	r.PUT("/dept", a.Update)
	r.GET("/dept", a.Index)
	r.DELETE("/dept/:id", a.Destroy)
	r.GET("/dept/:id", a.Show)
	r.GET("/list", a.List)
	r.POST("/bindmanager", a.BindManager)
	r.POST("/binddirector", a.BindDirector)
}

func (r *dept) Index(ctx *gin.Context) {
	index, err := r.Srv.Index()
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	logrus.WithFields(logrus.Fields{
		"len": len(index),
	}).Info("返回成功")
	response.OkWithData(ctx, index)
}

func (r *dept) List(ctx *gin.Context) {
	list, err := r.Srv.List()
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, list)
}

func (r *dept) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(idInt)
	response.OkWithData(ctx, show)
}

func (r *dept) Store(ctx *gin.Context) {
	var dpt models.Dept
	ctx.Bind(&dpt)
	store, err := r.Srv.Store(dpt)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, store)
}

func (r *dept) Update(ctx *gin.Context) {
	return
}

func (r *dept) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	destroy, err := r.Srv.Destroy(cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, destroy)
}
func (r *dept) BindManager(ctx *gin.Context) {
	type BindManagerReq struct {
		ManagerID int `json:"manager_id" form:"manager_id"`
		DeptID    int `json:"dept_id" form:"dept_id"`
	}
	var bindManagerReq BindManagerReq
	ctx.BindJSON(&bindManagerReq)
	r.Srv.BindManager(bindManagerReq.ManagerID, bindManagerReq.DeptID)
}

func (r *dept) BindDirector(ctx *gin.Context) {
	type BindDirectorReq struct {
		DirectorID int `json:"director_id" form:"director_id"`
		DeptID     int `json:"dept_id" form:"dept_id"`
	}
	var bindDirectorReq BindDirectorReq
	ctx.BindJSON(&bindDirectorReq)
	manager, err := r.Srv.BindManager(bindDirectorReq.DirectorID, bindDirectorReq.DeptID)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, manager)
}
