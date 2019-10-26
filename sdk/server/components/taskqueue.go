package components

import (
	"laatoo/sdk/server/core"
)

type Task struct {
	Queue string
	Data  []byte
	Token string
}

type TaskQueue interface {
	PushTask(ctx core.RequestContext, queue string, task *Task) error
}

type TaskServer interface {
	SubsribeQueue(ctx core.ServerContext, queue string) error
	UnsubsribeQueue(ctx core.ServerContext, queue string) error
}

type TaskProcessor func(ctx core.RequestContext, value *Task) (interface{}, error)
