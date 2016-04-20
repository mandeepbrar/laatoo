package objects

import (
	"laatoo/sdk/core"
)

//service method for doing various tasks
func NewObjectFactory(objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) core.ObjectFactory {
	if objectCreator != nil {
		return &objectFactory{objectCreator: objectCreator, objectCollectionCreator: objectCollectionCreator}
	} else {
		panic("Could not register object factory. Creator is nil.")
	}
	return nil
}

type objectFactory struct {
	objectCreator           core.ObjectCreator
	objectCollectionCreator core.ObjectCollectionCreator
}

//Creates object
func (fac *objectFactory) CreateObject(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return fac.objectCreator(ctx, args)
}

//Creates collection
func (fac *objectFactory) CreateObjectCollection(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return fac.objectCollectionCreator(ctx, args)
}
