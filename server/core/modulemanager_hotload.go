package core

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"time"

	"github.com/radovskyb/watcher"
	//"github.com/fsnotify/fsnotify"
)

func (modMgr *moduleManager) addWatch(ctx core.ServerContext, modName string, modDir string) error {
	ctx = ctx.SubContext("Watch " + modName)

	/*
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return errors.WrapError(ctx, err)
		}*/
	w := watcher.New()
	w.SetMaxEvents(1)

	go func() {
		for {
			select {
			case event := <-w.Event:
				log.Info(ctx, "Hot module changed", "modName", modName, " file ", event.Path)
				reloadCtx := ctx.SubContext("Reload module " + modName)
				err := modMgr.ReloadModule(reloadCtx, modName, modDir)
				if err != nil {
					log.Error(reloadCtx, "Error while reloading module", err)
				}
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Error(ctx, "Error while watching", err)
			case <-w.Closed:
				return
			}
		}
	}()

	//defer watcher.Close()
	modMgr.watchers = append(modMgr.watchers, w)

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(modDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	go func() {
		if err := w.Start(time.Millisecond * 800); err != nil {
			log.Error(ctx, "Watcher stopped watching")
		}
	}()

	log.Error(ctx, "Watching module directory for change ", "dir", modDir, "watchers", modMgr.watchers)

	return nil
}

func (modMgr *moduleManager) ReloadModule(ctx core.ServerContext, modName string, modDir string) error {
	var err error
	removedInstances := make(map[string]*serverModule)
	for name, modInstance := range modMgr.moduleInstances {
		if modInstance.moduleName == modName {
			log.Error(ctx, "Reload module instance", "name", name, "mod", modName)
			if err = modMgr.unloadInstanceLive(ctx, modInstance, false, removedInstances); err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}

	log.Info(ctx, "unloading module objects", "mod", modName)
	ldr := ctx.GetServerElement(core.ServerElementLoader).(*objectLoaderProxy).loader
	if err := ldr.unloadModuleObjects(ctx, modName); err != nil {
		return errors.WrapError(ctx, err)
	}

	log.Info(ctx, "Loading module objects", "mod", modName)

	modconf, ok := modMgr.moduleConf[modName]

	if ok {
		err = modMgr.loadModuleObjects(ctx, modName, modDir, modconf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	var createdInstanceNames []string

	log.Info(ctx, "Loading instances", "mod", modName)
	for insName, _ := range removedInstances {
		insConf := modMgr.moduleInstancesConfig[insName]
		createdInstanceNames, err = modMgr.loadLiveInstance(ctx, insName, modName, modDir, insConf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}

	createdInstances := make(map[string]*serverModule)
	for _, instanceName := range createdInstanceNames {
		modInstance := modMgr.moduleInstances[instanceName]
		createdInstances[instanceName] = modInstance
	}

	log.Info(ctx, "Starting instances", "mod", modName)
	if err = modMgr.startInstances(ctx, createdInstances); err != nil {
		return errors.WrapError(ctx, err)
	}

	log.Info(ctx, "Loading rules", "mod", modName)
	ruleManager := ctx.GetServerElement(core.ServerElementRulesManager).(*rulesManagerProxy).manager
	if _, err = ruleManager.loadRulesFromDirectory(ctx, modDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	log.Info(ctx, "Passing instances to plugins", "mod", modName)
	if err = modMgr.loadInstancesToPluginsforload(ctx, createdInstances); err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (modMgr *moduleManager) startInstances(ctx core.ServerContext, instances map[string]*serverModule) error {
	taskManager := ctx.GetServerElement(core.ServerElementTaskManager).(*taskManagerProxy).manager
	chnManager := ctx.GetServerElement(core.ServerElementChannelManager).(*channelManagerProxy).manager
	svcManager := ctx.GetServerElement(core.ServerElementServiceManager).(*serviceManagerProxy).manager
	facManager := ctx.GetServerElement(core.ServerElementFactoryManager).(*factoryManagerProxy).manager
	for _, modInstance := range instances {
		err := modMgr.startInstance(ctx, modInstance)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		if err := facManager.startModuleInstanceFactories(ctx, modInstance); err != nil {
			return err
		}
		if err := svcManager.startModuleInstanceServices(ctx, modInstance); err != nil {
			return err
		}
		if err := chnManager.startModuleInstanceChannels(ctx, modInstance); err != nil {
			return err
		}
		if err := taskManager.startModuleInstanceTasks(ctx, modInstance); err != nil {
			return err
		}

	}
	return nil
}

func (modMgr *moduleManager) unloadFromPluginsforReload(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Unload plugins ")
	for _, plugin := range modMgr.modulePlugins {
		err := plugin.Unloading(ctx, mod.name, mod.moduleName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	log.Error(ctx, "Unloaded instance from plugins", "instance", mod.name)
	return nil
}

func (modMgr *moduleManager) loadLiveInstance(ctx core.ServerContext, instance, moduleName, modDir string, instanceConf config.Config) ([]string, error) {
	newCtx, _, err := modMgr.setupInstanceContext(ctx, instance, instanceConf, modDir)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	pendingModuleInstances := make(map[string]config.Config)
	_, err = modMgr.createModuleInstance(newCtx, instance, moduleName, modDir, instanceConf, pendingModuleInstances)
	createdInstances := []string{instance}
	if len(pendingModuleInstances) > 0 {
		instancesCreated, err := modMgr.iterateAndLoadPendingModuleInstances(ctx, pendingModuleInstances)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		createdInstances = append(createdInstances, instancesCreated...)
	}

	return createdInstances, nil
}

func (modMgr *moduleManager) unloadInstanceLive(ctx core.ServerContext, mod *serverModule, unloadObjs bool, removedInstances map[string]*serverModule) error {
	ctx = ctx.SubContext("Unload live instance " + mod.name)
	var err error
	for ins, parentIns := range modMgr.parentModules {
		if parentIns == mod.name {
			childModIns, ok := modMgr.moduleInstances[ins]
			if ok {
				modMgr.unloadInstanceLive(ctx, childModIns, unloadObjs, removedInstances)
			}
		}
	}
	if err = modMgr.unloadFromPluginsforReload(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	if err = modMgr.unloadTasks(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	if err = modMgr.unloadRules(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	if err = modMgr.unloadChannels(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	if err = modMgr.unloadServices(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	if err = modMgr.unloadFactories(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	if unloadObjs {
		if err = modMgr.unloadObjects(ctx, mod); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	delete(modMgr.moduleInstances, mod.name)
	delete(modMgr.parentModules, mod.name)
	removedInstances[mod.name] = mod
	return nil
}

func (modMgr *moduleManager) unloadTasks(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Unload tasks ")
	taskManager := ctx.GetServerElement(core.ServerElementTaskManager).(*taskManagerProxy).manager
	if err := taskManager.unloadModuleTasks(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (modMgr *moduleManager) unloadRules(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Unload rules ")
	ruleManager := ctx.GetServerElement(core.ServerElementRulesManager).(*rulesManagerProxy).manager
	if err := ruleManager.unloadModuleRules(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (modMgr *moduleManager) unloadChannels(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Unload channels ")
	chnManager := ctx.GetServerElement(core.ServerElementChannelManager).(*channelManagerProxy).manager
	if err := chnManager.unloadModuleChannels(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (modMgr *moduleManager) unloadServices(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Unload services ")
	svcManager := ctx.GetServerElement(core.ServerElementServiceManager).(*serviceManagerProxy).manager
	if err := svcManager.unloadModuleServices(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (modMgr *moduleManager) unloadObjects(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Unload objects ")
	ldr := ctx.GetServerElement(core.ServerElementLoader).(*objectLoaderProxy).loader
	if err := ldr.unloadModuleObjects(ctx, mod.moduleName); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (modMgr *moduleManager) unloadFactories(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("Unload factories ")
	facManager := ctx.GetServerElement(core.ServerElementFactoryManager).(*factoryManagerProxy).manager
	if err := facManager.unloadModuleFactories(ctx, mod); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

/*cont, err := ioutil.ReadFile(file)
if err != nil {
	log.Error(ctx, "Error encountered during hot reload", "err", err)
}

err = actionF(svc.svrCtx, mod, file, dir, cont)
if err != nil {
	log.Error(ctx, "Error encountered during hot reload", "err", err)
}*/
// watch for errors

/*
//
go func() {
	for {
		select {
		// watch for events
		case event := <-watcher.Events:
			//fmt.Printf("EVENT! %#v\n", event)
			log.Error(ctx, "Module change event ", "event", event.Name, "op", event.Op)
			// watch for errors
		case err := <-watcher.Errors:
			fmt.Println("ERROR", err)
			log.Error(ctx, "Module change event error ", "err", err)
		}
	}
}()

// out of the box fsnotify can watch a single file, or a single directory

visit := func(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		log.Trace(ctx, "Watching directory for change ", "dir", path)
		if err := watcher.Add(path); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

err = filepath.Walk(modDir, visit)
if err != nil {
	log.Error(ctx, "Could not walk through hot directory")
}
*/
