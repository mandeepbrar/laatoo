package websocket

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
)

func (channel *wsChannel) handleMessage(ctx core.RequestContext, msg *rpcRequest) error {
	if channel.disabled {
		return nil
	}
	vals := msg.Params
	if channel.staticValues != nil {
		for name, val := range channel.staticValues {
			vals[name] = val
		}
	}
	log.Trace(ctx, "Handle Request", "info", vals)
	res, err := channel.svc.HandleRequest(ctx, vals)
	err = channel.respHandler.HandleResponse(ctx, res, err)
	return err
}
