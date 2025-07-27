package elements

import (
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type SessionManager interface {
	core.ServerElement
	GetSession(ctx core.RequestContext, sessionId string) (core.Session, error)
	RegisterUserNotifier(ctx core.ServerContext, userId string, notifier components.Notifier)
	GetUserSession(ctx core.RequestContext, userId string) (core.Session, error)
	Broadcast(ctx core.RequestContext, notif *core.Notification) error
}
