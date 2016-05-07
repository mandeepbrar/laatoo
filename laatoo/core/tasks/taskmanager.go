package tasks

import (
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type taskManager struct {
	parent         core.ServerElement
	proxy          server.TaskManager
	taskQueueSvcs  map[string]string
	taskQueueStore map[string]components.TaskQueue
}

func (tskMgr *taskManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	tskmgrInitializeCtx := tskMgr.createContext(ctx, "Initialize task manager")
	log.Logger.Trace(tskmgrInitializeCtx, "Create Task Manager queues")
	err := tskMgr.createTaskQueues(tskmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(tskmgrInitializeCtx, err)
	}
	return nil
}

func (tskMgr *taskManager) Start(ctx core.ServerContext) error {
	tskmgrStartCtx := tskMgr.createContext(ctx, "Start task manager")
	log.Logger.Trace(tskmgrStartCtx, "Start Task Manager queues")
	for queueName, svcName := range tskMgr.taskQueueSvcs {
		log.Logger.Trace(tskmgrStartCtx, "Starting queue", "queue", queueName)
		tqSvc, err := tskmgrStartCtx.GetService(svcName)
		if err != nil {
			return errors.WrapError(tskmgrStartCtx, err)
		}
		tq, ok := tqSvc.(components.TaskQueue)
		if !ok {
			return errors.ThrowError(tskmgrStartCtx, errors.CORE_ERROR_BAD_CONF)
		}
		tskMgr.taskQueueStore[queueName] = tq
	}
	return nil
}

func (tskMgr *taskManager) createTaskQueues(ctx core.ServerContext, conf config.Config) error {
	queuesConf, ok := conf.GetSubConfig(config.CONF_TASK_QUEUES)
	if ok {
		queueNames := queuesConf.AllConfigurations()
		for _, queueName := range queueNames {
			queueSvcName, _ := queuesConf.GetString(queueName)
			tskMgr.taskQueueSvcs[queueName] = queueSvcName
		}
	}
	return nil
}

func (tskMgr *taskManager) pushTask(ctx core.RequestContext, queue string, task interface{}) error {
	tq, ok := tskMgr.taskQueueStore[queue]
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "Missing Queue", queue)
	}
	log.Logger.Trace(ctx, "Pushing task to queue", "queue", queue)
	return tq.PushTask(ctx, queue, task)
}

//creates a context specific to factory manager
func (tskMgr *taskManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementTaskManager: tskMgr.proxy}, core.ServerElementTaskManager)
}
