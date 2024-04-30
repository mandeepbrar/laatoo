package core

import (
	"laatoo.io/sdk/config"
)

type ActionType string

type ActionExecutor func(ctx RequestContext, action *Action, params StringMap) (interface{}, error)

type Action struct {
	Type      ActionType
	Condition GenericExpression
	Config    config.GenericConfig
}
