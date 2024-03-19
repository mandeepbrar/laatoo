package elements

import (
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type ActivityManager interface {
	core.ServerElement
	GetActivity(ctx core.ServerContext, alias string) (components.Activity, error)
	InvokeActivity(ctx core.RequestContext, activity string, params ...interface{}) (interface{}, error)
}
