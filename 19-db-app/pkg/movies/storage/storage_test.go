package storage

import (
	"19-db-app/pkg/movies"
	"reflect"
	"testing"
)

func TestPgStorage_DeleteMovie(t *testing.T) {
	type fields struct {
		connString string
	}
	type args struct {
		ID int
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
				connString: "postgres://postgres:postgres@localhost:5434/movies_test",
			},
			args: args{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &PgStorage{
				connString: tt.fields.connString,
			}
			if err := pg.DeleteMovie(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPgStorage_GetMovies(t *testing.T) {
	type fields struct {
		connString string
	}
	type args struct {
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
				connString: "postgres://postgres:postgres@localhost:5434/movies_test",
			},
			args: args{
				companyId: []int{100},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test get movies with db in docker-compose with id company with one movie",
			fields: fields{
				connString: "postgres://postgres:postgres@localhost:5434/movies_test",
			},
			args: args{
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
				connString: tt.fields.connString,
			}
			got, err := pg.GetMovies(tt.args.companyId...)
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
		connString string
	}
	type args struct {
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
				connString: "postgres://postgres:postgres@localhost:5434/movies_test",
			},
			args: args{
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
				connString: "postgres://postgres:postgres@localhost:5434/movies_test",
			},
			args: args{
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
				connString: tt.fields.connString,
			}
			if err := pg.AddMovies(tt.args.movies); (err != nil) != tt.wantErr {
				t.Errorf("AddMovies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPgStorage_ChangeMovie(t *testing.T) {
	type fields struct {
		connString string
	}
	type args struct {
		ID int
		m  movies.Movie
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
				connString: "postgres://postgres:postgres@localhost:5434/movies_test",
			},
			args: args{
				ID: 7,
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
				connString: tt.fields.connString,
			}
			if err := pg.ChangeMovie(tt.args.ID, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("ChangeMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}