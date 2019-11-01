package core

import (
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

type taskManager struct {
	name                        string
	proxy                       elements.TaskManager
	authHeader                  string
	shandler                    elements.SecurityHandler
	taskPublishers              map[string]string
	taskPublisherSvcs           map[string]components.TaskQueue
	taskConsumerNames           map[string]string
	taskProcessorNames          map[string]string
	taskProcessors              map[string]elements.Service
	svrContext                  core.ServerContext
	defaultTaskPublisherSvcName string
	defaultTaskConsumerSvcName  string
}

func (tskMgr *taskManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(elements.SecurityHandler)
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

	taskMgrConf, err, taskConfok := common.ConfigFileAdapter(tskmgrInitializeCtx, conf, constants.CONF_TASKS)
	if err != nil {
		return errors.WrapError(tskmgrInitializeCtx, err)
	}
	log.Trace(tskmgrInitializeCtx, "Create Task Manager queues", "taskMgrConf", taskMgrConf, "conf", conf)

	if taskConfok {
		tskMgr.defaultTaskPublisherSvcName, _ = taskMgrConf.GetString(ctx, constants.CONF_TASK_DEFAULTPUBLISHER)

		tskMgr.defaultTaskConsumerSvcName, _ = taskMgrConf.GetString(ctx, constants.CONF_TASK_DEFAULTCONSUMER)
	}

	modManager := ctx.GetServerElement(core.ServerElementModuleManager).(*moduleManagerProxy).modMgr

	if err := modManager.loadTasks(tskmgrInitializeCtx, tskMgr.processTaskConf); err != nil {
		return err
	}

	if taskConfok {
		tasksConf, ok := taskMgrConf.GetSubConfig(ctx, constants.CONF_TASKS)
		if ok {
			taskNames := tasksConf.AllConfigurations(tskmgrInitializeCtx)
			for _, taskName := range taskNames {
				taskConf, _ := tasksConf.GetSubConfig(tskmgrInitializeCtx, taskName)
				tskCtx := tskmgrInitializeCtx.SubContext(taskName)
				if err := tskMgr.processTaskConf(tskCtx, taskConf, taskName); err != nil {
					return errors.WrapError(tskCtx, err)
				}
			}
		}
	}

	baseDir, _ := ctx.GetString(config.BASEDIR)

	return tskMgr.processTasksFromFolder(ctx, baseDir)
}

func (tskMgr *taskManager) processTasksFromFolder(ctx core.ServerContext, folder string) error {
	objs, err := tskMgr.loadTasksFromDirectory(ctx, folder)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, objs, tskMgr.processTaskConf); err != nil {
		return err
	}
	return nil
}

func (tskMgr *taskManager) loadTasksFromDirectory(ctx core.ServerContext, folder string) (map[string]config.Config, error) {
	return common.ProcessDirectoryFiles(ctx, folder, constants.CONF_TASKS, true)
}

func (tskMgr *taskManager) unloadModuleTasks(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("unload tasks")
	if err := common.ProcessObjects(ctx, mod.tasks, tskMgr.unloadTaskProcessor); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (tskMgr *taskManager) unloadTaskProcessor(ctx core.ServerContext, conf config.Config, taskName string) error {
	unloadCtx := ctx.SubContext("Unload task processor")
	log.Error(ctx, "Unloading task", "taskName", taskName)

	queueName, ok := conf.GetString(unloadCtx, constants.CONF_TASKS_QUEUE)
	if !ok {
		return errors.MissingConf(unloadCtx, constants.CONF_TASKS_QUEUE, "Task Name", taskName)
	}

	consumer, ok := conf.GetString(unloadCtx, constants.CONF_TASK_CONSUMER)
	if !ok {
		return errors.MissingConf(unloadCtx, constants.CONF_TASK_CONSUMER, "Task Name", taskName)
	}

	consumerSvc, err := unloadCtx.GetService(consumer)
	if err != nil {
		return errors.WrapError(unloadCtx, err)
	}
	ts, ok := consumerSvc.(components.TaskServer)
	if !ok {
		return errors.BadConf(unloadCtx, constants.CONF_TASK_PROCESSOR, "queue", queueName)
	}

	err = ts.UnsubsribeQueue(unloadCtx, queueName)
	if err != nil {
		return errors.WrapError(unloadCtx, err)
	}

	delete(tskMgr.taskConsumerNames, queueName)
	delete(tskMgr.taskProcessorNames, queueName)
	delete(tskMgr.taskPublishers, queueName)
	delete(tskMgr.taskPublisherSvcs, queueName)
	delete(tskMgr.taskProcessors, queueName)
	return nil
}

func (tskMgr *taskManager) Start(ctx core.ServerContext) error {
	tskmgrStartCtx := ctx.SubContext("Start task manager")
	log.Trace(tskmgrStartCtx, "Start Task Manager queues")
	for queueName, svcName := range tskMgr.taskPublishers {
		if err := tskMgr.startPublisher(tskmgrStartCtx, queueName, svcName); err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	for queueName, consumerName := range tskMgr.taskConsumerNames {
		if err := tskMgr.startConsumer(tskmgrStartCtx, queueName, consumerName); err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	for queueName, processorName := range tskMgr.taskProcessorNames {
		if err := tskMgr.startProcessor(tskmgrStartCtx, queueName, processorName); err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	return nil
}

func (tskMgr *taskManager) startPublisher(ctx core.ServerContext, queueName, svcName string) error {
	log.Trace(ctx, "Starting task producer ", "queue", queueName)
	tqSvc, err := ctx.GetService(svcName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Trace(ctx, "Got service for producer ", "queue", queueName, "svc", tqSvc)
	tq, ok := tqSvc.(components.TaskQueue)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF)
	}
	log.Trace(ctx, "Got task queue for producer ", "queue", queueName, "tq", tq)
	tskMgr.taskPublisherSvcs[queueName] = tq
	log.Trace(ctx, "Got task queue for producer ", "queue", queueName, "tskMgr.taskPublisherSvcs", tskMgr.taskPublisherSvcs)
	return nil
}

func (tskMgr *taskManager) startConsumer(ctx core.ServerContext, queueName, consumerName string) error {
	log.Trace(ctx, "Starting task consumer ", "queue", queueName)

	consumerSvc, err := ctx.GetService(consumerName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	ts, ok := consumerSvc.(components.TaskServer)
	if !ok {
		return errors.BadConf(ctx, constants.CONF_TASK_PROCESSOR, "queue", queueName)
	}

	//start task subscription in context of service context
	svcContext, _ := ctx.GetServiceContext(consumerName)
	consumerContext := svcContext.SubContext("Task Consumer: " + queueName)
	err = ts.SubsribeQueue(consumerContext, queueName)
	if err != nil {
		return errors.WrapError(consumerContext, err)
	}

	return nil
}

func (tskMgr *taskManager) startProcessor(ctx core.ServerContext, queueName, processorName string) error {
	svcMgr := ctx.GetServerElement(core.ServerElementServiceManager).(elements.ServiceManager)
	procSvc, err := svcMgr.GetService(ctx, processorName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	tskMgr.taskProcessors[queueName] = procSvc
	log.Debug(ctx, "Processor assigned to queue", "queue", queueName)
	return nil
}

func (tskMgr *taskManager) pushTask(ctx core.RequestContext, queue string, taskData interface{}) error {
	tq, ok := tskMgr.taskPublisherSvcs[queue]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Missing Queue", queue, "svcs", tskMgr.taskPublisherSvcs)
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
		/****TODO****error handling****/
		_, err := tskMgr.shandler.AuthenticateRequest(ctx, true)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		log.Error(ctx, "Processing background task", "Task", t)
		res, err := processor.HandleRequest(ctx, map[string]interface{}{tskMgr.authHeader: t.Token, "Task": t.Data})
		ctx.SetResponse(res)
		return err
	} else {
		return errors.BadConf(ctx, "No processor assigned to queue for processing task ", "queue", queue)
	}
	return nil
}

func (tskMgr *taskManager) startModuleInstanceTasks(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("load tasks")
	if err := common.ProcessObjects(ctx, mod.tasks, tskMgr.loadModuleInstanceTask); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (tskMgr *taskManager) loadModuleInstanceTask(ctx core.ServerContext, conf config.Config, taskName string) error {
	if err := tskMgr.loadTask(ctx, conf, taskName, true); err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (tskMgr *taskManager) processTaskConf(ctx core.ServerContext, conf config.Config, taskName string) error {
	if err := tskMgr.loadTask(ctx, conf, taskName, false); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (tskMgr *taskManager) loadTask(ctx core.ServerContext, conf config.Config, taskName string, start bool) error {
	queueName, ok := conf.GetString(ctx, constants.CONF_TASKS_QUEUE)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASKS_QUEUE, "Task Name", taskName)
	}

	consumer, consumerConfigured := conf.GetString(ctx, constants.CONF_TASK_CONSUMER)
	if !consumerConfigured {
		log.Warn(ctx, "Task  consumer not configured for task", "Task Name", taskName)
		//return errors.MissingConf(ctx, constants.CONF_TASK_CONSUMER, "Task Name", taskName)
		if tskMgr.defaultTaskConsumerSvcName != "" {
			tskMgr.taskConsumerNames[queueName] = tskMgr.defaultTaskConsumerSvcName
		}
	} else {
		tskMgr.taskConsumerNames[queueName] = consumer
	}

	processor, ok := conf.GetString(ctx, constants.CONF_TASK_PROCESSOR)
	if !ok {
		return errors.MissingConf(ctx, constants.CONF_TASK_PROCESSOR, "Task Name", taskName)
	}
	tskMgr.taskProcessorNames[queueName] = processor

	publisher, publisherConfigured := conf.GetString(ctx, constants.CONF_TASK_PUBLISHER)
	if !publisherConfigured {
		log.Warn(ctx, "Task publisher not configured for task.", "Task Name", taskName)
		//		return errors.MissingConf(ctx, constants.CONF_TASK_PUBLISHER, "Task Name", taskName)
		if tskMgr.defaultTaskPublisherSvcName != "" {
			tskMgr.taskPublishers[queueName] = tskMgr.defaultTaskPublisherSvcName
		}
	} else {
		tskMgr.taskPublishers[queueName] = publisher
	}
	if start {
		if publisherConfigured {
			if err := tskMgr.startPublisher(ctx, queueName, publisher); err != nil {
				return errors.WrapError(ctx, err)
			}
		}
		if consumerConfigured {
			if err := tskMgr.startConsumer(ctx, queueName, consumer); err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}
