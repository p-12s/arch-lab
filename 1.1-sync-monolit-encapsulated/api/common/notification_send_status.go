package common

type NotificationSendStatus int

const (
	NOTIFICATION_SEND_STATUS_UNKNOWN NotificationSendStatus = iota
	NOTIFICATION_SEND_STATUS_OK
	NOTIFICATION_SEND_STATUS_ERR
)
