package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"ppo/internal/utils"
)

type StorageActFieldSuite struct {
	suite.Suite
}

func (s *StorageActFieldSuite) Test_ActFieldStorageCreate(t provider.T) {
	t.Title("[ActFieldCreate] Успех")
	t.Tags("storage", "actField", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		model := utils.ActivityFieldMother{}.Default()
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("insert").WithArgs(model.Name, model.Description, model.Cost).WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(model.ID))

		repo := NewActivityFieldRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		res, err := repo.Create(ctx, &model)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&model, res)
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageCreate2(t provider.T) {
	t.Title("[ActFieldCreate] Ошибка в репозитории")
	t.Tags("storage", "actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		model := utils.NewActivityFieldBuilder().WithName("test").Build()
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("insert").WithArgs(model.Name, model.Description, model.Cost).WillReturnError(fmt.Errorf("sql error"))

		repo := NewActivityFieldRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err = repo.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("создание сферы деятельности: sql error").Error(), err.Error())
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageDeleteById(t provider.T) {
	t.Title("[ActFieldDeleteById] Успех")
	t.Tags("storage", "actField", "Success")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{0}
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("delete").WithArgs(id).WillReturnResult(pgxmock.NewResult("delete", 1))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		repo := NewActivityFieldRepository(mock)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func (s *StorageActFieldSuite) Test_ActFieldStorageDeleteById2(t provider.T) {
	t.Title("[ActFieldDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("actField", "Fail")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		id := uuid.UUID{1}
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("delete").WithArgs(id).WillReturnError(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		repo := NewActivityFieldRepository(mock)
		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление сферы деятельности по id: sql error").Error(), err.Error())
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(id).WillReturnRows(pgxmock.
			NewRows([]string{"name", "description", "cost"}).AddRow(returnedModel.Name, returnedModel.Description, returnedModel.Cost))

		repo := NewActivityFieldRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		model, err := repo.GetById(ctx, id)

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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(id).WillReturnError(fmt.Errorf("sql error"))

		repo := NewActivityFieldRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		_, err = repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение сферы деятельности по id: sql error").Error(), err.Error())
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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("update").WithArgs(model.Name, model.ID).WillReturnResult(pgxmock.NewResult("update", 1))

		repo := NewActivityFieldRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err = repo.Update(ctx, &model)

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

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("update").WithArgs(model.Name, model.ID).WillReturnError(fmt.Errorf("sql error"))

		repo := NewActivityFieldRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err = repo.Update(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление информации о сфере деятельности: sql error").Error(), err.Error())
	})
}
