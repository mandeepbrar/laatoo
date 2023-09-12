package elements

import (
	"laatoo.io/sdk/server/core"
)

type ModuleManager interface {
	core.ServerElement
	List(ctx core.ServerContext) map[string]string
	Describe(ctx core.ServerContext, mod string) (map[string]interface{}, error)
	ChangeLogger(ctx core.ServerContext, mod string, logLevel string, logFormat string) error
}
