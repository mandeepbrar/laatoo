package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/modulesrepository"
	"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/components/rules"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type ModuleDesignEntitiesRule struct {
	dataStore data.DataComponent
}

func (rule *ModuleDesignEntitiesRule) Initialize(ctx ctx.Context, conf config.Config) error {
	dataSvcName := "dataadapter.dataservice.ModuleDesignEntities"
	dataSvc, err := ctx.(core.ServerContext).GetService(dataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	rule.dataStore = dataSvc.(data.DataComponent)
	log.Error(ctx, "New module design rule initialized")
	return nil
}

func (rule *ModuleDesignEntitiesRule) Condition(ctx core.RequestContext, trigger *rules.Trigger) bool {

	if trigger.Message != nil {
		_, ok := trigger.Message.(*modulesrepository.Entitlement)
		if ok {
			return true
		}
	}
	return false
}

func (rule *ModuleDesignEntitiesRule) Action(ctx core.RequestContext, trigger *rules.Trigger) error {
	ent, _ := trigger.Message.(*modulesrepository.Entitlement)
	if ent.Local == true {
		mod := &ModuleDesignEntities{Name: ent.Name}
		mod.SetId(ent.Name)
		mod.SetTenant(ctx.GetUser().GetTenant())
		err := rule.dataStore.Put(ctx, ent.Name, mod)
		if err != nil {
			return err
		}
	}
	return nil
}
