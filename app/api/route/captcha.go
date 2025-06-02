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

// Get @BasePath /api
// @Summary 生成验证码
// @Description 生成验证码
// @Tags Captcha 验证码
// @Id CaptchaGet
// @Produce json
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /captcha/get [get]
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

// Validate @BasePath /api
// @Summary 验证码验证
// @Description 验证码验证
// @Tags Captcha 验证码
// @Id CaptchaValidate
// @Produce json
// @Param angle query string true "验证码角度"
// @Param captcha_key query string true "验证码key"
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Router /captcha/validate [post]
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
