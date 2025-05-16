package service

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"

	"gorm.io/gorm"

	"strings"
)

type AuthRoleService interface {
	All(auth *req.AuthReq) (res []resp.RoleSimpleResp, e error)
	List(page req.PageReq, auth *req.AuthReq) (res response.PageResp, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.RoleResp, e error)
	Add(addReq *req.RoleAddReq, auth *req.AuthReq) (e error)
	Edit(editReq *req.RoleEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
	Change(id uint, auth *req.AuthReq) error
}

type roleService struct {
	db      *gorm.DB
	permSrv AuthPermService
	cache   *cache.Redis
}

func (r roleService) All(auth *req.AuthReq) (res []resp.RoleSimpleResp, e error) {
	var roles []models.AuthRole
	sql := r.db.Model(&models.AuthRole{})
	err := sql.Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	response.Copy(&res, roles)
	return
}

func (r roleService) List(page req.PageReq, auth *req.AuthReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	sql := r.db.Model(&models.AuthRole{})

	var count int64
	err := sql.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var roles []models.AuthRole
	err = sql.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var roleResp []resp.RoleResp
	response.Copy(&roleResp, roles)

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    roleResp,
	}, nil
}

func (r roleService) Detail(id uint, auth *req.AuthReq) (res resp.RoleResp, e error) {
	var role models.AuthRole
	sql := r.db.Model(&models.AuthRole{})
	err := sql.Where("id = ?", id).Limit(1).First(&role).Error
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, role)
	res.Menus, e = r.permSrv.SelectMenuIdsByRoleId(role.ID)
	return
}

func (r roleService) Add(addReq *req.RoleAddReq, auth *req.AuthReq) (e error) {
	var role models.AuthRole
	sql := r.db.Model(&models.AuthRole{})
	var count int64
	sql.Where("name = ? and tenant_id = ?", addReq.Name, auth.TenantId).Limit(1).Count(&count)
	if count > 0 {
		return fmt.Errorf("角色名称已存在!")
	}
	response.Copy(&role, addReq)
	role.Name = strings.Trim(addReq.Name, " ")
	role.TenantId = auth.TenantId
	// 事务
	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&role).Error
		if err != nil {
			return err
		}
		if len(addReq.Menus) > 0 {
			te := r.permSrv.BatchSaveRoleMenusByMenuIds(role.ID, tx, addReq.Menus)
			if te != nil {
				return te
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("添加失败! %s", err.Error())
	}

	return nil
}

func (r roleService) Edit(editReq *req.RoleEditReq, auth *req.AuthReq) (e error) {
	sql := r.db.Model(&models.AuthRole{})
	err := sql.Where("id = ?", editReq.ID).Limit(1).First(&models.AuthRole{}).Error
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	var role models.AuthRole
	var count int64
	sql.Where("name = ? and tenant_id = ? and id <> ?", strings.Trim(editReq.Name, " "), auth.TenantId, editReq.ID).Limit(1).Count(&count)
	if count > 0 {
		return fmt.Errorf("角色名称已存在!")
	}
	response.Copy(&role, editReq)
	role.ID = editReq.ID
	roleMap := structs.Map(editReq)
	delete(roleMap, "ID")
	delete(roleMap, "Menus")
	roleMap["Name"] = strings.Trim(editReq.Name, " ")
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("你没有权限编辑此角色!")
	}
	// 事务
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&role).Updates(roleMap).Error; err != nil {
			return fmt.Errorf("修改失败! %s", err.Error())
		}
		r.permSrv.BatchDeleteRoleMenuByRoleId(editReq.ID, tx)
		r.permSrv.BatchSaveRoleMenusByMenuIds(editReq.ID, tx, editReq.Menus)
		r.permSrv.CacheRoleMenusByRoleId(editReq.ID)
		return nil
	})
	e = response.CheckErr(err, "Edit Transaction err")
	return
}

func (r roleService) Del(id uint, auth *req.AuthReq) (e error) {
	sql := r.db.Model(&models.AuthRole{})
	err := sql.Where("id = ?", id).Limit(1).First(&models.AuthRole{}).Error
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	var count int64
	r.db.Model(&models.User{}).Where("role_id = ?", id).Limit(1).Count(&count)
	if count > 0 {
		return fmt.Errorf("该角色下有用户, 不能删除!")
	}
	if !auth.IsAdmin {
		return fmt.Errorf("你没有权限删除此角色!")
	}
	var role models.AuthRole
	err = r.db.Where("id = ?", id).First(&role).Error
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	tenantID := role.TenantId
	// 事务
	err = r.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&models.AuthRole{}, "id = ?", id).Error
		var te error
		if te = response.CheckErr(txErr, "Del Delete in tx err"); te != nil {
			return te
		}
		if te = r.permSrv.BatchDeleteRoleMenuByRoleId(id, tx); te != nil {
			return te
		}
		cachekey := fmt.Sprintf("%d_%d", tenantID, id)
		r.cache.HDel(types.Admin.BackstageRolesKey, cachekey)
		return nil
	})
	e = response.CheckErr(err, "Del Transaction err")
	return
}

func (r roleService) Change(id uint, auth *req.AuthReq) error {
	var role models.AuthRole
	err := r.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		return fmt.Errorf("角色不存在!")
	}
	if !auth.IsAdmin {
		return fmt.Errorf("你没有权限修改此角色!")
	}
	if auth.RoleId == id {
		return fmt.Errorf("当前角色不能修改!")
	}
	role.IsDisable = 1 - role.IsDisable
	err = r.db.Save(&role).Error
	if err != nil {
		return fmt.Errorf("修改失败!")
	}
	return nil
}

func NewAuthRoleService(db *gorm.DB, rolePermSrv AuthPermService, cache *cache.Redis) AuthRoleService {
	return &roleService{db: db, permSrv: rolePermSrv, cache: cache}
}
