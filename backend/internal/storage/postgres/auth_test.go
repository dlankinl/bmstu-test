package postgres

import (
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"ppo/internal/utils"
)

type StorageAuthSuite struct {
	suite.Suite
}

func (s *StorageAuthSuite) Test_AuthStorageRegister2(t provider.T) {
	t.Title("[AuthRegister] Fail")
	t.Tags("storage", "auth", "register")
	t.Parallel()
	t.WithNewStep("Empty username", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		model := utils.UserAuthMother{}.WithoutUsernameUser()
		sCtx.WithNewParameters("ctx", ctx, "model", model)

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("insert").WithArgs(model.Username, model.HashedPass).
			WillReturnError(fmt.Errorf("sql error"))

		repo := NewAuthRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err = repo.Register(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("регистрация пользователя: sql error").Error(), err.Error())
	})
}

// TODO: fix
func (s *StorageAuthSuite) Test_AuthStorageGetByUsername(t provider.T) {
	t.Title("[AuthLogin] Success")
	t.Tags("storage", "auth", "login")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		model := utils.UserAuthMother{}.WithHashedPassUser()
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").
			WithArgs(model.Username).
			WillReturnRows(pgxmock.NewRows([]string{"id", "password", "role"}).
				AddRow(model.ID, model.Password, model.Role))

		repo := NewAuthRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		_, err = repo.GetByUsername(ctx, model.Username)

		sCtx.Assert().NoError(err)
	})
}
