//go:build unit

package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"ppo/domain"
	"ppo/internal/utils"
)

type StorageCompanySuite struct {
	suite.Suite
}

func (s *StorageCompanySuite) Test_CompanyStorageCreate(t provider.T) {
	t.Title("[CompanyCreate] Успех")
	t.Tags("storage", "company", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		model := utils.CompanyMother{}.Default()
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("insert").WithArgs(model.OwnerID, model.ActivityFieldId, model.Name, model.City).
			WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(model.ID))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		res, err := repo.Create(ctx, &model)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&model, res)
	})
}

func (s *StorageCompanySuite) Test_CompanyCreate2(t provider.T) {
	t.Title("[CompanyCreate] Пустое название компании")
	t.Tags("company", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		model := utils.NewCompanyBuilder().
			WithCity("ccc").
			WithActivityField(uuid.UUID{0}).
			WithOwner(uuid.UUID{0}).
			Build()
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("insert").WithArgs(model.OwnerID, model.ActivityFieldId, model.Name, model.City).
			WillReturnError(fmt.Errorf("sql error"))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err = repo.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("создание компании: sql error").Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageDeleteById(t provider.T) {
	t.Title("[CompanyDeleteById] Успешно")
	t.Tags("storage", "company", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{3}
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectBegin()
		mock.ExpectExec("delete").WithArgs(id).WillReturnResult(pgxmock.NewResult("delete", 1))
		mock.ExpectExec("delete").WithArgs(id).WillReturnResult(pgxmock.NewResult("delete", 1))
		mock.ExpectCommit()

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageDeleteById2(t provider.T) {
	t.Title("[CompanyDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("storage", "company", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{1}
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectBegin()
		mock.ExpectExec("delete").WithArgs(id).WillReturnError(fmt.Errorf("sql error"))
		mock.ExpectRollback()

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление компании по id: sql error").Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetAll(t provider.T) {
	t.Title("[CompanyGetAll] Успешно")
	t.Tags("storage", "company", "getAll")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		company1 := utils.NewCompanyBuilder().
			WithName("a").
			WithCity("a").
			WithID(uuid.UUID{1}).
			Build()
		company2 := utils.NewCompanyBuilder().
			WithName("b").
			WithCity("b").
			WithID(uuid.UUID{2}).
			Build()
		company3 := utils.NewCompanyBuilder().
			WithName("c").
			WithCity("c").
			WithID(uuid.UUID{3}).
			Build()
		expectedCompanies := []*domain.Company{&company1, &company2, &company3}
		page := 1

		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(page-1, 3).
			WillReturnRows(pgxmock.NewRows([]string{"id", "owner_id", "activity_field_id", "name", "city"}).
				AddRow(expectedCompanies[0].ID, expectedCompanies[0].OwnerID, expectedCompanies[0].ActivityFieldId, expectedCompanies[0].Name, expectedCompanies[0].City).
				AddRow(expectedCompanies[1].ID, expectedCompanies[1].OwnerID, expectedCompanies[1].ActivityFieldId, expectedCompanies[1].Name, expectedCompanies[1].City).
				AddRow(expectedCompanies[2].ID, expectedCompanies[2].OwnerID, expectedCompanies[2].ActivityFieldId, expectedCompanies[2].Name, expectedCompanies[2].City),
			)

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", page)

		res, err := repo.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedCompanies, res)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetAll2(t provider.T) {
	t.Title("[CompanyGetAll] Ошибка получения данных в репозитории")
	t.Tags("storage", "company", "getAll")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(page-1, 3).
			WillReturnError(fmt.Errorf("sql error"))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", page)

		_, err = repo.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение списка компаний: sql error").Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetById(t provider.T) {
	t.Title("[CompanyGetById] Успешно")
	t.Tags("storage", "company", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.UUID{4}
		compModel := utils.CompanyMother{}.WithID(id)

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(id).WillReturnRows(pgxmock.
			NewRows([]string{"owner_id", "activity_field_id", "name", "city"}).
			AddRow(compModel.OwnerID, compModel.ActivityFieldId, compModel.Name, compModel.City))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", compModel)

		model, err := repo.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(compModel, *model)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetById2(t provider.T) {
	t.Title("[CompanyGetById] Ошибка получения данных в репозитории")
	t.Tags("storage", "company", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.UUID{7}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(id).WillReturnError(fmt.Errorf("sql error"))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		_, err = repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение компании по id: sql error").Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetByOwnerId(t provider.T) {
	t.Title("[CompanyGetByOwnerId] Успешно")
	t.Tags("storage", "company", "getByOwnerId")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1
		ownerId := uuid.UUID{5}
		company1 := utils.NewCompanyBuilder().
			WithName("a").
			WithCity("a").
			WithID(uuid.UUID{6}).
			WithOwner(ownerId).
			Build()
		company2 := utils.NewCompanyBuilder().
			WithName("b").
			WithCity("b").
			WithID(uuid.UUID{7}).
			WithOwner(ownerId).
			Build()
		company3 := utils.NewCompanyBuilder().
			WithName("c").
			WithCity("c").
			WithID(uuid.UUID{8}).
			WithOwner(ownerId).
			Build()
		expectedCompanies := []*domain.Company{&company1, &company2, &company3}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(ownerId, page-1, 3).
			WillReturnRows(pgxmock.NewRows([]string{"id", "activity_field_id", "name", "city"}).
				AddRow(expectedCompanies[0].ID, expectedCompanies[0].ActivityFieldId, expectedCompanies[0].Name, expectedCompanies[0].City).
				AddRow(expectedCompanies[1].ID, expectedCompanies[1].ActivityFieldId, expectedCompanies[1].Name, expectedCompanies[1].City).
				AddRow(expectedCompanies[2].ID, expectedCompanies[2].ActivityFieldId, expectedCompanies[2].Name, expectedCompanies[2].City),
			)

		mock.ExpectQuery("select").WithArgs(ownerId).WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(3))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", ownerId)

		res, _, err := repo.GetByOwnerId(ctx, ownerId, page, true)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedCompanies, res)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetByOwnerId2(t provider.T) {
	t.Title("[CompanyGetByOwnerId] Ошибка получения данных в репозитории")
	t.Tags("storage", "company", "getByOwnerId")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ownerId := uuid.UUID{9}
		page := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(ownerId, page-1, 3).
			WillReturnError(fmt.Errorf("sql error"))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", page)

		_, _, err = repo.GetByOwnerId(ctx, ownerId, page, true)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение компаний: sql error").Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageUpdate(t provider.T) {
	t.Title("[CompanyUpdate] Успешно")
	t.Tags("storage", "company", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		updatedInfoCompany := utils.NewCompanyBuilder().
			WithID(uuid.UUID{13}).
			WithName("aaa").
			Build()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("update").WithArgs(updatedInfoCompany.Name, updatedInfoCompany.ID).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", updatedInfoCompany)

		err = repo.Update(ctx, &updatedInfoCompany)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageUpdate2(t provider.T) {
	t.Title("[CompanyUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("storage", "company", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		updatedInfoCompany := utils.NewCompanyBuilder().
			WithID(uuid.UUID{11}).
			WithName("aaa").
			Build()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("update").WithArgs(updatedInfoCompany.Name, updatedInfoCompany.ID).
			WillReturnError(fmt.Errorf("sql error"))

		repo := NewCompanyRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", updatedInfoCompany)

		err = repo.Update(ctx, &updatedInfoCompany)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление информации о компании: sql error").Error(), err.Error())
	})
}
