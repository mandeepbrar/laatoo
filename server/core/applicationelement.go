package core

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type applicationProxy struct {
	app *application
}

func (proxy *applicationProxy) GetApplet(name string) (server.Applet, bool) {
	applet, ok := proxy.app.applets[name]
	return applet, ok
}

func (proxy *applicationProxy) Reference() core.ServerElement {
	return &applicationProxy{app: proxy.app}
}
func (proxy *applicationProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *applicationProxy) GetName() string {
	return proxy.app.name
}
func (proxy *applicationProxy) GetType() core.ServerElementType {
	return core.ServerElementApplication
}
