package main

import (
	"encoding/json"
	"laatoo/sdk/server/components"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"

	"laatoo/sdk/server/log"

	"github.com/prep/beanstalk"
)

const (
	CONF_BEANSTALK_SERVER = "beanstalkserver"
)

type BeanstalkProducer struct {
	core.Service
	pool       *beanstalk.ProducerPool
	params     *beanstalk.PutParams
	authHeader string
}

func (svc *BeanstalkProducer) Describe(ctx core.ServerContext) error {
	svc.SetComponent(ctx, true)
	svc.SetDescription(ctx, "Beanstalk producer component")
	svc.AddStringConfigurations(ctx, []string{CONF_BEANSTALK_SERVER}, []string{":11300"})
	return nil
}

func (svc *BeanstalkProducer) Initialize(ctx core.ServerContext, conf config.Config) error {
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(elements.SecurityHandler)
		ah := shandler.GetProperty(config.AUTHHEADER)
		if ah == nil {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}
		svc.authHeader = ah.(string)
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
	}

	return nil
}

func (svc *BeanstalkProducer) PushTask(ctx core.RequestContext, queue string, t *components.Task) error {
	log.Debug(ctx, "Pushing task into queue", "queue", queue)
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = svc.pool.Put(queue, bytes, svc.params)
	return err
}

func (svc *BeanstalkProducer) Start(ctx core.ServerContext) error {
	addr, _ := svc.GetConfiguration(ctx, CONF_BEANSTALK_SERVER)

	pool, err := beanstalk.NewProducerPool([]string{addr.(string)}, nil)
	if err != nil {
		return errors.WrapError(ctx, err, "server", addr)
	} else {
		svc.pool = pool
	}
	// Reusable put parameters.
	svc.params = &beanstalk.PutParams{Priority: 0, Delay: 0, TTR: 5}
	return nil
}
