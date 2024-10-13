package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"math"
	"ppo/domain"
	"ppo/internal/storage"
	"strings"
)

type FinReportRepository struct {
	db storage.DBConn
}

func NewFinReportRepository(db storage.DBConn) domain.IFinancialReportRepository {
	return &FinReportRepository{
		db: db,
	}
}

func (r *FinReportRepository) Create(ctx context.Context, finReport *domain.FinancialReport) (report *domain.FinancialReport, err error) {
	query := `insert into ppo.fin_reports(company_id, revenue, costs, year, quarter) 
	values ($1, $2, $3, $4, $5) returning id`

	var id uuid.UUID
	err = r.db.QueryRow(
		ctx,
		query,
		finReport.CompanyID,
		finReport.Revenue,
		finReport.Costs,
		finReport.Year,
		finReport.Quarter,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("создание финансового отчета: %w", err)
	}
	finReport.ID = id

	return finReport, nil
}

func (r *FinReportRepository) GetById(ctx context.Context, id uuid.UUID) (report *domain.FinancialReport, err error) {
	query := `select company_id, revenue, costs, year, quarter from ppo.fin_reports where id = $1`

	report = new(domain.FinancialReport)
	err = r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&report.CompanyID,
		&report.Revenue,
		&report.Costs,
		&report.Year,
		&report.Quarter,
	)
	if err != nil {
		return nil, fmt.Errorf("получение отчета по id: %w", err)
	}

	report.ID = id
	return report, nil
}

func (r *FinReportRepository) GetByCompany(ctx context.Context, companyId uuid.UUID, period *domain.Period) (report *domain.FinancialReportByPeriod, err error) {
	query := `select id, company_id, revenue, costs, year, quarter
	from ppo.fin_reports 
	where company_id = $1 and year = $2 and quarter = $3`

	report = new(domain.FinancialReportByPeriod)
	report.Reports = make([]domain.FinancialReport, 0)

	for year := period.StartYear; year <= period.EndYear; year++ {
		startQtr := 1
		endQtr := 4

		if year == period.StartYear {
			startQtr = period.StartQuarter
		}
		if year == period.EndYear {
			endQtr = period.EndQuarter
		}

		for quarter := startQtr; quarter <= endQtr; quarter++ {
			tmp := new(domain.FinancialReport)

			err = r.db.QueryRow(
				ctx,
				query,
				companyId,
				year,
				quarter,
			).Scan(
				&tmp.ID,
				&tmp.CompanyID,
				&tmp.Revenue,
				&tmp.Costs,
				&tmp.Year,
				&tmp.Quarter,
			)

			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					continue
				} else {
					return nil, fmt.Errorf("сканирование записи: %w", err)
				}
			}

			report.Reports = append(report.Reports, *tmp)
		}
	}

	report.Period = period

	return report, nil
}

func (r *FinReportRepository) Update(ctx context.Context, finRep *domain.FinancialReport) (err error) {
	query := `update ppo.fin_reports set `

	args := make([]any, 0)
	i := 1
	equals := make([]string, 0)
	if finRep.CompanyID.ID() != 0 {
		equals = append(equals, fmt.Sprintf("company_id = $%d", i))
		i++
		args = append(args, finRep.CompanyID)
	}
	if math.Abs(float64(finRep.Revenue)) > 0 {
		equals = append(equals, fmt.Sprintf("revenue = $%d", i))
		i++
		args = append(args, finRep.Revenue)
	}
	if math.Abs(float64(finRep.Costs)) > 0 {
		equals = append(equals, fmt.Sprintf("costs = $%d", i))
		i++
		args = append(args, finRep.Costs)
	}
	if finRep.Year != 0 {
		equals = append(equals, fmt.Sprintf("year = $%d", i))
		i++
		args = append(args, finRep.Year)
	}
	if finRep.Quarter != 0 {
		equals = append(equals, fmt.Sprintf("quarter = $%d", i))
		i++
		args = append(args, finRep.Quarter)
	}

	query += strings.Join(equals, ", ")
	query += fmt.Sprintf(" where id = $%d", i)
	args = append(args, finRep.ID)

	_, err = r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о финансовом отчете: %w", err)
	}

	return nil
}

func (r *FinReportRepository) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	query := `delete from ppo.fin_reports where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("удаление отчета по id: %w", err)
	}

	return nil
}
