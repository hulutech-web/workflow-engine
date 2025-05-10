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

type user struct {
	fx.In
	Srv service.UserService
}

func userRoutes(t user, r *types.ApiRouter) {
	api := r.Group("/user")

	api.GET("/list", t.list)
	api.GET("/detail", t.detail)
	api.POST("/add", t.add)
	api.POST("/edit", t.edit)
	api.POST("/update", t.update)
	api.POST("/delete", t.delete)
	api.POST("/disable", t.disable)
}

func (t user) list(ctx *gin.Context) {
	var pageReq req.PageReq
	var listReq req.UserQueryReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &listReq, &pageReq)) {
		return
	}
	res, err := t.Srv.List(&pageReq, &listReq, req.GetAuth(ctx))
	response.CheckAndRespWithData(ctx, res, err)
}

func (t user) detail(ctx *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &idReq)) {
		return
	}
	res, err := t.Srv.Detail(idReq.ID)
	response.CheckAndRespWithData(ctx, res, err)
}

func (t user) add(ctx *gin.Context) {
	var userReq req.UserAddReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &userReq)) {
		return
	}
	err := t.Srv.Add(&userReq, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}

func (t user) edit(ctx *gin.Context) {
	var editReq req.UserEditReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &editReq)) {
		return
	}
	err := t.Srv.Edit(&editReq, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}

func (t user) update(ctx *gin.Context) {
	var updateReq req.UserUpdateReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &updateReq)) {
		return
	}
	err := t.Srv.Update(&updateReq)
	response.CheckAndResp(ctx, err)
}

func (t user) delete(ctx *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &idReq)) {
		return
	}
	err := t.Srv.Delete(idReq.ID, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}

func (t user) disable(ctx *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &idReq)) {
		return
	}
	err := t.Srv.Disable(idReq.ID, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}
