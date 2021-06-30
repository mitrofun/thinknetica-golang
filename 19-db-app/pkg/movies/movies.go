package movies

import "context"

type Movie struct {
	ID          int
	Name        string
	ReleaseYear int
	Gross       int
	Rating      string
	CompanyID   int
}

type Company struct {
	ID   int
	Name string
}

type IMovie interface {
	GetMovies(context.Context, ...int) ([]Movie, error)
	AddMovies(context.Context, []Movie) error
	ChangeMovie(context.Context, int, Movie) error
	DeleteMovie(context.Context, int) error
}
