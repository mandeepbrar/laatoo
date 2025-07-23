package components

import "laatoo.io/sdk/server/core"

type NotificationType string

const (
	INAPP    NotificationType = "INAPP"
	EMAIL    NotificationType = "EMAIL"
	SMS      NotificationType = "SMS"
	PUSH     NotificationType = "PUSH"
	WHATSAPP NotificationType = "WHATSAPP"
	WEBHOOK  NotificationType = "WEBHOOK"
)

type Notification struct {
	NotificationType NotificationType
	Subject          string
	Mime             string
	Attachments      []string
	Recipients       map[string]string
	Message          []byte
	Info             interface{}
}

type NotificationManager interface {
	SendNotification(ctx core.RequestContext, notification *Notification) error
}
