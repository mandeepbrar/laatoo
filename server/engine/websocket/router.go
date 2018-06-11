package websocket

import (
	"laatoo/sdk/core"

	"github.com/gorilla/websocket"
)

type router struct {
	routes map[string]*wsChannel
}

func newRouter(ctx core.ServerContext) (*router, error) {
	rtr := &router{make(map[string]*wsChannel)}
	return rtr, nil
}

func (rtr *router) addChannel(ctx core.ServerContext, route string, chn *wsChannel) error {
	rtr.routes[route] = chn
	return nil
}

func (rtr *router) routeMessage(ctx core.ServerContext, msg *rpcRequest, conn *websocket.Conn) error {
	reqctx := ctx.CreateNewRequest(msg.Method, conn)
	defer reqctx.CompleteRequest()
	reqctx.Set("__wsid", msg.Id)
	chn, ok := rtr.routes[msg.Method]
	if !ok {
		conn.WriteMessage(websocket.TextMessage, []byte("Missing channel: "+msg.Method))
		return nil
	}
	chn.handleMessage(reqctx, msg)
	return nil
}
