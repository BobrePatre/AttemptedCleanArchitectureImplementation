package repository

import (
	"cleanArchitecture/internal/user/entity"
)

type UserRepository struct {
	users []*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*entity.User, 0),
	}
}

func (r *UserRepository) Create(domain *entity.User) error {
	r.users = append(r.users, domain)
	return nil
}

func (r *UserRepository) GetAll() ([]*entity.User, error) {
	return r.users, nil
}
