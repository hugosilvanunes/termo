package config

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	DB *sqlx.DB
}

func NewDB() (*Repo, error) {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	return &Repo{DB: db}, nil
}

func (r Repo) RunMigrations() (sql.Result, error) {
	return r.DB.
		Exec("CREATE TABLE IF NOT EXISTS word(id INTEGER PRIMARY KEY, name VARCHAR(20) NOT NULL);")
}

func (r Repo) Create(word string) (sql.Result, error) {
	return r.DB.Exec("INSERT INTO word (name) VALUES (?)", word)
}

func (r Repo) FindOne(query string, args ...any) *sqlx.Row {
	if len(args) == 0 {
		return r.DB.QueryRowx(query)
	}

	return r.DB.QueryRowx(query, args)
}
