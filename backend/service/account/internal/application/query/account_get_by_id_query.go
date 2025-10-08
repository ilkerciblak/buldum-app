package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountGetByIdQuery struct {
	Id uuid.UUID
}

func NewAccountGetByIdQuery(id string) (*AccountGetByIdQuery, error) {
	userid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &AccountGetByIdQuery{
		Id: userid,
	}, nil
}

func (a AccountGetByIdQuery) Handler(r repository.IAccountRepository, ctx context.Context) (*model.Profile, coredomain.IApplicationError) {
	data, err := r.GetById(ctx, a.Id)
	if err != nil {
		return nil, err.(coredomain.IApplicationError)
	}

	return data, nil
}
