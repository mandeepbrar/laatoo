package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"reflect"
)

type serviceFactory struct {
	name       string
	objectName string
	factory    core.ServiceFactory
	conf       config.Config
	owner      *factoryManager
	svrContext *serverContext
	impl       *factoryImpl
}

func (fac *serviceFactory) loadMetaData(ctx core.ServerContext) error {
	//inject service implementation into
	//every service
	impl := newFactoryImpl(fac.name)
	fac.impl = impl
	var facval core.ServiceFactory
	facval = impl
	val := reflect.ValueOf(fac.factory)
	elem := val.Elem()
	fld := elem.FieldByName("ServiceFactory")
	if fld.CanSet() {
		fld.Set(reflect.ValueOf(facval))
	} else {
		return errors.TypeMismatch(ctx, "Factory does not inherit from core.ServiceFactory", fac.name)
	}

	ldr := ctx.GetServerElement(core.ServerElementLoader).(elements.ObjectLoader)
	md, _ := ldr.GetMetaData(ctx, fac.objectName)
	if md != nil {
		inf, ok := md.(*factoryInfo)
		if ok {
			impl.factoryInfo = inf.clone()
		}
	}
	fac.factory.Describe(ctx)
	log.Trace(ctx, "Factory info ", "Name", fac.name, "Info", fac.impl.factoryInfo.configurations)
	return nil
}

func (fac *serviceFactory) initialize(ctx core.ServerContext, conf config.Config) error {
	if err := fac.impl.processInfo(ctx, conf); err != nil {
		return errors.WrapError(ctx, err)
	}

	//inject configuration values
	confsToInject := fac.impl.getConfigurationsToInject()
	err := utils.SetObjectFields(ctx, fac.factory, confsToInject, nil, nil)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	err = fac.factory.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	fac.impl.state = Initialized
	return nil
}

func (fac *serviceFactory) start(ctx core.ServerContext) error {
	return fac.factory.Start(ctx)
}

func (fac *serviceFactory) stop(ctx core.ServerContext) error {
	return fac.factory.Stop(ctx)
}
func (fac *serviceFactory) unload(ctx core.ServerContext) error {
	return fac.factory.Unload(ctx)
}
