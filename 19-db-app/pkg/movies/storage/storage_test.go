package storage

import (
	"context"
	"dbapp/pkg/movies"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"reflect"
	"testing"
)

var (
	pgSt             *PgStorage
	ctx              = context.Background()
	connectionString = "postgres://postgres:postgres@localhost:5434/movies_test"
)

func TestMain(m *testing.M) {
	fmt.Println("Testing is running...")
	db, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	pgSt = New(db)
	defer db.Close()
	os.Exit(m.Run())
}

func TestPgStorage_GetMovies(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		ctx       context.Context
		companyId []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []movies.Movie
		wantErr bool
	}{
		{
			name: "Test get movies with db in docker-compose with id company without movies",
			fields: fields{
				pool: pgSt.pool,
			},
			args: args{
				ctx:       ctx,
				companyId: []int{100},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test get movies with db in docker-compose with id company with one movie",
			fields: fields{
				pool: pgSt.pool,
			},
			args: args{
				ctx:       ctx,
				companyId: []int{3},
			},
			want: []movies.Movie{
				{
					ID:          7,
					Name:        "Back to the Future Part II",
					ReleaseYear: 1989,
					Gross:       331950002,
					Rating:      "PG-10",
					CompanyID:   3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &PgStorage{
				pool: tt.fields.pool,
			}
			got, err := pg.GetMovies(tt.args.ctx, tt.args.companyId...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgStorage_AddMovies(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		movies []movies.Movie
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test add movies with db in docker-compose",
			fields: fields{
				pool: pgSt.pool,
			},
			args: args{
				ctx: ctx,
				movies: []movies.Movie{
					{
						Name:        "Test",
						ReleaseYear: 2014,
						Gross:       12,
						Rating:      "PG-13",
						CompanyID:   3},
				},
			},
			wantErr: false,
		},
		{
			name: "Test add exist movies with db in docker-compose",
			fields: fields{
				pool: pgSt.pool,
			},
			args: args{
				ctx: ctx,
				movies: []movies.Movie{
					{
						Name:        "Back to the Future Part II",
						ReleaseYear: 1989,
						Gross:       0,
						Rating:      "PG-13",
						CompanyID:   1,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &PgStorage{
				pool: tt.fields.pool,
			}
			if err := pg.AddMovies(tt.args.ctx, tt.args.movies); (err != nil) != tt.wantErr {
				t.Errorf("AddMovies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPgStorage_ChangeMovie(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		ID  int
		m   movies.Movie
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test rename movies name with db in docker-compose",
			fields: fields{
				pool: pgSt.pool,
			},
			args: args{
				ctx: ctx,
				ID:  7,
				m: movies.Movie{
					Name:        "Back to the Future Part 4",
					ReleaseYear: 2021,
					Gross:       331950002,
					Rating:      "PG-13",
					CompanyID:   2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &PgStorage{
				pool: tt.fields.pool,
			}
			if err := pg.ChangeMovie(tt.args.ctx, tt.args.ID, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("ChangeMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPgStorage_DeleteMovie(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test delete with db in docker-compose",
			fields: fields{
				pool: pgSt.pool,
			},
			args: args{
				ctx: ctx,
				ID:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &PgStorage{
				pool: tt.fields.pool,
			}
			if err := pg.DeleteMovie(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
