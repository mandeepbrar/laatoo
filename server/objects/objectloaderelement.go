package objects

import (
	"laatoo/server/common"
	"laatoo/sdk/core"
)

type objectLoaderProxy struct {
	*common.Context
	loader *objectLoader
}

func (ldr *objectLoaderProxy) Register(ctx core.Context, objectName string, object interface{}) {
	ldr.loader.register(ctx, objectName, object)
}
func (ldr *objectLoaderProxy) RegisterObjectFactory(ctx core.Context, objectName string, factory core.ObjectFactory) {
	ldr.loader.registerObjectFactory(ctx, objectName, factory)
}
func (ldr *objectLoaderProxy) RegisterObject(ctx core.Context, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) {
	ldr.loader.registerObject(ctx, objectName, objectCreator, objectCollectionCreator)
}
func (ldr *objectLoaderProxy) CreateCollection(ctx core.Context, objectName string, length int) (interface{}, error) {
	return ldr.loader.createCollection(ctx, objectName, length)
}
func (ldr *objectLoaderProxy) CreateObject(ctx core.Context, objectName string) (interface{}, error) {
	return ldr.loader.createObject(ctx, objectName)
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
