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

type StorageActFieldSuite struct {
	suite.Suite
	repo domain.IActivityFieldRepository
}

var createdField *domain.ActivityField

func (s *StorageActFieldSuite) BeforeAll(t provider.T) {
	t.Title("init test repository")
	s.repo = NewActivityFieldRepository(testDbInstance)
	t.Tags("fixture", "activityField")
}

func (s *StorageActFieldSuite) Test_ActFieldStorage_Create(t provider.T) {
	t.Title("[Create] Успех")
	t.Tags("storage", "postgres", "act_field")
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

func (s *StorageActFieldSuite) Test_ActFieldStorage_GetMaxCost(t provider.T) {
	t.Title("[GetMaxCost] Успех")
	t.Tags("storage", "postgres", "act_field")
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

func (s *StorageActFieldSuite) Test_ActFieldStorage_DeleteById(t provider.T) {
	t.Title("[DeleteById] Успех")
	t.Tags("storage", "postgres", "act_field")
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

func (s *StorageActFieldSuite) Test_ActFieldStorage_Update(t provider.T) {
	t.Title("[Update] Успех")
	t.Tags("storage", "postgres", "act_field")
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

//func (s *StorageActFieldSuite) Test_ActFieldStorage_GetAll(t provider.T) {
//	t.Title("[GetAll] Успех")
//	t.Tags("storage", "postgres", "act_field")
//	t.Parallel()
//	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
//		ctx := context.TODO()
//		id1, _ := uuid.Parse("f80426b8-27e7-4bfa-8721-23075f125165")
//		id2, _ := uuid.Parse("b9bacee6-3d2d-48f8-a7bc-493f44b0652a")
//		id3, _ := uuid.Parse("fa406cca-27d6-446e-8cfd-b1a71ed680a0")
//		id4, _ := uuid.Parse("a8fdf8a5-c539-4f74-9e67-c9853ff9994b")
//
//		actField1 := utils.NewActivityFieldBuilder().
//			WithID(id1).
//			WithCost(1.2).
//			WithName("field1").
//			WithDescription("field1_descr").
//			Build()
//
//		actField2 := utils.NewActivityFieldBuilder().
//			WithID(id2).
//			WithCost(0.2).
//			WithName("field2").
//			WithDescription("field2_descr").
//			Build()
//
//		actField3 := utils.NewActivityFieldBuilder().
//			WithID(id3).
//			WithCost(1.3).
//			WithName("field3").
//			WithDescription("field3_descr").
//			Build()
//
//		actField4 := utils.NewActivityFieldBuilder().
//			WithID(id4).
//			WithCost(0.3).
//			WithName("field4").
//			WithDescription("field4_descr").
//			Build()
//
//		expActFields := []*domain.ActivityField{&actField1, &actField2, &actField3, &actField4}
//
//		sCtx.WithNewParameters("ctx", ctx, "model", expActFields)
//
//		all, _, err := s.repo.GetAll(ctx, 1, false)
//
//		sCtx.Assert().NoError(err)
//		sCtx.Assert().Equal(expActFields, all)
//	})
//}
