package coredomain

import (
	"strconv"
)

type CommonQueryParameters struct {
	Pagination
	Sorting
}

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

type Sorting struct {
	Sort  string
	Order string
}

func (p *CommonQueryParameters) setpage(page string) {
	if parsed, err := strconv.Atoi(page); err == nil {
		p.Page = parsed
	}

	// //return p
}

func (p *CommonQueryParameters) setlimit(limit string) {
	if parsed, err := strconv.Atoi(limit); err == nil {
		p.Limit = parsed
	}

	//return p
}

func (p *CommonQueryParameters) SetPagination(limit string, page string) {

	p.setlimit(limit)
	p.setpage(page)

	if p.Page > 1 {
		p.Offset = p.Limit * (p.Page - 1)
	}
}

func (p *CommonQueryParameters) SetSortBy(sortBy string, sortingWhiteList map[string]bool) {
	if sortingWhiteList[sortBy] {
		p.Sort = sortBy
	}
	//return p
}

func (p *CommonQueryParameters) SetOrder(orderBy string) {

	orderWhiteList := map[string]bool{
		"asc":  true,
		"ASC":  true,
		"desc": true,
		"DESC": true,
	}

	if orderWhiteList[orderBy] {
		p.Order = orderBy
	}
	//return p
}

func DefaultCommonQueryParameters() *CommonQueryParameters {
	return &CommonQueryParameters{
		Pagination: Pagination{
			Page:   1,
			Limit:  30,
			Offset: 0,
		},
		Sorting: Sorting{
			Sort:  "created_at",
			Order: "asc",
		},
	}

}
