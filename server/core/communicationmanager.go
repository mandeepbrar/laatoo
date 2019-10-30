package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type communicator struct {
	name        string
	commSvcName string
	commSvc     components.Communicator
	proxy       elements.Communicator
	svrContext  core.ServerContext
}

func (comm *communicator) Initialize(ctx core.ServerContext, conf config.Config) error {
	commCtx := ctx.SubContext("Initialize communicator")
	log.Trace(commCtx, "Initialize Communicator")
	return nil
}

func (comm *communicator) Start(ctx core.ServerContext) error {
	commSvc, err := ctx.GetService(comm.commSvcName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc, ok := commSvc.(components.Communicator)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF)
	}
	log.Info(ctx, "Communication service configured")
	comm.commSvc = svc
	return nil
}

func (comm *communicator) sendCommunication(ctx core.RequestContext, communication *components.Communication) error {
	if comm.commSvc != nil {
		err := comm.commSvc.SendCommunication(ctx, communication)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		return nil
	}
	return errors.BadConf(ctx, "No communicator service has been configured")

}
