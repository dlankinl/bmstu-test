package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"ppo/domain"
	"ppo/pkg/logger"
)

type Service struct {
	userRepo     domain.IUserRepository
	companyRepo  domain.ICompanyRepository
	actFieldRepo domain.IActivityFieldRepository
	logger       logger.ILogger
}

func NewService(
	userRepo domain.IUserRepository,
	companyRepo domain.ICompanyRepository,
	actFieldRepo domain.IActivityFieldRepository,
	logger logger.ILogger,
) domain.IUserService {
	return &Service{
		userRepo:     userRepo,
		companyRepo:  companyRepo,
		actFieldRepo: actFieldRepo,
		logger:       logger,
	}
}

func (s *Service) GetByUsername(ctx context.Context, username string) (user *domain.User, err error) {
	prompt := "UserGetByUsername"

	user, err = s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		s.logger.Infof("%s: получение пользователя по username: %v", prompt, err)
		return nil, fmt.Errorf("получение пользователя по username: %w", err)
	}

	return user, nil
}

func (s *Service) GetById(ctx context.Context, userId uuid.UUID) (user *domain.User, err error) {
	prompt := "UserGetById"

	user, err = s.userRepo.GetById(ctx, userId)
	if err != nil {
		s.logger.Infof("%s: получение пользователя по id: %v", prompt, err)
		return nil, fmt.Errorf("получение пользователя по id: %w", err)
	}

	return user, nil
}

func (s *Service) GetAll(ctx context.Context, page int) (users []*domain.User, numPages int, err error) {
	prompt := "UserGetAll"

	users, numPages, err = s.userRepo.GetAll(ctx, page)
	if err != nil {
		s.logger.Infof("%s: получение списка всех пользователей: %v", prompt, err)
		return nil, 0, fmt.Errorf("получение списка всех пользователей: %w", err)
	}

	return users, numPages, nil
}

func (s *Service) Update(ctx context.Context, user *domain.User) (err error) {
	prompt := "UserUpdate"

	if user.Gender != "" && user.Gender != "m" && user.Gender != "w" {
		s.logger.Infof("%s: неизвестный пол", prompt)
		return fmt.Errorf("неизвестный пол")
	}

	if user.Role != "" && user.Role != "admin" && user.Role != "user" {
		s.logger.Infof("%s: невалидная роль", prompt)
		return fmt.Errorf("невалидная роль")
	}

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		s.logger.Infof("%s: обновление информации о пользователе: %v", prompt, err)
		return fmt.Errorf("обновление информации о пользователе: %w", err)
	}

	return nil
}

func (s *Service) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	prompt := "UserDeleteById"

	err = s.userRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Infof("%s: удаление пользователя по id: %v", prompt, err)
		return fmt.Errorf("удаление пользователя по id: %w", err)
	}

	return nil
}
