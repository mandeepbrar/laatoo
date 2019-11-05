package websocket

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"

	"github.com/gorilla/websocket"
)

type defaultResponseHandler struct {
	codec      core.Codec
	svrContext core.ServerContext
}

func DefaultResponseHandler(ctx core.ServerContext, codec core.Codec) *defaultResponseHandler {
	return &defaultResponseHandler{codec: codec}
}

func (rh *defaultResponseHandler) Initialize(ctx core.ServerContext, conf config.Config) error {
	rh.svrContext = ctx
	return nil
}

func (rh *defaultResponseHandler) HandleResponse(ctx core.RequestContext, resp *core.Response, handlingError error) error {
	conn := ctx.EngineRequestContext().(*websocket.Conn)
	log.Trace(ctx, "Returning request with default response handler")
	if resp != nil {
		switch resp.Status {
		case core.StatusSuccess:
			rh.sendResponse(ctx, conn, resp.Data)
		case core.StatusInternalError:
			rh.sendResponse(ctx, conn, resp.Error)
		case core.StatusFunctionalError:
			rh.sendResponse(ctx, conn, resp.Error)
		default:
			log.Error(ctx, "HandleResponse status not implemented", "resp", resp)
		}
	}
	return nil
}

func (rh *defaultResponseHandler) sendResponse(ctx core.RequestContext, conn *websocket.Conn, dat interface{}) {
	wsid, _ := ctx.GetString("__wsid")
	resp := &rpcResponse{"2.0", dat, wsid}
	byts, err := rh.codec.Marshal(ctx, resp)
	if err != nil {
		err = rh.HandleResponse(ctx, core.InternalErrorResponse(err.Error()), err)
	}
	err = conn.WriteMessage(websocket.TextMessage, byts)
	if err != nil {
		log.Error(ctx, "Failed to write bytes to connection")
	}
}

func (proxy *defaultResponseHandler) Reference() core.ServerElement {
	anotherref := proxy
	return anotherref
}

func (proxy *defaultResponseHandler) GetProperty(name string) interface{} {
	return nil
}

func (proxy *defaultResponseHandler) GetName() string {
	return "DefaultWebsocketResponseHandler"
}
func (proxy *defaultResponseHandler) GetType() core.ServerElementType {
	return core.ServerElementServiceResponseHandler
}
func (proxy *defaultResponseHandler) GetContext() core.ServerContext {
	return proxy.svrContext
}
