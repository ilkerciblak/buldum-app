package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
)

type ContactInformationServiceInterface interface {
	Create(dto.ContactInformationCreateDTO, context.Context) error
	Update(uuid.UUID, dto.ContactInformationUpdateDTO, context.Context) error
	Archive(uuid.UUID, context.Context) error
	GetSingleByAccountByType(accountId uuid.UUID, ciType string, ctx context.Context) (*dto.AccountResultDTO, error)
	GetAllByAccount(accountId uuid.UUID, ctx context.Context) []*dto.AccountResultDTO
}
