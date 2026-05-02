package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	v1 "github.com/unixtensor/monolith/pkg/api/v1"
)

type Config struct {
	Port  string
	Token string
	Debug bool
	Redis *redis.Client
}

func V1(bg_ctx context.Context, c Config) {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	api_v1 := v1.V1{
		Token: c.Token,
		Redis: c.Redis,
	}
	api_v1.V1(bg_ctx, c.Port)
}
