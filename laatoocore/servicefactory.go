package laatoocore

import (
	"fmt"
	"laatoosdk/log"
	"laatoosdk/utils"
)

var (
	//global provider register
	//service factory register exists for every server
	ServiceFactoryRegister = utils.NewMemoryStorer()
)

//every service should register a factory func
//this factory function provides a service if called
type ServiceFactory func(alias string, conf map[string]interface{}) (interface{}, error)

//register the service factory in the global register
func RegisterServiceProvider(serviceName string, factory ServiceFactory) {
	log.Logger.Infof("Registering service %s", serviceName)
	ServiceFactoryRegister.PutObject(serviceName, factory)
}

//Provides a service with a given name and alias
func GetService(serviceName string, alias string, confdata map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Getting service %s with alias %s", serviceName, alias)
	//get the factory from the register
	factoryInt, err := ServiceFactoryRegister.GetObject(serviceName)
	if err != nil {
		return nil, fmt.Errorf("Service Not Found : %s", err)
	}
	//cast to a creatory func
	factoryFunc, ok := factoryInt.(ServiceFactory)
	if !ok {
		return nil, fmt.Errorf("Incorrect provider registered %s", serviceName)
	}
	log.Logger.Debugf("Creating service %s from factory for alias %s", serviceName, alias)
	//call factory method for creating an object
	obj, err := factoryFunc(alias, confdata)
	if err != nil {
		return nil, err
	}
	log.Logger.Infof("Created service %s", alias)
	//return service by calling factory func
	return obj, nil
}
