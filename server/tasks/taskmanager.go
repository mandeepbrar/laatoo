package tasks

import (
	"encoding/json"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	"laatoo/server/constants"
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

	tskmgrInitializeCtx := tskMgr.createContext(ctx, "Initialize task manager")
	log.Logger.Trace(tskmgrInitializeCtx, "Create Task Manager queues")
	taskMgrConf, err, ok := common.ConfigFileAdapter(tskmgrInitializeCtx, conf, constants.CONF_TASKS)
	if err != nil {
		return errors.WrapError(tskmgrInitializeCtx, err)
	}
	if ok {
		taskNames := taskMgrConf.AllConfigurations()
		for _, taskName := range taskNames {
			taskConf, _ := taskMgrConf.GetSubConfig(taskName)
			tskCtx := tskmgrInitializeCtx.SubContext(taskName)
			if err := tskMgr.processTaskConf(tskCtx, taskConf, taskName); err != nil {
				return errors.WrapError(tskCtx, err)
			}
		}
	}

	if err := common.ProcessDirectoryFiles(tskmgrInitializeCtx, constants.CONF_TASKS, tskMgr.processTaskConf); err != nil {
		return err
	}

	return nil
}

func (tskMgr *taskManager) processTaskConf(ctx core.ServerContext, conf config.Config, taskName string) error {
	queueName, ok := conf.GetString(constants.CONF_TASKS_QUEUE)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASKS_QUEUE, "Task Name", taskName)
	}

	receiver, ok := conf.GetString(constants.CONF_TASK_RECEIVER)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASK_RECEIVER, "Task Name", taskName)
	}
	tskMgr.taskReceiverNames[queueName] = receiver

	processor, ok := conf.GetString(constants.CONF_TASK_PROCESSOR)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASK_PROCESSOR, "Task Name", taskName)
	}
	tskMgr.taskProcessorNames[queueName] = processor

	producer, ok := conf.GetString(constants.CONF_TASK_PRODUCER)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASK_PRODUCER, "Task Name", taskName)
	}
	tskMgr.taskProducers[queueName] = producer
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
			return errors.BadConf(tskmgrStartCtx, constants.CONF_TASK_RECEIVER, "queue", queueName)
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
		tskMgr.shandler.AuthenticateRequest(ctx, true)
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
