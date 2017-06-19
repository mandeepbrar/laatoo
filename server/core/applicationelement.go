package core

import (
	"laatoo/sdk/server"
	"laatoo/server/common"
)

type applicationProxy struct {
	*common.Context
	app *application
}

func (proxy *applicationProxy) GetApplet(name string) (server.Applet, bool) {
	applet, ok := proxy.app.applets[name]
	return applet, ok
}
