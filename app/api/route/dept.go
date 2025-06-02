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

// Index @BasePath /api
// @Summary 部门
// @Description 部门分页
// @Tags Dept 部门
// @Id DeptIndex
// @Produce json
// @Success 200 {object} service.PageResult "成功"
// @Router /dept [get]
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
	response.OkWithData(ctx, index)
}

// List @BasePath /api
// @Summary 部门
// @Description 部门列表
// @Tags Dept 部门
// @Id DeptList
// @Produce json
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept/list [get]
func (r *dept) List(ctx *gin.Context) {
	list, err := r.Srv.List(ctx)
	if err != nil {
		response.Fail(ctx, response.Failed)
	}
	response.OkWithData(ctx, list)
}

// Show @BasePath /api
// @Summary 部门
// @Description 单个部门
// @Tags Dept 部门
// @Id DeptShow
// @Produce json
// @Param id path int true "部门ID"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept/{id} [get]
func (r *dept) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt := cast.ToInt(id)
	show := r.Srv.Show(ctx, idInt)
	response.OkWithData(ctx, show)
}

// Store @BasePath /api
// @Summary 部门
// @Description 新增部门
// @Tags Dept 部门
// @Id DeptStore
// @Produce json
// @Param request body models.Dept true "部门信息"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept [post]
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

// Update @BasePath /api
// @Summary 部门
// @Description 新增部门
// @Tags Dept 部门
// @Id DeptUpdate
// @Produce json
// @Param id path int true "部门ID"
// @Param request body models.Dept true "部门信息"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept/{id} [put]
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

// Destroy @BasePath /api
// @Summary 部门
// @Description 新增部门
// @Tags Dept 部门
// @Id DeptDestroy
// @Produce json
// @Param id path int true "部门ID"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept/{id} [delete]
func (r *dept) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	err := r.Srv.Destroy(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, "操作成功")
}

// BindManager @BasePath /api
// @Summary 部门
// @Description 新增部门
// @Tags Dept 部门
// @Id DeptBindManager
// @Produce json
// @Param request body req.BindManagerReq true "绑定参数"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept/bind_manager [post]
func (r *dept) BindManager(ctx *gin.Context) {
	var bindManagerReq req.BindManagerReq
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

// BindDirector @BasePath /api
// @Summary 部门
// @Description 新增部门
// @Tags Dept 部门
// @Id DeptBindDirector
// @Produce json
// @Param request body req.BindDirectorReq true "绑定参数"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept/bind_director [post]
func (r *dept) BindDirector(ctx *gin.Context) {
	var bindDirectorReq req.BindDirectorReq
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

// DisplayTree @BasePath /api
// @Summary 部门
// @Description 新增部门
// @Tags Dept 部门
// @Id DeptDisplayTree
// @Param id path int true "部门ID"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /dept/{id}/tree [get]
func (r *dept) DisplayTree(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := r.Srv.DisplayTree(ctx, cast.ToInt(id))
	if err != nil {
		response.Fail(ctx, response.Failed)
		return
	}
	response.OkWithData(ctx, res)
}
