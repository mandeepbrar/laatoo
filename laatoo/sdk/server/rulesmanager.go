package server

import (
	"laatoo/sdk/core"
	"laatoo/sdk/services"
)

type RulesManager interface {
	core.ServerElement
	SubscribeEvent(ctx core.ServerContext, eventType string, eventObject string, rule services.Rule)
	FireEvent(ctx core.RequestContext, eventType string, eventObject string, data map[string]interface{})
}
