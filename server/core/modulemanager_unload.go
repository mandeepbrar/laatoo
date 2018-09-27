package core

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
)

func (modMgr *moduleManager) UnloadModule(ctx core.ServerContext, modName string) error {
	for name, mod := range modMgr.modules {
		log.Error(ctx, "Unload module ", "name", name, "mod", mod)
	}
	return nil
}

func (modMgr *moduleManager) UnloadInstance(ctx core.ServerContext, instance string) error {
	for name, mod := range modMgr.modules {
		log.Error(ctx, "Unload module ", "name", name, "mod", mod)
	}
	return nil
}