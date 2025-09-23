package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountGetByIdQuery struct {
	Id uuid.UUID `path:"id"`
}

func NewAccountGetByIdQuery(m map[string]string) (*AccountGetByIdQuery, error) {

	userId, err := uuid.Parse(m["id"])
	if err != nil {
		return nil, err
	}

	return &AccountGetByIdQuery{
		Id: userId,
	}, nil
}

func (a AccountGetByIdQuery) Handler(r repository.AccountRepository, ctx context.Context) (*model.Profile, coredomain.IApplicationError) {
	data, err := r.GetById(ctx, a.Id)
	if err != nil {
		return nil, err.(coredomain.IApplicationError)
	}

	return data, nil
}
