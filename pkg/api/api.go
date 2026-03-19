package api

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var game ConnectedGame

type Config struct {
	Port      string
	Secret    string
	Debugging bool
}

type ConnectedGame struct {
	CreatorId int    `json:"CreatorId"`
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
}

func is_connected(ctx *gin.Context) {
	if game.Id != 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No game linked"})
		return
	}
	ctx.Next()
}

func connect(ctx *gin.Context) {
	if game.Id != 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func upload(ctx *gin.Context) {
	gz, err := gzip.NewReader(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decompress"})
		return
	}
	defer gz.Close()

	bytes, err := io.ReadAll(gz)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var payload map[string]any
	if err := json.Unmarshal(bytes, &payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	log.Println(payload)
	ctx.JSON(http.StatusOK, gin.H{"Success": true})
}

func v1(api_root *gin.Engine) {
	api_v1 := api_root.Group("/api/v1")
	api_v1.POST("/connect", connect)

	connected_api := api_v1.Group("/connected")
	connected_api.Use(is_connected)
	connected_api.POST("/upload", upload)
}

func Start(cfg *Config) {
	if !cfg.Debugging {
		gin.SetMode(gin.ReleaseMode)
	}

	api_root := gin.Default()
	v1(api_root)
	log.Fatal(api_root.Run(":" + cfg.Port))
}
