package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/internal/storage"
)

type AuthRepository struct {
	db storage.DBConn
}

func NewAuthRepository(db storage.DBConn) domain.IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(ctx context.Context, authInfo *domain.UserAuth) (err error) {
	query := `insert into ppo.users (username, password, role) values ($1, $2, 'user')`

	_, err = r.db.Exec(
		ctx,
		query,
		authInfo.Username,
		authInfo.HashedPass,
	)
	if err != nil {
		return fmt.Errorf("регистрация пользователя: %w", err)
	}

	return nil
}

func (r *AuthRepository) GetByUsername(ctx context.Context, username string) (data *domain.UserAuth, err error) {
	query := `select id, password, role from ppo.users where username = $1`

	var id uuid.UUID
	var hashedPass, role string
	tmp := new(UserAuth)
	err = r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&id,
		&hashedPass,
		&role,
	)
	if err != nil {
		return nil, fmt.Errorf("получение пользователя по username: %w", err)
	}

	tmp.ID = id
	tmp.HashedPass.String = hashedPass
	tmp.HashedPass.Valid = true
	tmp.Role.String = role
	tmp.Role.Valid = true

	return UserAuthDbToUserAuth(tmp), nil
}
