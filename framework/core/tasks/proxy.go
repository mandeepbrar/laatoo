package tasks

import (
	"laatoo/framework/core/common"
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
