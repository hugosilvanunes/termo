package oneword

import (
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hugosilvanunes/termo/config"
)

func NewGame(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		word, err := FindOne(cfg.Repo)
		if err != nil {
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

func Attempt(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
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
				"error": "attempt not found",
			})
			return
		}

		word, err := FindOne(cfg.Repo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(wordAttempt) != len(word.Name) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Length not equal",
			})
			return
		}

		count := make(map[int]int, 0)
		for i, v := range word.Name {
			vs := string(v)
			count[i] = strings.Count(word.Name, vs)
		}

		corrects := make([]AttemptCharsResponse, 0)
		for i, v := range wordAttempt {
			charAt := string(v)
			for j, k := range word.Name {
				charWord := string(k)
				if charWord == charAt {
					if i == j {
						corrects = append(corrects, AttemptCharsResponse{
							Index:  i,
							Char:   charWord,
							Status: CORRECT,
						})
						count[j] = count[j] - 1
						break
					}
				}
			}
		}

		res := make([]AttemptCharsResponse, 0)
		for i, v := range wordAttempt {
			charAt := string(v)
			ok := false
			for j, k := range word.Name {
				charWord := string(k)
				if charWord == charAt {
					if i == j {
						ok = true
						break
					}

					if count[j] != 0 {
						ok = true
						count[j] = count[j] - 1
						res = append(res, AttemptCharsResponse{
							Index:  i,
							Char:   charAt,
							Status: SEMI,
						})
					}
				}
			}

			if ok {
				continue
			}

			res = append(res, AttemptCharsResponse{
				Index:  i,
				Char:   charAt,
				Status: WRONG,
			})
		}

		res = append(res, corrects...)

		sort.SliceStable(res, func(i, j int) bool {
			return res[i].Index < res[j].Index
		})

		c.JSON(http.StatusOK, res)
	}
}
