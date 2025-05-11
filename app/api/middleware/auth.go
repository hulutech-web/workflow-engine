package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/resp"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func AuthCheck(db *gorm.DB, cr *cache.Redis) gin.HandlerFunc {
	permSrv := service.NewAuthPermService(db, cr)
	userSrv := service.NewUserService(db, cr, permSrv)
	admin := types.Admin
	return func(c *gin.Context) {
		auths := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/api/", "", 1), "/", "_")
		// 免登录接口
		if util.ToolsUtil.Contains(admin.NotLoginUri, auths) {
			c.Next()
			return
		}
		// 获取token，从header、query、form中获取
		token := getToken(c)
		if len(token) == 0 {
			response.Fail(c, response.TokenEmpty)
			c.Abort()
			return
		}
		// 验证token
		tk := admin.BackstageTokenKey + token
		existCnt := cr.Exists(tk)
		if existCnt < 0 {
			response.Fail(c, response.SystemError)
			c.Abort()
			return
		} else if existCnt == 0 {
			response.Fail(c, response.TokenInvalid)
			c.Abort()
			return
		}
		// 用户信息缓存
		uidStr := cr.Get(token)
		var uid uint
		if uidStr != "" {
			i, err := strconv.ParseUint(uidStr, 10, 32)
			if err != nil {
				zap.S().Errorf("验证token失败，uidStr[%s]转换失败，err[%+v]", uidStr, err)
				response.Fail(c, response.TokenInvalid)
				c.Abort()
				return
			}
			uid = uint(i)
		}
		// 验证用户是否存在
		if !cr.HExists(admin.BackstageManageKey, uidStr) {
			err := userSrv.CacheUserById(uid)
			if err != nil {
				zap.S().Errorf("验证token失败，用户不存在，uid[%d],err[%+v]", uid, err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
		}
		var mapping resp.UserResp
		err := util.ToolsUtil.JsonToObj(cr.HGet(admin.BackstageManageKey, uidStr), &mapping)
		if err != nil {
			zap.S().Errorf("TokenAuth Unmarshal err: err=[%+v]", err)
			response.Fail(c, response.SystemError)
			c.Abort()
			return
		}
		// 校验用户被禁用
		if mapping.IsDisable == 1 {
			response.Fail(c, response.LoginDisableError)
			c.Abort()
			return
		}
		// 令牌剩余30分钟自动续签
		if cr.TTL(tk) > 1800 {
			cr.Expire(tk, 7200)
		}
		// 单次请求信息保存
		auth := req.AuthReq{
			UserId:        mapping.ID,
			TenantId:      mapping.Tenant.ID,
			RoleId:        mapping.Role.ID,
			IsSuperTenant: mapping.Tenant.ID == 1,
			IsAdmin:       mapping.Role.IsAdmin == 1,
		}
		// 免权限验证接口
		if util.ToolsUtil.Contains(admin.NotAuthUri, auths) || uid == 1 {
			c.Next()
			return
		}
		// 校验角色权限是否存在
		roleId := fmt.Sprintf("%d", mapping.Role.ID)
		if !cr.HExists(admin.BackstageRolesKey, roleId) {
			i, err := strconv.ParseUint(roleId, 10, 32)
			if err != nil {
				zap.S().Errorf("鉴权失败，roleId[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
			err = permSrv.CacheRoleMenusByRoleId(uint(i))
			if err != nil {
				zap.S().Errorf("鉴权失败，[%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}
		}

		// 验证是否有权限操作
		menus := cr.HGet(admin.BackstageRolesKey, roleId)
		if !(menus != "" && util.ToolsUtil.Contains(strings.Split(menus, ","), auths)) {
			response.Fail(c, response.NoPermission)
			c.Abort()
			return
		}

		c.Set("auth", &auth)
		c.Next()
	}
}

func getToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	if token == "" {
		token = c.PostForm("token")
	}
	return token
}
