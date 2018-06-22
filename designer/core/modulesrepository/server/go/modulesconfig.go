package main

/*
import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

//"path/filepath"

type ModulesConfiguration struct {
	core.Service
	dataStore data.DataComponent
}

const (
	CONF_DATASTORE = "datastore"
	PARAM_MOD      = "module"
)

func (svc *ModulesConfiguration) Describe(ctx core.ServerContext) {
	svc.AddStringParam(ctx, PARAM_MOD)
}

func (svc *ModulesConfiguration) Start(ctx core.ServerContext) error {
	dataSvcName := "repository.modules.database"
	dataSvc, err := ctx.GetService(dataSvcName)
	if err != nil {
		return errors.MissingService(ctx, dataSvcName)
	}
	svc.dataStore = dataSvc.(data.DataComponent)
	return nil
}

func (svc *ModulesConfiguration) Invoke(ctx core.RequestContext) error {
	ctx = ctx.SubContext("ModulesConfiguration")
	mod, ok := ctx.GetStringParam(PARAM_MOD)
	log.Info(ctx, " Get Module Configuration", "Module", mod)
	if ok {
		mod, svc.dataStore.GetById(ctx, mod)
		_, err := svc.processModule(ctx, mod)
		return err
	}
	return nil
}
*/
