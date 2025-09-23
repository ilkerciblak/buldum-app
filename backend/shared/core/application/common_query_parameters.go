package application

type CommonQueryParameters struct {
	Pagination
	Sorting
}

type Pagination struct {
	Page   int `query:"page"`
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type Sorting struct {
	Sort  string `query:"sort"`
	Order string `query:"order"`
}
