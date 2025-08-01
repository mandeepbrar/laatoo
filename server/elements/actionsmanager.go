package elements

import (
	"laatoo.io/sdk/server/core"
	"laatoo.io/sdk/utils"
)

type ActionsManager interface {
	core.ServerElement
	RegisterAction(ctx core.ServerContext, actionType core.ActionType, executor core.ActionExecutor) error
	ExecuteAction(ctx core.RequestContext, action core.ActionType, params utils.StringMap) (interface{}, error)
}
