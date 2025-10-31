package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type ContactInformationType int

const (
	Email ContactInformationType = iota
	PhoneNumber
)

var (
	typeNames = map[ContactInformationType]string{
		Email:       "Email",
		PhoneNumber: "PhoneNumber",
	}

	fromString = map[string]ContactInformationType{
		"Email":       Email,
		"Phone":       PhoneNumber,
		"PhoneNumber": PhoneNumber,
	}
)

func (t ContactInformationType) String() string {
	val, exists := typeNames[t]
	if !exists {
		return "Undefined"
	}

	return val
}

func ContactInformationTypeFromString(s string) ContactInformationType {
	return fromString[s]
}

type ContactInformation struct {
	Id          uuid.UUID
	UserID      uuid.UUID
	Type        ContactInformationType
	Publicity   bool
	ContactInfo string
	IsArchived  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewContactInformation(userId uuid.UUID, infoType ContactInformationType, publicity bool, contactInfo string) *ContactInformation {
	return &ContactInformation{
		Id:          uuid.New(),
		UserID:      userId,
		Type:        infoType,
		Publicity:   publicity,
		ContactInfo: contactInfo,
		CreatedAt:   time.Now(),
		IsArchived:  false,
	}
}

func (c *ContactInformation) SetContactInfo(i string) error {
	// TODO: Validation for each type
	c.ContactInfo = i
	return nil
}

func (c *ContactInformation) SetPublicity(p bool) {
	c.Publicity = p
}

func (c *ContactInformation) Archive() {
	c.IsArchived = true
}

func (c *ContactInformation) BeforeUpdate() error {
	if c.IsArchived {
		return coredomain.Conflict.WithMessage("Instance is archived")
	}

	c.UpdatedAt = time.Now()
	return nil
}
