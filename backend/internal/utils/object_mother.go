package utils

import (
	"github.com/google/uuid"
	"ppo/domain"
)

type UserAuthMother struct{}

func (m UserAuthMother) DefaultUser() domain.UserAuth {
	return domain.UserAuth{
		Username: "test",
		Password: "test",
	}
}

func (m UserAuthMother) WithoutPasswordUser() domain.UserAuth {
	return domain.UserAuth{
		Username: "test",
	}
}

func (m UserAuthMother) WithoutUsernameUser() domain.UserAuth {
	return domain.UserAuth{
		Password: "test",
	}
}

func (m UserAuthMother) WithHashedPassUser() domain.UserAuth {
	return domain.UserAuth{
		Username:   "test",
		Password:   "test",
		HashedPass: "pass123",
	}
}

type ActivityFieldMother struct{}

func (m ActivityFieldMother) Default() domain.ActivityField {
	return domain.ActivityField{
		Name:        "aaa",
		Description: "aaa",
		Cost:        0.3,
	}
}

func (m ActivityFieldMother) WithoutName() domain.ActivityField {
	return domain.ActivityField{
		Description: "aaa",
		Cost:        0.3,
	}
}

type CompanyMother struct{}

func (m CompanyMother) Default() domain.Company {
	return domain.Company{
		Name:            "aaa",
		City:            "aaa",
		ActivityFieldId: uuid.UUID{0},
		OwnerID:         uuid.UUID{0},
	}
}

type FinReportMother struct{}

func (m FinReportMother) ForBigPeriod(startYear, startQuarter, endYear, endQuarter int, revenues, costs []float32) []domain.FinancialReport {
	reps := make([]domain.FinancialReport, 0)
	curQuarter := startQuarter
	curYear := startYear
	i := 0
	for {
		reps = append(reps, domain.FinancialReport{
			CompanyID: uuid.UUID{byte(i + 1)},
			Revenue:   revenues[i],
			Costs:     costs[i],
			Year:      curYear,
			Quarter:   curQuarter,
		})
		if curYear == endYear && curQuarter == endQuarter {
			break
		}
		curQuarter++
		if curQuarter > 4 {
			curQuarter = 0
			curYear++
		}
	}

	return reps
}
