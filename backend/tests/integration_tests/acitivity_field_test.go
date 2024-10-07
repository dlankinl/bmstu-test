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

type ITActFieldSuite struct {
	suite.Suite
	repo domain.IActivityFieldRepository
}

var createdField *domain.ActivityField

func (s *ITActFieldSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = postgres.NewActivityFieldRepository(TestDbInstance)
	t.Tags("fixture", "activityField")
}

func (s *ITActFieldSuite) Test_ActFieldStorage_Create(t provider.T) {
	t.Title("[Create] Успех")
	t.Tags("integration test", "postgres", "act_field")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		actField := utils.NewActivityFieldBuilder().
			WithName("a").
			WithDescription("a").
			WithCost(1).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", actField)
		res, err := s.repo.Create(ctx, &actField)
		createdField = res

		sCtx.Assert().NoError(err)
	})
}

func (s *ITActFieldSuite) Test_ActFieldStorage_GetMaxCost(t provider.T) {
	t.Title("[GetMaxCost] Успех")
	t.Tags("integration test", "postgres", "act_field")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()

		expected := float32(1.3)

		sCtx.WithNewParameters("ctx", ctx, "model", 0)
		res, err := s.repo.GetMaxCost(ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expected, res)
	})
}

func (s *ITActFieldSuite) Test_ActFieldStorage_DeleteById(t provider.T) {
	t.Title("[DeleteById] Успех")
	t.Tags("integration test", "postgres", "act_field")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id, _ := uuid.Parse("a8fdf8a5-c539-4f74-9e67-c9853ff9994b")
		sCtx.WithNewParameters("ctx", ctx, "model", id)

		err := s.repo.DeleteById(ctx, id)

		_, getErr := s.repo.GetById(ctx, id)
		sCtx.Assert().NoError(err)
		sCtx.Assert().EqualError(getErr, fmt.Errorf("получение сферы деятельности по id: %w", pgx.ErrNoRows).Error())
	})
}

func (s *ITActFieldSuite) Test_ActFieldStorage_Update(t provider.T) {
	t.Title("[Update] Успех")
	t.Tags("integration test", "postgres", "act_field")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		id, _ := uuid.Parse("f80426b8-27e7-4bfa-8721-23075f125165")
		actField := utils.NewActivityFieldBuilder().
			WithID(id).
			WithCost(1.2).
			Build()

		expectedField := utils.NewActivityFieldBuilder().
			WithID(id).
			WithName("field1").
			WithDescription("field1_descr").
			WithCost(1.2).
			Build()

		sCtx.WithNewParameters("ctx", ctx, "model", actField)

		err := s.repo.Update(ctx, &actField)

		res, getErr := s.repo.GetById(ctx, id)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NoError(getErr)
		sCtx.Assert().Equal(&expectedField, res)
	})
}
