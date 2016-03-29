package registry

import (
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

var (
	//global provider register
	//objects factory register exists for every server
	ServiceFactoryProviders = make(map[string]core.ServiceFactoryProvider, 50)
)

//register the invokable method in the global register
func RegisterServiceFactoryProvider(factory string, factoryProvider core.ServiceFactoryProvider) {
	_, ok := ServiceFactoryProviders[factory]
	if !ok {
		log.Logger.Info(__regContext__, "Registering service factory provider", "Factory", factory)
		ServiceFactoryProviders[factory] = factoryProvider
	}
}
