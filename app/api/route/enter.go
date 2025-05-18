package route

import (
	"github.com/hulutech-web/workflow-engine/app/api/middleware"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/core/cache"
	"github.com/hulutech-web/workflow-engine/core/http"
	"github.com/hulutech-web/workflow-engine/pkg/log"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("api.route",
	fx.Provide(NewRoutes),
	fx.Invoke(deptRoutes),
	fx.Invoke(accountRoutes),
	fx.Invoke(userRoutes),
	fx.Invoke(templateRoutes),
	fx.Invoke(templateFormRoutes),
	fx.Invoke(empRoutes),
	fx.Invoke(tenantRoutes),
	fx.Invoke(roleRoutes),
	fx.Invoke(menuRoutes),
	fx.Invoke(flowRoutes),
	fx.Invoke(processRoutes),
	fx.Invoke(flowlinkRoutes),
	fx.Invoke(flowTypeRoutes),
	fx.Invoke(entryRoutes),
)

type Routes struct {
	fx.In
	Http *http.Service
}

func NewRoutes(deps Routes, db *gorm.DB, cache *cache.Redis) *types.ApiRouter {
	return &types.ApiRouter{RouterGroup: deps.Http.Gin.Group("/api", middleware.AuthCheck(db, cache)), LogCollector: log.NewLogCollector(db, 1024)}
}
