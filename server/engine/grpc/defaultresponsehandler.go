package grpc

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
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

func (rh *defaultResponseHandler) HandleResponse(ctx core.RequestContext, resp *core.Response) error {
	/*conn := ctx.EngineRequestContext().(*websocket.Conn)
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
	}*/
	return nil
}

func (proxy *defaultResponseHandler) Reference() core.ServerElement {
	anotherref := proxy
	return anotherref
}

func (proxy *defaultResponseHandler) GetProperty(name string) interface{} {
	return nil
}

func (proxy *defaultResponseHandler) GetName() string {
	return "GRPCResponseHandler"
}
func (proxy *defaultResponseHandler) GetType() core.ServerElementType {
	return core.ServerElementServiceResponseHandler
}
func (proxy *defaultResponseHandler) GetContext() core.ServerContext {
	return proxy.svrContext
}
