package components

import (
	"context"

	"laatoo.io/sdk/config"
)

type Workflow interface {
	Spec(ctx context.Context) interface{}
	GetName() string
}

type WorkflowInstance interface {
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
