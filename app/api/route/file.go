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

type file struct {
	fx.In
	Srv service.FileService
}

func fileRoutes(t file, r *types.ApiRouter) {
	fileGroup := r.RouterGroup.Group("/common/file")

	fileGroup.GET("/fileList", t.fileList)
	fileGroup.POST("/fileRename", t.fileRename, r.Log("附件重命名"))
	fileGroup.POST("/fileMove", t.fileMove, r.Log("附件移动"))
	fileGroup.POST("/fileDelete", t.fileDelete, r.Log("附件删除"))
	fileGroup.GET("/cateList", t.cateList)
	fileGroup.POST("/cateAdd", t.cateAdd, r.Log("附件分类添加"))
	fileGroup.POST("/cateRename", t.cateRename, r.Log("附件分类重命名"))
	fileGroup.POST("/cateDelete", t.cateDelete, r.Log("附件分类删除"))
}

// @BasePath /api
// @Summary 附件列表
// @Description 获取附件列表
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request query req.PageReq true "分页参数"
// @Param request body req.FileListReq true "附件列表参数"
// @Success 200 {object} response.Response{data=response.PageResp{Count=int64,PageNo=int,PageSize=int,Lists=[]resp.FileResp}} "成功"
// @Router /common/file/fileList [get]
func (t file) fileList(c *gin.Context) {
	var page req.PageReq
	var listReq req.FileListReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &page, &listReq)) {
		return
	}
	res, err := t.Srv.FileList(&page, &listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 附件重命名
// @Description 附件重命名
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.FileRenameReq true "附件重命名参数"
// @Success 200 {object} response.Response "成功"
// @Router /common/file/fileRename [post]
func (t file) fileRename(c *gin.Context) {
	var renameReq req.FileRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &renameReq)) {
		return
	}
	err := t.Srv.FileRename(&renameReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 附件移动
// @Description 附件移动
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.FileMoveReq true "附件移动参数"
// @Success 200 {object} response.Response "成功"
// @Router /common/file/fileMove [post]
func (t file) fileMove(c *gin.Context) {
	var moveReq req.FileMoveReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &moveReq)) {
		return
	}
	err := t.Srv.FileMove(&moveReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 附件删除
// @Description 附件删除
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdListReq true "附件删除参数"
// @Success 200 {object} response.Response "成功"
// @Router /common/file/fileDelete [post]
func (t file) fileDelete(c *gin.Context) {
	var deleteReq req.IdListReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &deleteReq)) {
		return
	}
	err := t.Srv.FileDelete(deleteReq.Ids, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 附件分类列表
// @Description 获取附件分类列表
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.FileCateListReq true "附件分类列表参数"
// @Success 200 {object} response.Response{data=[]interface{}} "成功"
// @Router /common/file/cateList [get]
func (t file) cateList(c *gin.Context) {
	var listReq req.FileCateListReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &listReq)) {
		return
	}
	res, err := t.Srv.CateList(&listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

// @BasePath /api
// @Summary 附件分类添加
// @Description 附件分类添加
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.FileCateAddReq true "附件分类添加参数"
// @Success 200 {object} response.Response "成功"
// @Router /common/file/cateAdd [post]
func (t file) cateAdd(c *gin.Context) {
	var cateReq req.FileCateAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &cateReq)) {
		return
	}
	err := t.Srv.CateAdd(&cateReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 附件分类重命名
// @Description 附件分类重命名
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.FileCateRenameReq true "附件分类重命名参数"
// @Success 200 {object} response.Response "成功"
// @Router /common/file/cateRename [post]
func (t file) cateRename(c *gin.Context) {
	var cateReq req.FileCateRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &cateReq)) {
		return
	}
	err := t.Srv.CateRename(&cateReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

// @BasePath /api
// @Summary 附件分类删除
// @Description 附件分类删除
// @Tags 附件管理
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.IdListReq true "附件分类删除参数"
// @Success 200 {object} response.Response "成功"
// @Router /common/file/cateDelete [post]
func (t file) cateDelete(c *gin.Context) {
	var deleteReq req.IdListReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &deleteReq)) {
		return
	}
	err := t.Srv.CateDelete(deleteReq.Ids, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
