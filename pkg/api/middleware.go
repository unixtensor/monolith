package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken(token string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token_cookie, _ := ctx.Cookie("session_token")
		if token_cookie == token {
			ctx.Next()
			return
		}
		token_key := ctx.GetHeader("X-API-Key")
		if token_key != token {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}
