package game

import "github.com/gin-gonic/gin"

var Game ConnectedGame
var GameInstance = gin.H{}

type ConnectedGame struct {
	CreatorId int    `json:"CreatorId"`
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
}

func Connected() bool {
	return Game.Id != 0
}
