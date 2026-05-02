package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (v1 *V1) disconnect(bg_ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		place_id := ctx.Param("placeId")
		if err := v1.Redis.Del(bg_ctx, place_id, place_id+"-jobs").Err(); err != nil {
			InternalError(ctx, err)
			return
		}
		ctx.Status(http.StatusOK)
	}
}
