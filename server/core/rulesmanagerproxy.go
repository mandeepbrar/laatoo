package core

import (
	"laatoo/sdk/server/components/rules"
	"laatoo/sdk/server/core"
)

type rulesManagerProxy struct {
	manager *rulesManager
}

func (rm *rulesManagerProxy) SubscribeSynchronousMessage(ctx core.ServerContext, msgType string, rule rules.Rule, ruleName string) {
	rm.manager.subscribeSynchronousMessage(ctx, msgType, rule, ruleName)
}

func (rm *rulesManagerProxy) SendSynchronousMessage(ctx core.RequestContext, msgType string, data interface{}) error {
	return rm.manager.sendSynchronousMessage(ctx, msgType, data)
}

func (proxy *rulesManagerProxy) Reference() core.ServerElement {
	return &rulesManagerProxy{manager: proxy.manager}
}
func (proxy *rulesManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *rulesManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *rulesManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementRulesManager
}
