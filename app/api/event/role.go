package event

import (
	"context"
	"github.com/hulutech-web/workflow-engine/core/event"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RoleService 这里是单例模式的实现.
// 发布事件在service/auth_role.go中实现
type RoleService struct {
	e *event.Service
}

// NewRoleService 创建RoleService实例
func NewRoleService(e *event.Service) *RoleService {
	return &RoleService{e: e}
}

// RoleCreated 处理角色创建事件
func (u *RoleService) RoleCreated(e event.Event) {
	zap.S().Info("处理角色创建事件", "event", e.Name, "data", e.Data)
	// 在这里添加业务逻辑：发送邮件、初始化角色资料等
	return
}

// RegisterRoleEvent 注册角色事件
func RegisterRoleEvent(lc fx.Lifecycle, srv *RoleService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			srv.e.Subscribe("role.created", srv.RoleCreated)
			return nil
		},
	})

}
