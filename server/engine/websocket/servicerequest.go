package websocket

import (
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

func (channel *wsChannel) handleMessage(ctx core.RequestContext, msg *rpcRequest) {
	if channel.disabled {
		return
	}
	vals := msg.Params
	if channel.staticValues != nil {
		for name, val := range channel.staticValues {
			vals[name] = val
		}
	}
	log.Trace(ctx, "Handle Request", "info", vals)
	res, err := channel.svc.HandleRequest(ctx, vals)
	if err != nil {
		channel.respHandler.HandleResponse(ctx, core.FunctionalErrorResponse(err))
	} else {
		err = channel.respHandler.HandleResponse(ctx, res)
	}
}
