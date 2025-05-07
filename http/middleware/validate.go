package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/auth/http"
	"github.com/hulutech-web/workflow-engine/auth/validation"
	"reflect"
)

func Validate(schema interface{}, validator validation.Validator) http.MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(ctx http.Context) error {
			// 创建schema实例
			s := reflect.New(reflect.TypeOf(schema).Elem()).Interface()

			if err := ctx.Bind(s); err != nil {
				return ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			if err := validator.Validate(s); err != nil {
				return ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			ctx.Set("validated", s)
			return next(ctx)
		}
	}
}
