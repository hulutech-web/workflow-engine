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
	r.POST("/login", a.login)
	r.GET("/refresh", a.refresh)
	r.GET("/logout", a.logout)
	r.GET("/info", a.info)
}

func (a account) login(c *gin.Context) {
	var loginReq req.AccountLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &loginReq)) {
		return
	}
	res, err := a.Srv.Login(&loginReq)
	response.CheckAndRespWithData(c, res, err)
}

func (a account) refresh(c *gin.Context) {
	var refreshReq req.AccountTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &refreshReq)) {
		return
	}
	res, err := a.Srv.RefreshToken(refreshReq.Token)
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

func (a account) info(c *gin.Context) {
	var infoReq req.AccountTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &infoReq)) {
		return
	}
	res, err := a.Srv.Info(infoReq.Token)
	response.CheckAndRespWithData(c, res, err)
}
