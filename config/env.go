package config

import "github.com/kelseyhightower/envconfig"

type EnvVars struct {
	Port     int    `default:"8080"`
	DicioURL string `required:"true"`
}

func NewEnv() (*EnvVars, error) {
	env := new(EnvVars)
	err := envconfig.Process("", env)

	return env, err
}
