package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id   uuid.UUID
	Name string
	Age  int32
}

func NewExample(name string, age int32) *User {
	return &User{
		Id:   uuid.New(),
		Name: name,
		Age:  age,
	}
}

func (domain User) Rename(newName string) {
	domain.Name = newName
}

func (domain User) Birthday() {
	domain.Age = domain.Age + 1
}
