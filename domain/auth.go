package domain

import (
	"context"
	"github.com/google/uuid"
)

type UserAuth struct {
	ID         uuid.UUID
	Username   string
	Password   string
	HashedPass string
	Role       string
}

//go:generate mockgen -source=auth.go -destination=../mocks/auth.go -package=mocks
type IAuthRepository interface {
	Register(context.Context, *UserAuth) (uuid.UUID, error)
	GetByUsername(context.Context, string) (*UserAuth, error)
}

type IAuthService interface {
	Login(context.Context, *UserAuth) (string, error)
	Register(context.Context, *UserAuth) (uuid.UUID, error)
}
