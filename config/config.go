package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Configuration struct {
	Port     string `env:"PORT" envDefault:":5656"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

func NewConfiguration() (*Configuration, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Info("failed to parse .env file, trying to load default env values")
	}

	cfg := &Configuration{}

	err = env.Parse(cfg)
	if err != nil {
		log.Error("failed to parse envs")
		return nil, err
	}

	return cfg, nil
}
