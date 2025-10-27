package dto

import "github.com/google/uuid"

type ContactInformationGetSingleDTO struct {
	AccountID uuid.UUID
	Type      string
}
