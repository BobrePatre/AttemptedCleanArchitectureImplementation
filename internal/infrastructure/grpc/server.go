package grpc

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"strconv"
	"time"
)

var grpcServerTag = slog.String("server", "grpc_server")

type Config struct {
	Host string `json:"host" env-default:"0.0.0.0" env:"HOST"`
	Port int    `json:"port" env-default:"50051" env:"PORT"`
}

func (cfg *Config) Address() string {
	return net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
}

func LoadConfig() (*Config, error) {
	var cfg struct {
		Config Config `json:"grpc" env-prefix:"GRPC_"`
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

func NewGrpcServer(logger *slog.Logger, unaryInterceptors []grpc.UnaryServerInterceptor) *grpc.Server {
	logger.Info("Initializing gRPC server")
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(server)
	return server
}

func RunGrpcServer(lc fx.Lifecycle, srv *grpc.Server, logger *slog.Logger, cfg *Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting server", grpcServerTag, "address", cfg.Address())
			listener, err := net.Listen("tcp", cfg.Address())
			if err != nil {
				logger.Error("cannot start server", "error", err.Error(), grpcServerTag)
				return err
			}
			go func() {
				err := srv.Serve(listener)
				if err != nil {
					logger.Error("cannot start server", "error", err.Error(), grpcServerTag)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("shutting down", grpcServerTag)
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			srv.GracefulStop()
			return nil
		},
	})
}
