package cadence

import (
	"context"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"

	"go.uber.org/cadence/workflow"
)

type Workflow func(ctx workflow.Context, input interface{}) error

type WorkflowRegisterar interface {
	RegisterWorkflow(ctx core.ServerContext, name string, workflowToRegister Workflow)
}

type TaskProcessor func(ctx context.Context, value *components.Task) (interface{}, error)
