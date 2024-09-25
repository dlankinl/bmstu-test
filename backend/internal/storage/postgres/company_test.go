package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"ppo/domain"
	"ppo/internal/utils"
)

type StorageCompanySuite struct {
	suite.Suite
	repo domain.ICompanyRepository
}

func (s *StorageCompanySuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = NewCompanyRepository(testDbInstance)
	t.Tags("fixture", "finReport")
}

func (s *StorageCompanySuite) Test_CompanyStorage_Create(t provider.T) {
	t.Title("[Create] Успех")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		ownerId, _ := uuid.Parse("8a7fe516-600e-4c01-a55e-423fac892250")
		actFieldId, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
		company := utils.NewCompanyBuilder().
			WithName("test").
			WithCity("test").
			WithOwner(ownerId).
			WithActivityField(actFieldId).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", company)
		res, err := s.repo.Create(ctx, &company)

		_, getErr := s.repo.GetById(ctx, res.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NoError(getErr)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorage_Create2(t provider.T) {
	t.Title("[Create] Несуществующий пользователь")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		ownerId, _ := uuid.Parse("8a7fe516-600e-4c02-a55e-423fac892250")
		actFieldId, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
		company := utils.NewCompanyBuilder().
			WithName("test").
			WithCity("test").
			WithOwner(ownerId).
			WithActivityField(actFieldId).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", company)
		_, err := s.repo.Create(ctx, &company)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf(`создание компании: ERROR: insert or update on table "companies" violates foreign key constraint "fk_owner" (SQLSTATE 23503)`).Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorage_GetById(t provider.T) {
	t.Title("[GetById] Успех")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		ownerId, _ := uuid.Parse("b384ea3b-df18-4bae-b459-fb96e2518fe7")
		actFieldId, _ := uuid.Parse("b9bacee6-3d2d-48f8-a7bc-493f44b0652a")
		compId, _ := uuid.Parse("c4f2abf1-e80c-4c31-bc77-fe5a8e5fab40")
		expCompany := utils.NewCompanyBuilder().
			WithID(compId).
			WithName("Company2").
			WithCity("Voronezh").
			WithOwner(ownerId).
			WithActivityField(actFieldId).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", compId)
		res, err := s.repo.GetById(ctx, compId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&expCompany, res)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorage_GetById2(t provider.T) {
	t.Title("[GetById] Несуществующая компания")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		compId, _ := uuid.Parse("c4f2abf1-e80c-5c31-bc77-fe5a8e5fab40")

		sCtx.WithNewParameters("ctx", ctx, "model", compId)
		_, err := s.repo.GetById(ctx, compId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение компании по id: %w", pgx.ErrNoRows).Error(), err.Error())
	})
}

func (s *StorageCompanySuite) Test_CompanyStorage_GetByOwnerId(t provider.T) {
	t.Title("[GetByOwnerId] Успех")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		compId, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
		ownerId, _ := uuid.Parse("bc3ab9bf-6a26-4212-941d-05a985fc0978")
		actFieldId, _ := uuid.Parse("f80426b8-27e7-4bfa-8721-23075f125165")

		expComp1 := utils.NewCompanyBuilder().
			WithID(compId).
			WithName("Company1").
			WithCity("Moscow").
			WithActivityField(actFieldId).
			WithOwner(ownerId).
			Build()

		expectedCompanies := []*domain.Company{&expComp1}

		sCtx.WithNewParameters("ctx", ctx, "model", ownerId)
		companies, _, err := s.repo.GetByOwnerId(ctx, ownerId, 1, false)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedCompanies, companies)
	})
}

func (s *StorageCompanySuite) Test_CompanyStorage_Update(t provider.T) {
	t.Title("[Update] Успех")
	t.Tags("storage", "postgres", "company")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		compId, _ := uuid.Parse("0714a5e6-b52a-4e92-b88f-ef27d21acd49")

		comp := utils.NewCompanyBuilder().
			WithID(compId).
			WithName("Company Renamed").
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", compId)
		err := s.repo.Update(ctx, &comp)

		sCtx.Assert().NoError(err)
	})
}

//func (s *StorageCompanySuite) Test_CompanyStorage_GetAll(t provider.T) {
//	t.Title("[GetByOwnerId] Успех")
//	t.Tags("storage", "postgres", "company")
//	t.Parallel()
//	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
//		ctx := context.TODO()
//
//		compId, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
//		ownerId, _ := uuid.Parse("bc3ab9bf-6a26-4212-941d-05a985fc0978")
//		actFieldId, _ := uuid.Parse("f80426b8-27e7-4bfa-8721-23075f125165")
//
//		expComp1 := utils.NewCompanyBuilder().
//			WithID(compId).
//			WithName("Company1").
//			WithCity("Moscow").
//			WithActivityField(actFieldId).
//			WithOwner(ownerId).
//			Build()
//
//		expectedCompanies := []*domain.Company{&expComp1}
//
//		sCtx.WithNewParameters("ctx", ctx, "model", ownerId)
//		companies, err := s.repo.GetAll(ctx, 0)
//
//		sCtx.Assert().NoError(err)
//		sCtx.Assert().Equal(expectedCompanies, companies)
//	})
//}
