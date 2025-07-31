package components

import "laatoo.io/sdk/server/core"

type NotificationManager interface {
	SendNotification(ctx core.RequestContext, notification *core.Notification) error
	Broadcast(ctx core.RequestContext, notif *core.Notification) error
	GetQueue() string
}
