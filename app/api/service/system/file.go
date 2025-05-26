package service

import (
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
)

type FileService interface {
	FileList(page *req.PageReq, listReq *req.FileListReq, auth *req.AuthReq) (res response.PageResp, err error)
	FileRename(renameReq *req.FileRenameReq, auth *req.AuthReq) (err error)
	FileMove(moveReq *req.FileMoveReq, auth *req.AuthReq) (err error)
	FileDelete(ids []uint, auth *req.AuthReq) (err error)
	CateList(listReq *req.FileCateListReq, auth *req.AuthReq) (res []interface{}, err error)
	CateAdd(cateReq *req.FileCateAddReq, auth *req.AuthReq) (err error)
	CateRename(cateReq *req.FileCateRenameReq, auth *req.AuthReq) (err error)
	CateDelete(ids []uint, auth *req.AuthReq) (err error)
}

type fileService struct {
	db *gorm.DB
}

func (f fileService) FileList(page *req.PageReq, listReq *req.FileListReq, auth *req.AuthReq) (res response.PageResp, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	chain := f.db.Model(&models.File{}).Where("tenant_id =?", auth.TenantId)
	if listReq.Cid > 0 {
		chain = chain.Where("cid =?", listReq.Cid)
	}
	if listReq.Type > 0 {
		chain = chain.Where("type =?", listReq.Type)
	}
	if listReq.Name != "" {
		chain = chain.Where("name like ?", "%"+listReq.Name+"%")
	}
	var files []models.File
	var count int64
	chain.Count(&count)
	chain.Limit(limit).Offset(offset).Find(&files)
	var result []resp.FileResp
	response.Copy(&result, files)
	res.Count = count
	res.Lists = result
	res.PageNo = page.PageNo
	res.PageSize = page.PageSize
	return res, nil
}

func (f fileService) FileRename(renameReq *req.FileRenameReq, auth *req.AuthReq) (err error) {
	var file models.File
	if err := f.db.Where("id =? and tenant_id =?", renameReq.ID, auth.TenantId).First(&file).Error; err != nil {
		return fmt.Errorf("未找到该文件")
	}
	file.Name = renameReq.Name
	if err := f.db.Save(&file).Error; err != nil {
		return fmt.Errorf("文件重命名失败")
	}
	return nil
}

func (f fileService) FileMove(moveReq *req.FileMoveReq, auth *req.AuthReq) (err error) {
	var files []uint
	f.db.Model(&models.File{}).Where("id in ? and tenant_id =?", moveReq.Ids, auth.TenantId).Pluck("id", &files)
	if len(files) == 0 {
		return fmt.Errorf("未找到该文件")
	}
	if moveReq.Cid > 0 {
		var count int64
		f.db.Model(&models.FileCate{}).Where("id =? and tenant_id =?", moveReq.Cid, auth.TenantId).Count(&count)
		if count == 0 {
			return fmt.Errorf("未找到该类目")
		}
	}
	err = f.db.Model(&models.File{}).Where("id in ?", files).Update("cid", moveReq.Cid).Error
	if err != nil {
		return fmt.Errorf("文件移动失败")
	}
	return nil
}

func (f fileService) FileDelete(ids []uint, auth *req.AuthReq) (err error) {
	var files []uint
	f.db.Model(&models.File{}).Where("id in ? and tenant_id =?", ids, auth.TenantId).Pluck("id", &files)
	if len(files) == 0 {
		return fmt.Errorf("未找到该文件")
	}
	err = f.db.Model(&models.File{}).Where("id in ?", files).Delete(&models.File{}).Error
	if err != nil {
		return fmt.Errorf("文件删除失败")
	}
	return nil
}

func (f fileService) CateList(listReq *req.FileCateListReq, auth *req.AuthReq) (res []interface{}, err error) {
	var cates []models.FileCate
	chain := f.db.Model(&models.FileCate{}).Where("tenant_id =?", auth.TenantId)
	if listReq.Type > 0 {
		chain = chain.Where("type =?", listReq.Type)
	}
	if listReq.Name != "" {
		chain = chain.Where("name like ?", "%"+listReq.Name+"%")
	}
	chain.Order("sort asc").Find(&cates)
	var result []resp.FileCateResp
	response.Copy(&result, cates)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(result), "id", "pid", "children"), nil
}

func (f fileService) CateAdd(cateReq *req.FileCateAddReq, auth *req.AuthReq) (err error) {
	var cate models.FileCate
	response.Copy(&cate, cateReq)
	cate.TenantId = auth.TenantId
	if err := f.db.Create(&cate).Error; err != nil {
		return fmt.Errorf("类目添加失败")
	}
	return nil
}

func (f fileService) CateRename(cateReq *req.FileCateRenameReq, auth *req.AuthReq) (err error) {
	var cate models.FileCate
	if err := f.db.Where("id =? and tenant_id =?", cateReq.ID, auth.TenantId).First(&cate).Error; err != nil {
		return fmt.Errorf("未找到该类目")
	}
	cate.Name = cateReq.Name
	if err := f.db.Save(&cate).Error; err != nil {
		return fmt.Errorf("类目重命名失败")
	}
	return nil
}

func (f fileService) CateDelete(ids []uint, auth *req.AuthReq) (err error) {
	var cates []models.FileCate
	f.db.Model(&models.FileCate{}).Where("id in ? and tenant_id =?", ids, auth.TenantId).Find(&cates)
	if len(cates) == 0 {
		return fmt.Errorf("未找到该类目")
	}
	err = f.db.Model(&models.FileCate{}).Where("id in ?", ids).Delete(&models.FileCate{}).Error
	if err != nil {
		return fmt.Errorf("类目删除失败")
	}
	return nil
}

func NewFileService(db *gorm.DB) FileService {
	return &fileService{db: db}
}
