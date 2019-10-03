package core

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
)

type applicationProxy struct {
	app *application
}

func (proxy *applicationProxy) GetApplet(name string) (elements.Applet, bool) {
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
func (proxy *applicationProxy) GetContext() core.ServerContext {
	return proxy.app.svrContext
}
