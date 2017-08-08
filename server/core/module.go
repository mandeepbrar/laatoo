package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"reflect"

	"github.com/blang/semver"
)

type serverModule struct {
	userModule   core.Module
	impl         *moduleImpl
	name         string
	version      semver.Version
	dependencies map[string]semver.Range
	svrContext   *serverContext
}

func (mod *serverModule) initialize(ctx core.ServerContext, conf config.Config) error {
	//inject service implementation into
	//every service
	impl := newModuleImpl()
	mod.impl = impl
	var modval core.Module
	modval = impl
	if mod.userModule != nil {
		val := reflect.ValueOf(mod.userModule)
		elem := val.Elem()
		fld := elem.FieldByName("Module")
		if fld.CanSet() {
			fld.Set(reflect.ValueOf(modval))
		} else {
			return errors.TypeMismatch(ctx, "Module does not inherit from core.Module", mod.name)
		}
		err := mod.userModule.Initialize(ctx)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	if err := mod.impl.processInfo(ctx, conf); err != nil {
		return err
	}

	impl.state = Initialized

	if mod.userModule != nil {
	}
	return nil
}

func (mod *serverModule) start(ctx core.ServerContext) error {
	if mod.userModule != nil {
		return mod.userModule.Start(ctx)
	} else {
		return nil
	}
}

func (mod *serverModule) factories(ctx core.ServerContext) map[string]config.Config {
	return mod.userModule.Factories(ctx)
}

func (mod *serverModule) services(ctx core.ServerContext) map[string]config.Config {
	return mod.userModule.Services(ctx)
}

func (mod *serverModule) rules(ctx core.ServerContext) map[string]config.Config {
	return mod.userModule.Rules(ctx)
}

func (mod *serverModule) channels(ctx core.ServerContext) map[string]config.Config {
	return mod.userModule.Channels(ctx)
}

func (mod *serverModule) tasks(ctx core.ServerContext) map[string]config.Config {
	return mod.userModule.Tasks(ctx)
}
