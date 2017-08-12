package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
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
	objectName   string
	version      semver.Version
	dependencies map[string]semver.Range
	svrContext   *serverContext
	services     map[string]config.Config
	factories    map[string]config.Config
	channels     map[string]config.Config
	tasks        map[string]config.Config
	rules        map[string]config.Config
	modSettings  config.Config
}

func newServerModule(ctx core.ServerContext, name, dirpath string, modconf config.Config) *serverModule {
	mod := &serverModule{svrContext: ctx.(*serverContext), name: name, dir: dirpath}
	mod.services = make(map[string]config.Config)
	mod.factories = make(map[string]config.Config)
	mod.channels = make(map[string]config.Config)
	mod.tasks = make(map[string]config.Config)
	mod.rules = make(map[string]config.Config)
	return mod
}

func (mod *serverModule) loadMetaData(ctx core.ServerContext) error {
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

		ldr := ctx.GetServerElement(core.ServerElementLoader).(server.ObjectLoader)
		md, _ := ldr.GetMetaData(ctx, mod.objectName)
		if md != nil {
			inf, ok := md.(*moduleInfo)
			if ok {
				impl.moduleInfo = inf
			}
		}
		mod.userModule.Describe(ctx)
	}
	log.Trace(ctx, "Module info ", "Name", mod.name, "Info", mod.impl.moduleInfo.configurations)
	return nil
}

func (mod *serverModule) initialize(ctx core.ServerContext, conf config.Config, env config.Config) error {
	if conf != nil {
		mod.modSettings = conf
	} else {
		mod.modSettings = make(config.GenericConfig)
	}

	if env != nil {
		envvars := env.AllConfigurations()
		for _, varname := range envvars {
			varvalue, _ := env.GetString(varname)
			ctx.SetVariable(varname, varvalue)
		}
	}

	if err := mod.impl.processInfo(ctx, conf); err != nil {
		return err
	}

	if mod.userModule != nil {
		err := mod.userModule.Initialize(ctx, conf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		if err := mod.loadModuleFromObj(ctx); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	if err := mod.loadModuleDir(ctx); err != nil {
		return errors.WrapError(ctx, err)
	}

	mod.impl.state = Initialized
	return nil
}

func (mod *serverModule) start(ctx core.ServerContext) error {
	if mod.userModule != nil {
		if err := mod.userModule.Start(ctx); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (mod *serverModule) loadModuleDir(ctx core.ServerContext) error {
	factoriesEnabled, ok := mod.modSettings.GetBool(constants.CONF_SERVICEFACTORIES)

	if !ok || factoriesEnabled {
		factories, err := common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_SERVICEFACTORIES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.factories = common.MergeConfigMaps(mod.factories, factories)
	}

	servicesEnabled, ok := mod.modSettings.GetBool(constants.CONF_SERVICES)

	if !ok || servicesEnabled {
		services, err := common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_SERVICES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.services = common.MergeConfigMaps(mod.services, services)
	}

	channelsEnabled, ok := mod.modSettings.GetBool(constants.CONF_CHANNELS)

	if !ok || channelsEnabled {
		channels, err := common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_CHANNELS, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.channels = common.MergeConfigMaps(mod.channels, channels)
	}

	rulesEnabled, ok := mod.modSettings.GetBool(constants.CONF_RULES)

	if !ok || rulesEnabled {
		rules, err := common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_RULES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.rules = common.MergeConfigMaps(mod.rules, rules)
	}

	tasksEnabled, ok := mod.modSettings.GetBool(constants.CONF_TASKS)
	if !ok || tasksEnabled {
		tasks, err := common.ProcessDirectoryFiles(ctx, mod.dir, constants.CONF_TASKS, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.tasks = common.MergeConfigMaps(mod.tasks, tasks)
	}
	return nil
}

func (mod *serverModule) loadModuleFromObj(ctx core.ServerContext) error {

	factoriesEnabled, ok := mod.modSettings.GetBool(constants.CONF_SERVICEFACTORIES)

	if !ok || factoriesEnabled {
		factories := mod.userModule.Factories(ctx)
		mod.factories = common.MergeConfigMaps(mod.factories, factories)
	}

	servicesEnabled, ok := mod.modSettings.GetBool(constants.CONF_SERVICES)

	if !ok || servicesEnabled {
		services := mod.userModule.Services(ctx)
		mod.services = common.MergeConfigMaps(mod.services, services)
	}

	channelsEnabled, ok := mod.modSettings.GetBool(constants.CONF_CHANNELS)

	if !ok || channelsEnabled {
		channels := mod.userModule.Channels(ctx)
		mod.channels = common.MergeConfigMaps(mod.channels, channels)
	}

	rulesEnabled, ok := mod.modSettings.GetBool(constants.CONF_RULES)

	if !ok || rulesEnabled {
		rules := mod.userModule.Rules(ctx)
		mod.rules = common.MergeConfigMaps(mod.rules, rules)
	}

	tasksEnabled, ok := mod.modSettings.GetBool(constants.CONF_TASKS)
	if !ok || tasksEnabled {
		tasks := mod.userModule.Tasks(ctx)
		mod.tasks = common.MergeConfigMaps(mod.tasks, tasks)
	}
	return nil
}
