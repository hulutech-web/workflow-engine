package workflow

import (
	"go.uber.org/fx"
)

var Module = fx.Module("engin",
	fx.Provide(NewEngin),
)
