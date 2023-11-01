package components

import (
	"context"

	"laatoo.io/sdk/config"
	"laatoo.io/sdk/server/core"
)

type Workflow interface {
	Spec(ctx context.Context) interface{}
	Type() string
	GetName() string
}

type WorkflowInstance interface {
	GetId() string
	InstanceDetails() core.StringMap
	GetWorkflow() string
	GetStatus() core.StringMap
	InitData() core.StringMap
}

type WorkflowActivityType int

const (
	MANUAL WorkflowActivityType = iota
	AUTOMATIC
)

type WorkflowManager interface {
	LoadWorkflows(ctx core.ServerContext, dir string) (map[string]Workflow, error)
	StartWorkflow(ctx core.RequestContext, workflowName string, initVal core.StringMap) (WorkflowInstance, error)
	IsWorkflowRegistered(ctx core.ServerContext, name string) bool
	SendSignal(ctx core.RequestContext, workflowref WorkflowInstance, signal string, signalVal core.StringMap) error
	CompleteActivity(ctx core.RequestContext, workflowRef WorkflowInstance, actId string,, data core.StringMap, err error) error
}
