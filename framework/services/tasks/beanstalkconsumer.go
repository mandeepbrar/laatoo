package tasks

import (
	"encoding/json"
	"github.com/prep/beanstalk"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type beanstalkConsumer struct {
	queues map[string]*taskQueue
	pool   *beanstalk.ConsumerPool
	addr   string
}

func (svc *beanstalkConsumer) Initialize(ctx core.ServerContext, conf config.Config) error {
	addr, ok := conf.GetString(CONF_BEANSTALK_SERVER)
	if !ok {
		addr = ":11300"
	}
	svc.addr = addr
	queuesConf, ok := conf.GetSubConfig(config.CONF_TASK_QUEUES)
	if ok {
		queueNames := queuesConf.AllConfigurations()
		for _, queueName := range queueNames {
			qCtx := ctx.SubContext("Creating Queue" + queueName)
			queueProcessorName, _ := queuesConf.GetString(queueName)
			processor, err := qCtx.GetService(queueProcessorName)
			if err != nil {
				return errors.WrapError(qCtx, err)
			}
			svc.createQueue(qCtx, queueName, processor)
		}
		svc.pool = beanstalk.NewConsumerPool([]string{svc.addr}, queueNames, nil)
		if svc.pool == nil {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "server", svc.addr)
		}
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_TASK_QUEUES)
	}
	return nil
}

func (svc *beanstalkConsumer) createQueue(ctx core.ServerContext, queue string, lstnr core.Service) error {
	_, ok := svc.queues[queue]
	if ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "queuename", queue)
	}
	tq := &taskQueue{name: queue, qRef: nil, lstnr: lstnr}
	svc.queues[queue] = tq
	return nil
}

func (svc *beanstalkConsumer) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *beanstalkConsumer) Start(ctx core.ServerContext) error {
	svc.pool.Play()
	go svc.work(ctx)
	return nil
}
func (svc *beanstalkConsumer) work(ctx core.ServerContext) {
	for {
		select {
		case job := <-svc.pool.C:
			go func(ctx core.ServerContext, job *beanstalk.Job) {
				t := &task{}
				err := json.Unmarshal(job.Body, t)
				if err != nil {
					log.Logger.Error(ctx, "Error in background process", "job", job.ID, "err", err)
					job.Bury()
				} else {
					log.Logger.Debug(ctx, "Received job", "id", job.ID, "task", t, "body", job.Body)
					req := ctx.CreateNewRequest("Task", nil)
					req.SetRequest(t.Data)
					req.Set("User", t.User)
					queueName := t.Queue
					log.Logger.Debug(ctx, "Received job", "id", job.ID, "task", t, "body", job.Body, "queue", queueName)
					q, ok := svc.queues[queueName]
					if ok {
						err := q.lstnr.Invoke(req)
						if err != nil {
							log.Logger.Error(ctx, "Error in background process", "job", job.ID, "err", err)
						} else {
							job.Delete()
						}
					} else {
						job.Bury()
					}
				}
			}(ctx, job)
		}
	}
}
