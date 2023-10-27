package components

import (
	"context"

	"laatoo.io/sdk/config"
	"laatoo.io/sdk/server/core"
)

type Workflow interface {
	GetId() string
	Spec(ctx context.Context) interface{}
	Type() string
	GetName() string
}

type WorkflowInstance interface {
	GetId() string
	InstanceDetails() config.Config
	GetWorkflow() string
	GetStatus() map[string]interface{}
}

type WorkflowActivityType int

const (
	MANUAL WorkflowActivityType = iota
	AUTOMATIC
)

type WorkflowActivity interface {
	GetId() string
	GetName() string
	GetActivityType() WorkflowActivityType
	GetWorkflowInstance() WorkflowInstance
}

type WorkflowManager interface {
	LoadWorkflows(ctx core.ServerContext, dir string) (map[string]Workflow, error)
	StartWorkflow(ctx core.RequestContext, workflowName string, initVal config.Config) (WorkflowInstance, error)
	IsWorkflowRegistered(ctx core.ServerContext, name string) bool
	RegisterWorkflow(ctx core.ServerContext, name string, workflowToRegister Workflow) error
	SendSignal(ctx core.RequestContext, workflowref WorkflowInstance, signal string, signalVal config.Config) error
	CompleteActivity(ctx core.RequestContext, workflowRef WorkflowInstance, act WorkflowActivity, data config.Config, err error) error
}
