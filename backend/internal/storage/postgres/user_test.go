package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pashagolub/pgxmock/v4"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/internal/utils"
	"time"
)

type StorageUserSuite struct {
	suite.Suite
}

func (s *StorageUserSuite) Test_UserStorageDeleteById(t provider.T) {
	t.Title("[UserDeleteById] Успех")
	t.Tags("storage", "user", "deleteById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		id := uuid.UUID{1}
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("delete").WithArgs(id).WillReturnResult(pgxmock.NewResult("delete", 1))

		repo := NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err = repo.DeleteById(ctx, id)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageUserSuite) Test_UserStorageDeleteById2(t provider.T) {
	t.Title("[UserDeleteById] Ошибка выполнения запроса в репозитории")
	t.Tags("storage", "user", "deleteById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		id := uuid.UUID{2}
		ctx := context.TODO()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("delete").WithArgs(id).WillReturnError(fmt.Errorf("sql error"))

		repo := NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err = repo.DeleteById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("удаление пользователя по id: sql error").Error(), err.Error())
	})
}

func (s *StorageUserSuite) Test_UserStorageGetAll(t provider.T) {
	t.Title("[UserGetAll] Успех")
	t.Tags("storage", "user", "getAll")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		user1 := utils.NewUserBuilder().
			WithId(uuid.UUID{1}).
			WithUsername("a").
			WithFullName("a").
			WithGender("m").
			WithBirthday(time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)).
			WithCity("a").
			Build()
		user2 := utils.NewUserBuilder().
			WithId(uuid.UUID{2}).
			WithUsername("b").
			WithFullName("b").
			WithGender("w").
			WithBirthday(time.Date(2, 2, 2, 2, 2, 2, 2, time.UTC)).
			WithCity("b").
			Build()
		user3 := utils.NewUserBuilder().
			WithId(uuid.UUID{3}).
			WithUsername("c").
			WithFullName("c").
			WithGender("m").
			WithBirthday(time.Date(3, 3, 3, 3, 3, 3, 3, time.UTC)).
			WithCity("c").
			Build()

		users := []*domain.User{&user1, &user2, &user3}
		page := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").
			WithArgs(
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnRows(
				pgxmock.NewRows([]string{"id", "username", "full_name", "birthday", "gender", "city"}).
					AddRow(users[0].ID, users[0].Username, users[0].FullName, users[0].Birthday, users[0].Gender, users[0].City).
					AddRow(users[1].ID, users[1].Username, users[1].FullName, users[1].Birthday, users[1].Gender, users[1].City).
					AddRow(users[2].ID, users[2].Username, users[2].FullName, users[2].Birthday, users[2].Gender, users[2].City),
			)

		mock.ExpectQuery("select").WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(3))

		sCtx.WithNewParameters("ctx", ctx, "model", users)

		repo := NewUserRepository(mock)

		got, _, err := repo.GetAll(ctx, page)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(users, got)
	})
}

func (s *StorageUserSuite) Test_UserStorageGetAll2(t provider.T) {
	t.Title("[UserGetAll] Ошибка получения данных в репозитории")
	t.Tags("storage", "user", "getAll")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		page := 1

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").
			WithArgs(
				config.PageSize*(page-1),
				config.PageSize,
			).
			WillReturnError(fmt.Errorf("sql error"))

		sCtx.WithNewParameters("ctx", ctx, "model", page)

		repo := NewUserRepository(mock)

		_, _, err = repo.GetAll(ctx, page)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение предпринимателей: sql error").Error(), err.Error())
	})
}

func (s *StorageUserSuite) Test_UserGetById(t provider.T) {
	t.Title("[UserGetById] Успех")
	t.Tags("user", "getById")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.UUID{1}
		model := utils.NewUserBuilder().
			WithId(id).
			WithUsername("a").
			WithFullName("a b c").
			WithGender("m").
			WithBirthday(time.Date(1, 1, 1, 1, 1, 1, 1, time.Local)).
			WithCity("a").
			Build()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(id).WillReturnRows(pgxmock.
			NewRows([]string{"username", "full_name", "birthday", "gender", "city", "role"}).
			AddRow(model.Username, model.FullName, model.Birthday, model.Gender, model.City, model.Role))

		repo := NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		user, err := repo.GetById(ctx, id)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(model, *user)
	})
}

func (s *StorageUserSuite) Test_UserStorageGetById2(t provider.T) {
	t.Title("[UserGetById] Ошибка получения данных в репозитории")
	t.Tags("storage", "user", "getById")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.UUID{4}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectQuery("select").WithArgs(id).WillReturnError(fmt.Errorf("sql error"))

		repo := NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", id)

		_, err = repo.GetById(ctx, id)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("получение пользователя по id: sql error").Error(), err.Error())
	})
}

func (s *StorageUserSuite) Test_UserStorageUpdate(t provider.T) {
	t.Title("[UserUpdate] Успех")
	t.Tags("storage", "user", "update")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id := uuid.UUID{1}
		model := utils.NewUserBuilder().
			WithId(id).
			WithCity("a").
			WithRole("admin").
			Build()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("update").WithArgs(model.City, model.Role, model.ID).
			WillReturnResult(pgxmock.NewResult("update", 1))

		repo := NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err = repo.Update(ctx, &model)

		sCtx.Assert().NoError(err)
	})
}

func (s *StorageUserSuite) Test_UserUpdate2(t provider.T) {
	t.Title("[UserUpdate] Ошибка выполнения запроса в репозитории")
	t.Tags("user", "update")
	t.Parallel()
	t.WithNewStep("Fail", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		uId := uuid.UUID{10}
		model := utils.NewUserBuilder().
			WithId(uId).
			WithCity("a").
			WithRole("admin").
			Build()

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close()

		mock.ExpectExec("update").WithArgs(model.City, model.Role, model.ID).
			WillReturnError(fmt.Errorf("sql error"))

		repo := NewUserRepository(mock)

		sCtx.WithNewParameters("ctx", ctx, "model", model)

		err = repo.Update(ctx, &model)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(fmt.Errorf("обновление информации о пользователе: sql error").Error(), err.Error())
	})
}
