package model

type NotificationSettings struct {
	SendEmail            bool
	SendPushNotification bool
	// SendLocationBasedNotifications bool
}

type ConfigureNotifications func(n *NotificationSettings) *NotificationSettings

func AllowEmails(n *NotificationSettings) *NotificationSettings {
	n.SendEmail = true
	return n
}

func AllowAll(n *NotificationSettings) *NotificationSettings {
	n.SendEmail = true
	n.SendPushNotification = true
	return n
}

func defaultNotificationSettings() *NotificationSettings {
	return &NotificationSettings{
		SendEmail:            false,
		SendPushNotification: false,
	}
}

func NewNotificationSettings(opt ...ConfigureNotifications) *NotificationSettings {
	n := defaultNotificationSettings()
	for _, f := range opt {
		f(n)
	}

	return n
}
