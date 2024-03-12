package components

import (
	"context"

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

type WorkflowActivityType string

const (
	MANUAL    WorkflowActivityType = "manual"
	AUTOMATIC WorkflowActivityType = "automatic"
	DECISION  WorkflowActivityType = "decision"
)

type WorkflowManager interface {
	LoadWorkflows(ctx core.ServerContext, dir string) (map[string]Workflow, error)
	StartWorkflow(ctx core.RequestContext, workflowName string, initVal core.StringMap, insconf core.StringMap) (WorkflowInstance, error)
	IsWorkflowRegistered(ctx core.ServerContext, name string) bool
	SendSignal(ctx core.RequestContext, workflowId string, workflowIns string, actId string, signal string, signalVal core.StringMap) error
	CompleteActivity(ctx core.RequestContext, workflowId string, workflowIns string, actId string, data core.StringMap, err error) error
}
