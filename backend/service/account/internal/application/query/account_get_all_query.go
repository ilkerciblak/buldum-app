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
}

func NewAccountGetAllQuery(m map[string]any) (*AccountGetAllQuery, error) {
	cqp, err := application.NewCommonQueryParameters(m)
	if err != nil {
		return nil, err
	}
	return &AccountGetAllQuery{
		CommonQueryParameters: *cqp,
	}, nil
}

func (a AccountGetAllQuery) Handler(r repository.AccountRepository, ctx context.Context) ([]*model.Profile, coredomain.IApplicationError) {

	data, err := r.GetAll(ctx, a.CommonQueryParameters)

	if err != nil {
		return nil, coredomain.BadRequest.WithMessage(err.Error())
	}

	return data, nil
}
