package api

import (
	"github.com/hulutech-web/workflow-engine/app/api/event"
	"github.com/hulutech-web/workflow-engine/app/api/queue"
	"github.com/hulutech-web/workflow-engine/app/api/route"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	service.Module,
	route.Module,
	event.Module,
	queue.Module,
)
