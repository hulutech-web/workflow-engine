package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hulutech-web/workflow-engine/core/http/validator"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"io/ioutil"
	"mime/multipart"
)

var VerifyUtil = verifyUtil{
	validator: validator.NewService(),
}

// verifyUtil 参数验证工具类
type verifyUtil struct {
	validator *validator.Service
}

func (vu verifyUtil) VerifyJSON(c *gin.Context, obj any) (e error) {
	_ = c.ShouldBindBodyWith(obj, binding.JSON)
	msgs := vu.validator.Validate(obj)
	if len(msgs) > 0 {
		e = response.ParamsValidError.MakeData(msgs[0])
	}
	return
}

func (vu verifyUtil) VerifyJSONArray(c *gin.Context, obj any) (e error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	msgs := vu.validator.Validate(obj)
	if len(msgs) > 0 {
		e = response.ParamsValidError.MakeData(msgs[0])
	}
	return
}

func (vu verifyUtil) VerifyBody(c *gin.Context, obj any) (e error) {
	_ = c.ShouldBind(obj)
	msgs := vu.validator.Validate(obj)
	if len(msgs) > 0 {
		e = response.ParamsValidError.MakeData(msgs[0])
	}
	return
}

func (vu verifyUtil) VerifyHeader(c *gin.Context, obj any) (e error) {
	_ = c.ShouldBindHeader(obj)
	msgs := vu.validator.Validate(obj)
	if len(msgs) > 0 {
		e = response.ParamsValidError.MakeData(msgs[0])
	}
	return
}

func (vu verifyUtil) VerifyQuery(c *gin.Context, obj any) (e error) {
	_ = c.ShouldBindQuery(obj)
	msgs := vu.validator.Validate(obj)
	if len(msgs) > 0 {
		e = response.ParamsValidError.MakeData(msgs[0])
	}
	return
}

func (vu verifyUtil) VerifyFile(c *gin.Context, name string) (file *multipart.FileHeader, e error) {
	file, err := c.FormFile(name)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

// VerifyData 验证请求数据
func (vu verifyUtil) Verify(c *gin.Context, obj ...any) (e error) {
	for _, o := range obj {
		switch c.Request.Method {
		case "POST":
			return VerifyUtil.VerifyJSON(c, o)
		case "GET":
			return VerifyUtil.VerifyQuery(c, o)
		case "PUT":
			return VerifyUtil.VerifyJSON(c, o)
		case "DELETE":
			return VerifyUtil.VerifyQuery(c, o)
		default:
			return response.ParamsValidError.MakeData("请求方式错误")
		}
	}
	return
}
