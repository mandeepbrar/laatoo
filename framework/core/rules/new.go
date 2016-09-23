package rules

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/components/rules"
	"laatoo/sdk/core"
)

func NewRulesManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*rulesManager, *rulesManagerProxy) {
	rulesMgr := &rulesManager{registeredRules: make(map[string][]rules.Rule, 10)}
	rulesElemCtx := parentElem.NewCtx("Rules Manager" + name)
	rulesElem := &rulesManagerProxy{Context: rulesElemCtx.(*common.Context), manager: rulesMgr}
	rulesMgr.proxy = rulesElem
	return rulesMgr, rulesElem
}