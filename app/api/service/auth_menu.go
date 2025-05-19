package service

import (
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
)

type AuthMenuService interface {
	SelectMenuByRoleId(auth *req.AuthReq) (mapList interface{}, e error)
	List(auth *req.AuthReq) (res interface{}, e error)
	Detail(id uint, auth *req.AuthReq) (res resp.MenuResp, e error)
	Add(addReq *req.MenuAddReq, auth *req.AuthReq) (e error)
	Edit(editReq *req.MenuEditReq, auth *req.AuthReq) (e error)
	Del(id uint, auth *req.AuthReq) (e error)
}

type authMenuServiceImpl struct {
	db      *gorm.DB
	cache   *cache.Redis
	permSrv AuthPermService
}

func (a authMenuServiceImpl) SelectMenuByRoleId(auth *req.AuthReq) (mapList interface{}, e error) {
	var role models.AuthRole
	if err := a.db.Where("id = ?", auth.RoleId).First(&role).Error; err != nil {
		return nil, fmt.Errorf("角色不存在")
	}

	var menuIds []uint
	if role.TenantId > 1 {
		tenantMenuIds, err := a.permSrv.SelectMenuIdsByTenantId(role.TenantId)
		if err != nil {
			return nil, err
		}
		if role.IsAdmin == 1 {
			menuIds = tenantMenuIds
		} else {
			roleMenuIds, err := a.permSrv.SelectMenuIdsByRoleId(role.ID)
			if err != nil {
				return nil, err
			}
			menuIds = commonIds(roleMenuIds, tenantMenuIds)
		}
	} else {
		var err error
		if role.IsAdmin == 1 {
			var mIds []uint
			a.db.Model(&models.AuthMenu{}).Where("menu_type in (?)", []string{"menu"}).Order("sort asc, id desc").Pluck("id", &mIds)
			menuIds = mIds
		} else {
			menuIds, err = a.permSrv.SelectMenuIdsByRoleId(role.ID)
			if err != nil {
				return nil, err
			}
		}
	}

	var menus []models.AuthMenu
	if err := a.db.Where("id in (?)", menuIds).Order("sort asc, id desc").Find(&menus).Error; err != nil || len(menus) == 0 {
		return nil, fmt.Errorf("菜单不存在")
	}

	var respList []resp.MenuResp
	response.Copy(&respList, menus)

	return util.ArrayUtil.ListToTree(util.ConvertUtil.StructsToMaps(respList), "id", "pid", "children"), nil
}

func (a authMenuServiceImpl) List(auth *req.AuthReq) (res interface{}, e error) {
	var menus []models.AuthMenu
	sql := a.db.Order("id desc")
	if auth.TenantId > 1 {
		tenantMenuIds, err := a.permSrv.SelectMenuIdsByTenantId(auth.TenantId)
		if err != nil {
			return nil, err
		}
		sql = sql.Where("id in (?)", tenantMenuIds)
	}
	if err := sql.Order("sort asc, id desc").Find(&menus).Error; err != nil {
		return nil, err
	}
	var respList []resp.MenuResp
	response.Copy(&respList, menus)
	return util.ArrayUtil.ListToTree(
		util.ConvertUtil.StructsToMaps(respList), "id", "pid", "children"), nil
	// return respList, nil
}

func (a authMenuServiceImpl) Detail(id uint, auth *req.AuthReq) (res resp.MenuResp, e error) {
	var menu models.AuthMenu
	if err := a.db.Where("id = ?", id).First(&menu).Error; err != nil {
		return res, fmt.Errorf("菜单不存在")
	}
	response.Copy(&res, menu)
	if menu.MenuType == "page" {
		var buttons []models.AuthMenu
		a.db.Model(&models.AuthMenu{}).Where("pid = ? and menu_type = ?", id, "action").Order("sort asc, id desc").Find(&buttons)
		var buttonList []resp.MenuButton
		if len(buttons) > 0 {
			response.Copy(&buttonList, buttons)
		}
		res.Button = buttonList
	}
	return res, nil
}

func (a authMenuServiceImpl) Add(addReq *req.MenuAddReq, auth *req.AuthReq) (e error) {
	var count int64
	a.db.Model(&models.AuthMenu{}).Where("name = ?", addReq.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("菜单标识或权限已存在")
	}
	var menu models.AuthMenu
	response.Copy(&menu, addReq)
	if err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&menu).Error; err != nil {
			return err
		}
		if menu.MenuType == "page" {
			for _, button := range addReq.Button {
				b := models.AuthMenu{
					Pid:        menu.ID,
					Title:      button.Title,
					MenuType:   "action",
					Name:       button.Name,
					RenderMenu: false,
				}
				if err := tx.Create(&b).Error; err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	a.cache.Del(types.Admin.BackstageRolesKey)
	return nil
}

func (a authMenuServiceImpl) Edit(editReq *req.MenuEditReq, auth *req.AuthReq) (e error) {
	var count int64
	a.db.Model(&models.AuthMenu{}).Where("name = ? and id != ?", editReq.Name, editReq.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("菜单标识或权限已存在")
	}
	var menu models.AuthMenu
	if err := a.db.Where("id = ?", editReq.ID).First(&menu).Error; err != nil {
		return fmt.Errorf("菜单不存在")
	}
	response.Copy(&menu, editReq)
	if err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&menu).Error; err != nil {
			return err
		}
		if menu.MenuType == "page" {
			if len(editReq.Button) > 0 {
				var btnIds []uint
				tx.Model(&models.AuthMenu{}).Where("pid = ? and menu_type = ?", menu.ID, "action").Pluck("id", &btnIds)
				var nowBtnIds []uint
				for _, button := range editReq.Button {
					var b models.AuthMenu
					if button.ID > 0 {
						if err := tx.Where("id = ?", button.ID).First(&b).Error; err != nil {
							return err
						}
						response.Copy(&b, button)
					} else {
						b = models.AuthMenu{
							Pid:        menu.ID,
							Title:      button.Title,
							MenuType:   "action",
							Name:       button.Name,
							RenderMenu: false,
						}
					}
					if err := tx.Save(&b).Error; err != nil {
						return err
					}
					nowBtnIds = append(nowBtnIds, b.ID)
				}
				if len(nowBtnIds) > 0 {
					var delBtnIds []uint
					for _, btnId := range btnIds {
						if !util.ToolsUtil.Contains(nowBtnIds, btnId) {
							delBtnIds = append(delBtnIds, btnId)
						}
					}
					if len(delBtnIds) > 0 {
						tx.Where("id in (?)", delBtnIds).Delete(&models.AuthMenu{})
					}
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	a.cache.Del(types.Admin.BackstageRolesKey)
	return nil
}

func (a authMenuServiceImpl) Del(id uint, auth *req.AuthReq) (e error) {
	var menu models.AuthMenu
	if err := a.db.Where("id = ?", id).First(&menu).Error; err != nil {
		return fmt.Errorf("菜单不存在")
	}
	if err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&menu).Error; err != nil {
			return err
		}
		var count int64
		tx.Model(&models.AuthMenu{}).Where("pid = ?", id).Count(&count)
		if count > 0 {
			return fmt.Errorf("请先删除子菜单")
		}
		err := a.permSrv.BatchDeleteRoleMenuByMenuId(id, tx)
		if err != nil {
			return err
		}
		err = a.permSrv.BatchDeleteTenantMenuByMenuId(id, tx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func NewAuthMenuService(db *gorm.DB, cache *cache.Redis, permSrv AuthPermService) AuthMenuService {
	return &authMenuServiceImpl{db: db, cache: cache, permSrv: permSrv}
}

func commonIds(ids []uint, ids2 []uint) []uint {
	var res []uint
	for _, id := range ids {
		for _, id2 := range ids2 {
			if id == id2 {
				res = append(res, id)
				break
			}
		}
	}
	return res
}
