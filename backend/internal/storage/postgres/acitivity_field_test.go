package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"ppo/internal/utils"
	"ppo/mocks"
)

type StorageActFieldSuite struct {
	suite.Suite
	repo *mocks.MockIActivityFieldRepository
	ctrl *gomock.Controller
}

func (s *StorageActFieldSuite) BeforeAll(t provider.T) {
	s.ctrl = gomock.NewController(t)
	s.repo = mocks.NewMockIActivityFieldRepository(s.ctrl)
}

func (s *StorageActFieldSuite) AfterAll(t provider.T) {
	s.ctrl.Finish()
}

func (s *StorageActFieldSuite) Test_ActFieldStorageCreate(t provider.T) {
	t.Title("[ActFieldCreate] Успех")
	t.Tags("storage", "actField", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		model := utils.ActivityFieldMother{}.Default()
		ctx := context.TODO()

		s.repo.EXPECT().
			Create(
				ctx,
				&model,
			).Return(&model, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		res, err := s.repo.Create(ctx, &model)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&model, res)
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageCreate2(t provider.T) {
	t.Title("[ActFieldCreate] Пустое название сферы деятельности")
	t.Tags("storage", "actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		model := utils.ActivityFieldMother{}.WithoutName()
		ctx := context.TODO()

		s.repo.EXPECT().
			Create(
				ctx,
				&model,
			).Return(&model, nil).
			AnyTimes()

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		res, err := s.repo.Create(ctx, &model)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&model, res)
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageCreate3(t provider.T) {
	t.Title("[ActFieldCreate] Ошибка в репозитории")
	t.Tags("storage", "actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		model := utils.NewActivityFieldBuilder().WithName("test").Build()
		ctx := context.TODO()

		s.repo.EXPECT().
			Create(
				ctx,
				&model,
			).Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err := s.repo.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageDeleteById(t provider.T) {
	t.Title("[ActFieldDeleteById] Успех")
	t.Tags("storage", "actField", "Success")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{0}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(ctx, id).
			Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := s.repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageDeleteById2(t provider.T) {
	t.Title("[ActFieldDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		id := uuid.UUID{1}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(ctx, id).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := s.repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageGetById(t provider.T) {
	t.Title("[ActFieldGetById] Успех")
	t.Tags("storage", "actField", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{2}
		ctx := context.TODO()
		returnedModel := utils.NewActivityFieldBuilder().
			WithName("a").
			WithDescription("a").
			WithID(id).
			Build()

		s.repo.EXPECT().
			GetById(
				ctx,
				id,
			).
			Return(&returnedModel, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		model, err := s.repo.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(returnedModel, *model)
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageGetById2(t provider.T) {
	t.Title("[ActFieldGetById] Ошибка получения данных в репозитории")
	t.Tags("storage", "actField", "getById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		id := uuid.UUID{3}
		ctx := context.TODO()

		s.repo.EXPECT().
			GetById(
				ctx,
				id,
			).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		_, err := s.repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageUpdate(t provider.T) {
	t.Title("[ActFieldUpdate] Успех")
	t.Tags("actField", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{0}
		ctx := context.TODO()
		model := utils.NewActivityFieldBuilder().
			WithName("aaa").
			WithID(id).
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

func (s *StorageActFieldSuite) Test_ActFieldStorageUpdate2(t provider.T) {
	t.Title("[ActFieldUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("actField", "update")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		id := uuid.UUID{1}
		ctx := context.TODO()
		model := utils.NewActivityFieldBuilder().
			WithName("aaa").
			WithID(id).
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
