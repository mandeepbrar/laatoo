package elements

import (
	"laatoo.io/sdk/server/core"
)

type ActionsManager interface {
	core.ServerElement
	RegisterAction(actionType core.ActionType, executor core.ActionExecutor) error
	ExecuteAction(action *core.Action, params core.StringMap) (interface{}, error)
}
