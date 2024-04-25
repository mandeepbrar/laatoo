package elements

import (
	"laatoo.io/sdk/server/core"
)

type ActionsManager interface {
	core.ServerElement
	RegisterAction(ctx core.ServerContext, actionType core.ActionType, executor core.ActionExecutor) error
	ExecuteAction(ctx core.RequestContext, action *core.Action, params core.StringMap) (interface{}, error)
}
