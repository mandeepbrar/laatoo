package components

import (
	"laatoo.io/sdk/server/core"
)

type Activity interface {
	GetName(ctx core.ServerContext) string
	GetParams(ctx core.ServerContext) map[string]core.Param
}

type ActivityManager interface {
	Load(ctx core.ServerContext, dir string) error
	GetActivity(ctx core.ServerContext, alias string) (Activity, error)
	RegisterActivity(ctx core.ServerContext, alias string, act Activity) error
	InvokeActivity(ctx core.RequestContext, act Activity, args ...interface{}) (interface{}, error)
}
