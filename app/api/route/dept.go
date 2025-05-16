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
	r.GET("/dept/list", a.List)
	r.POST("/dept/bind_manager", a.BindManager)
	r.POST("/dept/bind_director", a.BindDirector)
	r.GET("/dept/:id/tree", a.DisplayTree)
}

func (r *dept) Index(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	index, err := r.Srv.Index(ctx, query)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	logrus.WithFields(logrus.Fields{
		"index": index,
	}).Info("返回成功")
	ctx.JSON(http.StatusOK, index)
}

func (r *dept) List(ctx *gin.Context) {
	list, err := r.Srv.List(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, list)
}

func (r *dept) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(ctx, idInt)
	response.OkWithData(ctx, show)
}

func (r *dept) Store(ctx *gin.Context) {
	var dpt models.Dept
	ctx.Bind(&dpt)
	err := r.Srv.Store(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *dept) Update(ctx *gin.Context) {
	var dpt models.Dept
	ctx.Bind(&dpt)
	err := r.Srv.Update(ctx, dpt)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *dept) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.Srv.Destroy(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}
func (r *dept) BindManager(ctx *gin.Context) {
	type BindManagerReq struct {
		ManagerID int `json:"manager_id" form:"manager_id"`
		DeptID    int `json:"dept_id" form:"dept_id"`
	}
	var bindManagerReq BindManagerReq
	if err2 := ctx.ShouldBindJSON(&bindManagerReq); err2 != nil {
		response.FailWithMsg(ctx, response.Failed, err2.Error())
		return
	}

	err := r.Srv.BindManager(ctx, bindManagerReq.ManagerID, bindManagerReq.DeptID)
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *dept) BindDirector(ctx *gin.Context) {
	type BindDirectorReq struct {
		DirectorID int `json:"director_id" form:"director_id"`
		DeptID     int `json:"dept_id" form:"dept_id"`
	}
	var bindDirectorReq BindDirectorReq
	if err2 := ctx.ShouldBindJSON(&bindDirectorReq); err2 != nil {
		response.FailWithMsg(ctx, response.Failed, err2.Error())
		return
	}

	err := r.Srv.BindDirector(ctx, bindDirectorReq.DirectorID, bindDirectorReq.DeptID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	response.OkWithData(ctx, "操作成功")
}

func (r *dept) DisplayTree(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := r.Srv.DisplayTree(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, res)
}
