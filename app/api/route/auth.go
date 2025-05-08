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

type auth struct {
	fx.In
	Srv service.AuthService
}

func authRoutes(a auth, r *types.ApiRouter) {
	r.POST("/login", a.login)
	r.GET("/refresh", a.refresh)
	r.GET("/logout", a.logout)
	r.GET("/info", a.info)
}

func (a auth) login(c *gin.Context) {
	var loginReq req.AuthLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &loginReq)) {
		return
	}
	res, err := a.Srv.Login(&loginReq)
	response.CheckAndRespWithData(c, res, err)
}

func (a auth) refresh(c *gin.Context) {
	var refreshReq req.AuthTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &refreshReq)) {
		return
	}
	res, err := a.Srv.RefreshToken(refreshReq.Token)
	response.CheckAndRespWithData(c, res, err)
}

func (a auth) logout(c *gin.Context) {
	var logoutReq req.AuthTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &logoutReq)) {
		return
	}
	err := a.Srv.Logout(logoutReq.Token)
	response.CheckAndResp(c, err)
}

func (a auth) info(c *gin.Context) {
	var infoReq req.AuthTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &infoReq)) {
		return
	}
	res, err := a.Srv.Info(infoReq.Token)
	response.CheckAndRespWithData(c, res, err)
}
