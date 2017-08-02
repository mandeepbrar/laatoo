package core

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
	name               string
	proxy              server.TaskManager
	authHeader         string
	shandler           server.SecurityHandler
	taskPublishers     map[string]string
	taskPublisherSvcs  map[string]components.TaskQueue
	taskConsumerNames  map[string]string
	taskProcessorNames map[string]string
	taskProcessors     map[string]server.Service
}

func (tskMgr *taskManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(server.SecurityHandler)
		tskMgr.shandler = shandler

		val := shandler.GetProperty(config.AUTHHEADER)
		if val == nil {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}
		tskMgr.authHeader = val.(string)
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
	}

	tskmgrInitializeCtx := ctx.SubContext("Initialize task manager")
	log.Trace(tskmgrInitializeCtx, "Create Task Manager queues")
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

	baseDir, _ := ctx.GetString(constants.CONF_BASE_DIR)

	return tskMgr.loadTasksFromDirectory(ctx, baseDir)
}

func (tskMgr *taskManager) loadTasksFromDirectory(ctx core.ServerContext, baseDir string) error {
	return common.ProcessDirectoryFiles(ctx, baseDir, constants.CONF_TASKS, tskMgr.processTaskConf, true)
}

func (tskMgr *taskManager) processTaskConf(ctx core.ServerContext, conf config.Config, taskName string) error {
	queueName, ok := conf.GetString(constants.CONF_TASKS_QUEUE)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASKS_QUEUE, "Task Name", taskName)
	}

	consumer, ok := conf.GetString(constants.CONF_TASK_CONSUMER)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASK_CONSUMER, "Task Name", taskName)
	}
	consumer = common.FillVariables(ctx, consumer)
	tskMgr.taskConsumerNames[queueName] = consumer

	processor, ok := conf.GetString(constants.CONF_TASK_PROCESSOR)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASK_PROCESSOR, "Task Name", taskName)
	}
	processor = common.FillVariables(ctx, processor)
	tskMgr.taskProcessorNames[queueName] = processor

	publisher, ok := conf.GetString(constants.CONF_TASK_PUBLISHER)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASK_PUBLISHER, "Task Name", taskName)
	}
	publisher = common.FillVariables(ctx, publisher)
	tskMgr.taskPublishers[queueName] = publisher
	return nil
}

func (tskMgr *taskManager) Start(ctx core.ServerContext) error {
	tskmgrStartCtx := ctx.SubContext("Start task manager")
	svcMgr := ctx.GetServerElement(core.ServerElementServiceManager).(server.ServiceManager)
	log.Trace(tskmgrStartCtx, "Start Task Manager queues")
	for queueName, svcName := range tskMgr.taskPublishers {
		log.Trace(tskmgrStartCtx, "Starting task producer ", "queue", queueName)
		tqSvc, err := tskmgrStartCtx.GetService(svcName)
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}
		tq, ok := tqSvc.(components.TaskQueue)
		if !ok {
			return errors.ThrowError(tskmgrStartCtx, errors.CORE_ERROR_BAD_CONF)
		}
		tskMgr.taskPublisherSvcs[queueName] = tq
	}

	for queueName, consumerName := range tskMgr.taskConsumerNames {
		log.Trace(tskmgrStartCtx, "Starting task consumer ", "queue", queueName)
		consumerSvc, err := tskmgrStartCtx.GetService(consumerName)
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}
		ts, ok := consumerSvc.(components.TaskServer)
		if !ok {
			return errors.BadConf(tskmgrStartCtx, constants.CONF_TASK_PROCESSOR, "queue", queueName)
		}

		err = ts.SubsribeQueue(tskmgrStartCtx, queueName)
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}

		procSvc, err := svcMgr.GetService(tskmgrStartCtx, tskMgr.taskProcessorNames[queueName])
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}
		tskMgr.taskProcessors[queueName] = procSvc
	}

	return nil
}

func (tskMgr *taskManager) pushTask(ctx core.RequestContext, queue string, taskData interface{}) error {
	tq, ok := tskMgr.taskPublisherSvcs[queue]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Missing Queue", queue)
	}
	log.Trace(ctx, "Pushing task to queue", "queue", queue)
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
		log.Debug(req, "Received job ")
		req.SetRequest(t.Data)
		req.Set(svc.authHeader, t.Token)
		svc.shandler.AuthenticateRequest(req)
		queueName := t.Queue
		q, ok := svc.queues[queueName]
		if ok {
			err := q.lstnr.Invoke(req)
			if err != nil {
				log.Error(req, "Error in background process", "err", err)
				return err
			}
		}*/
	queue := t.Queue
	processor, ok := tskMgr.taskProcessors[queue]
	if ok {
		/*req := ctx.CreateRequest()
		req.SetBody(t.Data)
		req.AddParam(tskMgr.authHeader, t.Token)*/
		tskMgr.shandler.AuthenticateRequest(ctx, true)
		log.Trace(ctx, "Processing background task")
		res, err := processor.HandleRequest(ctx, map[string]interface{}{tskMgr.authHeader: t.Token}, t.Data)
		ctx.SetResponse(res)
		return err
	}
	return nil
}
