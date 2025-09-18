package model

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	Id         uuid.UUID
	UserID     uuid.UUID
	Username   string
	AvatarUrl  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	IsArchived bool
}

func NewProfile(username, avatarUrl string) *Profile {
	return &Profile{
		Id:         uuid.Must(uuid.NewV7()),
		UserID:     uuid.Must(uuid.NewV7()),
		Username:   username,
		AvatarUrl:  avatarUrl,
		CreatedAt:  time.Now(),
		IsArchived: false,
	}
}

// func (u *Profile) UpdateUserName(username string)
