package tasks

import (
	"encoding/json"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type taskManager struct {
	parent             core.ServerElement
	proxy              server.TaskManager
	authHeader         string
	shandler           server.SecurityHandler
	taskProducers      map[string]string
	taskProducerSvcs   map[string]components.TaskQueue
	taskReceiverNames  map[string]string
	taskProcessorNames map[string]string
	taskProcessors     map[string]core.Service
}

func (tskMgr *taskManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	tskmgrInitializeCtx := tskMgr.createContext(ctx, "Initialize task manager")

	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(server.SecurityHandler)
		tskMgr.shandler = shandler

		ah, ok := shandler.GetString(config.AUTHHEADER)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}
		tskMgr.authHeader = ah
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
	}

	log.Logger.Trace(tskmgrInitializeCtx, "Create Task Manager queues")
	err := tskMgr.createProducerQueues(tskmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(tskmgrInitializeCtx, err)
	}

	err = tskMgr.createConsumerQueues(tskmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(tskmgrInitializeCtx, err)
	}
	return nil
}

func (tskMgr *taskManager) createProducerQueues(ctx core.ServerContext, conf config.Config) error {
	queuesConf, ok := conf.GetSubConfig(config.CONF_TASK_PRODUCERS)
	if ok {
		queueNames := queuesConf.AllConfigurations()
		for _, queueName := range queueNames {
			queueSvcName, _ := queuesConf.GetString(queueName)
			tskMgr.taskProducers[queueName] = queueSvcName
		}
	}
	return nil
}

func (tskMgr *taskManager) createConsumerQueues(ctx core.ServerContext, conf config.Config) error {
	queuesConf, ok := conf.GetSubConfig(config.CONF_TASK_CONSUMERS)
	if ok {
		queueNames := queuesConf.AllConfigurations()
		for _, queueName := range queueNames {
			queueConf, _ := queuesConf.GetSubConfig(queueName)
			receiver, ok := queueConf.GetString(config.CONF_TASK_RECEIVER)
			if !ok {
				return errors.MissingConf(ctx, config.CONF_TASK_RECEIVER, "queue", queueName)
			}
			tskMgr.taskReceiverNames[queueName] = receiver
			processor, ok := queueConf.GetString(config.CONF_TASK_PROCESSOR)
			if !ok {
				return errors.MissingConf(ctx, config.CONF_TASK_PROCESSOR, "queue", queueName)
			}
			tskMgr.taskProcessorNames[queueName] = processor
		}
	}
	return nil
}

func (tskMgr *taskManager) Start(ctx core.ServerContext) error {
	tskmgrStartCtx := tskMgr.createContext(ctx, "Start task manager")
	log.Logger.Trace(tskmgrStartCtx, "Start Task Manager queues")
	for queueName, svcName := range tskMgr.taskProducers {
		log.Logger.Trace(tskmgrStartCtx, "Starting task producer ", "queue", queueName)
		tqSvc, err := tskmgrStartCtx.GetService(svcName)
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}
		tq, ok := tqSvc.(components.TaskQueue)
		if !ok {
			return errors.ThrowError(tskmgrStartCtx, errors.CORE_ERROR_BAD_CONF)
		}
		tskMgr.taskProducerSvcs[queueName] = tq
	}

	for queueName, receiverName := range tskMgr.taskReceiverNames {
		log.Logger.Trace(tskmgrStartCtx, "Starting task consumer ", "queue", queueName)
		receiverSvc, err := tskmgrStartCtx.GetService(receiverName)
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}
		ts, ok := receiverSvc.(components.TaskServer)
		if !ok {
			return errors.BadConf(tskmgrStartCtx, config.CONF_TASK_RECEIVER, "queue", queueName)
		}

		err = ts.SubsribeQueue(tskmgrStartCtx, queueName)
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}

		procSvc, err := tskmgrStartCtx.GetService(tskMgr.taskProcessorNames[queueName])
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}
		tskMgr.taskProcessors[queueName] = procSvc
	}

	return nil
}

func (tskMgr *taskManager) pushTask(ctx core.RequestContext, queue string, taskData interface{}) error {
	tq, ok := tskMgr.taskProducerSvcs[queue]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Missing Queue", queue)
	}
	log.Logger.Trace(ctx, "Pushing task to queue", "queue", queue)
	token, _ := ctx.GetString(tskMgr.authHeader)
	data, err := json.Marshal(taskData)
	if err != nil {
		return err
	}
	t := &components.Task{Queue: queue, Data: data, Token: token}
	return tq.PushTask(ctx, queue, t)
}

func (tskMgr *taskManager) processTask(ctx core.RequestContext, t *components.Task) error {
	/*
		req := ctx.SubContext("Gae background task " + t.Queue)
		log.Logger.Debug(req, "Received job ")
		req.SetRequest(t.Data)
		req.Set(svc.authHeader, t.Token)
		svc.shandler.AuthenticateRequest(req)
		queueName := t.Queue
		q, ok := svc.queues[queueName]
		if ok {
			err := q.lstnr.Invoke(req)
			if err != nil {
				log.Logger.Error(req, "Error in background process", "err", err)
				return err
			}
		}*/
	queue := t.Queue
	processor, ok := tskMgr.taskProcessors[queue]
	if ok {
		ctx.SetRequest(t.Data)
		ctx.Set(tskMgr.authHeader, t.Token)
		tskMgr.shandler.AuthenticateRequest(ctx)
		log.Logger.Trace(ctx, "Processing background task")
		return processor.Invoke(ctx)
	}
	return nil
}

//creates a context specific to factory manager
func (tskMgr *taskManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementTaskManager: tskMgr.proxy}, core.ServerElementTaskManager)
}
