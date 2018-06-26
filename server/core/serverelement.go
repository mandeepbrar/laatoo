package core

import "laatoo/sdk/server/core"

type serverProxy struct {
	server *serverObject
}

func (proxy *serverProxy) Reference() core.ServerElement {
	return &serverProxy{server: proxy.server}
}
func (proxy *serverProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *serverProxy) GetName() string {
	return proxy.server.name
}
func (proxy *serverProxy) GetType() core.ServerElementType {
	return core.ServerElementServer
}
