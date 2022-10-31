package main

import (
	"log"
	"strings"
	"unicode"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/termo/config"
	"github.com/hugosilvanunes/termo/dicioapi"
	"github.com/hugosilvanunes/termo/oneword"
	"go.uber.org/zap"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("cannot initialize config")
	}
	defer cfg.Log.Sync()

	dicioCli := dicioapi.NewClient(*cfg)

	res, err := dicioCli.GetRandomWord()
	if err != nil {
		cfg.Log.Fatal("cannot get random word")
	}

	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}), norm.NFC)
	result, _, _ := transform.String(t, res.Word)

	if _, err := cfg.Repo.Create(strings.ToLower(result)); err != nil {
		cfg.Log.Fatal("cannot create word seed", zap.Error(err))
	}

	r := gin.Default()
	r.Use(cors.Default())

	oneword.Router(r, *cfg)

	if err := r.Run(); err != nil {
		cfg.Log.Fatal("error in run web server", zap.Error(err))
	}
}
