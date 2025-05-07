package http

type Context interface {
	JSON(code int, obj interface{})
	Bind(obj interface{}) error
	GetHeader(key string) string
	Param(key string) string
	Query(key string) string
	AbortWithStatus(code int)
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HandlerFunc func(Context) error

type Router interface {
	GET(path string, handler HandlerFunc, middleware ...MiddlewareFunc)
	POST(path string, handler HandlerFunc, middleware ...MiddlewareFunc)
	PUT(path string, handler HandlerFunc, middleware ...MiddlewareFunc)
	DELETE(path string, handler HandlerFunc, middleware ...MiddlewareFunc)
	Group(prefix string, middleware ...MiddlewareFunc) Router
	Use(middleware ...MiddlewareFunc)
}

type MiddlewareFunc func(HandlerFunc) HandlerFunc
