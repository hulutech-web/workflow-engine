package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/api/workflow/common"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.uber.org/fx"
)

type process struct {
	fx.In
	Srv service.ProcessService
}

func processRoutes(a process, r *types.ApiRouter) {
	r.POST("/process", a.Store)
	r.PUT("/process/:id", a.Update)
	r.GET("/process", a.Index)
	r.DELETE("/process/:id", a.Destroy)
	r.GET("/process/:id", a.Show)
	r.POST("/process/list", a.List)
	r.GET("/process/attribute", a.Attribute)
	r.POST("/process/con", a.Condition)
}

func (r *process) Index(ctx *gin.Context) {
	index, err := r.Srv.Index(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	logrus.WithFields(logrus.Fields{
		"index": index,
	}).Info("返回成功")
	response.OkWithData(ctx, index)
}

func (r *process) List(ctx *gin.Context) {
	req := req.ProcessReq{}
	list, err := r.Srv.List(ctx, req)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, list)
}

func (r *process) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(ctx, idInt)
	response.OkWithData(ctx, show)
}

func (r *process) Store(ctx *gin.Context) {

	var prc req.ProReq
	if err2 := ctx.ShouldBindJSON(&prc); err2 != nil {
		response.FailWithMsg(ctx, response.Failed, err2.Error())
		return
	}

	err := r.Srv.Store(ctx, prc)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *process) Update(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Info("id值")
	var processRequest common.ProcessRequest
	ctx.Bind(&processRequest)

	err := r.Srv.Update(ctx, id, processRequest)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *process) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.Srv.Destroy(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}
func (r *process) Attribute(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	err, data := r.Srv.Attribute(ctx, cast.ToInt(id))
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, data)
}

func (r *process) Condition(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.Srv.Destroy(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}
