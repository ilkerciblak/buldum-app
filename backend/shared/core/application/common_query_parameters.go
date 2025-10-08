package application

import (
	"strconv"
)

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

type WithParamFunc func(commonQueryParameters *CommonQueryParameters) *CommonQueryParameters

func SetPage(page string) WithParamFunc {
	return func(commonQueryParameters *CommonQueryParameters) *CommonQueryParameters {
		if parsed, err := strconv.Atoi(page); err == nil {
			commonQueryParameters.Page = parsed
		}

		return commonQueryParameters
	}
}

func SetLimit(limit string) WithParamFunc {
	return func(commonQueryParameters *CommonQueryParameters) *CommonQueryParameters {
		if parsed, err := strconv.Atoi(limit); err == nil {
			commonQueryParameters.Limit = parsed
		}
		if commonQueryParameters.Page > 1 {
			commonQueryParameters.Offset = commonQueryParameters.Limit * (commonQueryParameters.Page - 1)
		}
		return commonQueryParameters
	}
}

func SetSortBy(sortBy string, sortingWhiteList map[string]bool) WithParamFunc {
	return func(commonQueryParameters *CommonQueryParameters) *CommonQueryParameters {
		if len(sortingWhiteList) == 0 || sortingWhiteList[sortBy] {
			commonQueryParameters.Sort = sortBy
		}
		return commonQueryParameters
	}
}

func SetOrder(orderBy string) WithParamFunc {

	orderWhiteList := map[string]bool{
		"asc":  true,
		"ASC":  true,
		"desc": true,
		"DESC": true,
	}

	return func(commonQueryParameters *CommonQueryParameters) *CommonQueryParameters {
		if orderWhiteList[orderBy] {
			commonQueryParameters.Order = orderBy
		}
		return commonQueryParameters
	}
}

func NewCommonQueryParameters(setters ...WithParamFunc) *CommonQueryParameters {
	pagination, sorting := defaultCommonQueryParameters()

	c := &CommonQueryParameters{
		Pagination: *pagination,
		Sorting:    *sorting,
	}

	for _, setter := range setters {
		setter(c)
	}

	return c
}

func defaultCommonQueryParameters() (*Pagination, *Sorting) {
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
