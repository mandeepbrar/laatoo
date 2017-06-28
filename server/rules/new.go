package rules

import (
	"laatoo/sdk/components/rules"
	"laatoo/sdk/core"
)

func NewRulesManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*rulesManager, *rulesManagerProxy) {
	rulesMgr := &rulesManager{parent: parentElem, registeredRules: make(map[string][]rules.Rule, 10), name: name}
	rulesElem := &rulesManagerProxy{manager: rulesMgr}
	rulesMgr.proxy = rulesElem
	return rulesMgr, rulesElem
}
