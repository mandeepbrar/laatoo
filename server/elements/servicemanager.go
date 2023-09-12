package elements

import (
	"laatoo/sdk/server/core"
)

type ServiceManager interface {
	core.ServerElement
	GetService(ctx core.ServerContext, alias string) (Service, error)
	GetServiceContext(ctx core.ServerContext, alias string) (core.ServerContext, error)
	List(ctx core.ServerContext) map[string]string
	Describe(ctx core.ServerContext, svc string) (map[string]interface{}, error)
	ChangeLogger(ctx core.ServerContext, svc string, logLevel string, logFormat string) error
}
