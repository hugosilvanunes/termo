package oneword

import "github.com/hugosilvanunes/termo/config"

const FindOneSQL = "SELECT * FROM word LIMIT 1"

func FindOne(repo *config.Repo) (*Word, error) {
	word := new(Word)
	err := repo.FindOne(FindOneSQL).StructScan(word)
	return word, err
}
