package tasks

import (
	"laatoo/core/common"
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

func NewTaskManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*taskManager, *taskManagerProxy) {
	tskMgr := &taskManager{parent: parentElem, taskQueueStore: make(map[string]components.TaskQueue, 10), taskQueueSvcs: make(map[string]string, 10)}
	tskElemCtx := parentElem.NewCtx(name)
	tskElem := &taskManagerProxy{Context: tskElemCtx.(*common.Context), manager: tskMgr}
	tskMgr.proxy = tskElem
	return tskMgr, tskElem
}

func ChildTaskManager(ctx core.ServerContext, name string, parentTaskMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	tskMgrProxy := parentTaskMgr.(*taskManagerProxy)
	tskMgr := tskMgrProxy.manager
	taskQueueStore := make(map[string]components.TaskQueue, len(tskMgr.taskQueueSvcs))
	taskQueueSvcs := make(map[string]string, len(tskMgr.taskQueueSvcs))
	for k, v := range tskMgr.taskQueueStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			taskQueueStore[k] = v
			taskQueueSvcs[k] = tskMgr.taskQueueSvcs[k]
		}
	}
	childTaskMgr := &taskManager{parent: parent, taskQueueStore: taskQueueStore, taskQueueSvcs: taskQueueSvcs}
	childtskMgrElemCtx := parentTaskMgr.NewCtx(name)
	childtskMgrElem := &taskManagerProxy{Context: childtskMgrElemCtx.(*common.Context), manager: childTaskMgr}
	childTaskMgr.proxy = childtskMgrElem
	return childTaskMgr, childtskMgrElem
}
