package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	_, err := ctx.Cookie("session_token")
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.SetCookieData(&http.Cookie{
		Name:     "session_token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	ctx.Status(http.StatusOK)
}
