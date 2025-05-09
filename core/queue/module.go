package queue

import "go.uber.org/fx"

var Module = fx.Provide(NewQueue, NewDelayQueue)
