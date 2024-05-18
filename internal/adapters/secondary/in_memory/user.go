package in_memory

import "cleanArchitecture/internal/domain"

type UserRepository struct {
	users []*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*domain.User, 0),
	}
}

func (r *UserRepository) Create(domain *domain.User) error {
	r.users = append(r.users, domain)
	return nil
}

func (r *UserRepository) GetAll() ([]*domain.User, error) {
	return r.users, nil
}
