package main

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

type UIPlugin interface {
	GetRegistry(ctx core.ServerContext) config.Config
}
