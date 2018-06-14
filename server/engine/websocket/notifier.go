package websocket

import "laatoo/sdk/core"

type NotifierService struct {
	core.Service
	engine *wsEngine
}

func (svc *NotifierService) Start(ctx core.ServerContext) error {
	return nil
}

func (svc *NotifierService) Notify(ctx core.ServerContext, sessionidentifier interface{}, msg interface{}, dataid string) error {
	return svc.engine.rtr.notify(ctx, sessionidentifier, msg, dataid)
}

func (svc *NotifierService) Broadcast(ctx core.ServerContext, msg interface{}, dataid string) error {
	return svc.engine.rtr.broadcast(ctx, msg, dataid)
}
