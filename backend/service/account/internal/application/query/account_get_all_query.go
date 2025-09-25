package query

import (
	"context"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/application"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountGetAllQuery struct {
	application.CommonQueryParameters
	repository.ProfileGetAllQueryFilter
}

func NewAccountGetAllQuery(m map[string]any) (*AccountGetAllQuery, error) {
	whiteList := map[string]map[string]bool{
		"sort": {
			"user_name":  true,
			"created_at": true,
			"updated_at": true,
			"id":         true,
		},
	}
	cqp, err := application.NewCommonQueryParameters(m, whiteList)
	if err != nil {
		return nil, err
	}
	filter, err := repository.NewAccountGetAllQueryFilter(m)
	if err != nil {
		return nil, err
	}

	return &AccountGetAllQuery{
		CommonQueryParameters:    *cqp,
		ProfileGetAllQueryFilter: *filter,
	}, nil
}

func (a AccountGetAllQuery) Handler(r repository.AccountRepository, ctx context.Context) ([]*model.Profile, coredomain.IApplicationError) {

	data, err := r.GetAll(ctx, a.CommonQueryParameters, a.ProfileGetAllQueryFilter)

	if err != nil {
		return nil, coredomain.BadRequest.WithMessage(err.Error())
	}

	return data, nil
}
