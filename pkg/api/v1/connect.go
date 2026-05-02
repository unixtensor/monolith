package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Game struct {
	Name       string `json:"Name" binding:"required"`
	CreatorId  uint   `json:"CreatorId" binding:"required"`
	MaxPlayers uint   `json:"MaxPlayers" binding:"required"`
	JobID      string
}
type Connector struct {
	Game  Game
	Redis *redis.Client
}

func (c *Connector) add_game_db(bg_ctx context.Context, place_id string) error {
	if game_set_err := c.Redis.JSONSet(bg_ctx, place_id, "$", c.Game).Err(); game_set_err != nil {
		return game_set_err
	}
	if set_jobs_err := c.Redis.LPush(bg_ctx, place_id+"-jobs", c.Game.JobID).Err(); set_jobs_err != nil {
		return set_jobs_err
	}
	return nil
}

func (v1 *V1) connect(bg_ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		place_id := ctx.Param("placeId")
		job_id := ctx.Param("jobId")

		var connector = Connector{
			Game:  Game{JobID: job_id},
			Redis: v1.Redis,
		}
		if json_bind_err := ctx.ShouldBindJSON(&connector.Game); json_bind_err != nil {
			InternalError(ctx, json_bind_err)
			return
		}
		if db_added := connector.add_game_db(bg_ctx, place_id); db_added != nil {
			InternalError(ctx, db_added)
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (v1 *V1) connected(bg_ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		place_id := ctx.Param("placeId")

		exists, err := v1.Redis.Exists(bg_ctx, place_id).Result()
		if err != nil {
			InternalError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, exists != 0)
	}
}
