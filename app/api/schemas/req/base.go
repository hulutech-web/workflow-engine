package req

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// PageReq 分页请求参数
type PageReq struct {
	Page  int `form:"page,default=1" validate:"omitempty,gte=1"`         // 页码
	Limit int `form:"limit,default=20" validate:"omitempty,gt=0,lte=60"` // 每页大小
}

type IdReq struct {
	ID uint `form:"id" validate:"required" json:"id"` // 主键ID
}

type IdListReq struct {
	Ids []string `form:"ids" validate:"required,dive" json:"ids"` // 主键ID列表
}

type KeyReq struct {
	Key string `form:"key" validate:"required" json:"key"` // 关键字
}

type AuthReq struct {
	UserId        uint `json:"user_id"`
	TenantId      uint `json:"tenant_id"`
	IsSuperTenant bool `json:"is_super_tenant"`
	IsAdmin       bool `json:"is_admin"`
}

func GetAuth(c *gin.Context) (*AuthReq, error) {
	auth, exists := c.Get("auth")
	if !exists {
		return nil, fmt.Errorf("获取认证信息失败")
	}
	return auth.(*AuthReq), nil
}
