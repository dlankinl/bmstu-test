package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"ppo/domain"
	"ppo/internal/utils"
	"ppo/mocks"
)

type StorageCompanySuite struct {
	suite.Suite
	repo *mocks.MockICompanyRepository
	ctrl *gomock.Controller
}

func (s *StorageCompanySuite) BeforeAll(t provider.T) {
	s.ctrl = gomock.NewController(t)
	s.repo = mocks.NewMockICompanyRepository(s.ctrl)
}

func (s *StorageCompanySuite) AfterAll(t provider.T) {
	s.ctrl.Finish()
}

func (s *StorageCompanySuite) Test_CompanyStorageCreate(t provider.T) {
	t.Title("[CompanyCreate] Успех")
	t.Tags("storage", "company", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		model := utils.CompanyMother{}.Default()
		//actFieldModel := utils.ActivityFieldMother{}.Default()
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

func (s *StorageCompanySuite) Test_CompanyStorageDeleteById(t provider.T) {
	t.Title("[CompanyDeleteById] Успешно")
	t.Tags("storage", "company", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{0}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(
				ctx,
				id,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := s.repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageDeleteById2(t provider.T) {
	t.Title("[CompanyDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("storage", "company", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{0}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(
				ctx,
				id,
			).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := s.repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
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

		ctx := context.TODO()

		s.repo.EXPECT().
			GetAll(
				ctx,
				0,
			).
			Return(expectedCompanies, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		companies, err := s.repo.GetAll(ctx, 0)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedCompanies, companies)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetAll2(t provider.T) {
	t.Title("[CompanyGetAll] Ошибка получения данных в репозитории")
	t.Tags("storage", "company", "getAll")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		s.repo.EXPECT().
			GetAll(ctx, 2).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 2)

		_, err := s.repo.GetAll(ctx, 2)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetById(t provider.T) {
	t.Title("[CompanyGetById] Успешно")
	t.Tags("storage", "company", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.UUID{4}
		compModel := utils.CompanyMother{}.Default()

		s.repo.EXPECT().
			GetById(
				ctx,
				id,
			).
			Return(&compModel, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		comp, err := s.repo.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&compModel, comp)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetById2(t provider.T) {
	t.Title("[CompanyGetById] Ошибка получения данных в репозитории")
	t.Tags("storage", "company", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.UUID{7}

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

func (s *StorageCompanySuite) Test_CompanyStorageGetByOwnerId(t provider.T) {
	t.Title("[CompanyGetByOwnerId] Успешно")
	t.Tags("storage", "company", "getByOwnerId")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
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

		s.repo.EXPECT().
			GetByOwnerId(
				ctx,
				ownerId,
				5,
				true,
			).
			Return(expectedCompanies, 0, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", 5)

		companies, _, err := s.repo.GetByOwnerId(ctx, ownerId, 5, true)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedCompanies, companies)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageGetByOwnerId2(t provider.T) {
	t.Title("[CompanyGetByOwnerId] Ошибка получения данных в репозитории")
	t.Tags("storage", "company", "getByOwnerId")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		ownerId := uuid.UUID{9}

		s.repo.EXPECT().
			GetByOwnerId(
				ctx,
				ownerId,
				7,
				true,
			).
			Return(nil, 0, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 7)

		_, _, err := s.repo.GetByOwnerId(ctx, ownerId, 7, true)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorageUpdate(t provider.T) {
	t.Title("[CompanyUpdate] Успешно")
	t.Tags("storage", "company", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		updatedInfoCompany := utils.NewCompanyBuilder().
			WithID(uuid.UUID{1}).
			WithName("aaa").
			Build()

		s.repo.EXPECT().
			Update(
				ctx,
				&updatedInfoCompany,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		err := s.repo.Update(ctx, &updatedInfoCompany)

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
			WithID(uuid.UUID{1}).
			WithName("aaa").
			Build()

		s.repo.EXPECT().
			Update(
				ctx,
				&updatedInfoCompany,
			).Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		err := s.repo.Update(ctx, &updatedInfoCompany)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}
