package middleware

import (
	"github.com/hulutech-web/workflow-engine/auth"
	"github.com/hulutech-web/workflow-engine/http"
	"net/http"
)

func Auth(auth auth.AuthManager) http.MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(ctx http.Context) error {
			token := ctx.GetHeader("Authorization")
			if token == "" {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return nil
			}

			user, err := auth.Authenticate(token)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return nil
			}

			ctx.Set("user", user)
			return next(ctx)
		}
	}
}

func RBAC(permission string) http.MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(ctx http.Context) error {
			user, ok := ctx.Get("user")
			if !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return nil
			}

			if !user.(auth.User).Can(permission) {
				ctx.AbortWithStatus(http.StatusForbidden)
				return nil
			}

			return next(ctx)
		}
	}
}
