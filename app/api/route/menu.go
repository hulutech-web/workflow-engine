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

type menu struct {
	fx.In
	MenuSrv service.AuthMenuService
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

func (t menu) route(c *gin.Context) {
	res, err := t.MenuSrv.SelectMenuByRoleId(req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t menu) list(c *gin.Context) {
	res, err := t.MenuSrv.List(req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t menu) detail(c *gin.Context) {
	var detailReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := t.MenuSrv.Detail(detailReq.ID, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t menu) add(c *gin.Context) {
	var addReq req.MenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := t.MenuSrv.Add(&addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t menu) edit(c *gin.Context) {
	var editReq req.MenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := t.MenuSrv.Edit(&editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t menu) delete(c *gin.Context) {
	var delReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := t.MenuSrv.Del(delReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
