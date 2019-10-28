package cadence

import (
	"context"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"

	"go.uber.org/cadence/workflow"
)

type Workflow func(ctx workflow.Context, input interface{}) error

func RegisterWorkflow(ctx core.ServerContext, name string, workflowToRegister Workflow) {
	workflow.RegisterWithOptions(workflowToRegister, workflow.RegisterOptions{Name: name})
}

type TaskProcessor func(ctx context.Context, value *components.Task) (interface{}, error)

type CadenceWorkflowInitiator interface {
	StartWorkflow(ctx core.RequestContext, workflowName string, initVal interface{}) error
	StartWorkflowOnTasklist(ctx core.RequestContext, workflowName, tasklistName string, initVal interface{}) error
}