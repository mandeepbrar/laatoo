package components

import (
	"laatoo.io/sdk/datatypes"
	"laatoo.io/sdk/server/auth"
	"laatoo.io/sdk/server/core"
)

type Task struct {
	datatypes.Serializable
	Queue  string
	Data   []byte
	Id     string
	User   auth.User
	Tenant auth.TenantInfo
}

type TaskManager interface {
	PushTask(ctx core.RequestContext, task *Task) error
	SubsribeQueue(ctx core.ServerContext, queue string) error
	UnsubsribeQueue(ctx core.ServerContext, queue string) error
}
