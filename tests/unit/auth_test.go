//go:build unit

package unit

import (
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"go.uber.org/mock/gomock"
	"ppo/internal/services/auth"
	"ppo/internal/utils"
	"ppo/mocks"
	"ppo/pkg/base"
)

type AuthSuite struct {
	suite.Suite
}

func (s *AuthSuite) Test_AuthRegister2(t provider.T) {
	t.Title("[AuthRegister] Fail")
	t.Tags("auth", "register")
	t.Parallel()
	t.WithNewStep("Empty username", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIAuthRepository(ctrl)
		crypto := mocks.NewMockIHashCrypto(ctrl)
		log := mocks.NewMockILogger(ctrl)

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

		svc := auth.NewService(repo, crypto, "abcdefgh123", log)

		ctx := context.TODO()

		model := utils.UserAuthMother{}.WithoutUsernameUser()
		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err := svc.Register(ctx, &model)

		sCtx.Assert().Error(err, fmt.Errorf("должно быть указано имя пользователя"))
	})
}

func (s *AuthSuite) Test_AuthLogin2(t provider.T) {
	t.Title("[AuthLogin] Success")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIAuthRepository(ctrl)
		crypto := mocks.NewMockIHashCrypto(ctrl)
		log := mocks.NewMockILogger(ctrl)

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

		returnedModel := utils.UserAuthMother{}.WithHashedPassUser()
		repo.EXPECT().
			GetByUsername(
				context.TODO(),
				"test",
			).
			Return(&returnedModel, nil)

		crypto.EXPECT().
			CheckPasswordHash("test", "pass123").
			Return(true)

		svc := auth.NewService(repo, crypto, "abcdefgh123", log)

		ctx := context.TODO()

		model := utils.UserAuthMother{}.DefaultUser()
		sCtx.WithNewParameters("ctx", ctx, "model", model)

		token, err := svc.Login(ctx, &model)
		_, verifErr := base.VerifyAuthToken(token, "abcdefgh123")

		sCtx.Assert().NoError(err)
		sCtx.Assert().NoError(verifErr)
	})
}

func (s *AuthSuite) Test_AuthLogin(t provider.T) {
	t.Title("[AuthLogin] Success")
	t.Tags("auth", "login")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mocks.NewMockIAuthRepository(ctrl)
		crypto := mocks.NewMockIHashCrypto(ctrl)
		log := mocks.NewMockILogger(ctrl)

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

		svc := auth.NewService(repo, crypto, "abcdefgh123", log)

		ctx := context.TODO()

		model := utils.UserAuthMother{}.WithoutPasswordUser()
		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err := svc.Login(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("должен быть указан пароль"), err)
	})
}
