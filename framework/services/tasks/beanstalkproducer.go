package tasks

import (
	"encoding/json"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"

	"github.com/prep/beanstalk"
	//	"laatoo/sdk/log"
)

const (
	CONF_BEANSTALK_SERVER = "server"
)

type beanstalkProducer struct {
	pool       *beanstalk.ProducerPool
	params     *beanstalk.PutParams
	authHeader string
}

func (svc *beanstalkProducer) Initialize(ctx core.ServerContext, conf config.Config) error {
	addr, ok := conf.GetString(CONF_BEANSTALK_SERVER)
	if !ok {
		addr = ":11300"
	}

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

	svc.pool = beanstalk.NewProducerPool([]string{addr}, nil)
	if svc.pool == nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "server", addr)
	}
	// Reusable put parameters.
	svc.params = &beanstalk.PutParams{Priority: 0, Delay: 0, TTR: 5}
	return nil
}

func (svc *beanstalkProducer) PushTask(ctx core.RequestContext, queue string, t *components.Task) error {
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = svc.pool.Put(queue, bytes, svc.params)
	return err
}

func (svc *beanstalkProducer) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *beanstalkProducer) Start(ctx core.ServerContext) error {
	return nil
}
