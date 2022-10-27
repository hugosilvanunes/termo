package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Config struct {
	Log *zap.Logger
	Env EnvVars
	DB  *sqlx.DB
}

func New() (*Config, error) {
	envVars, err := NewEnv()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	logger, err := NewLogger()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	db, err := NewDB()
	if err != nil {
		logger.Error("cannot connect to database", zap.Error(err))
		return nil, err
	}

	return &Config{
		Env: *envVars,
		DB:  db,
		Log: logger,
	}, nil
}
