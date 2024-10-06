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
	"time"
)

type StorageUserSuite struct {
	suite.Suite
	repo domain.IUserRepository
}

func (s *StorageUserSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = postgres.NewUserRepository(TestDbInstance)
	t.Tags("fixture", "finReport")
}

func (s *StorageUserSuite) Test_UserStorage_GetById(t provider.T) {
	t.Title("[GetById] Успех")
	t.Tags("storage", "postgres", "user")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		userId, _ := uuid.Parse("bc3ab9bf-6a26-4212-941d-05a985fc0978")
		expUser := utils.NewUserBuilder().
			WithId(userId).
			WithUsername("user1").
			WithFullName("User First 1").
			WithBirthday(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).
			WithGender("m").
			WithCity("Moscow").
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", userId)

		user, err := s.repo.GetById(ctx, userId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&expUser, user)
	})
}

func (s *StorageUserSuite) Test_UserStorage_GetById2(t provider.T) {
	t.Title("[GetById] Несуществующий пользователь")
	t.Tags("storage", "postgres", "user")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		userId, _ := uuid.Parse("bc3ab9bf-6a28-4212-941d-05a985fc0978")

		sCtx.WithNewParameters("ctx", ctx, "model", userId)

		_, err := s.repo.GetById(ctx, userId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение пользователя по id: %w", pgx.ErrNoRows).Error(), err.Error())
	})
}

func (s *StorageUserSuite) Test_UserStorage_Update(t provider.T) {
	t.Title("[Update] Успех")
	t.Tags("storage", "postgres", "user")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		userId, _ := uuid.Parse("b384ea3b-df18-4bae-b459-fb96e2518fe7")
		user := utils.NewUserBuilder().
			WithId(userId).
			WithGender("m").
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", userId)

		err := s.repo.Update(ctx, &user)

		getUser, getErr := s.repo.GetById(ctx, userId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NoError(getErr)
		sCtx.Assert().Equal(user.Gender, getUser.Gender)
	})
}

func (s *StorageUserSuite) Test_UserStorage_DeleteById(t provider.T) {
	t.Title("[DeleteById] Успех")
	t.Tags("storage", "postgres", "user")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		userId, _ := uuid.Parse("e129e5bd-0568-4568-a562-209b3161800e")

		sCtx.WithNewParameters("ctx", ctx, "model", userId)

		err := s.repo.DeleteById(ctx, userId)

		_, getErr := s.repo.GetById(ctx, userId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Error(getErr)
		sCtx.Assert().Equal(fmt.Errorf("получение пользователя по id: %w", pgx.ErrNoRows).Error(), getErr.Error())
	})
}

func (s *StorageUserSuite) Test_UserStorage_GetByUsername(t provider.T) {
	t.Title("[GetByUsername] Успех")
	t.Tags("storage", "postgres", "user")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		username := "user3"
		userId, _ := uuid.Parse("8a7fe516-600e-4c01-a55e-423fac892250")
		expUser := utils.NewUserBuilder().
			WithId(userId).
			WithUsername(username).
			WithCity("Moscow").
			WithBirthday(time.Date(2002, 3, 3, 0, 0, 0, 0, time.UTC)).
			WithGender("m").
			WithFullName("User Third 3").
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", username)

		user, err := s.repo.GetByUsername(ctx, username)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&expUser, user)
	})
}
