package elements

import (
	"laatoo.io/sdk/server/core"
)

type ModuleManager interface {
	core.ServerElement
	List(ctx core.ServerContext) core.StringsMap
	Describe(ctx core.ServerContext, mod string) (core.StringMap, error)
	ChangeLogger(ctx core.ServerContext, mod string, logLevel string, logFormat string) error
}
