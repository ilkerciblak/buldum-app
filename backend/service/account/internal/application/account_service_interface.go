package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountServiceInterface interface {
	CreateAccount(dto dto.AccountCreateDTO, ctx context.Context) error
	UpdateAccount(userId uuid.UUID, dto dto.AccountUpdateDTO, ctx context.Context) error
	ArchiveAccount(userId uuid.UUID, ctx context.Context) error
	GetAccountById(userId uuid.UUID, ctx context.Context) (*dto.AccountResultDTO, error)
	GetAllAccount(query coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter, ctx context.Context) ([]*dto.AccountResultDTO, error)
}
