package main

import (
	"encoding/json"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"

	"google.golang.org/appengine/taskqueue"
)

const (
	GAE_PATH = "path"
)

type gaeProducer struct {
	path       string
	authHeader string
}

func (svc *gaeProducer) Initialize(ctx core.ServerContext, conf config.Config) error {
	path, ok := conf.GetString(GAE_PATH)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", GAE_PATH)
	}

	svc.path = path
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(server.SecurityHandler)
		ah, ok := shandler.GetString(config.AUTHHEADER)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}
		svc.authHeader = ah
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
	}

	return nil
}

func (svc *gaeProducer) PushTask(ctx core.RequestContext, queue string, t *components.Task) error {
	appEngineContext := ctx.GetAppengineContext()
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	gaetask := taskqueue.NewPOSTTask(svc.path, map[string][]string{})
	gaetask.Payload = bytes
	_, err = taskqueue.Add(appEngineContext, gaetask, "")
	return err
}

func (svc *gaeProducer) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *gaeProducer) Start(ctx core.ServerContext) error {
	return nil
}
