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

func (a account) login(c *gin.Context) {
	var loginReq req.AccountLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &loginReq)) {
		return
	}
	res, err := a.Srv.Login(&loginReq)
	response.CheckAndRespWithData(c, res, err)
}

func (a account) logout(c *gin.Context) {
	var logoutReq req.AccountTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &logoutReq)) {
		return
	}
	err := a.Srv.Logout(logoutReq.Token)
	response.CheckAndResp(c, err)
}

func (a account) tenant(c *gin.Context) {
	res, err := a.Srv.TenantList()
	response.CheckAndRespWithData(c, res, err)
}

func (a account) register(c *gin.Context) {
	var registerReq req.AccountRegisterReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &registerReq)) {
		return
	}
	res, err := a.Srv.Register(&registerReq)
	response.CheckAndRespWithData(c, res, err)
}
