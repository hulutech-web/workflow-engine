package service

import "go.uber.org/fx"

var Module = fx.Module("service",
	fx.Provide(NewAccountService),
	fx.Provide(NewPaginatorService),
	fx.Provide(NewDeptService),
	fx.Provide(NewTemplateService),
	fx.Provide(NewUserService),
	fx.Provide(NewTemplateFormService),
	fx.Provide(NewEmpService),
)
