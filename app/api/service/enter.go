package service

import (
	"github.com/hulutech-web/workflow-engine/app/api/service/account"
	"github.com/hulutech-web/workflow-engine/app/api/service/auth"
	"github.com/hulutech-web/workflow-engine/app/api/service/common"
	"github.com/hulutech-web/workflow-engine/app/api/service/org"
	"github.com/hulutech-web/workflow-engine/app/api/service/system"
	"github.com/hulutech-web/workflow-engine/app/api/service/user"
	"github.com/hulutech-web/workflow-engine/app/api/service/workflow"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(account.NewAccountService),
	fx.Provide(auth.NewAuthPermService),
	fx.Provide(auth.NewAuthMenuService),
	fx.Provide(auth.NewAuthTenantService),
	fx.Provide(auth.NewAuthRoleService),
	fx.Provide(common.NewPaginatorService),
	fx.Provide(org.NewDeptService),
	fx.Provide(workflow.NewTemplateService),
	fx.Provide(user.NewUserService),
	fx.Provide(workflow.NewTemplateFormService),
	fx.Provide(org.NewEmpService),
	fx.Provide(workflow.NewFlowService),
	fx.Provide(workflow.NewProcessService),
	fx.Provide(workflow.NewFlowlinkService),
	fx.Provide(workflow.NewFlowTypeService),
	fx.Provide(workflow.NewEntryService),
	fx.Provide(workflow.NewProcService),
	fx.Provide(common.NewCaptchaService),
	fx.Provide(system.NewConfigService),
	fx.Provide(common.NewUploadService),
	fx.Provide(system.NewFileService),
	fx.Provide(workflow.NewEngin),
)
