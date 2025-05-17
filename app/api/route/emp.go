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
	"net/http"
)

type emp struct {
	fx.In
	Srv service.EmpService
}

func empRoutes(a emp, r *types.ApiRouter) {
	r.POST("/emp", a.Store)
	r.PUT("/emp", a.Update)
	r.GET("/emp", a.Index)
	r.DELETE("/emp/:id", a.Destroy)
	r.GET("/emp/:id", a.Show)
	r.GET("/emp/list", a.List)
	r.POST("/emp/bind_user", a.BindUser)
}

func (r *emp) Index(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	index, err := r.Srv.Index(ctx, query)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	logrus.WithFields(logrus.Fields{
		"index": index,
	}).Info("返回成功")
	response.OkWithData(ctx, index)
}

func (r *emp) List(ctx *gin.Context) {
	list, err := r.Srv.List(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, list)
}

func (r *emp) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(ctx, idInt)
	response.OkWithData(ctx, show)
}

func (r *emp) Store(ctx *gin.Context) {
	var dpt models.Emp
	ctx.Bind(&dpt)
	err := r.Srv.Store(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *emp) Update(ctx *gin.Context) {
	var dpt models.Emp
	ctx.Bind(&dpt)
	err := r.Srv.Update(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *emp) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.Srv.Destroy(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *emp) BindUser(ctx *gin.Context) {
	type BindUserReq struct {
		ID     int `json:"id" form:"id"`
		UserID int `json:"user_id" form:"user_id"`
	}
	var bindUserReq BindUserReq
	if err2 := ctx.ShouldBindJSON(&bindUserReq); err2 != nil {
		response.FailWithMsg(ctx, response.Failed, err2.Error())
		return
	}

	err := r.Srv.BindUser(ctx, bindUserReq.ID, bindUserReq.UserID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	response.OkWithData(ctx, "操作成功")
}
