package event

import "go.uber.org/fx"

var (
	Module = fx.Module("event",
		fx.Provide(NewRoleService),
		fx.Invoke(RegisterRoleEvent),
	)
)
