package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type ContactInformation struct {
	UserID    uuid.UUID
	Type      ContactInformationType
	Publicity bool
}

func NewContactInformation(userid uuid.UUID, ty ContactInformationType, public bool) *ContactInformation {
	return &ContactInformation{
		UserID:    userid,
		Type:      ty,
		Publicity: public,
	}
}

type ContactInformationType string

const (
	PhoneNumber ContactInformationType = "PhoneNumber"
	Email       ContactInformationType = "Email"
)

func (c *ContactInformationType) Validate() (*ContactInformationType, error) {
	switch *c {
	case Email:
		fallthrough
	case PhoneNumber:
		return c, nil
	default:
		return nil, coredomain.RequestValidationError.WithMessage(fmt.Sprintf("Invalid Contact Information Type %v", c))
	}
}
