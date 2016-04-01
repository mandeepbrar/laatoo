package registry

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
)

//service method for doing various tasks
func NewObjectFactory(objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) core.ObjectFactory {
	if objectCreator != nil && objectCollectionCreator != nil {
		return &objectFactory{objectCreator: objectCreator, objectCollectionCreator: objectCollectionCreator}
	}
	return nil
}

type objectFactory struct {
	objectCreator           core.ObjectCreator
	objectCollectionCreator core.ObjectCollectionCreator
}

//Initialize the object factory
func (fac *objectFactory) Initialize(ctx core.ServerContext, config config.Config) error {
	return nil
}

//Creates object
func (fac *objectFactory) CreateObject(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return fac.objectCreator(ctx, args)
}

//Creates collection
func (fac *objectFactory) CreateObjectCollection(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	return fac.objectCollectionCreator(ctx, args)
}
