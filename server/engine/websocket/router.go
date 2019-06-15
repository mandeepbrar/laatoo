package websocket

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"

	"github.com/gorilla/websocket"
)

type router struct {
	routes      map[string]*wsChannel
	engine      *wsEngine
	connections map[string]*websocket.Conn
}

func newRouter(ctx core.ServerContext, eng *wsEngine) (*router, error) {
	rtr := &router{make(map[string]*wsChannel), eng, make(map[string]*websocket.Conn)}
	return rtr, nil
}

func (rtr *router) addChannel(ctx core.ServerContext, route string, chn *wsChannel) error {
	rtr.routes[route] = chn
	return nil
}

func (rtr *router) addConnection(ctx core.ServerContext, sessionId string, conn *websocket.Conn) error {
	rtr.connections[sessionId] = conn
	return nil
}

func (rtr *router) routeMessage(ctx core.ServerContext, msg *rpcRequest, conn *websocket.Conn) error {
	reqctx, err := ctx.CreateNewRequest(msg.Method, rtr.engine.proxy, conn, "")
	if err != nil {
		return errors.WrapError(ctx, err)
	}
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

func (rtr *router) notify(ctx core.ServerContext, identifier interface{}, msg interface{}, dataid string) error {
	return nil
}

func (rtr *router) broadcast(ctx core.ServerContext, msg interface{}, dataid string) error {
	for _, conn := range rtr.connections {
		not := &rpcNotification{"2.0", msg, dataid}
		byts, err := rtr.engine.codec.Marshal(ctx, not)
		if err != nil {
			return err
		}
		err = conn.WriteMessage(websocket.TextMessage, byts)
		if err != nil {
			return err
		}
	}
	return nil
}
