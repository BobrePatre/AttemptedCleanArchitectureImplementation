package interactors

import (
	"cleanArchitecture/internal/user/entity"
	"cleanArchitecture/internal/user/service/dto"
	"log/slog"
)

type UserRepository interface {
	Create(example *entity.User) error
	GetAll() ([]*entity.User, error)
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
	example := entity.NewExample(dto.Name, dto.Age)
	err := interactor.repo.Create(example)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *UserInteractor) GetAll() ([]*entity.User, error) {
	return interactor.repo.GetAll()
}
