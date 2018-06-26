package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
)

type UIPlugin interface {
	GetRegistry(ctx core.ServerContext) config.Config
}
