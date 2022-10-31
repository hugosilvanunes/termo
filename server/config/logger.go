package config

import "go.uber.org/zap"

func NewLogger(env string) (*zap.Logger, error) {
	switch env {
	case "production":
		return zap.NewProduction()
	default:
		return zap.NewDevelopment()
	}
}
