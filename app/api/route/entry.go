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

type entry struct {
	fx.In
	Srv workflow.EntryService
}

func entryRoutes(a entry, r *types.ApiRouter) {
	r.GET("/flow/:id/entry", a.Create)
	r.POST("/entry", a.Store)
	r.GET("/entry/:id", a.Show)
	r.GET("/entry/:id/entrydata", a.EntryData)
	r.POST("/entry/:id/resend", a.Resend)
}

func (r *entry) Create(ctx *gin.Context) {
	index, err := r.Srv.Create(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	logrus.WithFields(logrus.Fields{
		"index": index,
	}).Info("返回成功")
	response.OkWithData(ctx, index)
}

func (r *entry) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(ctx, idInt)
	response.OkWithData(ctx, show)
}

func (r *entry) Store(ctx *gin.Context) {

	var flowID req.FlowIDReq
	ctx.Bind(&flowID)
	err := r.Srv.Store(ctx, flowID)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *entry) EntryData(ctx *gin.Context) {

	idInt := cast.ToInt(ctx.Param("id"))
	err, mp := r.Srv.EntryData(ctx, idInt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, mp)
}

func (r *entry) Resend(ctx *gin.Context) {
	var entryReq req.EntryIDReq
	ctx.Bind(&entryReq)
	err := r.Srv.Resend(ctx, entryReq)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}
