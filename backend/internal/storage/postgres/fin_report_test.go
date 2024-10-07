//go:build unit

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

type StorageFinReportSuite struct {
	suite.Suite
	repo *mocks.MockIFinancialReportRepository
	ctrl *gomock.Controller
}

func (s *StorageFinReportSuite) BeforeAll(t provider.T) {
	s.ctrl = gomock.NewController(t)
	s.repo = mocks.NewMockIFinancialReportRepository(s.ctrl)
}

func (s *StorageFinReportSuite) AfterAll(t provider.T) {
	s.ctrl.Finish()
}

func (s *StorageFinReportSuite) Test_FinReportStorageCreate(t provider.T) {
	t.Title("[FinReportCreate] Успех")
	t.Tags("storage", "finReport", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{1}
		model := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(1).
			WithCosts(1).
			WithYear(1).
			WithQuarter(1).
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

func (s *StorageFinReportSuite) Test_FinReportStorageCreate2(t provider.T) {
	t.Title("[FinReportCreate] Отрицательная выручка")
	t.Tags("storage", "finReport", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{1}
		model := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(-1).
			WithCosts(1).
			WithYear(1).
			WithQuarter(1).
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

func (s *StorageFinReportSuite) Test_FinReportStorageDeleteById(t provider.T) {
	t.Title("[FinReportDeleteById] Успешно")
	t.Tags("storage", "finReport", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{100}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(ctx, reportId).
			Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := s.repo.DeleteById(ctx, reportId)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageFinReportSuite) Test_FinReportDeleteById2(t provider.T) {
	t.Title("[FinReportDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("storage", "finReport", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{1}
		ctx := context.TODO()

		s.repo.EXPECT().
			DeleteById(ctx, reportId).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := s.repo.DeleteById(ctx, reportId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageFinReportSuite) Test_FinReportStorageGetByCompany(t provider.T) {
	t.Title("[FinReportGetByCompany] Успешно")
	t.Tags("storage", "finReport", "getByCompany")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		compId := uuid.UUID{0}
		period := utils.NewPeriodBuilder().
			WithStartYear(2021).
			WithEndYear(2023).
			WithStartQuarter(2).
			WithEndQuarter(4).
			Build()

		reps := utils.FinReportMother{}.ForBigPeriod(2021, 2, 2023, 4,
			[]float32{1432523, 7435235, 65742, 43635325, 50934123, 78902453, 64352357, 32532513, 6743634, 46754124, 14385253},
			[]float32{75423, 125654, 7845634, 12362332, 13543623, 15326443, 23534252, 5436438, 9876967, 24367653, 7546424})
		repByPeriod := utils.NewFinReportByPeriodBuilder().
			WithReports(reps).
			WithPeriod(period).
			Build()

		ctx := context.TODO()

		s.repo.EXPECT().
			GetByCompany(
				ctx,
				compId,
				&period,
			).
			Return(&repByPeriod, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", compId)

		rep, err := s.repo.GetByCompany(ctx, compId, &period)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&repByPeriod, rep)
	})
}

func (s *StorageFinReportSuite) Test_FinReportStorageGetByCompany2(t provider.T) {
	t.Title("[FinReportGetByCompany] Год начала периода больше года конца периода")
	t.Tags("storage", "finReport", "getByCompany")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		compId := uuid.UUID{1}
		period := utils.NewPeriodBuilder().
			WithStartYear(2).
			WithEndYear(1).
			WithStartQuarter(1).
			WithEndQuarter(1).
			Build()

		expected := domain.FinancialReportByPeriod{
			Reports: []domain.FinancialReport{},
			Period:  &period,
		}

		s.repo.EXPECT().
			GetByCompany(
				ctx,
				compId,
				&period,
			).Return(&expected, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", compId)

		res, err := s.repo.GetByCompany(ctx, compId, &period)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&expected, res)
	})
}

func (s *StorageFinReportSuite) Test_FinReportStorageGetById(t provider.T) {
	t.Title("[FinReportGetById] Успешно")
	t.Tags("storage", "finReport", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{1}
		compId := uuid.UUID{2}
		expectedReport := utils.NewFinReportBuilder().
			WithID(reportId).
			WithCompanyID(compId).
			WithRevenue(1).
			WithCosts(1).
			WithYear(1).
			WithQuarter(1).
			Build()
		ctx := context.TODO()

		s.repo.EXPECT().
			GetById(
				ctx,
				reportId,
			).
			Return(&expectedReport, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		report, err := s.repo.GetById(ctx, reportId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&expectedReport, report)
	})
}

func (s *StorageFinReportSuite) Test_FinReportStorageGetById2(t provider.T) {
	t.Title("[FinReportGetById] Ошибка получения данных в репозитории")
	t.Tags("storage", "finReport", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{2}
		ctx := context.TODO()

		s.repo.EXPECT().
			GetById(
				ctx,
				reportId,
			).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		_, err := s.repo.GetById(ctx, reportId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}

func (s *StorageFinReportSuite) Test_FinReportStorageUpdate(t provider.T) {
	t.Title("[FinReportUpdate] Успешно")
	t.Tags("storage", "finReport", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{4}
		report := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(2).
			Build()
		ctx := context.TODO()

		s.repo.EXPECT().
			Update(
				ctx,
				&report,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := s.repo.Update(ctx, &report)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageFinReportSuite) Test_FinReportStorageUpdate2(t provider.T) {
	t.Title("[FinReportUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("storage", "finReport", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		reportId := uuid.UUID{5}
		report := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(2).
			Build()
		ctx := context.TODO()

		s.repo.EXPECT().
			Update(
				ctx,
				&report,
			).Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := s.repo.Update(ctx, &report)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("sql error").Error(), err.Error())
	})
}
