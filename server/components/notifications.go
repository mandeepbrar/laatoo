package components

import "laatoo.io/sdk/server/core"

type NotificationManager interface {
	SendNotification(ctx core.RequestContext, notification *core.Notification) error
}
