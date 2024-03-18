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
	LoadActivities(ctx core.ServerContext, dir string) (map[string]Activity, error)
	GetActivity(ctx core.ServerContext, alias string) (Activity, error)
	InvokeActivity(ctx core.RequestContext, activity string, params core.StringMap) (interface{}, error)
}
