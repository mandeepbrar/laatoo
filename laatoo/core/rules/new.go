package rules

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/rules"
)

func NewRulesManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*rulesManager, *rulesManagerProxy) {
	rulesMgr := &rulesManager{registeredRules: make(map[string][]rules.Rule, 10)}
	rulesElemCtx := parentElem.NewCtx("Rules Manager" + name)
	rulesElem := &rulesManagerProxy{Context: rulesElemCtx.(*common.Context), manager: rulesMgr}
	rulesMgr.proxy = rulesElem
	return rulesMgr, rulesElem
}
