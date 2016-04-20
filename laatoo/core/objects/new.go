package objects

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	//"laatoo/sdk/log"
	"laatoo/sdk/server"
)

func NewObjectLoader(ctx core.ServerContext, name string, parentElem core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	ldr := &objectLoader{make(map[string]core.ObjectFactory, 30)}
	ldrElemCtx := parentElem.NewCtx(name)
	ldrElem := &objectLoaderProxy{Context: ldrElemCtx.(*common.Context), loader: ldr}
	return ldr, ldrElem
}

func ChildLoader(ctx core.ServerContext, name string, parentLdr core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	objLdrProxy := parentLdr.(*objectLoaderProxy)
	objLoader := objLdrProxy.loader
	registry := make(map[string]core.ObjectFactory, len(objLoader.objectsFactoryRegister))
	for k, v := range objLoader.objectsFactoryRegister {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			registry[k] = v
		}
	}
	ldr := &objectLoader{registry}
	ldrElemCtx := parentLdr.NewCtx(name)
	ldrElem := &objectLoaderProxy{Context: ldrElemCtx.(*common.Context), loader: ldr}
	return ldr, ldrElem
}
