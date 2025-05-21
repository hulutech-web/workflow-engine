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

type upload struct {
	fx.In
	Srv service.UploadService
}

func uploadRoutes(t upload, r *types.ApiRouter) {
	api := r.RouterGroup.Group("/upload")
	api.POST("/image", t.image)
	api.POST("/file", t.file)
	api.POST("/audio", t.audio)
	api.POST("/video", t.video)
}

// @BasePath /api
// @Summary 上传图片
// @Description 上传图片
// @Tags 上传
// @Accept  multipart/form-data
// @Produce  json
// @Param token header string true "access_token"
// @Param file formData file true "上传文件"
// @Param cid formData int true "分类ID"
// @Success 200 {object} response.Response{data=resp.FileResp} "成功"
// @Router /upload/image [post]
func (t upload) image(c *gin.Context) {
	var uReq req.UploadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, err := util.VerifyUtil.VerifyFile(c, "file")
	if err != nil {
		response.FailWithMsg(c, response.ParamsValidError, err.Error())
	}
	res, err := t.Srv.UploadImage(file, &uReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 上传文件
// @Description 上传文件
// @Tags 上传
// @Accept  multipart/form-data
// @Produce  json
// @Param token header string true "access_token"
// @Param file formData file true "上传文件"
// @Param cid formData int true "分类ID"
// @Success 200 {object} response.Response{data=resp.FileResp} "成功"
// @Router /upload/file [post]
func (t upload) file(c *gin.Context) {
	var uReq req.UploadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, err := util.VerifyUtil.VerifyFile(c, "file")
	if err != nil {
		response.FailWithMsg(c, response.ParamsValidError, err.Error())
	}
	res, err := t.Srv.UploadFile(file, &uReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 上传音频
// @Description 上传音频
// @Tags 上传
// @Accept  multipart/form-data
// @Produce  json
// @Param token header string true "access_token"
// @Param file formData file true "上传文件"
// @Param cid formData int true "分类ID"
// @Success 200 {object} response.Response{data=resp.FileResp} "成功"
// @Router /upload/audio [post]
func (t upload) audio(c *gin.Context) {
	var uReq req.UploadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, err := util.VerifyUtil.VerifyFile(c, "file")
	if err != nil {
		response.FailWithMsg(c, response.ParamsValidError, err.Error())
	}
	res, err := t.Srv.UploadAudio(file, &uReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 上传视频
// @Description 上传视频
// @Tags 上传
// @Accept  multipart/form-data
// @Produce  json
// @Param token header string true "access_token"
// @Param file formData file true "上传文件"
// @Param cid formData int true "分类ID"
// @Success 200 {object} response.Response{data=resp.FileResp} "成功"
// @Router /upload/video [post]
func (t upload) video(c *gin.Context) {
	var uReq req.UploadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, err := util.VerifyUtil.VerifyFile(c, "file")
	if err != nil {
		response.FailWithMsg(c, response.ParamsValidError, err.Error())
	}
	res, err := t.Srv.UploadVideo(file, &uReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}
