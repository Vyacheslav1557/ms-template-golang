package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

var configInstance *Config

type Config struct {
	ENV  string `env:"ENV" env-required:"true"`
	HOST string `env:"HOST" env-required:"true"`
	PORT int    `env:"PORT" env-required:"true"`
}

func mustLoad() {
	var (
		err error
		cfg Config
	)
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		slog.Error("error reading env: %s", err)
		os.Exit(1)
	}

	configInstance = &cfg
}

func SetupConfig() {
	mustLoad()
}

func Cfg() Config {
	if configInstance == nil {
		slog.Error("config was not initialized")
		os.Exit(1)
	}
	return *configInstance
}
