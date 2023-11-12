package components

import (
	"laatoo.io/sdk/server/core"
)

type Activity interface {
	GetName(ctx core.ServerContext) string
	GetParams(ctx core.ServerContext) map[string]core.Param
	Invoke(core.RequestContext) error
}

type ActivityManager interface {
	GetActivity(ctx core.ServerContext, alias string) (Activity, error)
}
