package logging

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

const (
	prodMode = "prod"
	devMode  = "dev"
)

type Config struct {
	MODE string `json:"mode"`
}

func LoggerConfig() (*Config, error) {
	var cfg struct {
		Config Config `json:"logger"`
	}
	err := cleanenv.ReadConfig("config.json", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg.Config, nil
}

func Logger(cfg *Config) *slog.Logger {
	switch cfg.MODE {
	case prodMode:
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	case devMode:
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	default:
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
}
