package tunnymemtasks

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/log"
)

const (
	CONF_TUNNY_TASKPROCESSOR = "tunnytasks"
)

type MemTaskProcessor struct {
	core.Service
	queuePools  map[string]*Pool
	taskManager elements.TaskManager
	PoolSize    int
}

func (processor *MemTaskProcessor) Initialize(ctx core.ServerContext, conf config.Config) error {
	processor.queuePools = make(map[string]*Pool)
	processor.taskManager = ctx.GetServerElement(core.ServerElementTaskManager).(elements.TaskManager)
	return nil
}

func (processor *MemTaskProcessor) PushTask(ctx core.RequestContext, queue string, t *components.Task) error {
	pool, ok := processor.queuePools[queue]
	if ok {
		_ = pool.Process(t)
	}
	return nil
}

func (processor *MemTaskProcessor) processTask(ctx core.ServerContext, task *components.Task) {
	req, err := ctx.CreateNewRequest("Tunny task "+task.Queue, nil, nil, "")
	if err != nil {
		log.Error(req, "Error in background process", "task", task.Queue, "err", err)
		//				return err
	}

	err = processor.taskManager.ProcessTask(req, task)
	if err != nil {
		log.Error(req, "Error in background process", "task", task.Queue, "err", err)
		log.Trace(req, "Error in processing task", "task", task)
		//	return err
	}
}

func (processor *MemTaskProcessor) SubsribeQueue(ctx core.ServerContext, queue string) error {

	pool := NewFunc(processor.PoolSize, func(payload interface{}) interface{} {
		t := payload.(*components.Task)
		log.Info(ctx, "Recieved Task")
		go processor.processTask(ctx, t)
		return nil
	})
	processor.queuePools[queue] = pool
	return nil
}

func (processor *MemTaskProcessor) UnsubsribeQueue(ctx core.ServerContext, queue string) error {
	delete(processor.queuePools, queue)
	return nil
}
