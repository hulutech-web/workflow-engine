package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.uber.org/fx"
)

type flow struct {
	fx.In
	Srv service.FlowService
}

func flowRoutes(a flow, r *types.ApiRouter) {
	r.POST("/flow", a.Store)
	r.PUT("/flow/:id", a.Update)
	r.GET("/flow/create", a.Create)
	r.GET("/flow", a.Index)
	r.DELETE("/flow/:id", a.Destroy)
	r.GET("/flow/:id", a.Show)
	r.GET("/flow/list", a.List)
	r.GET("/flow/flowchart/:id", a.FlowDesign)
	r.POST("/flow/publish", a.Publish)
}

func (r *flow) Index(ctx *gin.Context) {
	index, err := r.Srv.Index(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	logrus.WithFields(logrus.Fields{
		"index": index,
	}).Info("返回成功")
	response.OkOnlyData(ctx, index)
}
func (r *flow) Create(ctx *gin.Context) {
	err, templates, flowtypes := r.Srv.Create(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	logrus.WithFields(logrus.Fields{
		"templates": templates,
		"flowtypes": flowtypes,
	}).Warn("返回成功")
	response.OkOnlyData(ctx, map[string]interface{}{
		"templates": templates,
		"flowtypes": flowtypes,
	})
}

func (r *flow) List(ctx *gin.Context) {
	list, err := r.Srv.List(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, list)
}

func (r *flow) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(ctx, idInt)
	response.OkOnlyData(ctx, show)
}

func (r *flow) Store(ctx *gin.Context) {
	var dpt models.Flow
	ctx.Bind(&dpt)
	err := r.Srv.Store(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *flow) Update(ctx *gin.Context) {
	var dpt models.Flow
	ctx.Bind(&dpt)
	err := r.Srv.Update(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.Ok(ctx)
}

func (r *flow) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.Srv.Destroy(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}
func (r *flow) FlowDesign(ctx *gin.Context) {
	id := ctx.Param("id")
	err, m := r.Srv.FlowDesign(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkOnlyData(ctx, m)
}
func (r *flow) Publish(ctx *gin.Context) {

	var req req.FlowReq
	ctx.Bind(&req)
	err := r.Srv.Publish(ctx, req.FlowID)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "发布成功")
}
