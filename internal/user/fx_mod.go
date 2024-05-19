package user

import (
	"cleanArchitecture/internal/user/delivery"
	"cleanArchitecture/internal/user/repository"
	"cleanArchitecture/internal/user/service/interactors"
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Module(
	"User Domain",
	fx.Provide(
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(interactors.UserRepository)),
		),
		interactors.NewUserInteractor,
	),

	fx.Invoke(
		delivery.RegisterUserHandlers,
		func(logger *slog.Logger) {
			logger.Info("User Domain Connected")
		},
	),
)
