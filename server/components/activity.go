package components

import (
	"laatoo.io/sdk/server/core"
)

type Activity interface {
	GetName(ctx core.ServerContext) string
	GetParams(ctx core.ServerContext) map[string]core.Param
}

type ActivityManager interface {
	LoadActivities(ctx core.ServerContext, dir string) (map[string]Activity, error)
	GetActivity(ctx core.ServerContext, alias string) (Activity, error)
	InvokeActivity(ctx core.RequestContext, act Activity, args ...interface{}) (interface{}, error)
}
