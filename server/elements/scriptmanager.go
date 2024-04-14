package elements

import (
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type ScriptManager interface {
	core.ServerElement
	GetScript(ctx core.ServerContext, alias string) (components.Script, error)
	RegisterScript(ctx core.ServerContext, alias string, act components.Script) error
	InvokeScript(ctx core.RequestContext, script string, params core.StringMap) (interface{}, error)
}
