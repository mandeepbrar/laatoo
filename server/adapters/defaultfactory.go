package adapters

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"

	//	"laatoo/sdk/server/log"
	"laatoo/sdk/server/errors"
)

type DefaultFactory struct {
	core.ServiceFactory
}

//Create the services configured for factory.
func (mi *DefaultFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	obj, err := ctx.CreateObject(method)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Error(ctx, "Created service object", "obj", obj)
	svc, ok := obj.(core.Service)
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", method, "obj", obj)
	}
	return svc, nil
}
