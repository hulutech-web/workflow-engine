package service

import "go.uber.org/fx"

var Module = fx.Module("service",
	fx.Provide(NewAccountService),
	fx.Provide(NewAuthPermService),
	fx.Provide(NewAuthMenuService),
	fx.Provide(NewAuthTenantService),
	fx.Provide(NewAuthRoleService),
	fx.Provide(NewPaginatorService),
	fx.Provide(NewDeptService),
	fx.Provide(NewTemplateService),
	fx.Provide(NewUserService),
	fx.Provide(NewTemplateFormService),
	fx.Provide(NewEmpService),
	fx.Provide(NewFlowService),
	fx.Provide(NewProcessService),
	fx.Provide(NewFlowlinkService),
	fx.Provide(NewFlowTypeService),
)
