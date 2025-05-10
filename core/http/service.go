package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/core/config"
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
	//eng.Use(middleware.Cors()).Use(logging.GinLogging(), logging.GinRecovery(true))
	// 设置静态资源
	eng.StaticFS("/static", http.Dir("./public/webroot/static"))
	//engine.StaticFS("/resource", http.Dir("./webroot/resource"))
	eng.StaticFile("/favicon.ico", "./public/webroot/favicon.ico")
	eng.GET("/", func(c *gin.Context) {
		c.File("./public/webroot/index.html")
	})
	eng.NoRoute(func(c *gin.Context) {
		c.File("./public/webroot/index.html")
	})
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
