package service

import (
	"errors"
	"fmt"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/core/config"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

type AccountService interface {
	Login(loginReq *req.AccountLoginReq) (*resp.AccountLoginResp, error)
	Logout(token string) error
	TenantList() ([]resp.SelectOption, error)
	Register(r *req.AccountRegisterReq) (*resp.AccountLoginResp, error)
}

type authService struct {
	db      *gorm.DB
	cfg     *config.Config
	cache   *cache.Redis
	userSrv UserService
}

func (a authService) Login(loginReq *req.AccountLoginReq) (*resp.AccountLoginResp, error) {
	user, err := a.userSrv.FindByUsername(loginReq.Username, loginReq.TenantId)
	if err != nil {
		return nil, err
	}
	md5Pwd := util.ToolsUtil.MakeMd5(loginReq.Password + user.Salt)
	if md5Pwd != user.Password {
		return nil, fmt.Errorf("用户名或密码错误")
	}
	if user.IsDisable == 1 {
		return nil, fmt.Errorf("用户已被禁用")
	}

	token := util.ToolsUtil.MakeToken()
	key := fmt.Sprintf("%d", user.ID)
	// 不是多点登录
	if user.IsMultipoint == 0 {
		sysAdminSetKey := types.Admin.BackstageTokenSet + key
		ts := a.cache.SGet(sysAdminSetKey)
		if len(ts) > 0 {
			var keys []string
			for _, t := range ts {
				keys = append(keys, t)
			}
			a.cache.Del(keys...)
		}
		a.cache.Del(sysAdminSetKey)
		a.cache.SSet(sysAdminSetKey, token)
	}
	// 缓存用户信息
	t, _ := time.ParseDuration(a.cfg.Jwt.AccessExpiry)
	a.cache.Set(types.Admin.BackstageTokenKey+token, key, cast.ToInt(t.Seconds()))
	_ = a.userSrv.CacheUserById(user.ID)
	// 返回用户信息
	var userResp resp.UserResp
	response.Copy(&userResp, user)
	return &resp.AccountLoginResp{
		Token:    token,
		UserInfo: userResp,
	}, nil
}

func (a authService) Logout(token string) error {
	a.cache.Del(types.Admin.BackstageTokenKey + token)
	return nil
}

func (a authService) TenantList() ([]resp.SelectOption, error) {
	var list []models.AuthTenant
	a.db.Order("id desc").Find(&list)
	var res []resp.SelectOption
	for _, v := range list {
		res = append(res, resp.SelectOption{
			Value:    v.ID,
			Label:    v.Name,
			Disabled: false,
		})
	}
	return res, nil
}

func (a authService) Register(r *req.AccountRegisterReq) (*resp.AccountLoginResp, error) {
	var tenant models.AuthTenant
	if err := a.db.Where("id = ?", r.TenantId).First(&tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("租户不存在")
		}
		return nil, fmt.Errorf("获取租户信息失败")
	}
	var user models.User
	response.Copy(&user, r)
	user.Salt = util.ToolsUtil.RandomString(5)
	user.Password = util.ToolsUtil.MakeMd5(r.Password + user.Salt)
	if err := a.db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("创建用户失败")
	}
	token := util.ToolsUtil.MakeToken()
	key := fmt.Sprintf("%d", user.ID)
	// 不是多点登录
	if user.IsMultipoint == 0 {
		sysAdminSetKey := types.Admin.BackstageTokenSet + key
		ts := a.cache.SGet(sysAdminSetKey)
		if len(ts) > 0 {
			var keys []string
			for _, t := range ts {
				keys = append(keys, t)
			}
			a.cache.Del(keys...)
		}
		a.cache.Del(sysAdminSetKey)
		a.cache.SSet(sysAdminSetKey, token)
	}
	// 缓存用户信息
	t, _ := time.ParseDuration(a.cfg.Jwt.AccessExpiry)
	a.cache.Set(types.Admin.BackstageTokenKey+token, key, cast.ToInt(t.Seconds()))
	_ = a.userSrv.CacheUserById(user.ID)
	// 返回用户信息
	var userResp resp.UserResp
	response.Copy(&userResp, user)
	return &resp.AccountLoginResp{
		Token:    token,
		UserInfo: userResp,
	}, nil
}

func NewAccountService(db *gorm.DB, cfg *config.Config, cache *cache.Redis, userSrv UserService) AccountService {
	return &authService{db: db, cfg: cfg, cache: cache, userSrv: userSrv}
}
