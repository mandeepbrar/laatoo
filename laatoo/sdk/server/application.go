package server

import (
	"laatoo/sdk/core"
)

type Application interface {
	core.ServerElement
	GetService(ctx core.ServerContext, alias string) (Service, error)
}
