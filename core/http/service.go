package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/core/config"
	"github.com/hulutech-web/workflow-engine/core/http/middleware"
	"github.com/hulutech-web/workflow-engine/core/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"net/http"
	"time"
)

type Service struct {
	Gin    *gin.Engine
	Server *http.Server
}

func NewService(c *config.Config) *Service {
	gin.SetMode(c.Server.Mode)
	eng := gin.New()
	eng.Use(middleware.Cors())
	eng.Use(logging.GinLogging(), logging.GinRecovery(true))
	// 设置静态资源
	if c.Storage.LocalPath == "" {
		c.Storage.LocalPath = "./public/uploads"
	}
	if c.Storage.PublicPrefix == "" {
		c.Storage.PublicPrefix = "/uploads"
	}

	eng.StaticFS("/static", http.Dir("./public/webroot/static"))
	eng.StaticFile("/favicon.ico", "./public/webroot/favicon.ico")

	eng.GET("/", func(c *gin.Context) {
		c.File("./public/webroot/index.html")
	})
	eng.NoRoute(func(c *gin.Context) {
		c.File("./public/webroot/index.html")
	})
	eng.StaticFS(c.Storage.PublicPrefix, http.Dir(c.Storage.LocalPath))
	// 添加swagger路由
	//eng.GET("/docs/*any", gin.BasicAuth(gin.Accounts{
	//	"admin": "123456",
	//}), ginSwagger.WrapHandler(swaggerFiles.Handler))

	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      eng,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return &Service{
		Gin:    eng,
		Server: server,
	}
}

var Module = fx.Provide(
	NewService,
)
