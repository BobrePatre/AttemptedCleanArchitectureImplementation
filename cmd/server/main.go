package main

import (
	"cleanArchitecture/internal/infrastructure/http"
	"cleanArchitecture/internal/infrastructure/logging"
	"cleanArchitecture/internal/user"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func main() {
	app := fx.New(

		// Modules
		user.Module,

		// Infrastructure
		fx.Provide(
			logging.LoggerConfig,
			logging.Logger,
			http.NewHttpServer,
		),

		// Configure logger for uber fx
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{
				Logger: logger,
			}
		}),

		// EntryPoint
		fx.Invoke(
			http.RunHttpServer,
		),
	)

	app.Run()
}
