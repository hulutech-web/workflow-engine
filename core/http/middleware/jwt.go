package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/hulutech-web/workflow-engine/pkg/util"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			response.Fail(ctx, response.TokenEmpty)
			ctx.Abort()
			return
		}

		claims, err := util.JwtUtil.ParseToken(token)
		if err != nil {
			response.Fail(ctx, response.TokenInvalid)
			ctx.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			response.Fail(ctx, response.TokenExpired)
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
