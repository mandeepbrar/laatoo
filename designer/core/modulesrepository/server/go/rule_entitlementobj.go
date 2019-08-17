package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/components/rules"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type NewEntitlementRule struct {
	dataStore data.DataComponent
}

func (rule *NewEntitlementRule) Initialize(ctx ctx.Context, conf config.Config) error {
	dataSvcName := "repository.modules.database"
	dataSvc, err := ctx.(core.ServerContext).GetService(dataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	rule.dataStore = dataSvc.(data.DataComponent)
	log.Error(ctx, "New entitlement rule initialized")
	return nil
}

func (rule *NewEntitlementRule) Condition(ctx core.RequestContext, trigger *rules.Trigger) bool {

	if trigger.Message != nil {
		_, ok := trigger.Message.(*Entitlement)
		if ok {
			return true
		}
	}
	return false
}

func (rule *NewEntitlementRule) Action(ctx core.RequestContext, trigger *rules.Trigger) error {
	ent, _ := trigger.Message.(*Entitlement)
	if ent.Module.Id == "" {
		mod := &ModuleDefinition{Name: ent.Name, Version: "0.0.1"}
		mod.SetId(ent.Name)
		err := rule.dataStore.Put(ctx, ent.Name, mod)
		if err != nil {
			return err
		}
		modRef := data.StorableRef{Id: ent.Name, Name: ent.Name, Type: "ModuleDefinition"}
		ent.Module = modRef
	}
	return nil
}
