package objects

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/core"
)

type objectLoaderProxy struct {
	*common.Context
	loader *objectLoader
}

func (ldr *objectLoaderProxy) RegisterObjectFactory(ctx core.Context, objectName string, factory core.ObjectFactory) {
	ldr.loader.registerObjectFactory(ctx, objectName, factory)
}
func (ldr *objectLoaderProxy) RegisterObject(ctx core.Context, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) {
	ldr.loader.registerObject(ctx, objectName, objectCreator, objectCollectionCreator)
}
func (ldr *objectLoaderProxy) CreateCollection(ctx core.Context, objectName string, length int, args core.MethodArgs) (interface{}, error) {
	return ldr.loader.createCollection(ctx, objectName, length, args)
}
func (ldr *objectLoaderProxy) CreateObject(ctx core.Context, objectName string, args core.MethodArgs) (interface{}, error) {
	return ldr.loader.createObject(ctx, objectName, args)
}
func (ldr *objectLoaderProxy) GetObjectCollectionCreator(ctx core.Context, objectName string) (core.ObjectCollectionCreator, error) {
	return ldr.loader.getObjectCollectionCreator(ctx, objectName)
}
func (ldr *objectLoaderProxy) GetObjectCreator(ctx core.Context, objectName string) (core.ObjectCreator, error) {
	return ldr.loader.getObjectCreator(ctx, objectName)
}
func (ldr *objectLoaderProxy) GetMethod(ctx core.Context, methodName string) (core.ServiceFunc, error) {
	return ldr.loader.getMethod(ctx, methodName)
}
