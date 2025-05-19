package api

import (
	ae "github.com/hulutech-web/workflow-engine/app/api/event"
	aq "github.com/hulutech-web/workflow-engine/app/api/queue"
	"github.com/hulutech-web/workflow-engine/app/api/route"
	"github.com/hulutech-web/workflow-engine/app/api/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	service.Module,
	route.Module,
	ae.Module,
	aq.Module,
)
