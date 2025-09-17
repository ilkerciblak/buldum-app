package model

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	UserID    uuid.UUID
	Username  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	NotificationSettings
}

func NewProfile(username, avatarUrl string, c ...ConfigureNotifications) *Profile {
	return &Profile{
		UserID:               uuid.Must(uuid.NewV7()),
		Username:             username,
		AvatarUrl:            avatarUrl,
		CreatedAt:            time.Now(),
		NotificationSettings: *NewNotificationSettings(c...),
	}
}

// func (u *Profile) UpdateUserName(username string)
