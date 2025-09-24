package application

type CommonQueryParameters struct {
	Pagination
	Sorting
}

func NewCommonQueryParameters(m map[string]any) (*CommonQueryParameters, error) {
	pagination := &Pagination{}
	sorting := &Sorting{}

	if page := m["page"]; page != nil && page != 0 {
		pagination.Page = page.(int)
	}

	if limit := m["limit"]; limit != nil && limit != 0 {
		pagination.Limit = limit.(int)
	}
	if offset := m["limit"]; offset != nil && offset != 0 {
		pagination.Limit = offset.(int)
	}

	if m["sort"] != nil && m["sort"] != "" {
		sorting.Sort = m["sort"].(string)
	}

	return &CommonQueryParameters{
		Pagination: *pagination,
		Sorting:    *sorting,
	}, nil
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
