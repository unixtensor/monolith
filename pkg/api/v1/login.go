package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, token string) {
	var s string
	if err := json.NewDecoder(ctx.Request.Body).Decode(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if s != token {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	ctx.SetCookieData(&http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	ctx.Status(http.StatusOK)
}
