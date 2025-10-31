package dto

import (
	"github.com/google/uuid"
)

type ContactInformationUpdateDTO struct {
	AccountID   uuid.UUID
	ContactInfo string `json:"info"`
	Publicity   bool   `json:"is_public"`
}

func (c ContactInformationUpdateDTO) Validate() error {
	//TODO:ContactInformationUpdateRequestValidation
	return nil
}

func (c *ContactInformationUpdateDTO) SetAccountID(id uuid.UUID) {
	c.AccountID = id
}
