package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/termo/config"
)

func NewGame(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		word := new(Word)
		if err := cfg.DB.QueryRowx("SELECT * FROM word LIMIT 1").StructScan(word); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"length": len(word.Name),
		})
	}
}
