package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

type ContactInformationRepositoryInterface interface {
	GetById(id uuid.UUID, ctx context.Context) (*model.ContactInformation, error)
	GetAll(filter ContactInformationQueryFilter, ctx context.Context) ([]*model.ContactInformation, error)
	CreateCI(model model.ContactInformation, ctx context.Context) error
	UpdateCI(id uuid.UUID, updated model.ContactInformation, ctx context.Context) error
	ArchiveCI(id uuid.UUID, ctx context.Context) error
}

type ContactInformationQueryFilter struct {
	UserID     uuid.UUID
	Type       string
	Publicity  bool
	IsArchived bool
}

func DefaultContactInformationQueryFilter() *ContactInformationQueryFilter {
	return &ContactInformationQueryFilter{
		Publicity:  true,
		IsArchived: false,
	}
}

func (c *ContactInformationQueryFilter) SetUserID(id uuid.UUID) {
	c.UserID = id
}

func (c *ContactInformationQueryFilter) SetType(s string) {
	c.Type = model.ContactInformationTypeFromString(s).String()
}

func (c *ContactInformationQueryFilter) SetPublicity(b bool) {
	c.Publicity = b
}

func (c *ContactInformationQueryFilter) SetIsArchived(b bool) {
	c.IsArchived = b
}
