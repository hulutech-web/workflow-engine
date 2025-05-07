package main

import (
	"github.com/hulutech-web/workflow-engine/auth"
	"github.com/hulutech-web/workflow-engine/core"
	"github.com/hulutech-web/workflow-engine/database"
	"github.com/hulutech-web/workflow-engine/http"
	"github.com/hulutech-web/workflow-engine/http/middleware"
	"github.com/hulutech-web/workflow-engine/logging"
	"github.com/hulutech-web/workflow-engine/validation"
	"github.com/hulutech-web/workflow-engine/workflow"
	"time"
)

func main() {
	// 初始化应用
	app := core.NewApplication()

	// 注册服务提供者
	app.Register(&database.GormProvider{
		DSN: "user:password@tcp(127.0.0.1:3306)/workflow?charset=utf8mb4&parseTime=True&loc=Local",
	})

	app.Register(&auth.AuthProvider{
		JwtSecret:     "your-secret-key",
		AccessExpiry:  time.Hour,
		RefreshExpiry: time.Hour * 24 * 7,
		TokenRotation: true,
	})

	app.Register(&logging.LoggingProvider{
		Production: false,
	})

	app.Register(&validation.ValidationProvider{})

	app.Register(&workflow.WorkflowProvider{})

	// 启动应用
	app.Boot()

	// 初始化路由
	router := http.NewGinRouter()

	// 公共路由
	router.POST("/login", auth.LoginHandler)
	router.POST("/refresh", auth.RefreshHandler)

	// 需要认证的路由
	authGroup := router.Group("/api", middleware.Auth(app.GetContainer()))

	// 工作流路由
	workflowHandlers := workflow.NewHandlers(app.GetContainer().Make("workflow.engine").(*workflow.Engine))
	authGroup.POST("/process", workflowHandlers.StartProcess)
	authGroup.POST("/approve", workflowHandlers.ApproveStep)

	// 启动服务器
	router.Run(":8080")
}
