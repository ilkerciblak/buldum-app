package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

type ContactInformationRepositoryInterface interface {
	GetAllByAccountId(accountId uuid.UUID, ctx context.Context) []*model.ContactInformation
	GetByAccountAndType(accountId uuid.UUID, t model.ContactInformationType, ctx context.Context) (*model.ContactInformation, error)
	CreateCI(model model.ContactInformation, ctx context.Context) error
	UpdateCI(id uuid.UUID, updated model.ContactInformation, ctx context.Context) error
	ArchiveCI(id uuid.UUID, ctx context.Context) error
}
