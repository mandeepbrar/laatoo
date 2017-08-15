package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
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
	impl := newFactoryImpl()
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

	ldr := ctx.GetServerElement(core.ServerElementLoader).(server.ObjectLoader)
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

	err := fac.factory.Initialize(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	fac.impl.state = Initialized
	return nil
}
