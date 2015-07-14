package ginauth

import (
	"fmt"
)

var registeredModules = make(map[string]interface{})

type Module interface {
	Initialize(app *App) error
	Serve() error
}

func RegisterModule(name string, mod interface{}) {
	Logger.Infoln("GinAuth: Registering Module", name)
	registeredModules[name] = mod
}

func (app *App) configModules() error {
	for name, moduleInt := range registeredModules {
		module := moduleInt.(Module)
		app.Logger.Infoln("Ginauth: Initializing module ", name)
		if err := module.Initialize(app); err != nil {
			return fmt.Errorf("GinAuth: [%s] Error Initializing: %v", name, err)
		}
	}

	for name, moduleInt := range registeredModules {
		module := moduleInt.(Module)
		app.Logger.Infoln("Ginauth: Starting module ", name)
		if err := module.Serve(); err != nil {
			return fmt.Errorf("GinAuth: [%s] Module Error: %v", name, err)
		}
	}
	return nil
}
