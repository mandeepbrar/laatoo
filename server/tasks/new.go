package tasks

import (
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

func NewTaskManager(ctx core.ServerContext, name string) (*taskManager, *taskManagerProxy) {
	tskMgr := &taskManager{name: name, taskPublisherSvcs: make(map[string]components.TaskQueue, 10), taskProcessors: make(map[string]server.Service, 10),
		taskPublishers: make(map[string]string, 10), taskConsumerNames: make(map[string]string, 10), taskProcessorNames: make(map[string]string, 10)}
	tskElem := &taskManagerProxy{manager: tskMgr}
	tskMgr.proxy = tskElem
	return tskMgr, tskElem
}
