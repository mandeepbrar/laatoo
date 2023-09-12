package elements

import (
	"laatoo/sdk/config"
	"laatoo/sdk/server/core"
)

type ServiceResponseHandler interface {
	core.ServerElement
	Initialize(ctx core.ServerContext, conf config.Config) error
	HandleResponse(ctx core.RequestContext, resp *core.Response, err error) error
}
