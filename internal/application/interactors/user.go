package interactors

import (
	"cleanArchitecture/internal/application/dto"
	"cleanArchitecture/internal/domain"
)

type NotifiyGateway interface {
	Notify(domain *domain.User) error
}

type UserRepository interface {
	Save(example *domain.User) error
	GetById(id string) (*domain.User, error)
}

type UserInteractor struct {
	repo     UserRepository
	notifier NotifiyGateway
	counter  int
}

func NewUserInteractor() *UserInteractor {
	return &UserInteractor{}
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
	interactor.counter++
	return interactor.counter
}
