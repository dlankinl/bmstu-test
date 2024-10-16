//go:build e2e

package e2e

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"os"
	"ppo/domain"
	"ppo/mocks"
	"time"
)

type E2ESuite struct {
	suite.Suite
	ctrl *gomock.Controller

	logger *mocks.MockILogger
	crypto *mocks.MockIHashCrypto
	aSvc   domain.IActivityFieldService
	uSvc   domain.IUserService
	auSvc  domain.IAuthService
	cSvc   domain.ICompanyService
	fSvc   domain.IFinancialReportService

	e httpexpect.Expect
}

func (s *E2ESuite) BeforeAll(t provider.T) {
	s.ctrl = gomock.NewController(t)

	t.Title("[e2e] init test repository")
	//aRepo := postgres.NewActivityFieldRepository(TestDbInstance)
	//uRepo := postgres.NewUserRepository(TestDbInstance)
	//auRepo := postgres.NewAuthRepository(TestDbInstance)
	//cRepo := postgres.NewCompanyRepository(TestDbInstance)
	//fRepo := postgres.NewFinReportRepository(TestDbInstance)
	//
	s.logger = mocks.NewMockILogger(s.ctrl)
	s.crypto = mocks.NewMockIHashCrypto(s.ctrl)
	//s.aSvc = activity_field.NewService(aRepo, cRepo, s.logger)
	//s.uSvc = user.NewService(uRepo, cRepo, aRepo, s.logger)
	//s.auSvc = auth.NewService(auRepo, s.crypto, "jwt123", s.logger)
	//s.cSvc = company.NewService(cRepo, aRepo, s.logger)
	//s.fSvc = fin_report.NewService(fRepo, s.logger)

	s.e = *httpexpect.WithConfig(httpexpect.Config{
		Client:   &http.Client{},
		BaseURL:  "http://localhost:8083",
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Tags("fixture", "e2e")
	done := make(chan os.Signal, 1)
	go RunTheApp(TestDbInstance, done)
}

//func (s *E2ESuite) Test(t provider.T) {
//	t.Title("[e2e] Test")
//	t.Tags("e2e", "postgres")
//	t.Parallel()
//	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
//		ctx := context.TODO()
//
//		s.logger.EXPECT().
//			Infof(gomock.Any()).
//			AnyTimes()
//		s.logger.EXPECT().
//			Infof(gomock.Any(), gomock.Any()).
//			AnyTimes()
//		s.logger.EXPECT().
//			Warnf(gomock.Any(), gomock.Any()).
//			AnyTimes()
//		s.logger.EXPECT().
//			Errorf(gomock.Any(), gomock.Any()).
//			AnyTimes()
//
//		s.crypto.EXPECT().
//			CheckPasswordHash("newUserPass", "pass123").
//			Return(true)
//
//		s.crypto.EXPECT().
//			GenerateHashPass("newUserPass").
//			Return("pass123", nil)
//
//		userAuth := utils.NewUserAuthBuilder().
//			WithUsername("newUser").
//			WithPassword("newUserPass").
//			Build()
//
//		loginUserAuth := utils.NewUserAuthBuilder().
//			WithUsername("newUser").
//			WithPassword("newUserPass").
//			WithHashedPass("pass123").
//			Build()
//
//		id, regErr := s.auSvc.Register(ctx, &userAuth)
//
//		_, logErr := s.auSvc.Login(ctx, &loginUserAuth)
//
//		actFieldId, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
//
//		company := utils.NewCompanyBuilder().
//			WithName("newCompany").
//			WithCity("newCity").
//			WithOwner(id).
//			WithActivityField(actFieldId).
//			Build()
//
//		comp, compCreateErr := s.cSvc.Create(ctx, &company)
//
//		report := utils.NewFinReportBuilder().
//			WithQuarter(1).
//			WithYear(1).
//			WithCompanyID(comp.ID).
//			WithCosts(100).
//			WithRevenue(200).
//			Build()
//
//		reportCreateErr := s.fSvc.Create(ctx, &report)
//
//		period := utils.NewPeriodBuilder().
//			WithStartYear(1).
//			WithEndYear(1).
//			WithStartQuarter(1).
//			WithEndQuarter(1).
//			Build()
//
//		reports, getByCompanyErr := s.fSvc.GetByCompany(ctx, comp.ID, &period)
//
//		expectedReports := utils.NewFinReportByPeriodBuilder().
//			WithReports([]domain.FinancialReport{report}).
//			WithPeriod(period).
//			Build()
//
//		sCtx.Assert().NoError(regErr)
//		sCtx.Assert().NoError(logErr)
//		sCtx.Assert().NoError(compCreateErr)
//		sCtx.Assert().NoError(reportCreateErr)
//		sCtx.Assert().NoError(getByCompanyErr)
//		sCtx.Assert().Equal(&expectedReports, reports)
//	})
//}

func (s *E2ESuite) Test2(t provider.T) {
	t.Title("[e2e] Test2")
	t.Tags("e2e", "postgres")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		s.logger.EXPECT().
			Infof(gomock.Any()).
			AnyTimes()
		s.logger.EXPECT().
			Infof(gomock.Any(), gomock.Any()).
			AnyTimes()
		s.logger.EXPECT().
			Warnf(gomock.Any(), gomock.Any()).
			AnyTimes()
		s.logger.EXPECT().
			Errorf(gomock.Any(), gomock.Any()).
			AnyTimes()

		s.crypto.EXPECT().
			CheckPasswordHash("newUserPass", "pass123").
			Return(true)

		s.crypto.EXPECT().
			GenerateHashPass("newUserPass").
			Return("pass123", nil)

		type Req struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		req := Req{"newUser2", "newUserPass"}

		s.e.POST("/signup").
			WithJSON(req).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			HasValue("status", "success")

		time.Sleep(time.Millisecond * 200)

		s.e.POST("/login").
			WithJSON(req).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			NotEmpty().
			HasValue("status", "success")
	})
}
