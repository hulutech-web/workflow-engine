package route

import (
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/core/http"
	"go.uber.org/fx"
)

var Module = fx.Module("api.route",
	fx.Provide(NewRoutes),
	fx.Invoke(accountRoutes),
	fx.Invoke(userRoutes),
)

type Routes struct {
	fx.In
	Http *http.Service
}

func NewRoutes(deps Routes) *types.ApiRouter {
	return &types.ApiRouter{
		Engine: deps.Http.Gin,
	}
}
