package system

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/zhoudm1743/Seven/app/admin/schemas/query"
	"github.com/zhoudm1743/Seven/app/admin/schemas/req"
	"github.com/zhoudm1743/Seven/app/admin/schemas/resp"
	"github.com/zhoudm1743/Seven/app/models/system"
	"github.com/zhoudm1743/Seven/pkg/common/config"
	"github.com/zhoudm1743/Seven/pkg/common/response"
	"github.com/zhoudm1743/Seven/pkg/util"
	"gorm.io/gorm"

	"strings"
)

type RoleService interface {
	All(auth *req.AuthReq) (res []resp.SystemRoleSimpleResp, e error)
	List(page req.PageReq, auth *req.AuthReq) (res response.PageResp, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.SystemRoleResp, e error)
	Add(addReq req.SystemRoleAddReq, auth *req.AuthReq) (e error)
	Edit(editReq req.SystemRoleEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
}

type roleService struct {
	db          *gorm.DB
	rolePermSrv RolePermService
	cfg         *config.Config
}

func (r roleService) All(auth *req.AuthReq) (res []resp.SystemRoleSimpleResp, e error) {
	var roles []system.Role
	sql := query.AuthQuery(r.db, auth)
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
	roleModel := query.AuthQuery(r.db.Model(&system.Role{}), auth)

	var count int64
	err := roleModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var roles []system.Role
	err = roleModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var roleResp []resp.SystemRoleResp
	response.Copy(&roleResp, roles)

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    roleResp,
	}, nil
}

func (r roleService) Detail(id uint, auth *req.AuthReq) (res resp.SystemRoleResp, e error) {
	var role system.Role
	sql := query.AuthQuery(r.db, auth)
	err := sql.Where("id = ?", id).Limit(1).First(&role).Error
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, role)
	res.Menus, e = r.rolePermSrv.SelectMenuIdsByRoleId(role.ID, auth)
	return
}

func (r roleService) Add(addReq req.SystemRoleAddReq, auth *req.AuthReq) (e error) {
	var role system.Role
	sql := query.AuthQuery(r.db, auth)
	if r := sql.Where("name = ?", strings.Trim(addReq.Name, " ")).Limit(1).First(&role); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色名称已存在!")
	}
	response.Copy(&role, addReq)
	role.Name = strings.Trim(addReq.Name, " ")
	role.TenantId = auth.TenantID
	// 事务
	err := r.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Create(&role).Error
		var te error
		if te = response.CheckErr(txErr, "Add Create in tx err"); te != nil {
			return te
		}
		te = r.rolePermSrv.BatchSaveByMenuIds(role.ID, tx, addReq.MenuIds, auth)
		return te
	})
	e = response.CheckErr(err, "Add Transaction err")
	return
}

func (r roleService) Edit(editReq req.SystemRoleEditReq, auth *req.AuthReq) (e error) {
	sql := query.AuthQuery(r.db, auth)
	err := sql.Where("id = ?", editReq.ID).Limit(1).First(&system.Role{}).Error
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	var role system.Role
	if r := sql.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).Limit(1).First(&role); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色名称已存在!")
	}
	role.ID = editReq.ID
	roleMap := structs.Map(editReq)
	delete(roleMap, "ID")
	delete(roleMap, "MenuIds")
	roleMap["Name"] = strings.Trim(editReq.Name, " ")
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("你没有权限编辑此角色!")
	}
	// 事务
	err = r.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&role).Updates(roleMap).Error
		var te error
		if te = response.CheckErr(txErr, "Edit Updates in tx err"); te != nil {
			return te
		}
		if te = r.rolePermSrv.BatchDeleteByRoleId(editReq.ID, tx, auth); te != nil {
			return te
		}
		if te = r.rolePermSrv.BatchSaveByMenuIds(editReq.ID, tx, editReq.MenuIds, auth); te != nil {
			return te
		}
		te = r.rolePermSrv.CacheRoleMenusByRoleId(editReq.ID)
		return te
	})
	e = response.CheckErr(err, "Edit Transaction err")
	return
}

func (r roleService) Del(id uint, auth *req.AuthReq) (e error) {
	sql := query.AuthQuery(r.db, auth)
	err := sql.Where("id = ?", id).Limit(1).First(&system.Role{}).Error
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	if r := sql.Where("role = ?", id).Limit(1).Find(&system.Admin{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色已被管理员使用,请先移除!")
	}
	if !auth.IsAdmin {
		return response.AssertArgumentError.Make("你没有权限删除此角色!")
	}
	var role system.Role
	err = r.db.Where("id = ?", id).First(&role).Error
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	tenantID := role.TenantId
	// 事务
	err = r.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&system.Role{}, "id = ?", id).Error
		var te error
		if te = response.CheckErr(txErr, "Del Delete in tx err"); te != nil {
			return te
		}
		if te = r.rolePermSrv.BatchDeleteByRoleId(id, tx, auth); te != nil {
			return te
		}
		cachekey := fmt.Sprintf("%d_%d", tenantID, id)
		util.RedisUtil.HDel(r.cfg.Admin.BackstageRolesKey, cachekey)
		return nil
	})
	e = response.CheckErr(err, "Del Transaction err")
	return
}

func NewRoleService(db *gorm.DB, rolePermSrv RolePermService, cfg *config.Config) RoleService {
	return &roleService{db: db, rolePermSrv: rolePermSrv, cfg: cfg}
}
