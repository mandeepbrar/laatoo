package server

import (
	"laatoo/core/common"
	//	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type applicationProxy struct {
	*common.Context
	app *application
}

func (app *applicationProxy) GetService(ctx core.ServerContext, alias string) (server.Service, error) {
	return app.app.serviceManager.GetService(ctx, alias)
}
