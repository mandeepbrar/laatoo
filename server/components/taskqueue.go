package components

import (
	"laatoo.io/sdk/server/components/data"
	"laatoo.io/sdk/server/core"
)

type Task struct {
	Queue  string
	Data   []byte
	Id     string
	User   core.Serializable
	Tenant data.TenantInfo
}

type TaskManager interface {
	PushTask(ctx core.RequestContext, task *Task) error
	SubsribeQueue(ctx core.ServerContext, queue string) error
	UnsubsribeQueue(ctx core.ServerContext, queue string) error
}
