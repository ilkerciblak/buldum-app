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
	GetAllFiltered(filter dto.ContactInformationGetAllDTO, ctx context.Context) ([]*dto.ContactInformationGetAllResultDTO, error)
}
