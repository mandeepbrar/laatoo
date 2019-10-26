package cadence

import (
	"laatoo/sdk/server/core"

	"go.uber.org/cadence/workflow"
)

type CadenceWorkflow func(ctx workflow.Context, reqCtx core.RequestContext, input interface{}) error

type CadenceWorkflowRegisterar interface {
	RegisterWorkflow(ctx core.ServerContext, name string, workflowToRegister CadenceWorkflow)
}
