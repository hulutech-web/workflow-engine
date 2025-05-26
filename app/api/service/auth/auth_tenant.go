package auth

import (
	"errors"
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"gorm.io/gorm"
)

type AuthTenantService interface {
	All() ([]resp.TenantResp, error)
	List(page req.PageReq, listReq req.TenantQueryReq) (response.PageResp, error)
	Detail(id uint) (resp.TenantResp, error)
	Add(addReq req.TenantAddReq, auth *req.AuthReq) error
	Edit(editReq req.TenantEditReq, auth *req.AuthReq) error
	Del(id uint, auth *req.AuthReq) error
}

type tenantServiceImpl struct {
	db      *gorm.DB
	cache   *cache.Redis
	permSrv AuthPermService
}

func (t tenantServiceImpl) All() ([]resp.TenantResp, error) {
	var tenants []models.AuthTenant
	if err := t.db.Order("id desc").Find(&tenants).Error; err != nil {
		return nil, err
	}
	var res []resp.TenantResp
	response.Copy(&res, tenants)
	return res, nil
}

func (t tenantServiceImpl) List(page req.PageReq, listReq req.TenantQueryReq) (response.PageResp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	sql := t.db.Model(&models.AuthTenant{})
	if len(listReq.Name) > 0 {
		sql = sql.Where("name LIKE ?", fmt.Sprintf("%%%s%%", listReq.Name))
	}
	var tenants []models.AuthTenant
	var count int64
	sql.Count(&count)
	if err := sql.Order("id desc").Limit(limit).Offset(offset).Find(&tenants).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("数据库查询错误: %v", err)
	}
	var res []resp.TenantResp
	response.Copy(&res, tenants)
	return response.PageResp{
		Count:    count,
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Lists:    res,
	}, nil
}

func (t tenantServiceImpl) Detail(id uint) (resp.TenantResp, error) {
	var tenant models.AuthTenant
	if err := t.db.First(&tenant, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.TenantResp{}, fmt.Errorf("租户不存在")
		}
		return resp.TenantResp{}, fmt.Errorf("数据库查询错误: %v", err)
	}
	var res resp.TenantResp
	response.Copy(&res, tenant)
	t.db.Model(&models.AuthPerm{}).Where("type = ? and type_id = ?", "tenant", id).Pluck("menu_id", &res.Menus)
	return res, nil
}

func (t tenantServiceImpl) Add(addReq req.TenantAddReq, auth *req.AuthReq) error {
	if !auth.IsSuperTenant {
		return fmt.Errorf("无权限操作")
	}
	var count int64
	t.db.Model(&models.AuthTenant{}).Where("name = ?", addReq.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("租户名称已存在")
	}
	var tenant models.AuthTenant
	response.Copy(&tenant, addReq)
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := t.db.Create(&tenant).Error; err != nil {
			return fmt.Errorf("数据库插入错误: %v", err)
		}
		if len(addReq.Menus) > 0 {
			if err := t.permSrv.BatchSaveTenantMenusByMenuIds(tenant.ID, tx, addReq.Menus); err != nil {
				return fmt.Errorf("菜单权限保存错误: %v", err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("数据库插入错误: %v", err)
	}

	return nil
}

func (t tenantServiceImpl) Edit(editReq req.TenantEditReq, auth *req.AuthReq) error {
	if !auth.IsSuperTenant {
		return fmt.Errorf("无权限操作")
	}
	var count int64
	t.db.Model(&models.AuthTenant{}).Where("name = ?", editReq.Name).Where("id <> ?", editReq.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("租户名称已存在")
	}
	var tenant models.AuthTenant
	if err := t.db.First(&tenant, editReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("租户不存在")
		}
		return fmt.Errorf("数据库查询错误: %v", err)
	}
	response.Copy(&tenant, editReq)
	if err := t.db.Save(&tenant).Error; err != nil {
		return fmt.Errorf("数据库更新错误: %v", err)
	}
	if len(editReq.Menus) > 0 {
		_ = t.permSrv.BatchDeleteByTenantId(editReq.ID, t.db)
		t.cache.HDel(types.Admin.BackstageTenantsKey, fmt.Sprintf("%d", editReq.ID))
		if err := t.permSrv.BatchSaveTenantMenusByMenuIds(tenant.ID, t.db, editReq.Menus); err != nil {
			return fmt.Errorf("菜单权限保存错误: %v", err)
		}
	}
	return nil
}

func (t tenantServiceImpl) Del(id uint, auth *req.AuthReq) error {
	if !auth.IsSuperTenant {
		return fmt.Errorf("无权限操作")
	}
	if err := t.db.Delete(&models.AuthTenant{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("数据库删除错误: %v", err)
	}
	_ = t.permSrv.BatchDeleteByTenantId(id, t.db)
	t.cache.HDel(types.Admin.BackstageTenantsKey, fmt.Sprintf("%d", id))
	return nil
}

func NewAuthTenantService(db *gorm.DB, cache *cache.Redis, permSrv AuthPermService) AuthTenantService {
	return &tenantServiceImpl{db: db, cache: cache, permSrv: permSrv}
}
