package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func scan_keys(redis *redis.Client, bg_ctx context.Context, filter func(string) bool) ([]string, error) {
	var cursor uint64
	var keys []string = []string{}

	batch, cursor, err := redis.Scan(bg_ctx, cursor, "*", 0).Result()
	for _, k := range batch {
		if filter(k) {
			keys = append(keys, k)
		}
	}
	return keys, err
}

func (v1 *V1) games(bg_ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keys, err := scan_keys(v1.Redis, bg_ctx, func(k string) bool {
			return !strings.HasSuffix(k, "-jobs")
		})
		if err != nil {
			InternalError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, keys)
	}
}

func (v1 *V1) servers(bg_ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keys, err := scan_keys(v1.Redis, bg_ctx, func(k string) bool {
			return strings.HasSuffix(k, "-jobs")
		})
		if err != nil {
			InternalError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, keys)
	}
}
