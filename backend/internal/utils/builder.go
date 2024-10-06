package utils

import (
	"github.com/google/uuid"
	"ppo/domain"
	"time"
)

type userAuthBuilder struct {
	userAuth domain.UserAuth
}

func NewUserAuthBuilder() userAuthBuilder {
	return userAuthBuilder{
		userAuth: domain.UserAuth{},
	}
}

func (b userAuthBuilder) WithUsername(username string) userAuthBuilder {
	b.userAuth.Username = username
	return b
}

func (b userAuthBuilder) WithPassword(password string) userAuthBuilder {
	b.userAuth.Password = password
	return b
}

func (b userAuthBuilder) WithHashedPass(hashedPass string) userAuthBuilder {
	b.userAuth.HashedPass = hashedPass
	return b
}

func (b userAuthBuilder) WithRole(role string) userAuthBuilder {
	b.userAuth.Role = role
	return b
}

func (b userAuthBuilder) WithID(id uuid.UUID) userAuthBuilder {
	b.userAuth.ID = id
	return b
}

func (b userAuthBuilder) Build() domain.UserAuth {
	return b.userAuth
}

type activityFieldBuilder struct {
	actField domain.ActivityField
}

func NewActivityFieldBuilder() activityFieldBuilder {
	return activityFieldBuilder{
		actField: domain.ActivityField{},
	}
}

func (b activityFieldBuilder) WithName(name string) activityFieldBuilder {
	b.actField.Name = name
	return b
}

func (b activityFieldBuilder) WithDescription(description string) activityFieldBuilder {
	b.actField.Description = description
	return b
}

func (b activityFieldBuilder) WithCost(cost float32) activityFieldBuilder {
	b.actField.Cost = cost
	return b
}

func (b activityFieldBuilder) WithID(id uuid.UUID) activityFieldBuilder {
	b.actField.ID = id
	return b
}

func (b activityFieldBuilder) Build() domain.ActivityField {
	return b.actField
}

type companyBuilder struct {
	company domain.Company
}

func NewCompanyBuilder() companyBuilder {
	return companyBuilder{
		company: domain.Company{},
	}
}

func (b companyBuilder) WithName(name string) companyBuilder {
	b.company.Name = name
	return b
}

func (b companyBuilder) WithCity(city string) companyBuilder {
	b.company.City = city
	return b
}

func (b companyBuilder) WithActivityField(actField uuid.UUID) companyBuilder {
	b.company.ActivityFieldId = actField
	return b
}

func (b companyBuilder) WithOwner(owner uuid.UUID) companyBuilder {
	b.company.OwnerID = owner
	return b
}

func (b companyBuilder) WithID(id uuid.UUID) companyBuilder {
	b.company.ID = id
	return b
}

func (b companyBuilder) Build() domain.Company {
	return b.company
}

type finReportBuilder struct {
	finReport domain.FinancialReport
}

func NewFinReportBuilder() finReportBuilder {
	return finReportBuilder{
		finReport: domain.FinancialReport{},
	}
}

func (b finReportBuilder) WithID(id uuid.UUID) finReportBuilder {
	b.finReport.ID = id
	return b
}

func (b finReportBuilder) WithCompanyID(id uuid.UUID) finReportBuilder {
	b.finReport.CompanyID = id
	return b
}

func (b finReportBuilder) WithRevenue(revenue float32) finReportBuilder {
	b.finReport.Revenue = revenue
	return b
}

func (b finReportBuilder) WithCosts(costs float32) finReportBuilder {
	b.finReport.Costs = costs
	return b
}

func (b finReportBuilder) WithYear(year int) finReportBuilder {
	b.finReport.Year = year
	return b
}

func (b finReportBuilder) WithQuarter(quarter int) finReportBuilder {
	b.finReport.Quarter = quarter
	return b
}

func (b finReportBuilder) Build() domain.FinancialReport {
	return b.finReport
}

type periodBuilder struct {
	period domain.Period
}

func NewPeriodBuilder() periodBuilder {
	return periodBuilder{
		period: domain.Period{},
	}
}

func (b periodBuilder) WithStartYear(year int) periodBuilder {
	b.period.StartYear = year
	return b
}

func (b periodBuilder) WithStartQuarter(quarter int) periodBuilder {
	b.period.StartQuarter = quarter
	return b
}

func (b periodBuilder) WithEndYear(year int) periodBuilder {
	b.period.EndYear = year
	return b
}

func (b periodBuilder) WithEndQuarter(quarter int) periodBuilder {
	b.period.EndQuarter = quarter
	return b
}

func (b periodBuilder) Build() domain.Period {
	return b.period
}

type finReportByPeriodBuilder struct {
	report domain.FinancialReportByPeriod
}

func NewFinReportByPeriodBuilder() finReportByPeriodBuilder {
	return finReportByPeriodBuilder{
		report: domain.FinancialReportByPeriod{},
	}
}

func (b finReportByPeriodBuilder) WithReports(reports []domain.FinancialReport) finReportByPeriodBuilder {
	b.report.Reports = reports
	return b
}

func (b finReportByPeriodBuilder) WithPeriod(period domain.Period) finReportByPeriodBuilder {
	b.report.Period = &period
	return b
}

func (b finReportByPeriodBuilder) Build() domain.FinancialReportByPeriod {
	return b.report
}

type userBuilder struct {
	user domain.User
}

func NewUserBuilder() userBuilder {
	return userBuilder{
		user: domain.User{},
	}
}

func (b userBuilder) WithId(id uuid.UUID) userBuilder {
	b.user.ID = id
	return b
}

func (b userBuilder) WithUsername(username string) userBuilder {
	b.user.Username = username
	return b
}

func (b userBuilder) WithFullName(fullName string) userBuilder {
	b.user.FullName = fullName
	return b
}

func (b userBuilder) WithGender(gender string) userBuilder {
	b.user.Gender = gender
	return b
}

func (b userBuilder) WithBirthday(birthday time.Time) userBuilder {
	b.user.Birthday = birthday
	return b
}

func (b userBuilder) WithCity(city string) userBuilder {
	b.user.City = city
	return b
}

func (b userBuilder) WithRole(role string) userBuilder {
	b.user.Role = role
	return b
}

func (b userBuilder) Build() domain.User {
	return b.user
}
