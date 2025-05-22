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

type account struct {
	fx.In
	Srv service.AccountService
}

func accountRoutes(a account, r *types.ApiRouter) {
	api := r.Group("/account")
	api.POST("/login", a.login)
	api.GET("/logout", a.logout)
	api.GET("/tenant", a.tenant)
	api.POST("/register", a.register)
}

// @BasePath /api
// @Summary 用户登录
// @Tags Account 账户管理
// @Accept json
// @Id Login
// @Produce json
// @Param request body req.AccountLoginReq true "登录参数"
// @Success 200 {object} response.Response{data=resp.AccountLoginResp}
// @Router /account/login [post]
func (a account) login(c *gin.Context) {
	var loginReq req.AccountLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &loginReq)) {
		return
	}
	res, err := a.Srv.Login(&loginReq)
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 用户登出
// @Tags Account
// @Tags 账户管理
// @Id Logout
// @Security ApiKeyAuth
// @Param token header string true "访问令牌"
// @Success 200 {object} response.Response
// @Router /account/logout [get]
func (a account) logout(c *gin.Context) {
	var logoutReq req.AccountTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &logoutReq)) {
		return
	}
	err := a.Srv.Logout(logoutReq.Token)
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 租户列表
// @Tags Account
// @Tags 账户管理
// @Id Tenant
// @Produce json
// @Success 200 {object} response.Response{data=[]resp.SelectOption}
// @Router /account/tenant [get]
func (a account) tenant(c *gin.Context) {
	res, err := a.Srv.TenantList()
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 用户注册
// @Tags Account
// @Tags 账户管理
// @Accept json
// @Id Register
// @Produce json
// @Param request body req.AccountRegisterReq true "注册参数"
// @Success 200 {object} response.Response{data=resp.AccountLoginResp}
// @Router /account/register [post]
func (a account) register(c *gin.Context) {
	var registerReq req.AccountRegisterReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &registerReq)) {
		return
	}
	res, err := a.Srv.Register(&registerReq)
	response.CheckAndRespWithData(c, res, err)
}
