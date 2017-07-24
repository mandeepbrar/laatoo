package objects

import (
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	//"laatoo/sdk/log"
	"laatoo/sdk/server"
)

func NewObjectLoader(ctx core.ServerContext, name string, parentElem core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	ldr := &objectLoader{objectsFactoryRegister: make(map[string]core.ObjectFactory, 30), name: name, parentElem: parentElem}
	ldrElem := &objectLoaderProxy{loader: ldr}
	return ldr, ldrElem
}

func ChildLoader(ctx core.ServerContext, name string, parentLdr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
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
	log.Trace(ctx, "carrying over the following objects to the child", "objects", registry)
	ldr := &objectLoader{objectsFactoryRegister: registry, name: name, parentElem: parent}
	ldrElem := &objectLoaderProxy{loader: ldr}
	return ldr, ldrElem
}
