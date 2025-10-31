package dto

import (
	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

type ContactInformationGetAllResultDTO struct {
	ID         uuid.UUID
	AccountId  uuid.UUID
	Type       string
	Info       string
	IsArchived bool
}

func FromModel(m *model.ContactInformation) *ContactInformationGetAllResultDTO {

	return &ContactInformationGetAllResultDTO{
		ID:         m.Id,
		AccountId:  m.UserID,
		Type:       m.Type.String(),
		Info:       m.ContactInfo,
		IsArchived: m.IsArchived,
	}
}

func FromModelListToList(m []*model.ContactInformation) []*ContactInformationGetAllResultDTO {
	res := make([]*ContactInformationGetAllResultDTO, len(m))

	for i, m := range m {
		res[i] = FromModel(m)
	}
	return res
}
