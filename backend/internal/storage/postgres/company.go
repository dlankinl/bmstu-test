package postgres

import (
	"context"
	"fmt"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/internal/storage"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CompanyRepository struct {
	db storage.DBConn
}

func NewCompanyRepository(db storage.DBConn) domain.ICompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) Create(ctx context.Context, company *domain.Company) (comp *domain.Company, err error) {
	query := `insert into ppo.companies(owner_id, activity_field_id, name, city) 
	values ($1, $2, $3, $4) returning id`

	var id uuid.UUID
	err = r.db.QueryRow(
		ctx,
		query,
		company.OwnerID,
		company.ActivityFieldId,
		company.Name,
		company.City,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("создание компании: %w", err)
	}
	company.ID = id

	return company, nil
}

func (r *CompanyRepository) GetById(ctx context.Context, id uuid.UUID) (company *domain.Company, err error) {
	query := `select owner_id, activity_field_id, name, city from ppo.companies where id = $1`

	company = new(domain.Company)
	err = r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&company.OwnerID,
		&company.ActivityFieldId,
		&company.Name,
		&company.City,
	)
	if err != nil {
		return nil, fmt.Errorf("получение компании по id: %w", err)
	}
	company.ID = id

	return company, nil
}

func (r *CompanyRepository) GetByOwnerId(ctx context.Context, id uuid.UUID, page int, isPaginated bool) (companies []*domain.Company, numPages int, err error) {
	query :=
		`select 
    		id, 
    		activity_field_id,
    		name,
    		city 
		from ppo.companies 
		where owner_id = $1`

	var rows pgx.Rows
	if !isPaginated {
		rows, err = r.db.Query(
			ctx,
			query,
			id,
		)
	} else {
		rows, err = r.db.Query(
			ctx,
			query+` offset $2 limit $3`,
			id,
			(page-1)*config.PageSize,
			config.PageSize,
		)
	}
	if err != nil {
		return nil, 0, fmt.Errorf("получение компаний: %w", err)
	}

	companies = make([]*domain.Company, 0)
	for rows.Next() {
		tmp := new(domain.Company)

		err = rows.Scan(
			&tmp.ID,
			&tmp.ActivityFieldId,
			&tmp.Name,
			&tmp.City,
		)
		tmp.OwnerID = id

		if err != nil {
			return nil, 0, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		companies = append(companies, tmp)
	}

	var numRecords int
	err = r.db.QueryRow(
		ctx,
		`select count(*) from ppo.companies where owner_id = $1`,
		id,
	).Scan(&numRecords)
	if err != nil {
		return nil, 0, fmt.Errorf("получение списка компаний предпринимателя: %w", err)
	}

	numPages = numRecords / config.PageSize
	if numRecords%config.PageSize != 0 {
		numPages++
	}

	return companies, numPages, nil
}

func (r *CompanyRepository) Update(ctx context.Context, company *domain.Company) (err error) {
	query := "update ppo.companies set "

	args := make([]any, 0)
	i := 1
	equals := make([]string, 0)
	if company.OwnerID.ID() != 0 {
		equals = append(equals, fmt.Sprintf("owner_id = $%d", i))
		i++
		args = append(args, company.OwnerID)
	}
	if company.ActivityFieldId.ID() != 0 {
		equals = append(equals, fmt.Sprintf("activity_field_id = $%d", i))
		i++
		args = append(args, company.ActivityFieldId)
	}
	if company.Name != "" {
		equals = append(equals, fmt.Sprintf("name = $%d", i))
		i++
		args = append(args, company.Name)
	}
	if company.City != "" {
		equals = append(equals, fmt.Sprintf("city = $%d", i))
		i++
		args = append(args, company.City)
	}
	query += strings.Join(equals, ", ")
	query += fmt.Sprintf(" where id = $%d", i)
	args = append(args, company.ID)

	_, err = r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о компании: %w", err)
	}

	return nil
}

func (r *CompanyRepository) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("открытие транзакции: %w", err)
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = fmt.Errorf("обработанная ошибка: %w\nоткат транзакции: %v", err, rollbackErr)
			}
		}
	}()

	_, err = tx.Exec(
		ctx,
		`delete from ppo.companies where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("удаление компании по id: %w", err)
	}

	_, err = tx.Exec(
		ctx,
		`delete from ppo.fin_reports where company_id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("удаление отчетов, связанных с компанией: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("закрытие транзакции: %w", err)
	}

	return nil
}

func (r *CompanyRepository) GetAll(ctx context.Context, page int) (companies []*domain.Company, err error) {
	query := `select id, owner_id, activity_field_id, name, city from ppo.companies offset $1 limit $2`

	rows, err := r.db.Query(
		ctx,
		query,
		(page-1)*config.PageSize,
		config.PageSize,
	)
	if err != nil {
		return nil, fmt.Errorf("получение списка компаний: %w", err)
	}

	companies = make([]*domain.Company, 0)
	for rows.Next() {
		tmp := new(domain.Company)

		err = rows.Scan(
			&tmp.ID,
			&tmp.OwnerID,
			&tmp.ActivityFieldId,
			&tmp.Name,
			&tmp.City,
		)

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
		companies = append(companies, tmp)
	}

	return companies, nil
}
