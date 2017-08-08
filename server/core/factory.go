package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"reflect"
)

type serviceFactory struct {
	name       string
	factory    core.ServiceFactory
	conf       config.Config
	owner      *factoryManager
	svrContext *serverContext
	impl       *factoryImpl
}

func (fac *serviceFactory) initialize(ctx core.ServerContext, conf config.Config) error {
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

	err := fac.factory.Initialize(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if err := fac.impl.processInfo(ctx, conf); err != nil {
		return err
	}

	impl.state = Initialized
	return nil
}

/*
func (fac *serviceFactory) processInfo(ctx core.ServerContext, facconf config.Config) error {
	confs := fac.impl.GetConfigurations()
	for name, configObj := range confs {
		configuration := configObj.(*configuration)
		val, ok := facconf.Get(name)

		if !ok && configuration.required {
			return errors.MissingConf(ctx, name)
		}
		if ok {
			switch configuration.conftype {
			case "", config.CONF_OBJECT_STRING:
				val, ok = svcconf.GetString(name)
				if ok {
					configuration.value = val
				}
			case config.CONF_OBJECT_STRINGMAP:
				val, ok = svcconf.GetSubConfig(name)
				if ok {
					configuration.value = val
				}
			case config.CONF_OBJECT_STRINGARR:
				val, ok = svcconf.GetStringArray(name)
				if ok {
					configuration.value = val
				}
			case config.CONF_OBJECT_CONFIG:
				val, ok = svcconf.GetSubConfig(name)
				if ok {
					configuration.value = val
				}
			case config.CONF_OBJECT_BOOL:
				val, ok = svcconf.GetBool(name)
				if ok {
					configuration.value = val
				}
			default:
				configuration.value = val
			}

			//check if value provided is a variable name..
			//allow assignment if its a variable name
			if !ok {
				strval, _ := svcconf.GetString(name)
				if strings.HasPrefix(strval, ":") {
					configuration.value = strval
					ok = true
				}
			}

			//configuration was there but wrong type
			if !ok && configuration.required {
				return errors.BadConf(ctx, name)
			}
		}
	}

	return nil
}
*/
