package infrastructure

import (
	"cleanArchitecture/internal/infrastructure/http"
	"cleanArchitecture/internal/infrastructure/logging"
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Module(
	"Infrastructure",
	fx.Provide(
		logging.LoggerConfig,
		logging.Logger,
		http.NewHttpServer,
	),
	fx.Invoke(
		http.RunHttpServer,
		func(logger *slog.Logger) {
			logger.Info("Infrastructure Initialized")
		},
	),
)
