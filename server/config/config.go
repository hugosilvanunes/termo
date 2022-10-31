package config

import (
	"log"

	"go.uber.org/zap"
)

type Config struct {
	Log  *zap.Logger
	Env  EnvVars
	Repo *Repo
}

func New() (*Config, error) {
	envVars, err := NewEnv()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	logger, err := NewLogger(envVars.ENV)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	db, err := NewDB()
	if err != nil {
		logger.Error("cannot connect to database", zap.Error(err))
		return nil, err
	}

	if _, err := db.RunMigrations(); err != nil {
		logger.Error("cannot run migrations", zap.Error(err))
		return nil, err
	}

	return &Config{
		Env:  *envVars,
		Repo: db,
		Log:  logger,
	}, nil
}
