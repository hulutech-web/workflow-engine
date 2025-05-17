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

func (t role) all(ctx *gin.Context) {
	res, err := t.Srv.All(req.GetAuth(ctx))
	response.CheckAndRespWithData(ctx, res, err)
}

func (t role) list(c *gin.Context) {
	var page req.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := t.Srv.List(page, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t role) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.Srv.Detail(detailReq.ID, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t role) add(c *gin.Context) {
	var addReq req.RoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := t.Srv.Add(&addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t role) edit(c *gin.Context) {
	var editReq req.RoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := t.Srv.Edit(&editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t role) delete(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := t.Srv.Del(delReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t role) change(c *gin.Context) {
	var changeReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := t.Srv.Change(changeReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
