package service

import (
	"fmt"
	"laatoosdk/log"
)

//Environment hosting an application
type ModuleHelper struct {
	//sub modules registered with a service
	RegisteredModules map[string]Service
}

//Initialize the service. Consumer of a service passes the data
func (mods *ModuleHelper) InitializeModules(ctx interface{}) error {
	for name, moduleInt := range mods.RegisteredModules {
		module := moduleInt.(Service)
		log.Logger.Infoln("ModuleHelper: Initializing module ", name)
		if err := module.Initialize(ctx); err != nil {
			return fmt.Errorf("ModuleHelper: [%s] Error initializing module: %v", name, err)
		}
	}
	return nil
}

//The service starts serving when this method is called
func (mods *ModuleHelper) Serve() error {
	for name, moduleInt := range mods.RegisteredModules {
		module := moduleInt.(Service)
		log.Logger.Infoln("ModuleHelper: Starting module ", name)
		if err := module.Serve(); err != nil {
			return fmt.Errorf("ModuleHelper: [%s] Module Error: %v", name, err)
		}
	}
	return nil
}

//Module support for any service
func NewModuleHelper() *ModuleHelper {
	mod := &ModuleHelper{}
	mod.RegisteredModules = make(map[string]Service)
	return mod
}
