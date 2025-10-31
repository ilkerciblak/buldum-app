package dto

import (
	"strconv"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type ContactInformationGetAllDTO struct {
	repository.ContactInformationQueryFilter
	coredomain.CommonQueryParameters
}

func NewContactInformationGetAllByAccountDTO() *ContactInformationGetAllDTO {
	return &ContactInformationGetAllDTO{
		ContactInformationQueryFilter: *repository.DefaultContactInformationQueryFilter(),
		CommonQueryParameters:         *coredomain.DefaultCommonQueryParameters(),
	}
}

func (d *ContactInformationGetAllDTO) SetSort(sort string) {
	whiteList := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"id":         true,
	}

	d.CommonQueryParameters.SetSortBy(sort, whiteList)

}

func (d *ContactInformationGetAllDTO) SetIsArchived(isArchived string) *ContactInformationGetAllDTO {
	if parsed, err := strconv.ParseBool(isArchived); err == nil {
		d.IsArchived = parsed
	}

	return d
}
func (d *ContactInformationGetAllDTO) SetIsPublic(isPublic string) *ContactInformationGetAllDTO {
	if parsed, err := strconv.ParseBool(isPublic); err == nil {
		d.Publicity = parsed
	}

	return d
}
