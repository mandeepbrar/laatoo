package laatoocore

import (
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
)

var (
	//global provider register
	//objects factory register exists for every server
	ObjectsFactoryRegister = utils.NewMemoryStorer()
)

//every object that needs to be created by configuration should register a factory func
//this factory function provides a object if called
type ObjectFactory func(conf map[string]interface{}) (interface{}, error)

//register the object factory in the global register
func RegisterObjectProvider(objectName string, factory ObjectFactory) {

	log.Logger.Infof("Registering object factory for %s", objectName)
	ObjectsFactoryRegister.PutObject(objectName, factory)
}

//Provides a object with a given name
func CreateObject(objectName string, confdata map[string]interface{}) (interface{}, error) {
	log.Logger.Debugf("Getting object %s", objectName)

	//get the factory from the register
	factoryInt, err := ObjectsFactoryRegister.GetObject(objectName)
	if err != nil {
		return nil, errors.RethrowError(CORE_ERROR_PROVIDER_NOT_FOUND, err, objectName)

	}
	//cast to a creatory func
	factoryFunc, ok := factoryInt.(ObjectFactory)
	if !ok {
		return nil, errors.RethrowError(CORE_ERROR_PROVIDER_NOT_FOUND, err, objectName)
	}

	log.Logger.Debugf("Creating object %s from factory", objectName)
	//call factory method for creating an object
	obj, err := factoryFunc(confdata)
	if err != nil {
		return nil, errors.RethrowError(CORE_ERROR_OBJECT_NOT_CREATED, err, objectName)
	}
	log.Logger.Debugf("Created object %s", objectName)
	//return object by calling factory func
	return obj, nil
}
