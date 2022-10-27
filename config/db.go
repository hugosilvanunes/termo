package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", ":memory:")
}

func RunMigrations(db *sqlx.DB) error {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS word(id INTEGER PRIMARY KEY, name VARCHAR(20) NOT NULL);"); err != nil {
		return err
	}

	return nil
}
