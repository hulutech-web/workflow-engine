package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginContext struct {
	*gin.Context
}

func (c *ginContext) Bind(obj interface{}) error {
	return c.Context.ShouldBind(obj)
}

func (c *ginContext) Get(key string) (interface{}, bool) {
	return c.Context.Get(key)
}

func (c *ginContext) Set(key string, value interface{}) {
	c.Context.Set(key, value)
}

type ginRouter struct {
	engine *gin.Engine
	group  *gin.RouterGroup
}

func NewGinRouter() Router {
	return &ginRouter{engine: gin.Default()}
}

func (r *ginRouter) GET(path string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.handle(http.MethodGet, path, handler, middleware...)
}

func (r *ginRouter) POST(path string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.handle(http.MethodPost, path, handler, middleware...)
}

func (r *ginRouter) PUT(path string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.handle(http.MethodPut, path, handler, middleware...)
}

func (r *ginRouter) DELETE(path string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.handle(http.MethodDelete, path, handler, middleware...)
}

func (r *ginRouter) handle(method, path string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	h := applyMiddleware(handler, middleware...)

	ginHandler := func(c *gin.Context) {
		ctx := &ginContext{Context: c}
		if err := h(ctx); err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}
	}

	if r.group != nil {
		r.group.Handle(method, path, ginHandler)
	} else {
		r.engine.Handle(method, path, ginHandler)
	}
}

func (r *ginRouter) Group(prefix string, middleware ...MiddlewareFunc) Router {
	var ginGroup *gin.RouterGroup
	if r.group != nil {
		ginGroup = r.group.Group(prefix)
	} else {
		ginGroup = r.engine.Group(prefix)
	}

	return &ginRouter{
		engine: r.engine,
		group:  ginGroup,
	}
}

func (r *ginRouter) Use(middleware ...MiddlewareFunc) {
	ginMiddleware := make([]gin.HandlerFunc, len(middleware))
	for i, m := range middleware {
		m := m
		ginMiddleware[i] = func(c *gin.Context) {
			ctx := &ginContext{Context: c}
			_ = m(func(ctx Context) error { return nil })(ctx)
		}
	}

	if r.group != nil {
		r.group.Use(ginMiddleware...)
	} else {
		r.engine.Use(ginMiddleware...)
	}
}

func applyMiddleware(handler HandlerFunc, middleware ...MiddlewareFunc) HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}
