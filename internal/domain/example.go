package domain

import (
	"github.com/google/uuid"
)

type Example struct {
	Id     uuid.UUID
	Field1 string
	Field2 int
}

func NewExample(field1 string, field2 int) *Example {
	return &Example{
		Id:     uuid.New(),
		Field1: field1,
		Field2: field2,
	}
}

func (domain Example) Rename(field1 string) {
	domain.Field1 = field1
}

func (domain Example) Birthday() {
	domain.Field2 = domain.Field2 + 1
}
