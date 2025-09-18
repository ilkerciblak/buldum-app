package model

import (
	"time"

	"github.com/google/uuid"
)

type NotificationSettings struct {
	Id                   uuid.UUID
	UserId               uuid.UUID
	SendEmail            bool
	SendPushNotification bool
	// SendLocationBasedNotifications bool
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
	IsArchived bool
}

type ConfigureNotifications func(n *NotificationSettings) *NotificationSettings

func AllowEmails(n *NotificationSettings) *NotificationSettings {
	n.SendEmail = true
	return n
}

func AllowPushNotifications(n *NotificationSettings) *NotificationSettings {
	n.SendPushNotification = true
	return n
}

func AllowAll(n *NotificationSettings) *NotificationSettings {
	n.SendEmail = true
	n.SendPushNotification = true
	return n
}

func SetUserId(userId uuid.UUID) func(n *NotificationSettings) *NotificationSettings {
	return func(n *NotificationSettings) *NotificationSettings {
		n.UserId = userId
		return n
	}
}

func defaultNotificationSettings() *NotificationSettings {
	return &NotificationSettings{
		Id:                   uuid.Must(uuid.NewV7()),
		SendEmail:            false,
		SendPushNotification: false,
		Created_at:           time.Now(),
		IsArchived:           false,
	}
}

func NewNotificationSettings(opt ...ConfigureNotifications) *NotificationSettings {
	n := defaultNotificationSettings()

	for _, f := range opt {
		f(n)
	}

	return n
}
