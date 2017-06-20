package main

import (
	"encoding/json"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"

	"github.com/prep/beanstalk"
)

const (
	CONF_TASKS_BEANSTALK_PRODUCER = "beanstalktaskpublisher"
	CONF_TASKS_BEANSTALK_CONSUMER = "beanstalktaskprocessor"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_TASKS_BEANSTALK_PRODUCER, Object: beanstalkConsumer{}},
		core.PluginComponent{Name: CONF_TASKS_BEANSTALK_CONSUMER, Object: beanstalkProducer{}}}
}

type beanstalkConsumer struct {
	addr        string
	taskManager server.TaskManager
	worker      func(ctx core.ServerContext, pool *beanstalk.ConsumerPool)
}

func (svc *beanstalkConsumer) Initialize(ctx core.ServerContext, conf config.Config) error {
	addr, ok := conf.GetString(CONF_BEANSTALK_SERVER)
	if !ok {
		addr = ":11300"
	}
	/*
		sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
		if sh != nil {
			shandler := sh.(server.SecurityHandler)
			svc.shandler = shandler
			ah, ok := shandler.GetString(config.AUTHHEADER)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
			}
			svc.authHeader = ah
		} else {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}*/

	svc.addr = addr

	svc.taskManager = ctx.GetServerElement(core.ServerElementTaskManager).(server.TaskManager)

	/*queuesConf, ok := conf.GetSubConfig(config.CONF_TASK_QUEUES)
	if ok {
		queueNames := queuesConf.AllConfigurations()
		for _, queueName := range queueNames {
			qCtx := ctx.SubContext("Creating Queue " + queueName)
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
	}*/
	return nil
}

func (svc *beanstalkConsumer) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *beanstalkConsumer) SubsribeQueue(ctx core.ServerContext, queue string) error {
	pool := beanstalk.NewConsumerPool([]string{svc.addr}, []string{queue}, nil)
	if pool == nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "server", svc.addr)
	}
	pool.Play()
	go svc.worker(ctx, pool)
	return nil
}

func (svc *beanstalkConsumer) Start(ctx core.ServerContext) error {
	svc.worker = func(workerctx core.ServerContext, pool *beanstalk.ConsumerPool) {
		for {
			select {
			case job := <-pool.C:
				go func(ctx core.ServerContext, job *beanstalk.Job) {

					t := &components.Task{}
					log.Logger.Info(ctx, "Recieved job", "Job Id", job.ID)
					err := json.Unmarshal(job.Body, t)
					if err != nil {
						log.Logger.Error(ctx, "Error in background process", "job", job.ID, "err", err)
						job.Bury()
					} else {
						req := ctx.CreateNewRequest("Beanstalk task "+t.Queue, nil)
						err := svc.taskManager.ProcessTask(req, t)
						if err != nil {
							log.Logger.Error(req, "Error in background process", "job", job.ID, "err", err)
							job.Bury()
						} else {
							job.Delete()
						}
					}
				}(ctx, job)
			}
		}
	}

	return nil
}