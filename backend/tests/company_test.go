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
	"ppo/internal/services/company"
	"ppo/internal/storage/postgres"
	"ppo/internal/utils"
	"ppo/mocks"
)

type CompanySuite struct {
	suite.Suite
}

func (s *CompanySuite) Test_CompanyCreate(t provider.T) {
	t.Title("[CompanyCreate] Успех")
	t.Tags("company", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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

		model := utils.CompanyMother{}.Default()
		actFieldModel := utils.ActivityFieldMother{}.Default()
		ctx := context.TODO()

		compRepo.EXPECT().
			Create(
				ctx,
				&model,
			).Return(&model, nil)

		repo.EXPECT().
			GetById(
				ctx,
				uuid.UUID{0},
			).Return(&actFieldModel, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err := svc.Create(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *CompanySuite) Test_CompanyCreate2(t provider.T) {
	t.Title("[CompanyCreate] Пустое название компании")
	t.Tags("company", "create")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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

		model := utils.NewCompanyBuilder().
			WithCity("ccc").
			WithActivityField(uuid.UUID{0}).
			WithOwner(uuid.UUID{0}).
			Build()
		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err := svc.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("должно быть указано название компании").Error(), err.Error())
	})
}

func (s *CompanySuite) Test_ClassicCompanyCreate(t provider.T) {
	t.Title("[ClassicCompanyCreate] Пустое название компании")
	t.Tags("classic", "company", "create")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		repo := postgres.NewActivityFieldRepository(mock)
		compRepo := postgres.NewCompanyRepository(mock)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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

		model := utils.NewCompanyBuilder().
			WithCity("ccc").
			WithActivityField(uuid.UUID{0}).
			WithOwner(uuid.UUID{0}).
			Build()
		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err = svc.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("должно быть указано название компании").Error(), err.Error())
	})
}

func (s *CompanySuite) Test_CompanyDeleteById(t provider.T) {
	t.Title("[CompanyDeleteById] Успешно")
	t.Tags("company", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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

		compRepo.EXPECT().
			DeleteById(
				ctx,
				id,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := svc.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *CompanySuite) Test_ClassicCompanyDeleteById(t provider.T) {
	t.Title("[ClassicCompanyDeleteById] Успех")
	t.Tags("classic", "company", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		repo := postgres.NewActivityFieldRepository(mock)
		compRepo := postgres.NewCompanyRepository(mock)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		mock.ExpectBegin()
		mock.ExpectExec("delete").WithArgs(id).WillReturnResult(pgxmock.NewResult("delete", 1))
		mock.ExpectExec("delete").WithArgs(id).WillReturnResult(pgxmock.NewResult("delete", 1))
		mock.ExpectCommit()

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err = svc.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *CompanySuite) Test_ClassicCompanyDeleteById2(t provider.T) {
	t.Title("[ClassicCompanyDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("classic", "company", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		repo := postgres.NewActivityFieldRepository(mock)
		compRepo := postgres.NewCompanyRepository(mock)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		mock.ExpectBegin()
		mock.ExpectExec("delete").WithArgs(id).WillReturnError(fmt.Errorf("sql error"))
		mock.ExpectRollback()

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err = svc.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление компании по id: удаление компании по id: sql error").Error(), err.Error())
	})
}

func (s *CompanySuite) Test_CompanyDeleteById2(t provider.T) {
	t.Title("[CompanyDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("company", "deleteById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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

		compRepo.EXPECT().
			DeleteById(
				ctx,
				id,
			).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := svc.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление компании по id: sql error").Error(), err.Error())
	})
}

func (s *CompanySuite) Test_CompanyGetAll(t provider.T) {
	t.Title("[CompanyGetAll] Успешно")
	t.Tags("company", "getAll")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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

		compRepo.EXPECT().
			GetAll(
				ctx,
				0,
			).
			Return(expectedCompanies, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		companies, err := svc.GetAll(ctx, 0)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedCompanies, companies)
	})
}

func (s *CompanySuite) Test_CompanyGetAll2(t provider.T) {
	t.Title("[CompanyGetAll] Ошибка получения данных в репозитории")
	t.Tags("company", "getAll")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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

		compRepo.EXPECT().
			GetAll(ctx, 0).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		_, err := svc.GetAll(ctx, 0)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение списка всех компаний: sql error").Error(), err.Error())
	})
}

func (s *CompanySuite) Test_CompanyGetById(t provider.T) {
	t.Title("[CompanyGetById] Успешно")
	t.Tags("company", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		id := uuid.UUID{0}
		compModel := utils.CompanyMother{}.Default()

		compRepo.EXPECT().
			GetById(
				ctx,
				id,
			).
			Return(&compModel, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		comp, err := svc.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&compModel, comp)
	})
}

func (s *CompanySuite) Test_CompanyGetById2(t provider.T) {
	t.Title("[CompanyGetById] Ошибка получения данных в репозитории")
	t.Tags("company", "getById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		id := uuid.UUID{0}

		compRepo.EXPECT().
			GetById(
				ctx,
				id,
			).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		_, err := svc.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение компании по id: sql error").Error(), err.Error())
	})
}

func (s *CompanySuite) Test_CompanyGetByOwnerId(t provider.T) {
	t.Title("[CompanyGetByOwnerId] Успешно")
	t.Tags("company", "getByOwnerId")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		ownerId := uuid.UUID{0}
		company1 := utils.NewCompanyBuilder().
			WithName("a").
			WithCity("a").
			WithID(uuid.UUID{1}).
			WithOwner(ownerId).
			Build()
		company2 := utils.NewCompanyBuilder().
			WithName("b").
			WithCity("b").
			WithID(uuid.UUID{2}).
			WithOwner(ownerId).
			Build()
		company3 := utils.NewCompanyBuilder().
			WithName("c").
			WithCity("c").
			WithID(uuid.UUID{3}).
			WithOwner(ownerId).
			Build()
		expectedCompanies := []*domain.Company{&company1, &company2, &company3}

		compRepo.EXPECT().
			GetByOwnerId(
				ctx,
				ownerId,
				0,
				true,
			).
			Return(expectedCompanies, 0, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		companies, _, err := svc.GetByOwnerId(ctx, ownerId, 0, true)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedCompanies, companies)
	})
}

func (s *CompanySuite) Test_CompanyGetByOwnerId2(t provider.T) {
	t.Title("[CompanyGetByOwnerId] Ошибка получения данных в репозитории")
	t.Tags("company", "getByOwnerId")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		ownerId := uuid.UUID{0}

		compRepo.EXPECT().
			GetByOwnerId(
				ctx,
				ownerId,
				0,
				true,
			).
			Return(nil, 0, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		_, _, err := svc.GetByOwnerId(ctx, ownerId, 0, true)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение списка компаний по id владельца: sql error").Error(), err.Error())
	})
}

func (s *CompanySuite) Test_CompanyUpdate(t provider.T) {
	t.Title("[CompanyUpdate] Успешно")
	t.Tags("company", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		actFieldId := uuid.UUID{0}
		updatedInfoCompany := utils.NewCompanyBuilder().
			WithID(uuid.UUID{1}).
			WithName("aaa").
			Build()

		actFieldModel := utils.NewActivityFieldBuilder().
			WithID(actFieldId).
			WithName("test").
			WithDescription("test").
			WithCost(1.0).
			Build()

		compRepo.EXPECT().
			Update(
				ctx,
				&updatedInfoCompany,
			).Return(nil)

		repo.EXPECT().
			GetById(
				ctx,
				actFieldId,
			).Return(&actFieldModel, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		err := svc.Update(ctx, &updatedInfoCompany)

		sCtx.Assert().NoError(err)
	})
}

func (s *CompanySuite) Test_CompanyUpdate2(t provider.T) {
	t.Title("[CompanyUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("company", "update")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIActivityFieldRepository(ctrl)
		compRepo := mocks.NewMockICompanyRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := company.NewService(compRepo, repo, log)

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
		actFieldId := uuid.UUID{0}
		updatedInfoCompany := utils.NewCompanyBuilder().
			WithID(uuid.UUID{1}).
			WithName("aaa").
			Build()

		actFieldModel := utils.NewActivityFieldBuilder().
			WithID(actFieldId).
			WithName("test").
			WithDescription("test").
			WithCost(1.0).
			Build()

		compRepo.EXPECT().
			Update(
				ctx,
				&updatedInfoCompany,
			).Return(fmt.Errorf("sql error"))

		repo.EXPECT().
			GetById(
				ctx,
				actFieldId,
			).Return(&actFieldModel, nil)
		sCtx.WithNewParameters("ctx", ctx, "model", 0)

		err := svc.Update(ctx, &updatedInfoCompany)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление информации о компании: sql error").Error(), err.Error())
	})
}
