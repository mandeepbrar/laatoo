package components

import "laatoo.io/sdk/server/core"

type NotificationManager interface {
	SendNotification(ctx core.ServerContext, notification *core.Notification) error
}
