package main

import (
	"github.com/hulutech-web/workflow-engine/boot"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	// 替你提交测试一下自动贡献者在Readme中展示
	fx.New(
		boot.Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	).Run()
}
