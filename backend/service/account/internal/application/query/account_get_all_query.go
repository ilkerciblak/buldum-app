package query

import (
	"context"
	"strconv"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountGetAllQuery struct {
	coredomain.CommonQueryParameters
	repository.ProfileGetAllQueryFilter
}

type WithParamFunc func(queryParams *AccountGetAllQuery) *AccountGetAllQuery

func SetLimit(limit string) WithParamFunc {
	return func(queryParams *AccountGetAllQuery) *AccountGetAllQuery {
		coredomain.SetLimit(limit)(&queryParams.CommonQueryParameters)
		return queryParams
	}
}

func SetPage(page string) WithParamFunc {
	return func(queryParams *AccountGetAllQuery) *AccountGetAllQuery {
		coredomain.SetPage(page)(&queryParams.CommonQueryParameters)
		return queryParams
	}
}

func SetOrderBy(orderBy string) WithParamFunc {
	return func(queryParams *AccountGetAllQuery) *AccountGetAllQuery {
		coredomain.SetOrder(orderBy)(&queryParams.CommonQueryParameters)
		return queryParams
	}
}

func SetSortBy(sortBy string) WithParamFunc {
	whiteList := map[string]bool{
		"user_name":  true,
		"created_at": true,
		"updated_at": true,
		"id":         true,
	}
	return func(queryParams *AccountGetAllQuery) *AccountGetAllQuery {
		coredomain.SetSortBy(sortBy, whiteList)(&queryParams.CommonQueryParameters)
		return queryParams
	}
}

func SetUsername(username string) WithParamFunc {
	return func(queryParams *AccountGetAllQuery) *AccountGetAllQuery {
		queryParams.Username = username
		return queryParams
	}
}
func SetIsArchived(isArchived string) WithParamFunc {
	return func(queryParams *AccountGetAllQuery) *AccountGetAllQuery {
		if k, err := strconv.ParseBool(isArchived); err == nil {
			queryParams.IsArchived = k
		}
		return queryParams
	}
}

func DefaultAccountGetAllQuery() *AccountGetAllQuery {
	return &AccountGetAllQuery{
		CommonQueryParameters:    *coredomain.NewCommonQueryParameters(),
		ProfileGetAllQueryFilter: *repository.DefaultAccountGetAllQueryFilter(),
	}
}

func NewAccountGetAllQuery(setters ...WithParamFunc) (*AccountGetAllQuery, error) {
	query := DefaultAccountGetAllQuery()

	for _, setter := range setters {
		setter(query)
	}

	return query, nil
}

func (a AccountGetAllQuery) Handler(r repository.IAccountRepository, ctx context.Context) ([]*model.Profile, coredomain.IApplicationError) {

	data, err := r.GetAll(ctx, a.CommonQueryParameters, a.ProfileGetAllQueryFilter)

	if err != nil {
		return nil, coredomain.BadRequest.WithMessage(err.Error())
	}

	return data, nil
}
