package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"ppo/domain"
	"ppo/internal/services/user"
	"ppo/internal/utils"
	"ppo/mocks"
	"time"
)

type StorageUserSuite struct {
	suite.Suite
	repo *mocks.MockIUserRepository
	ctrl *gomock.Controller
}

func (s *StorageUserSuite) BeforeAll(t provider.T) {
	s.ctrl = gomock.NewController(t)
	s.repo = mocks.NewMockIUserRepository(s.ctrl)
}

func (s *StorageUserSuite) AfterAll(t provider.T) {
	s.ctrl.Finish()
}

func (s *StorageUserSuite) Test_UserStorageDeleteById(t provider.T) {
	t.Title("[UserDeleteById] Успех")
	t.Tags("storage", "user", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		uID := uuid.UUID{1}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(ctx, uID).
			Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", uID)

		err := s.repo.DeleteById(ctx, uID)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageUserSuite) Test_UserStorageDeleteById2(t provider.T) {
	t.Title("[UserDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("storage", "user", "deleteById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		uID := uuid.UUID{2}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(ctx, uID).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", uID)

		err := s.repo.DeleteById(ctx, uID)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageUserSuite) Test_UserStorageGetAll(t provider.T) {
	t.Title("[UserGetAll] Успех")
	t.Tags("storage", "user", "getAll")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		user1 := utils.NewUserBuilder().
			WithId(uuid.UUID{1}).
			WithUsername("a").
			WithFullName("a").
			WithGender("m").
			WithBirthday(time.Date(1, 1, 1, 1, 1, 1, 1, time.Local)).
			WithCity("a").
			Build()
		user2 := utils.NewUserBuilder().
			WithId(uuid.UUID{2}).
			WithUsername("b").
			WithFullName("b").
			WithGender("w").
			WithBirthday(time.Date(2, 2, 2, 2, 2, 2, 2, time.Local)).
			WithCity("b").
			Build()
		user3 := utils.NewUserBuilder().
			WithId(uuid.UUID{3}).
			WithUsername("c").
			WithFullName("c").
			WithGender("m").
			WithBirthday(time.Date(3, 3, 3, 3, 3, 3, 3, time.Local)).
			WithCity("c").
			Build()

		users := []*domain.User{&user1, &user2, &user3}

		s.repo.EXPECT().
			GetAll(ctx, 1).
			Return(users, 1, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", users)

		got, _, err := s.repo.GetAll(ctx, 1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(users, got)
	})
}

func (s *StorageUserSuite) Test_UserStorageGetAll2(t provider.T) {
	t.Title("[UserGetAll] Ошибка получения данных в репозитории")
	t.Tags("storage", "user", "getAll")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		s.repo.EXPECT().
			GetAll(ctx, 1).
			Return(nil, 0, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		_, _, err := s.repo.GetAll(ctx, 1)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageUserSuite) Test_UserGetById(t provider.T) {
	t.Title("[UserGetById] Успех")
	t.Tags("user", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uRepo := mocks.NewMockIUserRepository(ctrl)
		cRepo := mocks.NewMockICompanyRepository(ctrl)
		aRepo := mocks.NewMockIActivityFieldRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := user.NewService(uRepo, cRepo, aRepo, log)

		log.EXPECT().
			Infof(gomock.Any()).
			AnyTimes()
		log.EXPECT().
			Infof(gomock.Any(), gomock.Any()).
			AnyTimes()
		log.EXPECT().
			Warnf(gomock.Any(), gomock.Any()).
			AnyTimes()
		log.EXPECT().
			Errorf(gomock.Any(), gomock.Any()).
			AnyTimes()

		ctx := context.TODO()
		uId := uuid.UUID{1}
		model := utils.NewUserBuilder().
			WithId(uId).
			WithUsername("a").
			WithFullName("a b c").
			WithGender("m").
			WithBirthday(time.Date(1, 1, 1, 1, 1, 1, 1, time.Local)).
			WithCity("a").
			Build()

		uRepo.EXPECT().
			GetById(
				ctx,
				uuid.UUID{1},
			).
			Return(&model, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		user, err := svc.GetById(ctx, uId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&model, user)
	})
}

func (s *StorageUserSuite) Test_UserStorageGetById2(t provider.T) {
	t.Title("[UserGetById] Ошибка получения данных в репозитории")
	t.Tags("storage", "user", "getById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		uId := uuid.UUID{4}

		s.repo.EXPECT().
			GetById(
				ctx,
				uId,
			).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", uId)

		_, err := s.repo.GetById(ctx, uId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageUserSuite) Test_UserStorageUpdate(t provider.T) {
	t.Title("[UserUpdate] Успех")
	t.Tags("storage", "user", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		uId := uuid.UUID{1}
		model := utils.NewUserBuilder().
			WithId(uId).
			WithCity("a").
			WithRole("admin").
			Build()

		s.repo.EXPECT().
			Update(
				ctx,
				&model,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := s.repo.Update(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageUserSuite) Test_UserUpdate2(t provider.T) {
	t.Title("[UserUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("user", "update")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		uId := uuid.UUID{10}
		model := utils.NewUserBuilder().
			WithId(uId).
			WithCity("a").
			WithRole("admin").
			Build()

		s.repo.EXPECT().
			Update(
				ctx,
				&model,
			).Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := s.repo.Update(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}
