package components

import (
	"context"

	"laatoo.io/sdk/server/core"
)

type Workflow func(ctx context.Context, input interface{}) error

type WorkflowService interface {
	RegisterWorkflow(ctx core.ServerContext, name string, workflowToRegister Workflow) error
	StartWorkflow(ctx core.RequestContext, workflowName string, initVal interface{}) (interface{}, error)
	SendSignal(ctx core.RequestContext, workflowref interface{}, signal string, signalVal interface{}) error
	CompleteActivity(ctx core.RequestContext, workflowRef interface{}, data interface{}, err error) error
}
