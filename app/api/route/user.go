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

// @BasePath /api
// @Summary 用户权限
// @Description 用户权限相关接口
// @Tags User 用户
// @Id UserSelf
// @Produce  json
// @Param token header string true "access_token"
// @Success 200 {object} response.Response{data=resp.UserSelfResp} "成功"
// @Router /user/self [get]
func (t user) self(ctx *gin.Context) {
	res, err := t.Srv.Self(req.GetAuth(ctx))
	response.CheckAndRespWithData(ctx, res, err)
}

// @BasePath /api
// @Summary 用户列表
// @Description 用户列表相关接口
// @Tags User 用户
// @Id UserList
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.UserQueryReq true "查询条件"
// @Success 200 {object} response.Response{data=response.PageResp} "成功"
// @Router /user/list [get]
func (t user) list(ctx *gin.Context) {
	var pageReq req.PageReq
	var listReq req.UserQueryReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &listReq, &pageReq)) {
		return
	}
	res, err := t.Srv.List(&pageReq, &listReq, req.GetAuth(ctx))
	response.CheckAndRespWithData(ctx, res, err)
}

// @BasePath /api
// @Summary 用户分页
// @Description 用户分页
// @Tags User 用户
// @Id UserIndex
// @Produce  json
// @Param token header string true "access_token"
// @Success 200 {object} response.Response{data=resp.UserResp} "成功"
// @Router /user [get]
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

// @BasePath /api
// @Summary 用户详情
// @Description 用户详情相关接口
// @Tags User 用户
// @Id UserDetail
// @Produce  json
// @Param token header string true "access_token"
// @Param id path uint true "用户ID"
// @Success 200 {object} response.Response{data=resp.UserResp} "成功"
// @Router /user/detail/{id} [get]
func (t user) detail(ctx *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &idReq)) {
		return
	}
	res, err := t.Srv.Detail(idReq.ID)
	response.CheckAndRespWithData(ctx, res, err)
}

// @BasePath /api
// @Summary 添加用户
// @Description 添加用户相关接口
// @Tags User 用户
// @Id UserAdd
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.UserAddReq true "用户信息"
// @Success 200 {object} response.Response "成功"
// @Router /user/add [post]
func (t user) add(ctx *gin.Context) {
	var userReq req.UserAddReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &userReq)) {
		return
	}
	err := t.Srv.Add(&userReq, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}

// @BasePath /api
// @Summary 编辑用户
// @Description 编辑用户相关接口
// @Tags User 用户
// @Id UserEdit
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.UserEditReq true "用户信息"
// @Success 200 {object} response.Response "成功"
// @Router /user/edit [post]
func (t user) edit(ctx *gin.Context) {
	var editReq req.UserEditReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &editReq)) {
		return
	}
	err := t.Srv.Edit(&editReq, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}

// @BasePath /api
// @Summary 更新用户
// @Description 更新用户相关接口
// @Tags User 用户
// @Id UserUpdate
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.UserUpdateReq true "用户信息"
// @Success 200 {object} response.Response "成功"
// @Router /user/update [post]
func (t user) update(ctx *gin.Context) {
	var updateReq req.UserUpdateReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &updateReq)) {
		return
	}
	err := t.Srv.Update(&updateReq)
	response.CheckAndResp(ctx, err)
}

// @BasePath /api
// @Summary 删除用户
// @Description 删除用户相关接口
// @Tags User 用户
// @Id UserDelete
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "用户ID"
// @Success 200 {object} response.Response "成功"
// @Router /user/delete/{id} [post]
func (t user) delete(ctx *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &idReq)) {
		return
	}
	err := t.Srv.Delete(idReq.ID, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}

// @BasePath /api
// @Summary 禁用用户
// @Description 禁用用户相关接口
// @Tags User 用户
// @Id UserDisable
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdReq true "用户ID"
// @Success 200 {object} response.Response "成功"
// @Router /user/disable [post]
func (t user) disable(ctx *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(ctx, util.VerifyUtil.Verify(ctx, &idReq)) {
		return
	}
	err := t.Srv.Disable(idReq.ID, req.GetAuth(ctx))
	response.CheckAndResp(ctx, err)
}
