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

type tenant struct {
	fx.In
	Srv auth.AuthTenantService
}

func tenantRoutes(t tenant, r *types.ApiRouter) {
	api := r.Group("/auth/tenant")
	api.GET("/all", t.all)
	api.POST("/list", t.list)
	api.GET("/detail", t.detail)
	api.POST("/add", t.add, r.Log("添加租户"))
	api.POST("/edit", t.edit, r.Log("编辑租户"))
	api.POST("/delete", t.delete, r.Log("删除租户"))
}

// @BasePath /api
// @Summary 获取所有租户
// @Description 获取所有租户
// @Tags 租户管理
// @Produce  json
// @Param token header string true "access_token"
// @Success 200 {object} response.PageResp{data=[]resp.TenantResp} "成功"
// @Router /auth/tenant/all [get]
func (t tenant) all(ctx *gin.Context) {
	res, err := t.Srv.All()
	response.CheckAndRespWithData(ctx, res, err)
}

// @BasePath /api
// @Summary 获取租户列表
// @Description 获取租户列表
// @Tags 租户管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.PageReq true "分页请求参数"
// @Success 200 {object} response.PageResp{data=[]resp.TenantResp} "成功"
// @Router /auth/tenant/list [post]
func (t tenant) list(c *gin.Context) {
	var page req.PageReq
	var listReq req.TenantQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &page, &listReq)) {
		return
	}
	res, err := t.Srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 获取租户详情
// @Description 获取租户详情
// @Tags 租户管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "请求参数"
// @Success 200 {object} resp.TenantResp "成功"
// @Router /auth/tenant/detail [get]
func (t tenant) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.Srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 添加租户
// @Description 添加租户
// @Tags 租户管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.TenantAddReq true "请求参数"
// @Success 200 {object} response.Response "成功"
// @Router /auth/tenant/add [post]
func (t tenant) add(c *gin.Context) {
	var addReq req.TenantAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := t.Srv.Add(addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 编辑租户
// @Description 编辑租户
// @Tags 租户管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.TenantEditReq true "请求参数"
// @Success 200 {object} response.Response "成功"
// @Router /auth/tenant/edit [post]
func (t tenant) edit(c *gin.Context) {
	var editReq req.TenantEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := t.Srv.Edit(editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 删除租户
// @Description 删除租户
// @Tags 租户管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "请求参数"
// @Success 200 {object} response.Response "成功"
// @Router /auth/tenant/delete [post]
func (t tenant) delete(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.Srv.Del(delReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
