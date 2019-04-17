package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"laatoo/server/constants"
	"path"
	"reflect"
)

type serverModule struct {
	userModule      core.Module
	impl            *moduleImpl
	dir             string
	name            string
	moduleName      string
	objectName      string
	svrContext      *serverContext
	services        map[string]config.Config
	factories       map[string]config.Config
	channels        map[string]config.Config
	tasks           map[string]config.Config
	rules           map[string]config.Config
	properties      map[string]interface{}
	modConf         config.Config
	isExtended      bool
	extendedMod     string
	extendedModDir  string
	extendedModConf config.Config
	modSettings     config.Config
}

func newServerModule(ctx core.ServerContext, name, moduleName, dirpath string, modconf config.Config, modMgr *moduleManager) *serverModule {
	mod := &serverModule{svrContext: ctx.(*serverContext), name: name, moduleName: moduleName, dir: dirpath, modConf: modconf}

	mod.extendedMod, mod.isExtended = modconf.GetString(ctx, constants.CONF_EXTENDED_MOD)

	if mod.isExtended {
		mod.extendedModConf = modMgr.moduleConf[mod.extendedMod]
		mod.extendedModDir = modMgr.availableModules[mod.extendedMod]
	}

	mod.services = make(map[string]config.Config)
	mod.factories = make(map[string]config.Config)
	mod.channels = make(map[string]config.Config)
	mod.tasks = make(map[string]config.Config)
	mod.rules = make(map[string]config.Config)
	mod.properties = make(map[string]interface{})
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

		ldr := ctx.GetServerElement(core.ServerElementLoader).(elements.ObjectLoader)
		md, _ := ldr.GetMetaData(ctx, mod.objectName)
		if md != nil {
			inf, ok := md.(*moduleInfo)
			if ok {
				impl.moduleInfo = inf.clone()
			}
		}
		mod.userModule.Describe(ctx)
	}
	log.Trace(ctx, "Module info ", "Name", mod.name, "Object", mod.objectName, "Info", mod.impl.moduleInfo.configurations)
	return nil
}

func (mod *serverModule) initialize(ctx core.ServerContext, conf config.Config) error {
	if conf != nil {
		mod.modSettings = conf
	} else {
		mod.modSettings = ctx.CreateConfig()
	}

	if mod.extendedModConf != nil {
		if err := mod.initWithConf(ctx, mod.extendedModConf, mod.extendedModDir); err != nil {
			return err
		}
	}
	return mod.initWithConf(ctx, conf, mod.dir)

}
func (mod *serverModule) initWithConf(ctx core.ServerContext, conf config.Config, dir string) error {

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
	if err := mod.loadModuleDir(ctx, dir); err != nil {
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

func (mod *serverModule) readProperties(ctx core.ServerContext, dir string) error {
	propsDir := path.Join(dir, constants.PROPERTIES_DIR)

	props, err := common.ReadProperties(ctx, propsDir)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	mod.properties = common.MergeProps(mod.properties, props)
	return nil
}

func (mod *serverModule) loadModuleDir(ctx core.ServerContext, dir string) error {

	modConfigDir := path.Join(dir, constants.CONF_CONFIG_DIR)

	if err := mod.readProperties(ctx, modConfigDir); err != nil {
		return err
	}

	factoriesEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_SERVICEFACTORIES)

	if !ok || factoriesEnabled {
		factories, err := common.ProcessDirectoryFiles(ctx, modConfigDir, constants.CONF_SERVICEFACTORIES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.factories = common.MergeConfigMaps(mod.factories, factories)
	}

	servicesEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_SERVICES)

	if !ok || servicesEnabled {
		services, err := common.ProcessDirectoryFiles(ctx, modConfigDir, constants.CONF_SERVICES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.services = common.MergeConfigMaps(mod.services, services)
	}

	channelsEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_CHANNELS)

	if !ok || channelsEnabled {
		channels, err := common.ProcessDirectoryFiles(ctx, modConfigDir, constants.CONF_CHANNELS, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.channels = common.MergeConfigMaps(mod.channels, channels)
	}

	rulesEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_RULES)

	if !ok || rulesEnabled {
		rules, err := common.ProcessDirectoryFiles(ctx, modConfigDir, constants.CONF_RULES, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.rules = common.MergeConfigMaps(mod.rules, rules)
	}

	tasksEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_TASKS)
	if !ok || tasksEnabled {
		tasks, err := common.ProcessDirectoryFiles(ctx, modConfigDir, constants.CONF_TASKS, true)
		if err != nil {
			return errors.WrapError(ctx, err, "Module", mod.name)
		}
		mod.tasks = common.MergeConfigMaps(mod.tasks, tasks)
	}
	return nil
}

func (mod *serverModule) loadModuleFromObj(ctx core.ServerContext) error {

	factoriesEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_SERVICEFACTORIES)

	if !ok || factoriesEnabled {
		factories := mod.userModule.Factories(ctx)
		mod.factories = common.MergeConfigMaps(mod.factories, factories)
	}

	servicesEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_SERVICES)

	if !ok || servicesEnabled {
		services := mod.userModule.Services(ctx)
		mod.services = common.MergeConfigMaps(mod.services, services)
	}

	channelsEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_CHANNELS)

	if !ok || channelsEnabled {
		channels := mod.userModule.Channels(ctx)
		mod.channels = common.MergeConfigMaps(mod.channels, channels)
	}

	rulesEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_RULES)

	if !ok || rulesEnabled {
		rules := mod.userModule.Rules(ctx)
		mod.rules = common.MergeConfigMaps(mod.rules, rules)
	}

	tasksEnabled, ok := mod.modSettings.GetBool(ctx, constants.CONF_TASKS)
	if !ok || tasksEnabled {
		tasks := mod.userModule.Tasks(ctx)
		mod.tasks = common.MergeConfigMaps(mod.tasks, tasks)
	}
	return nil
}

func (mod *serverModule) plugins(ctx core.ServerContext) map[string]config.Config {
	retVal := make(map[string]config.Config)
	for k, v := range mod.services {
		isPlugin, _ := v.GetBool(ctx, constants.MODULEMGR_PLUGIN)
		if isPlugin {
			retVal[k] = v
		}
	}
	return retVal
}
