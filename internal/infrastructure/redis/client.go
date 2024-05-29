package redis

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net"
	"strconv"
)

type Config struct {
	Host     string `env:"HOST" json:"host" env-default:"0.0.0.0""`
	Port     int    `env:"PORT" json:"port" env-default:"6379"`
	Password string `env:"PASSWORD" json:"password" env-default:"password"`
}

func (cfg *Config) Address() string {
	return net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
}

func LoadConfig(validator *validator.Validate) (*Config, error) {
	var cfg struct {
		Config Config `json:"redis" env-prefix:"REDIS_"`
	}
	err := cleanenv.ReadConfig("config.json", &cfg)

	if err != nil {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, err
		}
	}
	return &cfg.Config, nil
}

func NewClient(cfg *Config, logger *slog.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password,
	})
	res := client.Ping(context.Background())
	err := res.Err()
	if err != nil {
		logger.Error("Failed to connect to redis", "address", cfg.Address(), "error", err)
		return nil, err
	}
	return client, nil
}
