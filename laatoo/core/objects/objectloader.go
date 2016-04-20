package objects

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type objectLoader struct {
	objectsFactoryRegister map[string]core.ObjectFactory
}

func (objLoader *objectLoader) Initialize(ctx core.ServerContext, conf config.Config) error {
	objectNames, ok := conf.GetStringArray(config.CONF_OBJECTLDR_OBJECTS)
	if ok {
		for _, objectName := range objectNames {
			fac, ok := objectsFactoryRegister[objectName]
			if ok {
				objLoader.registerObjectFactory(ctx, objectName, fac)
			} else {
				return errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
			}
		}
	}
	return nil
}

func (objLoader *objectLoader) Start(ctx core.ServerContext) error {
	return nil
}
func (objLoader *objectLoader) registerObject(ctx core.Context, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) {
	objLoader.registerObjectFactory(ctx, objectName, NewObjectFactory(objectCreator, objectCollectionCreator))
}
func (objLoader *objectLoader) registerObjectFactory(ctx core.Context, objectName string, factory core.ObjectFactory) {
	_, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		log.Logger.Info(ctx, "Registering object factory ", "Object Name", objectName)
		objLoader.objectsFactoryRegister[objectName] = factory
	}
}

//returns a collection of the object type
func (objLoader *objectLoader) createCollection(ctx core.Context, objectName string, args core.MethodArgs) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection(ctx, args)
}

//Provides an object with a given name
func (objLoader *objectLoader) createObject(ctx core.Context, objectName string, args core.MethodArgs) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
	}
	return factory.CreateObject(ctx, args)
}

func (objLoader *objectLoader) getObjectCreator(ctx core.Context, objectName string) (core.ObjectCreator, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObject, nil
}

func (objLoader *objectLoader) getObjectCollectionCreator(ctx core.Context, objectName string) (core.ObjectCollectionCreator, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection, nil
}
