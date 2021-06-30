package storage

import (
	"context"
	"dbapp/pkg/movies"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgStorage struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PgStorage {
	db := PgStorage{
		pool: pool,
	}
	return &db
}


func (pg *PgStorage) DeleteMovie(ctx context.Context, ID int) error {

	sqlQuery := `DELETE FROM movies WHERE id=$1`
	_, err := pg.pool.Exec(ctx, sqlQuery, ID)
	if err != nil {
		return err
	}

	return nil
}


func (pg *PgStorage) GetMovies(ctx context.Context, companyId ...int) ([]movies.Movie, error) {
	cId := 0

	if len(companyId) > 0 {
		cId = companyId[0]
	}

	var (
		rows pgx.Rows
		err error
	)
	sqlQuery := `SELECT id, name, release_year, gross, rating, company_id FROM movies WHERE company_id = $1`
	rows, err = pg.pool.Query(ctx, sqlQuery, cId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []movies.Movie

	for rows.Next() {
		var m movies.Movie
		err := rows.Scan(&m.ID, &m.Name, &m.ReleaseYear, &m.Gross, &m.Rating, &m.CompanyID)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pg *PgStorage) AddMovies(ctx context.Context, movies []movies.Movie) error {

	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	batch := new(pgx.Batch)
	sqlQuery := `INSERT INTO movies (name, release_year, gross, rating, company_id) VALUES ($1, $2, $3, $4, $5)`

	for _, m := range movies {
		batch.Queue(sqlQuery, m.Name, m.ReleaseYear, m.Gross, m.Rating, m.CompanyID)
	}

	res := tx.SendBatch(ctx, batch)
	err = res.Close()
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (pg *PgStorage) ChangeMovie(ctx context.Context, ID int, m movies.Movie) error {

	sqlQuery := `UPDATE movies SET name=$1, release_year=$2, gross=$3, rating=$4, company_id=$5 where id=$6`
	_, err := pg.pool.Exec(ctx, sqlQuery, m.Name, m.ReleaseYear, m.Gross, m.Rating, m.CompanyID, ID)
	if err != nil {
		return err
	}

	return nil
}
