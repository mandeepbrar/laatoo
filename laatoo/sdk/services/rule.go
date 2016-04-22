package services

import (
	"laatoo/sdk/core"
)

type TriggerType int

const (
	Message TriggerType = iota
	Event
)

type Trigger struct {
	TriggerType TriggerType
	Event       string
	EventObject string
	Data        map[string]interface{}
}

type Rule interface {
	Condition(ctx core.RequestContext, trigger *Trigger) bool
	Action(ctx core.RequestContext, trigger *Trigger) error
}
