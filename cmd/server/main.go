package main

import (
	"cleanArchitecture/internal/adapters/primary/rest"
	"cleanArchitecture/internal/adapters/secondary/in_memory"
	"cleanArchitecture/internal/application/interactors"
	"cleanArchitecture/internal/infrastructure/http"
	"cleanArchitecture/internal/infrastructure/logging"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func main() {
	app := fx.New(
		// Logger Deps
		fx.Provide(
			logging.LoggerConfig,
			logging.Logger,
		),

		// Configure logger for uber fx
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{
				Logger: logger,
			}
		}),

		// Primary Adapters
		fx.Provide(
			http.NewHttpServer,
		),

		// User
		fx.Provide(
			fx.Annotate(
				in_memory.NewUserCounterRepository,
				fx.As(new(interactors.UserCounterRepository)),
			),
			interactors.NewUserInteractor,
		),

		// EntryPoint
		fx.Invoke(
			rest.RegisterUserHandlers,
			http.RunHttpServer,
		),
	)

	app.Run()
}
