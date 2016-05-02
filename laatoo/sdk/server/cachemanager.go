package server

import (
	"laatoo/sdk/core"
	"laatoo/sdk/services"
)

type CacheManager interface {
	core.ServerElement
	GetCache(ctx core.ServerContext, name string) services.Cache
}
