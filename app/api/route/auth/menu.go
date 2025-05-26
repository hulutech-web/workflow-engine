package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service/auth"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"go.uber.org/fx"
)

type menu struct {
	fx.In
	MenuSrv auth.AuthMenuService
}

func menuRoutes(t menu, r *types.ApiRouter) {
	menu := r.Group("/auth/menu")

	menu.GET("/route", t.route)
	menu.GET("/list", t.list)
	menu.GET("/detail", t.detail)
	menu.POST("/add", t.add, r.Log("添加菜单"))
	menu.POST("/edit", t.edit, r.Log("编辑菜单"))
	menu.POST("/delete", t.delete, r.Log("删除菜单"))
}

// @BasePath /api
// @Summary 获取当前用户的菜单权限
// @Description 获取当前用户的菜单权限
// @Tags 菜单权限
// @Produce  json
// @Param token header string true "access_token"
// @Success 200 {object} resp.MenuResp "成功"
// @Router /auth/menu/route [get]
func (t menu) route(c *gin.Context) {
	res, err := t.MenuSrv.SelectMenuByRoleId(req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 获取菜单列表
// @Description 获取菜单列表
// @Tags 菜单权限
// @Produce  json
// @Param token header string true "access_token"
// @Success 200 {object} []resp.MenuResp "成功"
// @Router /auth/menu/list [get]
func (t menu) list(c *gin.Context) {
	res, err := t.MenuSrv.List(req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 获取菜单详情
// @Description 获取菜单详情
// @Tags 菜单权限
// @Produce  json
// @Param token header string true "access_token"
// @Param id path uint true "菜单ID"
// @Success 200 {object} resp.MenuResp "成功"
// @Router /auth/menu/detail/{id} [get]
func (t menu) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.MenuSrv.Detail(detailReq.ID, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 添加菜单
// @Description 添加菜单
// @Tags 菜单权限
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.MenuAddReq true "菜单信息"
// @Success 200 {object} response.Response "成功"
// @Router /auth/menu/add [post]
func (t menu) add(c *gin.Context) {
	var addReq req.MenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := t.MenuSrv.Add(&addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 编辑菜单
// @Description 编辑菜单
// @Tags 菜单权限
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.MenuEditReq true "菜单信息"
// @Success 200 {object} response.Response "成功"
// @Router /auth/menu/edit [post]
func (t menu) edit(c *gin.Context) {
	var editReq req.MenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := t.MenuSrv.Edit(&editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 删除菜单
// @Description 删除菜单
// @Tags 菜单权限
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "菜单ID"
// @Success 200 {object} response.Response "成功"
// @Router /auth/menu/delete [post]
func (t menu) delete(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := t.MenuSrv.Del(delReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
