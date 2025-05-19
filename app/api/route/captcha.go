package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"go.uber.org/fx"
)

type captcha struct {
	fx.In
	Srv service.CaptchaService
}

func captchaRoutes(a captcha, r *types.ApiRouter) {
	r.GET("/captcha/get", a.Get)
	r.POST("/captcha/validate", a.Validate)
}

func (r captcha) Get(c *gin.Context) {
	captcha_key, code, image_base64, thumb_base64, err := r.Srv.Generate()
	if err != nil {
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}
	response.OkWithData(c, map[string]interface{}{
		"captcha_key":  captcha_key,
		"code":         code,
		"image_base64": image_base64,
		"thumb_base64": thumb_base64,
	})
}

func (a captcha) Validate(c *gin.Context) {
	var rq req.CaptchaReq
	if err := c.ShouldBind(&rq); err != nil {
		response.ParamsValidError.MakeData(err.Error())
		return
	}
	code, isOk := a.Srv.CheckAngle(fmt.Sprintf("%d", rq.Angle), rq.CaptchaKey)
	if !isOk {
		response.FailWithMsg(c, response.CaptchaError, "验证码错误")
		return
	}
	response.OkWithData(c, map[string]interface{}{
		"code": code,
	})
}
