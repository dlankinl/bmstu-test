//go:build integration

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
)

type ITFinReportSuite struct {
	suite.Suite
	repo domain.IFinancialReportRepository
}

func (s *ITFinReportSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = postgres.NewFinReportRepository(TestDbInstance)
	t.Tags("fixture", "finReport")
}

func (s *ITFinReportSuite) Test_FinReportStorage_Create(t provider.T) {
	t.Title("[Create] Успех")
	t.Tags("integration test", "postgres", "fin_report")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
		finReport := utils.NewFinReportBuilder().
			WithCompanyID(id).
			WithRevenue(1.32).
			WithCosts(1.23).
			WithYear(2024).
			WithQuarter(1).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", finReport)
		res, err := s.repo.Create(ctx, &finReport)

		_, getErr := s.repo.GetById(ctx, res.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NoError(getErr)
	})
}

func (s *ITFinReportSuite) Test_FinReportStorage_Create2(t provider.T) {
	t.Title("[Create] Несуществующий id компании")
	t.Tags("integration test", "postgres", "fin_report")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed380a0")
		finReport := utils.NewFinReportBuilder().
			WithCompanyID(id).
			WithRevenue(1.32).
			WithCosts(1.23).
			WithYear(2024).
			WithQuarter(1).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", finReport)
		_, err := s.repo.Create(ctx, &finReport)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf(`создание финансового отчета: ERROR: insert or update on table "fin_reports" violates foreign key constraint "fk_company" (SQLSTATE 23503)`).Error(), err.Error())
	})
}

func (s *ITFinReportSuite) Test_FinReportStorage_GetById(t provider.T) {
	t.Title("[GetById] Успех")
	t.Tags("integration test", "postgres", "fin_report")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id, _ := uuid.Parse("124b7c0e-5a5b-4489-a856-c8ad0a2328ad")
		compId, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
		expFinReport := utils.NewFinReportBuilder().
			WithID(id).
			WithCompanyID(compId).
			WithRevenue(1).
			WithCosts(0.5).
			WithYear(1).
			WithQuarter(1).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", id)
		report, err := s.repo.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&expFinReport, report)
	})
}

func (s *ITFinReportSuite) Test_FinReportStorage_GetById2(t provider.T) {
	t.Title("[GetById] Несуществующий отчет")
	t.Tags("integration test", "postgres", "fin_report")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id, _ := uuid.Parse("124b7c0e-5a5b-4489-a816-c8ad0a2328ad")

		sCtx.WithNewParameters("ctx", ctx, "model", id)
		_, err := s.repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение отчета по id: %w", pgx.ErrNoRows).Error(), err.Error())
	})
}

func (s *ITFinReportSuite) Test_FinReportStorage_DeleteById(t provider.T) {
	t.Title("[GetById] Успех")
	t.Tags("integration test", "postgres", "fin_report")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		id, _ := uuid.Parse("4b202e34-7e00-4eee-8cd7-436dd7f1298f")

		sCtx.WithNewParameters("ctx", ctx, "model", id)
		err := s.repo.DeleteById(ctx, id)

		_, getErr := s.repo.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Error(getErr)
		sCtx.Assert().Equal(fmt.Errorf("получение отчета по id: %w", pgx.ErrNoRows).Error(), getErr.Error())
	})
}

func (s *ITFinReportSuite) Test_FinReportStorage_GetByCompany(t provider.T) {
	t.Title("[GetByCompany] Успех")
	t.Tags("integration test", "postgres", "fin_report")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		repId, _ := uuid.Parse("124b7c0e-5a5b-4489-a856-c8ad0a2328ad")
		compId, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")

		period := utils.NewPeriodBuilder().
			WithStartYear(1).
			WithEndYear(2).
			WithStartQuarter(1).
			WithEndQuarter(2).
			Build()

		expReport := utils.NewFinReportBuilder().
			WithID(repId).
			WithCompanyID(compId).
			WithRevenue(1.0).
			WithCosts(0.5).
			WithYear(1).
			WithQuarter(1).
			Build()

		reportByPeriod := utils.NewFinReportByPeriodBuilder().
			WithReports([]domain.FinancialReport{expReport}).
			WithPeriod(period).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", compId)
		res, err := s.repo.GetByCompany(ctx, compId, &period)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&reportByPeriod, res)
	})
}

func (s *ITFinReportSuite) Test_FinReportStorage_Update(t provider.T) {
	t.Title("[Update] Успех")
	t.Tags("integration test", "postgres", "fin_report")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		repId, _ := uuid.Parse("594b555c-aa23-4adc-a67f-419124a60fd6")

		updReport := utils.NewFinReportBuilder().
			WithID(repId).
			WithRevenue(2.0).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", updReport)
		err := s.repo.Update(ctx, &updReport)

		sCtx.Assert().NoError(err)
	})
}
