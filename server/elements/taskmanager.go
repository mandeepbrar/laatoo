package elements

import (
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type TaskManager interface {
	core.ServerElement
	PushTask(ctx core.RequestContext, queue string, task interface{}) error
	ProcessTask(ctx core.ServerContext, task *components.Task) (interface{}, error)
	List(ctx core.ServerContext) map[string]string
}
