package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/api/service/auth"
	"github.com/hulutech-web/workflow-engine/app/api/service/common"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

type UserService interface {
	FindByUsername(username string, tenantId uint) (*models.User, error)
	List(page *req.PageReq, query *req.UserQueryReq, auth *req.AuthReq) (response.PageResp, error)
	Index(ctx *gin.Context, query url.Values) (*common.PageResult, error)
	Detail(userId uint) (resp.UserResp, error)
	Add(userReq *req.UserAddReq, auth *req.AuthReq) error
	Edit(userReq *req.UserEditReq, auth *req.AuthReq) error
	Update(userReq *req.UserUpdateReq) error
	Delete(userId uint, auth *req.AuthReq) error
	Disable(userId uint, auth *req.AuthReq) error
	CacheUserById(userId uint) error
	Self(auth *req.AuthReq) (*resp.UserSelfResp, error)
}

type userServiceImpl struct {
	db      *gorm.DB
	cache   *cache.Redis
	permSrv auth.AuthPermService
}

func (d userServiceImpl) Index(ctx *gin.Context, query url.Values) (*common.PageResult, error) {
	var tmpls []models.User
	paginatorService := common.NewPaginatorServiceImpl(d.db, ctx)

	err, result := paginatorService.SearchByParams(query, nil).ResultPagination(&tmpls)
	return result, err
}
func (u userServiceImpl) FindByUsername(username string, tenantId uint) (*models.User, error) {
	var user models.User
	if err := u.db.Where("username =? AND tenant_id =?", username, tenantId).Preload("Role").Preload("Tenant").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, fmt.Errorf("数据库查询错误: %v", err)
	}
	return &user, nil
}

func (u userServiceImpl) Detail(userId uint) (resp.UserResp, error) {
	var user models.User
	if err := u.db.Preload("Role").Preload("Tenant").First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.UserResp{}, fmt.Errorf("用户不存在")
		}
		return resp.UserResp{}, fmt.Errorf("数据库查询错误: %v", err)
	}
	var res resp.UserResp
	response.Copy(&res, user)
	return res, nil
}

func (u userServiceImpl) List(page *req.PageReq, query *req.UserQueryReq, auth *req.AuthReq) (response.PageResp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	sql := u.db.Model(&models.User{})
	if !auth.IsSuperTenant {
		sql = sql.Where("tenant_id = ?", auth.TenantId)
	}
	if len(query.Username) > 0 {
		sql = sql.Where("username LIKE ?", fmt.Sprintf("%%%s%%", query.Username))
	}
	if len(query.Email) > 0 {
		sql = sql.Where("email LIKE ?", fmt.Sprintf("%%%s%%", query.Email))
	}
	if len(query.Phone) > 0 {
		sql = sql.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", query.Phone))
	}
	if query.IsDisable > -1 {
		sql = sql.Where("is_disable = ?", query.IsDisable)
	}
	if query.RoleId > 0 {
		sql = sql.Where("role_id = ?", query.RoleId)
	}
	var users []models.User
	var count int64
	sql.Count(&count)
	if err := sql.Order("id desc").Limit(limit).Offset(offset).Preload("Role").Preload("Tenant").Find(&users).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("数据库查询错误: %v", err)
	}
	var res []resp.UserResp
	response.Copy(&res, users)
	return response.PageResp{
		Count:    count,
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Lists:    res,
	}, nil
}

func (u userServiceImpl) Add(userReq *req.UserAddReq, auth *req.AuthReq) error {
	var count int64
	u.db.Model(&models.User{}).Where("username =? AND tenant_id =?", userReq.Username, auth.TenantId).Count(&count)
	if count > 0 {
		return fmt.Errorf("用户名已存在")
	}
	var user models.User
	response.Copy(&user, userReq)
	// 哪个租户创建的用户，就默认是该租户的用户
	user.TenantId = auth.TenantId
	salt := util.ToolsUtil.RandomString(5)
	user.Password = util.ToolsUtil.MakeMd5(fmt.Sprintf("%s%s", userReq.Password, salt))
	user.Salt = salt
	if err := u.db.Create(&user).Error; err != nil {
		return fmt.Errorf("数据库插入错误: %v", err)
	}
	return nil
}

func (u userServiceImpl) Edit(userReq *req.UserEditReq, auth *req.AuthReq) error {
	var count int64
	u.db.Model(&models.User{}).Where("username =? AND tenant_id =? AND id <> ?", userReq.Username, auth.TenantId, userReq.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("用户名已存在")
	}
	var user models.User
	if err := u.db.First(&user, userReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return fmt.Errorf("数据库查询错误: %v", err)
	}
	// 不能修改租户ID
	userReq.TenantId = user.TenantId
	response.Copy(&user, userReq)
	if err := u.db.Save(&user).Error; err != nil {
		return fmt.Errorf("数据库更新错误: %v", err)
	}
	return nil
}

func (u userServiceImpl) Update(userReq *req.UserUpdateReq) error {
	var user models.User
	if err := u.db.First(&user, userReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return fmt.Errorf("数据库查询错误: %v", err)
	}
	response.Copy(&user, userReq)
	if err := u.db.Save(&user).Error; err != nil {
		return fmt.Errorf("数据库更新错误: %v", err)
	}
	return nil
}

func (u userServiceImpl) Delete(userId uint, auth *req.AuthReq) error {
	var user models.User
	if err := u.db.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return fmt.Errorf("数据库查询错误: %v", err)
	}
	if user.ID == auth.UserId {
		return fmt.Errorf("不能删除自己")
	}
	if !auth.IsSuperTenant {
		if !auth.IsAdmin {
			return fmt.Errorf("无权限删除用户")
		}
	}
	if err := u.db.Delete(&user).Error; err != nil {
		return fmt.Errorf("数据库删除错误: %v", err)
	}
	return nil
}

func (u userServiceImpl) Disable(userId uint, auth *req.AuthReq) error {
	var user models.User
	if userId == auth.UserId {
		return fmt.Errorf("不能禁用自己")
	}
	if err := u.db.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return fmt.Errorf("数据库查询错误: %v", err)
	}
	user.IsDisable = 1 - user.IsDisable
	if err := u.db.Save(&user).Error; err != nil {
		return fmt.Errorf("数据库更新错误: %v", err)
	}
	return nil
}

func (u userServiceImpl) CacheUserById(userId uint) error {
	user, err := u.Detail(userId)
	if err != nil {
		return fmt.Errorf("获取用户信息错误: %v", err)
	}
	str, err := util.ToolsUtil.ObjToJson(&user)
	if err != nil {
		return fmt.Errorf("缓存用户数据错误: %v", err)
	}
	u.cache.HSet(types.Admin.BackstageManageKey, fmt.Sprintf("%d", user.ID), str, 0)
	return nil
}

func (u userServiceImpl) Self(auth *req.AuthReq) (*resp.UserSelfResp, error) {
	var user models.User
	var e error

	if err := u.db.Preload("Role").Preload("Tenant").First(&user, auth.UserId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, fmt.Errorf("数据库查询错误: %v", err)
	}
	var auths []string
	if !auth.IsAdmin && !auth.IsSuperTenant {
		roleId := user.RoleId
		var menuIds []uint
		if menuIds, e = u.permSrv.SelectMenuIdsByRoleId(roleId); e != nil {
			return nil, fmt.Errorf("获取权限菜单ID错误: %v", e)
		}
		if len(menuIds) > 0 {
			var menus []models.AuthMenu
			err := u.db.Where(
				"id in ? AND menu_type in ?", menuIds, 0, []string{"menu", "page"}).Order(
				"sort asc, id desc").Find(&menus).Error
			if e = response.CheckErr(err, "Self SystemAuthMenu Find err"); e != nil {
				return nil, e
			}
			if len(menus) > 0 {
				for _, v := range menus {
					auths = append(auths, strings.Trim(v.Name, " "))
				}
			}
		}
		if !auth.IsSuperTenant {
			var tentMenus []uint
			var menus []models.AuthMenu
			tentMenus, e = u.permSrv.SelectMenuIdsByTenantId(auth.TenantId)
			err := u.db.Where(
				"id in ? AND menu_type in ?", tentMenus, 0, []string{"menu", "page"}).Order(
				"sort asc, id desc").Find(&menus).Error
			if e = response.CheckErr(err, "Self SystemAuthMenu Find err"); e != nil {
				return nil, e
			}
			if len(menus) > 0 {
				for _, v := range menus {
					auths = append(auths, strings.Trim(v.Name, " "))
				}
			}
		}
		if len(auths) > 0 {
			auths = append(auths, "")
		}
	} else {
		auths = append(auths, "*")
	}
	var userRes resp.UserResp
	response.Copy(&userRes, user)
	return &resp.UserSelfResp{
		User:        userRes,
		Permissions: auths,
	}, nil
}

func NewUserService(db *gorm.DB, cache *cache.Redis, permSrv auth.AuthPermService) UserService {
	return &userServiceImpl{db: db, cache: cache, permSrv: permSrv}
}
