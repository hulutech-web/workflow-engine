package service

import "go.uber.org/fx"

var Module = fx.Module("api.service",
	fx.Provide(NewUserService),
	fx.Provide(NewAccountService),
)
