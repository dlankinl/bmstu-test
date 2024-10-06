package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"math"
	"ppo/domain"
	"ppo/internal/config"
	"ppo/internal/storage"
	"strings"
)

type ActivityFieldRepository struct {
	db storage.DBConn
}

func NewActivityFieldRepository(db storage.DBConn) domain.IActivityFieldRepository {
	return &ActivityFieldRepository{
		db: db,
	}
}

func (r *ActivityFieldRepository) Create(ctx context.Context, data *domain.ActivityField) (res *domain.ActivityField, err error) {
	query := `insert into ppo.activity_fields(name, description, cost) 
	values ($1, $2, $3) returning id`

	var id uuid.UUID
	err = r.db.QueryRow(
		ctx,
		query,
		data.Name,
		data.Description,
		data.Cost,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("создание сферы деятельности: %w", err)
	}
	data.ID = id

	//var id uuid.UUID
	//rows, err := r.db.Query(
	//	ctx,
	//	query,
	//	data.Name,
	//	data.Description,
	//	data.Cost,
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("создание сферы деятельности: %w", err)
	//}
	//
	//id, err = pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (id uuid.UUID, err error) {
	//	err = row.Scan(&id)
	//	if err != nil {
	//		return uuid.UUID{}, fmt.Errorf("cканирование возвращенного идентификатора: %w", err)
	//	}
	//
	//	return id, err
	//})
	//if err != nil {
	//	return nil, fmt.Errorf("collect one row: %w", err)
	//}
	//data.ID = id

	return data, nil
}

func (r *ActivityFieldRepository) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	query := `delete from ppo.activity_fields where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("удаление сферы деятельности по id: %w", err)
	}

	return nil
}

func (r *ActivityFieldRepository) Update(ctx context.Context, data *domain.ActivityField) (err error) {
	query := `update ppo.activity_fields set `

	args := make([]any, 0)
	i := 1
	equals := make([]string, 0)
	if data.Name != "" {
		equals = append(equals, fmt.Sprintf("name = $%d", i))
		i++
		args = append(args, data.Name)
	}
	if data.Description != "" {
		equals = append(equals, fmt.Sprintf("description = $%d", i))
		i++
		args = append(args, data.Description)
	}
	if math.Abs(float64(data.Cost)) > 0 {
		equals = append(equals, fmt.Sprintf("cost = $%d", i))
		i++
		args = append(args, data.Cost)
	}
	query += strings.Join(equals, ", ")
	query += fmt.Sprintf(" where id = $%d", i)
	args = append(args, data.ID)

	_, err = r.db.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о сфере деятельности: %w", err)
	}

	return nil
}

func (r *ActivityFieldRepository) GetById(ctx context.Context, id uuid.UUID) (field *domain.ActivityField, err error) {
	query := `select name, description, cost from ppo.activity_fields where id = $1`

	field = new(domain.ActivityField)
	err = r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&field.Name,
		&field.Description,
		&field.Cost,
	)
	if err != nil {
		return nil, fmt.Errorf("получение сферы деятельности по id: %w", err)
	}

	field.ID = id

	return field, nil
}

func (r *ActivityFieldRepository) GetMaxCost(ctx context.Context) (cost float32, err error) {
	query := `select max(cost)
		from ppo.activity_fields`

	err = r.db.QueryRow(
		ctx,
		query,
	).Scan(&cost)

	if err != nil {
		return 0, fmt.Errorf("получение максимального веса сферы деятельности: %w", err)
	}

	return cost, nil
}

func (r *ActivityFieldRepository) GetAll(ctx context.Context, page int, isPaginated bool) (fields []*domain.ActivityField, numPages int, err error) {
	query :=
		`select
   		id,
   		name,
   		description,
   		cost
		from ppo.activity_fields`

	var rows pgx.Rows
	if !isPaginated {
		rows, err = r.db.Query(
			ctx,
			query,
		)
	} else {
		rows, err = r.db.Query(
			ctx,
			query+` offset $1 limit $2`,
			(page-1)*config.PageSize,
			config.PageSize,
		)
	}
	if err != nil {
		return nil, 0, fmt.Errorf("получение сфер деятельности: %w", err)
	}

	fields = make([]*domain.ActivityField, 0)
	for rows.Next() {
		tmp := new(domain.ActivityField)

		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Description,
			&tmp.Cost,
		)

		if err != nil {
			return nil, 0, fmt.Errorf("сканирование полученных строк: %w", err)
		}

		fields = append(fields, tmp)
	}

	var numRecords int
	err = r.db.QueryRow(
		ctx,
		`select count(*) from ppo.activity_fields`,
	).Scan(&numRecords)
	if err != nil {
		return nil, 0, fmt.Errorf("получение числа сфер деятельности: %w", err)
	}

	numPages = numRecords / config.PageSize
	if numRecords%config.PageSize != 0 {
		numPages++
	}

	return fields, numPages, nil
}
