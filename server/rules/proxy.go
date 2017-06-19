package rules

import (
	"laatoo/server/common"
	"laatoo/sdk/components/rules"
	"laatoo/sdk/core"
)

type rulesManagerProxy struct {
	*common.Context
	manager *rulesManager
}

func (rm *rulesManagerProxy) SubscribeSynchronousMessage(ctx core.ServerContext, msgType string, rule rules.Rule) {
	rm.manager.subscribeSynchronousMessage(ctx, msgType, rule)
}

func (rm *rulesManagerProxy) SendSynchronousMessage(ctx core.RequestContext, msgType string, data interface{}) error {
	return rm.manager.sendSynchronousMessage(ctx, msgType, data)
}
