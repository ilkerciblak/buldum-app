package application

import "log"

type CommonQueryParameters struct {
	Pagination
	Sorting
}

func NewCommonQueryParameters(m map[string]any, allowedCommonQueries map[string]map[string]bool) (*CommonQueryParameters, error) {
	pagination, sorting := DefaultCommonQueryParameters()

	orderWhiteList := map[string]bool{
		"asc":  true,
		"ASC":  true,
		"desc": true,
		"DESC": true,
	}

	if page := m["page"]; page != nil && page != 0 {
		pagination.Page = page.(int)
	}

	if limit := m["limit"]; limit != nil && limit != 0 {
		pagination.Limit = limit.(int)
	}

	if pagination.Page > 1 {
		pagination.Offset = pagination.Limit * (pagination.Page - 1)
	}

	if sortBy := m["sort"]; sortBy != nil && sortBy != "" {
		if allowedCommonQueries["sort"][sortBy.(string)] {
			sorting.Sort = sortBy.(string)
		}
		log.Printf("Sorting Query Not in the whitelist: %v", sortBy)
	}

	if order := m["order"]; order != nil && order != "" && orderWhiteList[order.(string)] {
		sorting.Order = order.(string)
	}

	return &CommonQueryParameters{
		Pagination: *pagination,
		Sorting:    *sorting,
	}, nil
}

func DefaultCommonQueryParameters() (*Pagination, *Sorting) {
	return &Pagination{
			Page:   1,
			Limit:  30,
			Offset: 0,
		},
		&Sorting{
			Sort:  "created_at",
			Order: "asc",
		}

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
