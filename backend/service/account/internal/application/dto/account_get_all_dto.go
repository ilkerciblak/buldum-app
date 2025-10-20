package dto

import (
	"strconv"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type GetAllAccountDTO struct {
	coredomain.CommonQueryParameters
	repository.ProfileGetAllQueryFilter
}

func (d *GetAllAccountDTO) SetLimit(limit string) *GetAllAccountDTO {
	coredomain.SetLimit(limit)(&d.CommonQueryParameters)
	return d
}

func (d *GetAllAccountDTO) SetPage(page string) *GetAllAccountDTO {
	coredomain.SetPage(page)(&d.CommonQueryParameters)
	return d
}

func (d *GetAllAccountDTO) SetOrder(order string) *GetAllAccountDTO {
	coredomain.SetOrder(order)(&d.CommonQueryParameters)
	return d
}

func (d *GetAllAccountDTO) SetSortBy(sort string) *GetAllAccountDTO {
	whiteList := map[string]bool{
		"user_name":  true,
		"created_at": true,
		"updated_at": true,
		"id":         true,
	}

	coredomain.SetSortBy(sort, whiteList)(&d.CommonQueryParameters)
	return d
}

func (d *GetAllAccountDTO) SetUsername(username string) *GetAllAccountDTO {
	d.Username = username
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
		CommonQueryParameters:    *coredomain.NewCommonQueryParameters(),
		ProfileGetAllQueryFilter: *repository.DefaultAccountGetAllQueryFilter(),
	}
}
