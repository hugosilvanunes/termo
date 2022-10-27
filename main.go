package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/hugosilvanunes/termo/config"
	"go.uber.org/zap"
)

type Word struct {
	ID   int
	Name string
}

type DicioResponse struct {
	ID        int    `json:"id,omitempty"`
	Word      string `json:"word,omitempty"`
	Count     int    `json:"count,omitempty"`
	Character string `json:"character,omitempty"`
}

type AttemptRequest struct {
	Word string `json:"word,omitempty" binding:"required"`
}

type AttemptCharsResponse struct {
	Index  int    `json:"index,omitempty"`
	Char   string `json:"character,omitempty"`
	Status string `json:"status,omitempty"`
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("cannot initialize config")
	}
	defer cfg.Log.Sync()

	if err := config.RunMigrations(cfg.DB); err != nil {
		cfg.Log.Fatal("cannot run migrations", zap.Error(err))
	}

	client := resty.New()

	resp, err := client.R().EnableTrace().SetResult(&DicioResponse{}).Get(cfg.Env.DicioURL)
	if err != nil {
		cfg.Log.Fatal("cannnot request to dicio api", zap.Error(err))
	}

	resBody, ok := resp.Result().(*DicioResponse)
	if !ok {
		cfg.Log.Fatal("cannot parse response body")
	}

	if _, err := cfg.DB.Exec("INSERT INTO word (name) VALUES (?)", resBody.Word); err != nil {
		cfg.Log.Fatal("cannot insert seed word", zap.Error(err))
	}

	r := gin.Default()
	r.GET("/new_game", func(c *gin.Context) {
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
	})

	r.POST("/attempt", func(c *gin.Context) {
		attempt := new(AttemptRequest)
		if err := c.ShouldBindJSON(attempt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		wordAttempt := strings.TrimSpace(attempt.Word)
		if wordAttempt == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		word := new(Word)
		if err := cfg.DB.QueryRowx("SELECT * FROM word LIMIT 1").StructScan(word); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(wordAttempt) != len(word.Name) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

	})

	if err := r.Run(); err != nil {
		cfg.Log.Fatal("error in run web server", zap.Error(err))
	}
}
