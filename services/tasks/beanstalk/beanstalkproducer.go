package main

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

type BeanstalkProducer struct {
	pool       *beanstalk.ProducerPool
	params     *beanstalk.PutParams
	authHeader string
}

func (svc *BeanstalkProducer) Initialize(ctx core.ServerContext, conf config.Config) error {
	addr, ok := conf.GetString(CONF_BEANSTALK_SERVER)
	if !ok {
		addr = ":11300"
	}

	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(server.SecurityHandler)
		ah := shandler.GetProperty(config.AUTHHEADER)
		if ah == nil {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}
		svc.authHeader = ah.(string)
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

func (svc *BeanstalkProducer) PushTask(ctx core.RequestContext, queue string, t *components.Task) error {
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = svc.pool.Put(queue, bytes, svc.params)
	return err
}

func (svc *BeanstalkProducer) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *BeanstalkProducer) Start(ctx core.ServerContext) error {
	return nil
}
