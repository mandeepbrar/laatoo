package components

import "laatoo.io/sdk/server/core"

type Notifier interface {
	Notify(ctx core.ServerContext, notificaiton *core.Notification) error
}
