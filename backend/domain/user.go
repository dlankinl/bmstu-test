package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	FullName string
	Gender   string
	Birthday time.Time
	City     string
	Role     string
}

//go:generate mockgen -source=user.go -destination=../mocks/user.go -package=mocks
type IUserRepository interface {
	GetByUsername(context.Context, string) (*User, error)
	GetById(context.Context, uuid.UUID) (*User, error)
	GetAll(context.Context, int) ([]*User, int, error)
	Update(context.Context, *User) error
	DeleteById(context.Context, uuid.UUID) error
}

type IUserService interface {
	GetByUsername(context.Context, string) (*User, error)
	GetById(context.Context, uuid.UUID) (*User, error)
	GetAll(context.Context, int) ([]*User, int, error)
	Update(context.Context, *User) error
	DeleteById(context.Context, uuid.UUID) error
}
