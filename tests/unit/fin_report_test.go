//go:build unit

package unit

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"ppo/internal/services/fin_report"
	"ppo/internal/utils"
	"ppo/mocks"
)

type FinReportSuite struct {
	suite.Suite
}

func (s *FinReportSuite) Test_FinReportCreate(t provider.T) {
	t.Title("[FinReportCreate] Успех")
	t.Tags("finReport", "create")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{10}
		model := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(1).
			WithCosts(1).
			WithYear(1).
			WithQuarter(1).
			Build()
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

func (s *FinReportSuite) Test_FinReportCreate2(t provider.T) {
	t.Title("[FinReportCreate] Отрицательная выручка")
	t.Tags("finReport", "create")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{20}
		model := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(-1).
			WithCosts(1).
			WithYear(1).
			WithQuarter(1).
			Build()
		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err := svc.Create(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("выручка не может быть отрицательной").Error(), err.Error())
	})
}

func (s *FinReportSuite) Test_FinReportDeleteById(t provider.T) {
	t.Title("[FinReportDeleteById] Успешно")
	t.Tags("finReport", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{1}
		ctx := context.TODO()

		repo.EXPECT().
			DeleteById(ctx, reportId).
			Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := svc.DeleteById(ctx, reportId)

		sCtx.Assert().NoError(err)
	})
}

func (s *FinReportSuite) Test_FinReportDeleteById2(t provider.T) {
	t.Title("[FinReportDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("finReport", "deleteById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{1}
		ctx := context.TODO()

		repo.EXPECT().
			DeleteById(ctx, reportId).
			Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := svc.DeleteById(ctx, reportId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление отчета по id: sql error").Error(), err.Error())
	})
}

func (s *FinReportSuite) Test_FinReportGetByCompany(t provider.T) {
	t.Title("[FinReportGetByCompany] Успешно")
	t.Tags("finReport", "getByCompany")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		repo.EXPECT().
			GetByCompany(
				ctx,
				compId,
				&period,
			).
			Return(&repByPeriod, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", compId)

		rep, err := svc.GetByCompany(ctx, compId, &period)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&repByPeriod, rep)
	})
}

func (s *FinReportSuite) Test_FinReportGetByCompany2(t provider.T) {
	t.Title("[FinReportGetByCompany] Год начала периода больше года конца периода")
	t.Tags("finReport", "getByCompany")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		compId := uuid.UUID{0}
		period := utils.NewPeriodBuilder().
			WithStartYear(2).
			WithEndYear(1).
			WithStartQuarter(1).
			WithEndQuarter(1).
			Build()

		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "model", compId)

		_, err := svc.GetByCompany(ctx, compId, &period)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("дата конца периода должна быть позже даты начала").Error(), err.Error())
	})
}

func (s *FinReportSuite) Test_FinReportGetById(t provider.T) {
	t.Title("[FinReportGetById] Успешно")
	t.Tags("finReport", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{0}
		compId := uuid.UUID{0}
		expectedReport := utils.NewFinReportBuilder().
			WithID(reportId).
			WithCompanyID(compId).
			WithRevenue(1).
			WithCosts(1).
			WithYear(1).
			WithQuarter(1).
			Build()
		ctx := context.TODO()

		repo.EXPECT().
			GetById(
				ctx,
				reportId,
			).
			Return(&expectedReport, nil)

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		report, err := svc.GetById(ctx, reportId)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&expectedReport, report)
	})
}

func (s *FinReportSuite) Test_FinReportGetById2(t provider.T) {
	t.Title("[FinReportGetById] Ошибка получения данных в репозитории")
	t.Tags("finReport", "getById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{0}
		ctx := context.TODO()

		repo.EXPECT().
			GetById(
				ctx,
				reportId,
			).
			Return(nil, fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		_, err := svc.GetById(ctx, reportId)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение финансового отчета по id: sql error").Error(), err.Error())
	})
}

func (s *FinReportSuite) Test_FinReportUpdate(t provider.T) {
	t.Title("[FinReportUpdate] Успешно")
	t.Tags("finReport", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{0}
		report := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(2).
			Build()
		ctx := context.TODO()

		repo.EXPECT().
			Update(
				ctx,
				&report,
			).Return(nil)

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := svc.Update(ctx, &report)

		sCtx.Assert().NoError(err)
	})
}

func (s *FinReportSuite) Test_FinReportUpdate2(t provider.T) {
	t.Title("[FinReportUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("finReport", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIFinancialReportRepository(ctrl)
		log := mocks.NewMockILogger(ctrl)
		svc := fin_report.NewService(repo, log)

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

		reportId := uuid.UUID{0}
		report := utils.NewFinReportBuilder().
			WithID(reportId).
			WithRevenue(2).
			Build()
		ctx := context.TODO()

		repo.EXPECT().
			Update(
				ctx,
				&report,
			).Return(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", reportId)

		err := svc.Update(ctx, &report)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление отчета: sql error").Error(), err.Error())
	})
}
