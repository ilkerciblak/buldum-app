package query

import (
	"context"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountGetAllQuery struct{}

func (a AccountGetAllQuery) Handler(r repository.AccountRepository, ctx context.Context) ([]*model.Profile, coredomain.IApplicationError) {
	data, err := r.GetAll(ctx)

	if err != nil {
		return nil, coredomain.BadRequest.WithMessage(err.Error())
	}

	return data, nil
}
