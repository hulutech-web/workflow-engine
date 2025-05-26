package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/service/workflow"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.uber.org/fx"
)

type templateForm struct {
	fx.In
	Srv workflow.TemplateFormService
}

func templateFormRoutes(a templateForm, r *types.ApiRouter) {
	r.POST("/templateform", a.Store)
	r.PUT("/templateform", a.Update)
	r.GET("/templateform", a.Index)
	r.DELETE("/templateform/:id", a.Destroy)
	r.GET("/templateform/:id", a.Show)
	r.GET("/templateform/list", a.List)
}

func (r *templateForm) Index(ctx *gin.Context) {
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

func (r *templateForm) List(ctx *gin.Context) {
	list, err := r.Srv.List(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, list)
}

func (r *templateForm) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(ctx, idInt)
	response.OkWithData(ctx, show)
}

func (r *templateForm) Store(ctx *gin.Context) {
	var dpt models.TemplateForm
	if err2 := ctx.ShouldBind(&dpt); err2 != nil {
		response.FailWithMsg(ctx, response.Failed, err2.Error())
		return
	}
	logrus.WithFields(logrus.Fields{
		"model": dpt,
	}).Info("返回@@成功")
	err := r.Srv.Store(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *templateForm) Update(ctx *gin.Context) {
	var dpt models.TemplateForm
	ctx.Bind(&dpt)
	err := r.Srv.Update(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *templateForm) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.Srv.Destroy(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}
