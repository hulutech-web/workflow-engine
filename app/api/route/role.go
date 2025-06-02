package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"go.uber.org/fx"
)

type role struct {
	fx.In
	Srv service.AuthRoleService
}

func roleRoutes(t role, r *types.ApiRouter) {
	api := r.Group("/auth/role")
	api.GET("/all", t.all)
	api.GET("/list", t.list)
	api.GET("/detail", t.detail)
	api.POST("/add", t.add, r.Log("添加角色"))
	api.POST("/edit", t.edit, r.Log("编辑角色"))
	api.POST("/delete", t.delete, r.Log("删除角色"))
	api.POST("/change", t.change, r.Log("角色状态修改"))
}

// @BasePath /api
// @Summary 获取所有角色
// @Description 获取所有角色
// @Tags Role 角色管理
// @Id RoleAll
// @Produce  json
// @Param token header string true "access_token"
// @Success 200 {object} response.PageResp{data=[]resp.RoleSimpleResp} "成功"
// @Router /auth/role/all [get]
func (t role) all(ctx *gin.Context) {
	res, err := t.Srv.All(req.GetAuth(ctx))
	response.CheckAndRespWithData(ctx, res, err)
}

// @BasePath /api
// @Summary 获取角色列表
// @Description 获取角色列表
// @Tags Role 角色管理
// @Id RoleList
// @Produce  json
// @Param token header string true "access_token"
// @Param request query req.PageReq true "分页请求参数"
// @Success 200 {object} response.PageResp{data=[]resp.RoleResp} "成功"
// @Router /auth/role/list [get]
func (t role) list(c *gin.Context) {
	var page req.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := t.Srv.List(page, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 获取角色详情
// @Description 获取角色详情
// @Tags Role 角色管理
// @Id RoleDetail
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "角色ID"
// @Success 200 {object} resp.RoleResp "成功"
// @Router /auth/role/detail/{id} [get]
func (t role) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.Srv.Detail(detailReq.ID, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 添加角色
// @Description 添加角色
// @Tags Role 角色管理
// @Id RoleAdd
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.RoleAddReq true "角色信息"
// @Success 200 {object} response.Response "成功"
// @Router /auth/role/add [post]
func (t role) add(c *gin.Context) {
	var addReq req.RoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := t.Srv.Add(&addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 编辑角色
// @Description 编辑角色
// @Tags Role 角色管理
// @Id RoleEdit
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.RoleEditReq true "角色信息"
// @Success 200 {object} response.Response "成功"
// @Router /auth/role/edit [post]
func (t role) edit(c *gin.Context) {
	var editReq req.RoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := t.Srv.Edit(&editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 删除角色
// @Description 删除角色
// @Tags Role 角色管理
// @Id RoleDelete
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "角色ID"
// @Success 200 {object} response.Response "成功"
// @Router /auth/role/delete [post]
func (t role) delete(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := t.Srv.Del(delReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 角色状态修改
// @Description 角色状态修改
// @Tags Role 角色管理
// @Id RoleChange
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "角色ID"
// @Success 200 {object} response.Response "成功"
// @Router /auth/role/change [post]
func (t role) change(c *gin.Context) {
	var changeReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := t.Srv.Change(changeReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
