package objects

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
)

var (
	__regContext__ = common.NewContext("__registry__")
	//global objects register
	//objects factory register exists for every server
	objectsFactoryRegister = make(map[string]core.ObjectFactory, 30)
)

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
