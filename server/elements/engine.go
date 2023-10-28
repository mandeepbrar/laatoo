package elements

import (
	//	"laatoo.io/sdk/config"
	"laatoo.io/sdk/server/core"
)

type Engine interface {
	core.ServerElement
	GetRootChannel(ctx core.ServerContext) Channel
	GetRequestParams(ctx core.RequestContext) core.StringMap
	GetDefaultResponseHandler(ctx core.ServerContext) ServiceResponseHandler
}
