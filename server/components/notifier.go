package components

import "laatoo.io/sdk/server/core"

type Notifier interface {
	GetSessionId() string
	GetUserId() string
	Notify(ctx core.RequestContext, notificaiton *core.Notification) error
}
