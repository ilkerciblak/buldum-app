package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountServiceInterface interface {
	CreateAccount(u model.Profile, ctx context.Context) error
	UpdateAccount(u model.Profile, ctx context.Context) error
	ArchiveAccount(userId uuid.UUID, ctx context.Context) error
	GetAccountById(userId uuid.UUID, ctx context.Context) (*model.Profile, error)
	GetAllAccount(query coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter, ctx context.Context) ([]*model.Profile, error)
}
