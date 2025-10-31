package dto

import (
	"github.com/google/uuid"
)

type ContactInformationCreateDTO struct {
	AccountID   uuid.UUID
	ContactInfo string `json:"info"`
	Type        string `json:"type"`
	Publicity   bool   `json:"is_public"`
}

func (c ContactInformationCreateDTO) Validate() error {
	//TODO:ContactInformationCreateRequestValidation
	return nil
}

func (c *ContactInformationCreateDTO) SetAccountID(id uuid.UUID) {
	c.AccountID = id
}
