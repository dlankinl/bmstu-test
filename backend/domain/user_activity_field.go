package domain

import (
	"context"
	"github.com/google/uuid"
)

//go:generate mockgen -source=user_activity_field.go -destination=../mocks/user_activity_field.go -package=mocks
type IInteractor interface {
	GetMostProfitableCompany(context.Context, *Period, []*Company) (*Company, error)
	CalculateUserRating(context.Context, uuid.UUID) (float32, error)
	GetUserFinancialReport(context.Context, uuid.UUID, *Period) (*FinancialReportByPeriod, error)
}
