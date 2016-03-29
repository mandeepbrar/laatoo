package registry

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

var (
	__regContext__ = common.NewContext("__registry__")
	//global objects register
	//objects factory register exists for every server
	ObjectsFactoryRegister = make(map[string]core.ObjectFactory, 30)
)

//register the object factory in the global register
func RegisterObjectFactory(objectName string, factory core.ObjectFactory) {
	_, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		log.Logger.Info(__regContext__, "Registering object factory ", "Object Name", objectName)
		ObjectsFactoryRegister[objectName] = factory
	}
}

//returns a collection of the object type
func CreateCollection(ctx core.Context, objectName string, args core.MethodArgs) (interface{}, error) {
	//get the factory object from the register
	factory, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection(ctx, args)
}

//Provides an object with a given name
func CreateObject(ctx core.Context, objectName string, args core.MethodArgs) (interface{}, error) {
	//get the factory object from the register
	factory, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObject(ctx, args)
}

func GetObjectCreator(ctx core.Context, objectName string) (core.ObjectCreator, error) {
	//get the factory object from the register
	factory, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObject, nil
}

func GetObjectCollectionCreator(ctx core.Context, objectName string) (core.ObjectCollectionCreator, error) {
	//get the factory object from the register
	factory, ok := ObjectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection, nil
}
