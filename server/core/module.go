package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/server/common"
	"laatoo/server/constants"
	"reflect"

	"github.com/blang/semver"
)

type serverModule struct {
	userModule   core.Module
	impl         *moduleImpl
	dir          string
	name         string
	version      semver.Version
	dependencies map[string]semver.Range
	svrContext   *serverContext
	services     map[string]config.Config
	factories    map[string]config.Config
	channels     map[string]config.Config
	tasks        map[string]config.Config
	rules        map[string]config.Config
}

func newServerModule(ctx core.ServerContext, name, dirpath string) *serverModule {
	mod := &serverModule{name: name, dir: dirpath}
	mod.services = make(map[string]config.Config)
	mod.factories = make(map[string]config.Config)
	mod.channels = make(map[string]config.Config)
	mod.tasks = make(map[string]config.Config)
	mod.rules = make(map[string]config.Config)
	return mod
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
		if err := mod.loadModuleFromObj(ctx, conf); err != nil {
			return err
		}
	}
	return mod.loadModuleDirs(ctx, conf)
}

func (mod *serverModule) start(ctx core.ServerContext) error {
	if mod.userModule != nil {
		return mod.userModule.Start(ctx)
	} else {
		return nil
	}
}

func (mod *serverModule) loadModuleDirs(ctx core.ServerContext, modSettings config.Config) error {
	var err error
	factoriesEnabled, ok := modSettings.GetBool(constants.CONF_SERVICEFACTORIES)

	if !ok || factoriesEnabled {
		mod.factories, err = common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_SERVICEFACTORIES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
	}

	servicesEnabled, ok := modSettings.GetBool(constants.CONF_SERVICES)

	if !ok || servicesEnabled {
		mod.services, err = common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_SERVICES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
	}

	channelsEnabled, ok := modSettings.GetBool(constants.CONF_CHANNELS)

	if !ok || channelsEnabled {
		mod.channels, err = common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_CHANNELS, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
	}

	rulesEnabled, ok := modSettings.GetBool(constants.CONF_RULES)

	if !ok || rulesEnabled {
		mod.rules, err = common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_RULES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
	}

	tasksEnabled, ok := modSettings.GetBool(constants.CONF_TASKS)
	if !ok || tasksEnabled {
		mod.tasks, err = common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_TASKS, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
	}
	return nil
}

func (mod *serverModule) loadModuleFromObj(ctx core.ServerContext, modSettings config.Config) error {

	factoriesEnabled, ok := modSettings.GetBool(constants.CONF_SERVICEFACTORIES)

	if !ok || factoriesEnabled {
		mod.factories = mod.userModule.Factories(ctx)
	}

	servicesEnabled, ok := modSettings.GetBool(constants.CONF_SERVICES)

	if !ok || servicesEnabled {
		mod.services = mod.userModule.Services(ctx)
	}

	channelsEnabled, ok := modSettings.GetBool(constants.CONF_CHANNELS)

	if !ok || channelsEnabled {
		mod.channels = mod.userModule.Channels(ctx)
	}

	rulesEnabled, ok := modSettings.GetBool(constants.CONF_RULES)

	if !ok || rulesEnabled {
		mod.rules = mod.userModule.Rules(ctx)
	}

	tasksEnabled, ok := modSettings.GetBool(constants.CONF_TASKS)
	if !ok || tasksEnabled {
		mod.tasks = mod.userModule.Tasks(ctx)
	}
	return nil
}
