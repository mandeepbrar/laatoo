package elements

import (
	"laatoo.io/sdk/config"
	"laatoo.io/sdk/server/core"
)

type Service interface {
	core.ServerElement
	Service() core.Service
	ServiceContext() core.ServerContext
	GetConfiguration() config.Config
	HandleRequest(ctx core.RequestContext, vals core.StringMap, encoding core.StringsMap) (*core.Response, error)
}
