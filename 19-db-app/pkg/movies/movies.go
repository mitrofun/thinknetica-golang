package movies

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
