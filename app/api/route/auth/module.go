package auth

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(menuRoutes),
	fx.Invoke(roleRoutes),
	fx.Invoke(tenantRoutes))
