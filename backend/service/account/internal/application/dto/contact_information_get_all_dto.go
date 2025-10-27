package dto

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type ContactInformationGetAllByAccountDTO struct {
	AccountID  uuid.UUID
	IsArchived bool
	coredomain.CommonQueryParameters
}

func NewContactInformationGetAllByAccountDTO(accountId uuid.UUID) *ContactInformationGetAllByAccountDTO {
	return &ContactInformationGetAllByAccountDTO{
		AccountID:             accountId,
		IsArchived:            false,
		CommonQueryParameters: *coredomain.DefaultCommonQueryParameters(),
	}
}

func (d *ContactInformationGetAllByAccountDTO) SetSort(sort string) {
	whiteList := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"id":         true,
	}

	d.CommonQueryParameters.SetSortBy(sort, whiteList)

}

func (d *ContactInformationGetAllByAccountDTO) SetIsArchived(isArchived string) *ContactInformationGetAllByAccountDTO {
	if parsed, err := strconv.ParseBool(isArchived); err == nil {
		d.IsArchived = parsed
	}

	return d
}
