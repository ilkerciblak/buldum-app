package dto

import (
	"strconv"
	"strings"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type GetAllAccountDTO struct {
	coredomain.CommonQueryParameters
	repository.ProfileGetAllQueryFilter
}

func (d *GetAllAccountDTO) SetSortBy(sort string) *GetAllAccountDTO {
	whiteList := map[string]bool{
		"user_name":  true,
		"created_at": true,
		"updated_at": true,
		"id":         true,
	}

	d.CommonQueryParameters.SetSortBy(sort, whiteList)
	return d
}

func (d *GetAllAccountDTO) SetUsername(username string) *GetAllAccountDTO {
	if len(strings.Trim(username, " ")) > 0 {
		d.Username = username
	}
	return d
}

func (d *GetAllAccountDTO) SetIsArchived(isArchived string) *GetAllAccountDTO {
	if parsed, err := strconv.ParseBool(isArchived); err == nil {
		d.IsArchived = parsed
	}

	return d
}

func NewGetAllAccountDTO() *GetAllAccountDTO {
	return &GetAllAccountDTO{
		CommonQueryParameters:    *coredomain.DefaultCommonQueryParameters(),
		ProfileGetAllQueryFilter: *repository.DefaultAccountGetAllQueryFilter(),
	}
}
