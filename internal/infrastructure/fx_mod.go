package infrastructure

import (
	"cleanArchitecture/internal/infrastructure/grpc"
	"cleanArchitecture/internal/infrastructure/http"
	"cleanArchitecture/internal/infrastructure/logging"
	"cleanArchitecture/internal/infrastructure/redis"
	"cleanArchitecture/internal/infrastructure/validate"
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Module(
	"Infrastructure",

	// Validator
	fx.Provide(
		validate.NewValidate,
	),

	// Logger
	fx.Provide(
		logging.LoadConfig,
		logging.Logger,
	),

	// Redis
	fx.Provide(
		redis.LoadConfig,
		redis.NewClient,
	),

	// Grpc
	fx.Provide(
		grpc.LoadConfig,
		grpc.AsUnaryServerInterceptor(grpc.NewLoggingInterceptor),
		fx.Annotate(
			grpc.NewGrpcServer,
			fx.ParamTags("", grpc.UnaryServerInterceptorGroup),
		),
	),

	// Http
	fx.Provide(
		http.LoadConfig,
		http.NewGatewayServer,
		http.NewHttpServer,
	),

	// Module Entrypoint
	fx.Invoke(
		grpc.RunGrpcServer,
		http.RunHttpServer,
		func(logger *slog.Logger) {
			logger.Info("Infrastructure Initialized")
		},
	),
)
