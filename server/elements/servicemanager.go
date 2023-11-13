package elements

import (
	"laatoo.io/sdk/config"
	"laatoo.io/sdk/server/core"
)

type ServiceManager interface {
	core.ServerElement
	GetService(ctx core.ServerContext, alias string) (Service, error)
	GetServiceContext(ctx core.ServerContext, alias string) (core.ServerContext, error)
	List(ctx core.ServerContext) core.StringsMap
	Describe(ctx core.ServerContext, svc string) (core.StringMap, error)
	ChangeLogger(ctx core.ServerContext, svc string, logLevel string, logFormat string) error
	CreateParams(ctx core.ServerContext, paramsConf config.Config) (map[string]core.Param, error)
}
