package game

import "github.com/gin-gonic/gin"

type Game struct {
	Metadata GameMetadata
	Instance gin.H
}

type GameMetadata struct {
	CreatorId int    `json:"CreatorId"`
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
}

var CurrentGame Game

func Connected() bool {
	return CurrentGame.Metadata.Id != 0
}
