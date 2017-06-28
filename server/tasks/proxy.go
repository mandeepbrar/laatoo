package tasks

import (
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	//	"laatoo/sdk/server"
)

type taskManagerProxy struct {
	manager *taskManager
}

func (mgr *taskManagerProxy) PushTask(ctx core.RequestContext, queue string, task interface{}) error {
	return mgr.manager.pushTask(ctx, queue, task)
}

func (mgr *taskManagerProxy) ProcessTask(ctx core.RequestContext, task *components.Task) error {
	return mgr.manager.processTask(ctx, task)
}
func (proxy *taskManagerProxy) Reference() core.ServerElement {
	return &taskManagerProxy{manager: proxy.manager}
}
func (proxy *taskManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *taskManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *taskManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementTaskManager
}
