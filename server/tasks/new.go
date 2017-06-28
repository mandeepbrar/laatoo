package tasks

import (
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

func NewTaskManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*taskManager, *taskManagerProxy) {
	tskMgr := &taskManager{name: name, parent: parentElem, taskPublisherSvcs: make(map[string]components.TaskQueue, 10), taskProcessors: make(map[string]core.Service, 10),
		taskPublishers: make(map[string]string, 10), taskConsumerNames: make(map[string]string, 10), taskProcessorNames: make(map[string]string, 10)}
	tskElem := &taskManagerProxy{manager: tskMgr}
	tskMgr.proxy = tskElem
	return tskMgr, tskElem
}

func ChildTaskManager(ctx core.ServerContext, name string, parentTaskMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	childTaskMgr := &taskManager{name: name, parent: parent, taskPublisherSvcs: make(map[string]components.TaskQueue, 10), taskProcessors: make(map[string]core.Service, 10),
		taskPublishers: make(map[string]string, 10), taskConsumerNames: make(map[string]string, 10), taskProcessorNames: make(map[string]string, 10)}
	childtskMgrElem := &taskManagerProxy{manager: childTaskMgr}
	childTaskMgr.proxy = childtskMgrElem
	return childTaskMgr, childtskMgrElem
}
