package service

import "go.uber.org/fx"

var Module = fx.Module("service",
	fx.Provide(NewAuthService),
	fx.Provide(NewPaginatorService),
	fx.Provide(NewDeptService),
)
