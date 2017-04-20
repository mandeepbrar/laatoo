package objects

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

var (
	__regContext__ = common.NewContext("__registry__")
	//global objects register
	//objects factory register exists for every server
	objectsFactoryRegister   = make(map[string]core.ObjectFactory, 30)
	invokableMethodsRegister = make(map[string]core.ServiceFunc, 30)
)

func init() {
	Register("string", "")
}

func Register(objectName string, obj interface{}) {
	RegisterObjectFactory(objectName, NewObjectType(obj))
}

//register the object factory in the global register
func RegisterObjectFactory(objectName string, factory core.ObjectFactory) {
	_, ok := objectsFactoryRegister[objectName]
	if !ok {
		log.Logger.Info(__regContext__, "Registering object factory ", "Object Name", objectName)
		objectsFactoryRegister[objectName] = factory
	}
}

//register the object factory in the global register
func RegisterObject(objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) {
	RegisterObjectFactory(objectName, NewObjectFactory(objectCreator, objectCollectionCreator))
}

//register the object factory in the global register
func RegisterInvokableMethod(methodName string, method core.ServiceFunc) {
	_, ok := invokableMethodsRegister[methodName]
	if !ok {
		log.Logger.Debug(__regContext__, "Registering method ", "Method Name", methodName)
		invokableMethodsRegister[methodName] = method
	}
}
