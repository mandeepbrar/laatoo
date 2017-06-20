package tasks

import (
	"laatoo/server/common"
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	//	"laatoo/sdk/server"
)

type taskManagerProxy struct {
	*common.Context
	manager *taskManager
}

func (mgr *taskManagerProxy) PushTask(ctx core.RequestContext, queue string, task interface{}) error {
	return mgr.manager.pushTask(ctx, queue, task)
}

func (mgr *taskManagerProxy) ProcessTask(ctx core.RequestContext, task *components.Task) error {
	return mgr.manager.processTask(ctx, task)
}