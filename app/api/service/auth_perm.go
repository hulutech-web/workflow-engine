package service

import (
	"errors"
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type AuthPermService interface {
	SelectMenuIdsByRoleId(roleId uint) (menuIds []uint, e error)
	CacheRoleMenusByRoleId(roleId uint) (e error)
	BatchSaveRoleMenusByMenuIds(roleId uint, db *gorm.DB, menuIds string) (e error)
	BatchDeleteRoleMenuByRoleId(roleId uint, db *gorm.DB) (e error)
	BatchDeleteRoleMenuByMenuId(menuId uint, db *gorm.DB) (e error)

	SelectMenuIdsByTenantId(tenantId uint) (menuIds []uint, e error)
	CacheTenantMenusByTenantId(tenantId uint) (e error)
	BatchSaveTenantMenusByMenuIds(tenantId uint, db *gorm.DB, menuIds string) (e error)
	BatchDeleteByTenantId(tenantId uint, db *gorm.DB) (e error)
	BatchDeleteTenantMenuByMenuId(menuId uint, db *gorm.DB) (e error)
}

type authPermImpl struct {
	db    *gorm.DB
	cache *cache.Redis
}

func (a authPermImpl) SelectMenuIdsByRoleId(roleId uint) (menuIds []uint, e error) {
	if roleId == 0 {
		return []uint{}, nil
	}
	if err := a.db.Model(&models.AuthPerm{}).Where("`type` = ? AND `type_id` = ?", "role", roleId).Pluck("menu_id", &menuIds).Error; err != nil {
		return nil, fmt.Errorf("查询角色权限失败: %v", err)
	}
	return menuIds, nil
}

func (a authPermImpl) CacheRoleMenusByRoleId(roleId uint) (e error) {
	if roleId == 0 {
		return fmt.Errorf("角色ID不能为空")
	}
	var role models.AuthRole
	if err := a.db.First(&role, roleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("角色不存在")
		}
		return fmt.Errorf("查询角色失败: %v", err)
	}
	menuIds, err := a.SelectMenuIdsByRoleId(roleId)
	if err != nil {
		return err
	}
	var menus []models.AuthMenu
	a.db.Where("id in (?) and menuType in (?)", menuIds, []string{"dir", "page"}).Order("id desc").Find(&menus)
	if len(menus) == 0 {
		return fmt.Errorf("角色没有权限")
	}
	var menuArr []string
	for _, menu := range menus {
		menuArr = append(menuArr, strings.Trim(menu.Name, ""))
	}
	if len(types.Admin.CommonUri) > 0 {
		menuArr = append(menuArr, types.Admin.CommonUri...)
	}
	key := fmt.Sprintf("%d", roleId)
	a.cache.HSet(types.Admin.BackstageRolesKey, key, strings.Join(menuArr, ","), 0)
	return nil
}

func (a authPermImpl) BatchSaveRoleMenusByMenuIds(roleId uint, db *gorm.DB, menuIds string) (e error) {
	if roleId == 0 {
		return fmt.Errorf("角色ID不能为空")
	}
	if menuIds == "" {
		return fmt.Errorf("角色权限不能为空")
	}
	if db == nil {
		db = a.db
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []models.AuthPerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.ParseUint(menuIdStr, 10, 32)
			perms = append(perms, models.AuthPerm{ID: util.ToolsUtil.MakeUuid(), Type: "role", TypeId: roleId, MenuId: uint(menuId)})
		}
		if err := tx.Create(&perms).Error; err != nil {
			return tx.Rollback().Error
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("保存角色权限失败: %v", err)
	}
	return nil
}

func (a authPermImpl) BatchDeleteRoleMenuByRoleId(roleId uint, db *gorm.DB) (e error) {
	if db == nil {
		db = a.db
	}
	if err := db.Where("`type` = ? AND `type_id` = ?", "role", roleId).Delete(&models.AuthPerm{}).Error; err != nil {
		return fmt.Errorf("删除角色权限失败: %v", err)
	}
	return nil
}

func (a authPermImpl) BatchDeleteRoleMenuByMenuId(menuId uint, db *gorm.DB) (e error) {
	if db == nil {
		db = a.db
	}
	if err := db.Where("`menu_id` = ? and `type` = ?", menuId, "role").Delete(&models.AuthPerm{}).Error; err != nil {
		return fmt.Errorf("删除角色权限失败: %v", err)
	}
	return nil
}

func (a authPermImpl) SelectMenuIdsByTenantId(tenantId uint) (menuIds []uint, e error) {
	if tenantId == 0 {
		return nil, nil
	}
	err := a.db.Model(&models.AuthPerm{}).Where("`type` = ? AND `type_id` = ?", "tenant", tenantId).Pluck("menu_id", &menuIds).Error
	if err != nil {
		return nil, fmt.Errorf("查询租户权限失败: %v", err)
	}
	return menuIds, nil
}

func (a authPermImpl) CacheTenantMenusByTenantId(tenantId uint) (e error) {
	if tenantId == 0 {
		return fmt.Errorf("租户ID不能为空")
	}
	var menuIds []uint
	if err := a.db.Model(&models.AuthPerm{}).Where("`type` = ? AND `type_id` = ?", "tenant", tenantId).Pluck("menu_id", &menuIds).Error; err != nil {
		return fmt.Errorf("查询租户权限失败: %v", err)
	}
	var menus []models.AuthMenu
	a.db.Where("id in (?) and menuType in (?)", menuIds, []string{"dir", "page"}).Order("id desc").Find(&menus)
	if len(menus) == 0 {
		return fmt.Errorf("租户没有权限")
	}
	var menuArr []string
	for _, menu := range menus {
		menuArr = append(menuArr, strings.Trim(menu.Name, ""))
	}
	if len(types.Admin.CommonUri) > 0 {
		menuArr = append(menuArr, types.Admin.CommonUri...)
	}
	key := fmt.Sprintf("%d", tenantId)
	a.cache.HSet(types.Admin.BackstageTenantsKey, key, strings.Join(menuArr, ","), 0)
	return nil
}

func (a authPermImpl) BatchSaveTenantMenusByMenuIds(tenantId uint, db *gorm.DB, menuIds string) (e error) {
	if tenantId == 0 {
		return fmt.Errorf("租户ID不能为空")
	}
	if menuIds == "" {
		return fmt.Errorf("租户权限不能为空")
	}
	if db == nil {
		db = a.db
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var perms []models.AuthPerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.ParseUint(menuIdStr, 10, 32)
			perms = append(perms, models.AuthPerm{ID: util.ToolsUtil.MakeUuid(), Type: "tenant", TypeId: tenantId, MenuId: uint(menuId)})
		}
		if err := tx.Create(&perms).Error; err != nil {
			return tx.Rollback().Error
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("保存租户权限失败: %v", err)
	}
	return nil
}

func (a authPermImpl) BatchDeleteByTenantId(tenantId uint, db *gorm.DB) (e error) {
	if db == nil {
		db = a.db
	}
	if err := db.Where("`type` = ? AND `type_id` = ?", "tenant", tenantId).Delete(&models.AuthPerm{}).Error; err != nil {
		return fmt.Errorf("删除租户权限失败: %v", err)
	}
	return nil
}

func (a authPermImpl) BatchDeleteTenantMenuByMenuId(menuId uint, db *gorm.DB) (e error) {
	if db == nil {
		db = a.db
	}
	if err := db.Where("`menu_id` = ? and `type` = ?", menuId, "tenant").Delete(&models.AuthPerm{}).Error; err != nil {
		return fmt.Errorf("删除租户权限失败: %v", err)
	}
	return nil
}

func NewAuthPerm(db *gorm.DB, cache *cache.Redis) AuthPermService {
	return &authPermImpl{db: db, cache: cache}
}
