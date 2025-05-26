package common

import (
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/api/service/system"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/storage"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
)

type UploadService interface {
	UploadImage(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error)
	UploadFile(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error)
	UploadAudio(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error)
	UploadVideo(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error)
}

type uploadService struct {
	db     *gorm.DB
	cfgSrv system.ConfigService
}

func (u uploadService) UploadImage(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error) {
	return u.upload(file, auth, "image", uReq.Cid)
}

func (u uploadService) UploadFile(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error) {
	return u.upload(file, auth, "file", uReq.Cid)
}

func (u uploadService) UploadAudio(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error) {
	return u.upload(file, auth, "audio", uReq.Cid)
}

func (u uploadService) UploadVideo(file *multipart.FileHeader, uReq *req.UploadReq, auth *req.AuthReq) (resp.FileResp, error) {
	return u.upload(file, auth, "video", uReq.Cid)
}

func NewUploadService(db *gorm.DB, cfgSrv system.ConfigService) UploadService {
	return &uploadService{db: db, cfgSrv: cfgSrv}
}

func (u uploadService) upload(file *multipart.FileHeader, auth *req.AuthReq, fileType string, cid uint) (resp.FileResp, error) {
	tenantId := auth.TenantId
	userId := auth.UserId
	var optReq req.FileStorageReq
	jstr, err := util.ToolsUtil.ObjToJson(optReq)
	if err != nil {
		return resp.FileResp{}, fmt.Errorf("解析配置失败:%v", err)
	}
	optStr, err := u.cfgSrv.GetVal(tenantId, "storage", "options", jstr)
	if err != nil {
		return resp.FileResp{}, fmt.Errorf("获取配置失败:%v", err)
	}
	var opt map[string]interface{}
	err = util.ToolsUtil.JsonToObj(optStr, &opt)
	if err != nil {
		return resp.FileResp{}, fmt.Errorf("解析配置失败:%v", err)
	}
	folder := fmt.Sprintf("%s/%d", fileType, tenantId)
	uploader := storage.NewStorageDriver(opt)
	fileAdd := models.File{
		Cid:      cid,
		UserId:   userId,
		TenantId: tenantId,
	}
	if auth.IsSuperTenant {
		up, err := uploader.Upload(file, folder, fileType, opt["engine_type"].(string), opt)
		if err != nil {
			return resp.FileResp{}, fmt.Errorf("上传文件失败:%v", err)
		}
		response.Copy(&fileAdd, up)
		if err := u.db.Create(&fileAdd).Error; err != nil {
			return resp.FileResp{}, fmt.Errorf("保存文件信息失败:%v", err)
		}
		var res resp.FileResp
		response.Copy(&res, fileAdd)
		return res, nil
	}
	log.Printf("auth.IsSuperTenant=%+v", auth.IsSuperTenant)
	if opt["engine_type"] == "local" {
		fileSize := file.Size
		// mock 上传文件大小限制
		var maxSize int64 = 50 * 1024 * 1024 // 50M
		var curSize int64
		// 合计该租户已上传的文件大小,字段size
		if err := u.db.Model(&models.File{}).Where("tenant_id = ?", tenantId).Select("SUM(size)").Scan(&curSize).Error; err != nil {
			return resp.FileResp{}, fmt.Errorf("获取已上传文件大小失败:%v", err)
		}
		if curSize+fileSize > maxSize {
			return resp.FileResp{}, fmt.Errorf("上传文件超出限制:%dm/%dm", (curSize+fileSize)/(1024*1024), maxSize/(1024*1024))
		}
	}
	up, err := uploader.Upload(file, folder, fileType, opt["engine_type"].(string), opt)
	if err != nil {
		return resp.FileResp{}, fmt.Errorf("上传文件失败:%v", err)
	}
	response.Copy(&fileAdd, up)
	if err := u.db.Create(&fileAdd).Error; err != nil {
		return resp.FileResp{}, fmt.Errorf("保存文件信息失败:%v", err)
	}
	var res resp.FileResp
	response.Copy(&res, fileAdd)

	res.Uri = util.UrlUtil.ToAbsoluteUrl(up.Uri, opt["engine_type"].(string), opt)

	return res, nil
}
