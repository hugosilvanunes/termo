package oneword

import (
	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/termo/config"
)

func Router(r *gin.Engine, cfg config.Config) {
	r.GET("/new_game", NewGame(cfg))
	r.POST("/attempt", Attempt(cfg))
}
