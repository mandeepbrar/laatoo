package elements

import (
	"laatoo.io/sdk/server/core"
)

type FactoryManager interface {
	core.ServerElement
	GetFactory(ctx core.ServerContext, factoryName string) (Factory, error)
	List(ctx core.ServerContext) map[string]string
	ChangeLogger(ctx core.ServerContext, chanName string, logLevel string, logFormat string) error
	Describe(ctx core.ServerContext, chanName string) (map[string]interface{}, error)
}
