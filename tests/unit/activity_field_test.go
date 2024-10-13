//go:build unit

package unit

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"ppo/internal/services/activity_field"
	"ppo/internal/utils"
	"ppo/mocks"
)

type ActFieldSuite struct {
	suite.Suite
}

func (s *ActFieldSuite) Test_ActFieldCreate(t provider.T) {
	t.Title("[ActFieldCreate] Успех")
	t.Tags("actField", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		model := utils.ActivityFieldMother{}.Default()
		ctx := context.TODO()

		repo.EXPECT().
			Create(
				ctx,
				&model,
			).Return(&model, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Create(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *ActFieldSuite) Test_ActFieldCreate2(t provider.T) {
	t.Title("[ActFieldCreate] Пустое название сферы деятельности")
	t.Tags("actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		model := utils.ActivityFieldMother{}.WithoutName()
		ctx := context.TODO()

		repo.EXPECT().
			Create(
				ctx,
				&model,
			).Return(&model, nil).
			AnyTimes()

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("должно быть указано название сферы деятельности").Error(), err.Error())
	})
}

func (s *ActFieldSuite) Test_ActFieldCreate3(t provider.T) {
	t.Title("[ActFieldCreate] Ошибка в репозитории")
	t.Tags("actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		model := utils.ActivityFieldMother{}.Default()
		ctx := context.TODO()

		repo.EXPECT().
			Create(
				ctx,
				&model,
			).Return(nil, fmt.Errorf("sql error")).
			AnyTimes()

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("создание сферы деятельности: sql error").Error(), err.Error())
	})
}

func (s *ActFieldSuite) Test_ActFieldDeleteById(t provider.T) {
	t.Title("[ActFieldDeleteById] Успех")
	t.Tags("actField", "Success")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		id := uuid.UUID{0}
		ctx := context.TODO()

		repo.EXPECT().
			DeleteById(ctx, id).
			Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := svc.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *ActFieldSuite) Test_ActFieldDeleteById2(t provider.T) {
	t.Title("[ActFieldDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		id := uuid.UUID{0}
		ctx := context.TODO()

		repo.EXPECT().
			DeleteById(ctx, id).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := svc.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление сферы деятельности по id: sql error").Error(), err.Error())
	})
}

func (s *ActFieldSuite) Test_ActFieldGetById(t provider.T) {
	t.Title("[ActFieldGetById] Успех")
	t.Tags("actField", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		id := uuid.UUID{0}
		ctx := context.TODO()
		returnedModel := utils.NewActivityFieldBuilder().
			WithName("a").
			WithDescription("a").
			WithID(id).
			Build()

		repo.EXPECT().
			GetById(
				ctx,
				id,
			).
			Return(&returnedModel, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		model, err := svc.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(returnedModel, *model)
	})
}

func (s *ActFieldSuite) Test_ActFieldGetById2(t provider.T) {
	t.Title("[ActFieldGetById] Ошибка получения данных в репозитории")
	t.Tags("actField", "getById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		id := uuid.UUID{0}
		ctx := context.TODO()

		repo.EXPECT().
			GetById(
				ctx,
				id,
			).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		_, err := svc.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение сферы деятельности по id: sql error").Error(), err.Error())
	})
}

func (s *ActFieldSuite) Test_ActFieldUpdate(t provider.T) {
	t.Title("[ActFieldUpdate] Успех")
	t.Tags("actField", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		id := uuid.UUID{0}
		ctx := context.TODO()
		model := utils.NewActivityFieldBuilder().
			WithName("aaa").
			WithID(id).
			Build()

		repo.EXPECT().
			Update(
				ctx,
				&model,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Update(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *ActFieldSuite) Test_ActFieldUpdate2(t provider.T) {
	t.Title("[ActFieldUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("actField", "update")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := activity_field.NewService(repo, compRepo, log)

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

		id := uuid.UUID{0}
		ctx := context.TODO()
		model := utils.NewActivityFieldBuilder().
			WithName("aaa").
			WithID(id).
			Build()

		repo.EXPECT().
			Update(
				ctx,
				&model,
			).Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Update(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление информации о cфере деятельности: sql error").Error(), err.Error())
	})
}
