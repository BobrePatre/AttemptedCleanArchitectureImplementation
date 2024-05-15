package domain

import (
	"context"
	"github.com/google/uuid"
)

type Example struct {
	Id     uuid.UUID
	Field1 string
	Field2 int
}

type ExampleRepository interface {
	CreateExample(ctx context.Context, example *Example) error
	GetExampleById(ctx context.Context, id uuid.UUID) (*Example, error)
}
