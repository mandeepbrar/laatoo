package objects

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type objectLoader struct {
	objectsFactoryRegister   map[string]core.ObjectFactory
	invokableMethodsRegister map[string]core.ServiceFunc
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
	methodNames, ok := conf.GetStringArray(config.CONF_OBJECTLDR_METHODS)
	if ok {
		for _, methodName := range methodNames {
			method, ok := invokableMethodsRegister[methodName]
			if ok {
				objLoader.registerInvokableMethod(ctx, methodName, method)
			} else {
				return errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Method Name", methodName)
			}
		}
	}
	return nil
}

func (objLoader *objectLoader) Start(ctx core.ServerContext) error {
	return nil
}
func (objLoader *objectLoader) register(ctx core.Context, objectName string, object interface{}) {
	objLoader.registerObjectFactory(ctx, objectName, NewObjectType(object))
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
func (objLoader *objectLoader) registerInvokableMethod(ctx core.Context, methodName string, method core.ServiceFunc) {
	_, ok := objLoader.invokableMethodsRegister[methodName]
	if !ok {
		log.Logger.Debug(ctx, "Registering method ", "Method Name", methodName)
		objLoader.invokableMethodsRegister[methodName] = method
	}
}

//returns a collection of the object type
func (objLoader *objectLoader) createCollection(ctx core.Context, objectName string, length int) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection(length), nil
}

//Provides an object with a given name
func (objLoader *objectLoader) createObject(ctx core.Context, objectName string) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
	}
	return factory.CreateObject(), nil
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

func (objLoader *objectLoader) getMethod(ctx core.Context, methodName string) (core.ServiceFunc, error) {
	method, ok := objLoader.invokableMethodsRegister[methodName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Method Name", methodName)

	}
	return method, nil
}
