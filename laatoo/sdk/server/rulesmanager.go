package server

import (
	"laatoo/sdk/core"
	"laatoo/sdk/rules"
)

type RulesManager interface {
	core.ServerElement
	SubscribeEvent(ctx core.ServerContext, eventType string, eventObject string, rule rules.Rule)
	FireEvent(ctx core.RequestContext, eventType string, eventObject string, data map[string]interface{})
}
