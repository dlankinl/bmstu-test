package domain

import (
	"context"

	"github.com/google/uuid"
)

type ActivityField struct {
	ID          uuid.UUID
	Name        string
	Description string
	Cost        float32
}

//go:generate mockgen -source=activity_field.go -destination=../mocks/activity_field.go -package=mocks
type IActivityFieldRepository interface {
	Create(context.Context, *ActivityField) (*ActivityField, error)
	DeleteById(context.Context, uuid.UUID) error
	Update(context.Context, *ActivityField) error
	GetById(context.Context, uuid.UUID) (*ActivityField, error)
	GetMaxCost(context.Context) (float32, error)
	GetAll(context.Context, int, bool) ([]*ActivityField, int, error)
}

type IActivityFieldService interface {
	Create(context.Context, *ActivityField) error
	DeleteById(context.Context, uuid.UUID) error
	Update(context.Context, *ActivityField) error
	GetById(context.Context, uuid.UUID) (*ActivityField, error)
	GetMaxCost(context.Context) (float32, error)
	GetAll(context.Context, int, bool) ([]*ActivityField, int, error)
}
