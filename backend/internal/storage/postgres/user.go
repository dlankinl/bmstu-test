package postgres

import (
	"context"
	"fmt"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/internal/storage"
	"strings"

	"github.com/google/uuid"
)

type UserRepository struct {
	db storage.DBConn
}

func NewUserRepository(db storage.DBConn) domain.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (user *domain.User, err error) {
	query := `select id, username, full_name, birthday, gender, city, role from ppo.users where username = $1`

	tmp := new(User)
	err = r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&tmp.ID,
		&tmp.Username,
		&tmp.FullName,
		&tmp.Birthday,
		&tmp.Gender,
		&tmp.City,
		&tmp.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("получение пользователя по username: %w", err)
	}

	return UserDbToUser(tmp), nil
}

func (r *UserRepository) GetById(ctx context.Context, userId uuid.UUID) (user *domain.User, err error) {
	query := `select username, full_name, birthday, gender, city, role from ppo.users where id = $1`

	tmp := new(User)
	err = r.db.QueryRow(
		ctx,
		query,
		userId,
	).Scan(
		&tmp.Username,
		&tmp.FullName,
		&tmp.Birthday,
		&tmp.Gender,
		&tmp.City,
		&tmp.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("получение пользователя по id: %w", err)
	}

	tmp.ID = userId
	return UserDbToUser(tmp), nil
}

func (r *UserRepository) GetAll(ctx context.Context, page int) (users []*domain.User, numPages int, err error) {
	query := `select 
    	id,
    	username,
    	full_name,
    	birthday,
    	gender,
    	city 
	from ppo.users
	where role = 'user'
	offset $1
	limit $2`

	rows, err := r.db.Query(
		ctx,
		query,
		(page-1)*config.PageSize,
		config.PageSize,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("получение предпринимателей: %w", err)
	}

	users = make([]*domain.User, 0)
	for rows.Next() {
		tmp := new(User)

		err = rows.Scan(
			&tmp.ID,
			&tmp.Username,
			&tmp.FullName,
			&tmp.Birthday,
			&tmp.Gender,
			&tmp.City,
		)

		if err != nil {
			return nil, 0, fmt.Errorf("сканирование полученных строк: %w", err)
		}
		users = append(users, UserDbToUser(tmp))
	}

	var numRecords int
	err = r.db.QueryRow(
		ctx,
		`select count(*) from ppo.users`,
	).Scan(&numRecords)
	if err != nil {
		return nil, 0, fmt.Errorf("получение количества предпринимателей: %w", err)
	}

	numPages = numRecords / config.PageSize
	if numRecords%config.PageSize != 0 {
		numPages++
	}

	return users, numPages, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) (err error) {
	query := `update ppo.users set `

	args := make([]any, 0)
	i := 1
	equals := make([]string, 0)
	if user.FullName != "" {
		equals = append(equals, fmt.Sprintf("full_name = $%d", i))
		i++
		args = append(args, user.FullName)
	}
	if !user.Birthday.IsZero() {
		equals = append(equals, fmt.Sprintf("birthday = $%d", i))
		i++
		args = append(args, user.Birthday)
	}
	if user.Gender != "" {
		equals = append(equals, fmt.Sprintf("gender = $%d", i))
		i++
		args = append(args, user.Gender)
	}
	if user.City != "" {
		equals = append(equals, fmt.Sprintf("city = $%d", i))
		i++
		args = append(args, user.City)
	}
	if user.Role != "" {
		equals = append(equals, fmt.Sprintf("role = $%d", i))
		i++
		args = append(args, user.Role)
	}
	if user.Username != "" {
		equals = append(equals, fmt.Sprintf("username = $%d", i))
		i++
		args = append(args, user.Username)
	}
	query += strings.Join(equals, ", ")
	query += fmt.Sprintf(" where id = $%d", i)
	args = append(args, user.ID)

	_, err = r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о пользователе: %w", err)
	}

	return nil
}

func (r *UserRepository) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	query := `delete from ppo.users where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("удаление пользователя по id: %w", err)
	}

	return nil
}
