package interactors

import (
	"cleanArchitecture/internal/application/dto"
	"cleanArchitecture/internal/domain"
	"log/slog"
)

type UserRepository interface {
	Create(example *domain.User) error
	GetAll() ([]*domain.User, error)
}

type UserInteractor struct {
	repo   UserRepository
	logger *slog.Logger
}

func NewUserInteractor(logger *slog.Logger, counterRepo UserRepository) *UserInteractor {
	return &UserInteractor{
		repo:   counterRepo,
		logger: logger,
	}
}

func (interactor *UserInteractor) CreateUser(dto dto.CreateUserRq) error {
	example := domain.NewExample(dto.Name, dto.Age)
	err := interactor.repo.Create(example)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *UserInteractor) GetAll() ([]*domain.User, error) {
	return interactor.repo.GetAll()
}
