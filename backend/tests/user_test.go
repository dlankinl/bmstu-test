package tests

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"go.uber.org/mock/gomock"
	"ppo/domain"
	"ppo/internal/services/user"
	"ppo/internal/storage/postgres"
	"ppo/internal/utils"
	"ppo/mocks"
	"time"
)

type UserSuite struct {
	suite.Suite
}

func (s *UserSuite) Test_UserDeleteById(t provider.T) {
	t.Title("[UserDeleteById] Успех")
	t.Tags("user", "deleteById")
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

		uID := uuid.UUID{1}
		ctx := context.TODO()

		uRepo.EXPECT().
			DeleteById(ctx, uID).
			Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", uID)

		err := svc.DeleteById(ctx, uID)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserSuite) Test_UserDeleteById2(t provider.T) {
	t.Title("[UserDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("user", "deleteById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
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

		uID := uuid.UUID{1}
		ctx := context.TODO()

		uRepo.EXPECT().
			DeleteById(ctx, uID).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", uID)

		err := svc.DeleteById(ctx, uID)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление пользователя по id: sql error").Error(), err.Error())
	})
}

func (s *UserSuite) Test_UserGetAll(t provider.T) {
	t.Title("[UserGetAll] Успех")
	t.Tags("user", "getAll")
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

		uRepo.EXPECT().
			GetAll(ctx, 1).
			Return(users, 1, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", users)

		got, _, err := svc.GetAll(ctx, 1)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(users, got)
	})
}

func (s *UserSuite) Test_UserGetAll2(t provider.T) {
	t.Title("[UserGetAll] Ошибка получения данных в репозитории")
	t.Tags("user", "getAll")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
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

		uRepo.EXPECT().
			GetAll(ctx, 1).
			Return(nil, 0, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		_, _, err := svc.GetAll(ctx, 1)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение списка всех пользователей: sql error").Error(), err.Error())
	})
}

func (s *UserSuite) Test_UserGetById(t provider.T) {
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

func (s *UserSuite) Test_UserGetById2(t provider.T) {
	t.Title("[UserGetById] Ошибка получения данных в репозитории")
	t.Tags("user", "getById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
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

		uRepo.EXPECT().
			GetById(
				ctx,
				uId,
			).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", uId)

		_, err := svc.GetById(ctx, uId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение пользователя по id: sql error").Error(), err.Error())
	})
}

func (s *UserSuite) Test_UserUpdate(t provider.T) {
	t.Title("[UserUpdate] Успех")
	t.Tags("user", "update")
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
			WithCity("a").
			WithRole("admin").
			Build()

		uRepo.EXPECT().
			Update(
				ctx,
				&model,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Update(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserSuite) Test_UserUpdate2(t provider.T) {
	t.Title("[UserUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("user", "update")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
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
			WithCity("a").
			WithRole("admin").
			Build()

		uRepo.EXPECT().
			Update(
				ctx,
				&model,
			).Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Update(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление информации о пользователе: sql error").Error(), err.Error())
	})
}

func (s *UserSuite) Test_ClassicUserUpdate(t provider.T) {
	t.Title("[ClassicUserUpdate] Успех")
	t.Tags("classic", "user", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		uRepo := postgres.NewUserRepository(mock)
		cRepo := postgres.NewCompanyRepository(mock)
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
			WithCity("a").
			WithRole("admin").
			Build()

		mock.ExpectExec("update").WithArgs(model.City, model.Role, model.ID).
			WillReturnResult(pgxmock.NewResult("update", 1))

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err = svc.Update(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserSuite) Test_ClassicUserUpdate2(t provider.T) {
	t.Title("[ClassicUserUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("classic", "user", "update")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		uRepo := postgres.NewUserRepository(mock)
		cRepo := postgres.NewCompanyRepository(mock)
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
			WithCity("a").
			WithRole("admin").
			Build()

		mock.ExpectExec("update").WithArgs(model.City, model.Role, model.ID).WillReturnError(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err = svc.Update(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление информации о пользователе: обновление информации о пользователе: sql error").Error(), err.Error())
	})
}
