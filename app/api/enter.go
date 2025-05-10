package api

import (
	"github.com/hulutech-web/workflow-engine/app/api/route"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/workflow"
	"go.uber.org/fx"
)

var Module = fx.Options(
	service.Module,
	route.Module,
	workflow.Module,
)
