package rules

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/rules"
)

type rulesManagerProxy struct {
	*common.Context
	manager *rulesManager
}

func (rm *rulesManagerProxy) SubscribeEvent(ctx core.ServerContext, eventType string, eventObject string, rule rules.Rule) {
	rm.manager.subscribeEvent(ctx, eventType, eventObject, rule)
}

func (rm *rulesManagerProxy) FireEvent(ctx core.RequestContext, eventType string, eventObject string, data map[string]interface{}) {
	rm.manager.fireEvent(ctx, eventType, eventObject, data)
}
