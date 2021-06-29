package storage

import (
	"19-db-app/pkg/movies"
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgStorage struct {
	connString string
}

func New(connString string) *PgStorage {
	var pg PgStorage
	pg.connString = connString
	return &pg
}

func (pg *PgStorage) connect() (context.Context, *pgxpool.Pool, error) {
	ctx := context.Background()
	db, err := pgxpool.Connect(context.Background(), pg.connString)
	if err != nil {
		return nil, nil, err
	}
	return ctx, db, nil
}

func (pg *PgStorage) GetMovies(companyId ...int) ([]movies.Movie, error) {
	cId := 0

	if len(companyId) > 0 {
		cId = companyId[0]
	}

	ctx, db, err := pg.connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var rows pgx.Rows
	if cId > 0 {
		sqlQuery := `SELECT id, name, release_year, gross, rating, company_id FROM movies WHERE company_id = $1`
		rows, err = db.Query(ctx, sqlQuery, cId)
	} else {
		sqlQuery := `SELECT id, name, release_year, gross, rating, company_id FROM movies`
		rows, err = db.Query(ctx, sqlQuery)
	}
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

func (pg *PgStorage) AddMovies(movies []movies.Movie) error {
	ctx, db, err := pg.connect()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin(ctx)
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

func (pg *PgStorage) ChangeMovie(ID int, m movies.Movie) error {
	ctx, db, err := pg.connect()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlQuery := `UPDATE movies SET name=$1, release_year=$2, gross=$3, rating=$4, company_id=$5 where id=$6`
	_, err = db.Exec(ctx, sqlQuery, m.Name, m.ReleaseYear, m.Gross, m.Rating, m.CompanyID, ID)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PgStorage) DeleteMovie(ID int) error {
	ctx, db, err := pg.connect()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlQuery := `DELETE FROM movies WHERE id=$1`
	_, err = db.Exec(ctx, sqlQuery, ID)
	if err != nil {
		return err
	}

	return nil
}
