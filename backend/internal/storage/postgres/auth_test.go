package postgres

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"ppo/internal/utils"
	"ppo/mocks"
)

type StorageAuthSuite struct {
	suite.Suite
	repo *mocks.MockIAuthRepository
	ctrl *gomock.Controller
}

func (s *StorageAuthSuite) BeforeAll(t provider.T) {
	s.ctrl = gomock.NewController(t)
	s.repo = mocks.NewMockIAuthRepository(s.ctrl)
}

func (s *StorageAuthSuite) AfterAll(t provider.T) {
	s.ctrl.Finish()
}

func (s *StorageAuthSuite) Test_AuthStorageRegister(t provider.T) {
	t.Title("[AuthRegister] Success")
	t.Tags("storage", "auth", "register")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		registerModel := utils.UserAuthMother{}.WithHashedPassUser()
		s.repo.EXPECT().
			Register(
				context.TODO(),
				&registerModel,
			).
			Return(nil)

		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "model", registerModel)

		err := s.repo.Register(ctx, &registerModel)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageAuthSuite) Test_AuthStorageRegister2(t provider.T) {
	t.Title("[AuthRegister] Fail")
	t.Tags("storage", "auth", "register")
	t.Parallel()
	t.WithNewStep("Empty username", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		model := utils.UserAuthMother{}.WithoutUsernameUser()
		sCtx.WithNewParameters("ctx", ctx, "model", model)

		s.repo.EXPECT().
			Register(
				ctx,
				&model,
			).Return(nil)

		err := s.repo.Register(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageAuthSuite) Test_AuthStorageGetByUsername(t provider.T) {
	t.Title("[AuthLogin] Success")
	t.Tags("storage", "auth", "login")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		model := utils.UserAuthMother{}.WithHashedPassUser()
		s.repo.EXPECT().
			GetByUsername(
				context.TODO(),
				"test",
			).
			Return(&model, nil)

		ctx := context.TODO()

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		userAuth, err := s.repo.GetByUsername(ctx, model.Username)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(&model, userAuth)
	})
}
