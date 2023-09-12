package components

import (
	"laatoo.io/sdk/server/auth"
	"laatoo.io/sdk/server/core"
)

type Task struct {
	Queue  string
	Data   []byte
	User   auth.User
	Tenant auth.TenantInfo
}

type TaskQueue interface {
	PushTask(ctx core.RequestContext, queue string, task *Task) error
}

type TaskServer interface {
	SubsribeQueue(ctx core.ServerContext, queue string) error
	UnsubsribeQueue(ctx core.ServerContext, queue string) error
}
