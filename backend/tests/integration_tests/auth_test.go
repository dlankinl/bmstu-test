package integration_tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/domain"
	"ppo/internal/storage/postgres"
	"ppo/internal/utils"
)

type StorageAuthSuite struct {
	suite.Suite
	repo domain.IAuthRepository
}

func (s *StorageAuthSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = postgres.NewAuthRepository(TestDbInstance)
	t.Tags("fixture", "auth")
}

func (s *StorageAuthSuite) Test_AuthStorage_Register(t provider.T) {
	t.Title("[Register] Успех")
	t.Tags("storage", "postgres", "auth")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		authInfo := utils.NewUserAuthBuilder().
			WithUsername("test123").
			WithHashedPass("test123").
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", authInfo)
		err := s.repo.Register(ctx, &authInfo)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageAuthSuite) Test_AuthStorage_Register2(t provider.T) {
	t.Title("[Register] Неуникальное имя пользователя")
	t.Tags("storage", "postgres", "auth")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		authInfo := utils.NewUserAuthBuilder().
			WithUsername("user1").
			WithHashedPass("test123").
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", authInfo)
		err := s.repo.Register(ctx, &authInfo)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf(`регистрация пользователя: ERROR: duplicate key value violates unique constraint "users_username_key" (SQLSTATE 23505)`).Error(), err.Error())
	})
}

func (s *StorageAuthSuite) Test_AuthStorage_GetByUsername(t provider.T) {
	t.Title("[GetByUsername] Успех")
	t.Tags("storage", "postgres", "auth")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id, _ := uuid.Parse("8a7fe516-600e-4c01-a55e-423fac892250")
		authInfo := utils.NewUserAuthBuilder().
			WithHashedPass("user3hehe").
			WithID(id).
			Build()
		username := "user3"

		sCtx.WithNewParameters("ctx", ctx, "model", authInfo)
		user, err := s.repo.GetByUsername(ctx, username)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&authInfo, user)
	})
}

func (s *StorageAuthSuite) Test_AuthStorage_GetByUsername2(t provider.T) {
	t.Title("[GetByUsername] Не найден")
	t.Tags("storage", "postgres", "auth")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		username := "undefined"

		sCtx.WithNewParameters("ctx", ctx, "model", username)
		_, err := s.repo.GetByUsername(ctx, username)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение пользователя по username: %w", pgx.ErrNoRows).Error(), err.Error())
	})
}
