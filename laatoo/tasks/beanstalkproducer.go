package tasks

import (
	"encoding/json"
	"github.com/prep/beanstalk"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
)

const (
	CONF_BEANSTALK_SERVER = "server"
)

type beanstalkProducer struct {
	pool   *beanstalk.ProducerPool
	params *beanstalk.PutParams
}

func (svc *beanstalkProducer) Initialize(ctx core.ServerContext, conf config.Config) error {
	addr, ok := conf.GetString(CONF_BEANSTALK_SERVER)
	if !ok {
		addr = ":11300"
	}
	svc.pool = beanstalk.NewProducerPool([]string{addr}, nil)
	if svc.pool == nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "server", addr)
	}
	// Reusable put parameters.
	svc.params = &beanstalk.PutParams{Priority: 0, Delay: 0, TTR: 5}
	return nil
}

func (svc *beanstalkProducer) PushTask(ctx core.RequestContext, queue string, taskData interface{}) error {
	data, err := json.Marshal(taskData)
	if err != nil {
		return err
	}
	t := task{Queue: queue, Data: data, User: ctx.GetUser().GetId()}
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
