package core

/*
import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

const (
	CONF_MODULESSERVICE = "MODULES_SVC"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_MODULESSERVICE, Object: ModulesService{}}}
}

type ModulesService struct {
	core.Service
}

func (ls *ModulesService) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

//Expects Local user to be provided inside the request
func (ls *ModulesService) Invoke(ctx core.RequestContext) error {
	ptyp, ok := ctx.GetStringParam("type")
  if(ok) {
    errors.BadRequest(ctx, "Missing Param", "type")
  }
	return nil
}
*/
