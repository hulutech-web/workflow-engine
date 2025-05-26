package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service/workflow"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.uber.org/fx"
)

type proc struct {
	fx.In
	Srv workflow.ProcService
}

func procRoutes(a proc, r *types.ApiRouter) {
	r.GET("/proc/:entry_id", a.Index)
	r.POST("/pass", a.Index)
	r.POST("/unpass", a.Index)
}

func (r *proc) Index(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("entry_id"))
	index, err := r.Srv.Index(ctx, id)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	logrus.WithFields(logrus.Fields{
		"index": index,
	}).Info("返回成功")
	response.OkWithData(ctx, index)
}

func (r *proc) Pass(ctx *gin.Context) {
	var reqPass req.ProcPass
	if err := ctx.ShouldBindJSON(&reqPass); err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	err := r.Srv.Pass(ctx, reqPass)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithMsg(ctx, "操作成功")
}

func (r *proc) UnPass(ctx *gin.Context) {
	var reqPass req.ProcUnPass
	if err := ctx.ShouldBindJSON(&reqPass); err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	err := r.Srv.UnPass(ctx, reqPass)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithMsg(ctx, "操作成功")
}
