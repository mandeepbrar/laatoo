package core

import (
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
)

type objectLoaderProxy struct {
	loader *objectLoader
}

func (proxy *objectLoaderProxy) Reference() core.ServerElement {
	return &objectLoaderProxy{loader: proxy.loader}
}
func (proxy *objectLoaderProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *objectLoaderProxy) GetName() string {
	return proxy.loader.name
}
func (proxy *objectLoaderProxy) GetType() core.ServerElementType {
	return core.ServerElementLoader
}

func (proxy *objectLoaderProxy) GetContext() core.ServerContext {
	return proxy.loader.svrContext
}

func (ldr *objectLoaderProxy) Register(ctx ctx.Context, objectName string, object interface{}, metadata core.Info) error {
	return ldr.loader.register(ctx, objectName, object, metadata)
}
func (ldr *objectLoaderProxy) RegisterObjectFactory(ctx ctx.Context, objectName string, factory core.ObjectFactory) error {
	return ldr.loader.registerObjectFactory(ctx, objectName, factory)
}
func (ldr *objectLoaderProxy) RegisterObject(ctx ctx.Context, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator, metadata core.Info) error {
	return ldr.loader.registerObject(ctx, objectName, objectCreator, objectCollectionCreator, metadata)
}
func (ldr *objectLoaderProxy) CreateCollection(ctx ctx.Context, objectName string, length int) (interface{}, error) {
	return ldr.loader.createCollection(ctx, objectName, length)
}
func (ldr *objectLoaderProxy) CreateObject(ctx ctx.Context, objectName string) (interface{}, error) {
	return ldr.loader.createObject(ctx, objectName)
}

func (ldr *objectLoaderProxy) GetObjectFactory(ctx ctx.Context, name string) (core.ObjectFactory, bool) {
	if ldr.loader.objectsFactoryRegister != nil {
		log.Error(ctx, "Getting object factory", "object", name)
		fac, ok := ldr.loader.objectsFactoryRegister[name]
		log.Error(ctx, "Getting object factory", "ok", ok)
		return fac, ok
	}
	return nil, false
}

/*func (ldr *objectLoaderProxy) GetObjectCollectionCreator(ctx ctx.Context, objectName string) (core.ObjectCollectionCreator, error) {
	return ldr.loader.getObjectCollectionCreator(ctx, objectName)
}
func (ldr *objectLoaderProxy) GetObjectCreator(ctx ctx.Context, objectName string) (core.ObjectCreator, error) {
	return ldr.loader.getObjectCreator(ctx, objectName)
}*/
func (ldr *objectLoaderProxy) GetMetaData(ctx ctx.Context, objectName string) (core.Info, error) {
	return ldr.loader.getMetaData(ctx, objectName)
}
