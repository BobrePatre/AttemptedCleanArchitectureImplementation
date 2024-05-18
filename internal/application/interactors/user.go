package interactors

import (
	"cleanArchitecture/internal/application/dto"
	"cleanArchitecture/internal/domain"
	"log/slog"
)

type NotifiyGateway interface {
	Notify(domain *domain.User) error
}

type UserRepository interface {
	Save(example *domain.User) error
	GetById(id string) (*domain.User, error)
}

type UserCounterRepository interface {
	Increment() int
}

type UserInteractor struct {
	repo     UserRepository
	notifier NotifiyGateway
	counter  UserCounterRepository
	logger   *slog.Logger
}

func NewUserInteractor(logger *slog.Logger, counterRepo UserCounterRepository) *UserInteractor {
	return &UserInteractor{
		counter: counterRepo,
		logger:  logger,
	}
}

func (interactor *UserInteractor) CreateExample(dto dto.CreateUserRq) error {
	example := domain.NewExample(dto.Name, dto.Age)
	err := interactor.repo.Save(example)
	if err != nil {
		return err
	}
	err = interactor.notifier.Notify(example)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *UserInteractor) Incr() int {
	interactor.logger.Info("Incr called")
	return interactor.counter.Increment()
}
