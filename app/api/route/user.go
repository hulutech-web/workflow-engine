package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type user struct {
	fx.In
	Srv service.UserService
}

func userRoutes(t user, r *types.ApiRouter) {
	api := r.Group("/user")

	api.GET("/self", t.self)
	api.GET("/list", t.list)
	api.GET("/index", t.index)
	api.GET("/detail", t.detail)
	api.POST("/add", t.add, r.Log("添加用户"))
	api.POST("/edit", t.edit, r.Log("编辑用户"))
	api.POST("/update", t.update, r.Log("更新用户"))
	api.POST("/delete", t.delete, r.Log("删除用户"))
	api.POST("/disable", t.disable, r.Log("禁用用户"))
}

func (t user) self(ctx *gin.Context) {
	res, err := t.Srv.Self(req.GetAuth(ctx))
	response.CheckAndRespWithData(ctx, res, err)
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
func (r *user) index(ctx *gin.Context) {
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
