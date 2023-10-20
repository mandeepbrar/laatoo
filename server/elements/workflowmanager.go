package elements

import (
	"laatoo.io/sdk/config"
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
)

type WorkflowManager interface {
	core.ServerElement
	StartWorkflow(ctx core.RequestContext, workflowName string, initVal config.Config) (components.WorkflowInstance, error)
	RegisterWorkflow(ctx core.ServerContext, name string, workflowToRegister components.Workflow) error
	SendSignal(ctx core.RequestContext, workflowref components.WorkflowInstance, signal string, signalVal config.Config) error
	CompleteActivity(ctx core.RequestContext, workflowRef components.WorkflowInstance, act components.WorkflowActivity, data config.Config, err error) error
}
