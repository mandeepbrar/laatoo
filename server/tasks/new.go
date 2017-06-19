package tasks

import (
	"laatoo/server/common"
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

func NewTaskManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*taskManager, *taskManagerProxy) {
	tskMgr := &taskManager{parent: parentElem, taskProducerSvcs: make(map[string]components.TaskQueue, 10), taskProcessors: make(map[string]core.Service, 10),
		taskProducers: make(map[string]string, 10), taskReceiverNames: make(map[string]string, 10), taskProcessorNames: make(map[string]string, 10)}
	tskElemCtx := parentElem.NewCtx(name)
	tskElem := &taskManagerProxy{Context: tskElemCtx.(*common.Context), manager: tskMgr}
	tskMgr.proxy = tskElem
	return tskMgr, tskElem
}

func ChildTaskManager(ctx core.ServerContext, name string, parentTaskMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	childTaskMgr := &taskManager{parent: parent, taskProducerSvcs: make(map[string]components.TaskQueue, 10), taskProcessors: make(map[string]core.Service, 10),
		taskProducers: make(map[string]string, 10), taskReceiverNames: make(map[string]string, 10), taskProcessorNames: make(map[string]string, 10)}
	childtskMgrElemCtx := parentTaskMgr.NewCtx(name)
	childtskMgrElem := &taskManagerProxy{Context: childtskMgrElemCtx.(*common.Context), manager: childTaskMgr}
	childTaskMgr.proxy = childtskMgrElem
	return childTaskMgr, childtskMgrElem
}
