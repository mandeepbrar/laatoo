package core

import (
	"fmt"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
	//"github.com/fsnotify/fsnotify"
)

func (modMgr *moduleManager) addWatch(ctx core.ServerContext, modName string, modDir string) error {
	ctx = ctx.SubContext("Watch " + modName)

	moduleInstallConf := modMgr.moduleInstallationConf[modName]

	hotmodcompiler, compilerok := moduleInstallConf.GetString(ctx, "hotmodcompiler")

	compilerWatcher := watcher.New()
	compilerWatcher.SetMaxEvents(1)
	nonCompilerWatcher := watcher.New()
	nonCompilerWatcher.SetMaxEvents(1)

	watchers := []*watcher.Watcher{compilerWatcher, nonCompilerWatcher}

	closeWatchers := func() {
		for _, w := range watchers {
			w.Close()
		}
	}

	if compilerok {
		log.Info(ctx, "Add compile watcher", "modName", modName)
		err := modMgr.watchFilesToCompile(ctx, modName, modDir, hotmodcompiler, compilerWatcher, closeWatchers)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {
		log.Info(ctx, "Skipping compile watcher", "hotmodcompiler", hotmodcompiler, "modName", modName, "moduleInstallConf", moduleInstallConf)
	}

	log.Info(ctx, "Add nonCompile watcher", "modName", modName)
	err := modMgr.watchNonCompileFileChanges(ctx, modName, modDir, nonCompilerWatcher, closeWatchers)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	modMgr.watchers = append(modMgr.watchers, watchers...)

	modMgr.startWatcher(ctx, compilerWatcher, time.Millisecond*800)
	modMgr.startWatcher(ctx, nonCompilerWatcher, time.Millisecond*800)

	//defer watcher.Close()

	log.Error(ctx, "Watching module directory for change ", "dir", modDir)

	return nil
}

func (modMgr *moduleManager) watchNonCompileFileChanges(ctx core.ServerContext, modName string, modDir string, w *watcher.Watcher, closeWatchers func()) error {
	log.Info(ctx, "Add nonCompileFileChanges", "modName", modName)

	addFoldersToWatcher := func(ctx core.ServerContext, modDir string, w *watcher.Watcher) error {
		log.Error(ctx, "adding non compile folders", "moddir", modDir)
		foldersToWatch := []string{"files", "ui/registry", "properties"} /*"uibuild", "objects",*/

		for _, name := range foldersToWatch {
			folderToWatch := path.Join(modDir, name)
			exists, _, _ := utils.FileExists(folderToWatch)
			if exists {
				if err := w.AddRecursive(folderToWatch); err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
		return nil
	}

	go func(ctx core.ServerContext, modName, modDir string) {
		for {
			select {
			case event := <-w.Event:
				log.Info(ctx, "Hot module changed", "modName", modName, " file ", event.Path, "event", event)
				reloadCtx := ctx.SubContext("Reload module " + modName)
				closeWatchers()
				err := modMgr.ReloadModule(reloadCtx, modName, modDir)
				if err != nil {
					log.Error(reloadCtx, "Error while reloading module", "Error", err)
				}
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Error(ctx, "Error while watching", err)
			case <-w.Closed:
				return
			}
		}
	}(ctx, modName, modDir)

	if err := addFoldersToWatcher(ctx, modDir, w); err != nil {
		return errors.WrapError(ctx, err)
	}

	return nil
}

func (modMgr *moduleManager) startWatcher(ctx core.ServerContext, w *watcher.Watcher, ms time.Duration) {
	go func(ctx core.ServerContext, w *watcher.Watcher, ms time.Duration) {
		if err := w.Start(ms); err != nil {
			log.Error(ctx, "Watcher stopped watching")
		}
	}(ctx, w, ms)
}
func (modMgr *moduleManager) watchFilesToCompile(ctx core.ServerContext, modName string, modDir, compilerCommand string, compileWatcher *watcher.Watcher, closeWatchers func()) error {

	autogen_skipper := func(info os.FileInfo, fullPath string) error {
		if strings.Contains(fullPath, "autogen_") {
			return watcher.ErrSkip
		}
		return nil
	}

	addFoldersToWatcher := func(ctx core.ServerContext, modDir string, w *watcher.Watcher) error {
		log.Error(ctx, "adding folders", "moddir", modDir)
		foldersToWatch := []string{"ui/js", "server/go", "build"}
		w.AddFilterHook(autogen_skipper)

		for _, name := range foldersToWatch {
			folderToWatch := path.Join(modDir, name)
			exists, _, _ := utils.FileExists(folderToWatch)
			if exists {
				if err := w.AddRecursive(folderToWatch); err != nil {
					return errors.WrapError(ctx, err)
				}
			}
		}
		return nil
	}

	compile := func(compileCtx core.ServerContext, modname, modDir, compilerCommand string) {
		compilerCmdArr := strings.Split(compilerCommand, " ")
		command := compilerCmdArr[0]
		compilerCmdArr = append(compilerCmdArr[1:], "--name", modname, "--packageFolder", modDir)

		cmd := exec.Command(command, compilerCmdArr...)
		//cmd.Env = env
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(compileCtx, "Error executing command ***********", "err", err, "stdoutStderr", string(stdoutStderr))
		} else {
			log.Error(compileCtx, "Compile success ***********", "stdoutStderr", string(stdoutStderr))
			reloadCtx := compileCtx.SubContext("Reload module " + modName)
			err = modMgr.ReloadModule(reloadCtx, modName, modDir)
			if err != nil {
				log.Error(reloadCtx, "Error while reloading module", "err", err)
			}
		}

	}
	/*
		removedFoldersFromWatcher := func(ctx core.ServerContext, modDir string, w *watcher.Watcher) error {
			log.Error(ctx, "removed folders", "moddir", modDir)
			if err := w.RemoveRecursive(path.Join(modDir, "ui")); err != nil {
				return err
			}
			if err := w.RemoveRecursive(path.Join(modDir, "server")); err != nil {
				return err
			}
			return nil
		}*/

	go func(ctx core.ServerContext, modname, modDir, compilerCommand string) {
		for {
			select {
			case event := <-compileWatcher.Event:
				compileCtx := ctx.SubContext("Module Compile" + modName)
				closeWatchers()
				log.Info(compileCtx, "Compile required", "modName", modName, " file ", event.Path, "event", event)
				compile(compileCtx, modname, modDir, compilerCommand)
				//env := os.Environ()
				//log.Error(compileCtx, "Environment", "env", env)
				/*err = removedFoldersFromWatcher(ctx, modDir, compileWatcher)
				if err != nil {
					log.Error(compileCtx, "Error reloading watchers", "err", err)
				}*/
				//modMgr.startWatcher(ctx, compileWatcher, time.Millisecond*5000)
			case err := <-compileWatcher.Error:
				log.Error(ctx, "Error while watching", err)
			case <-compileWatcher.Closed:
				return
			}
		}
	}(ctx, modName, modDir, compilerCommand)

	if err := addFoldersToWatcher(ctx, modDir, compileWatcher); err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Info(ctx, "Added Compile watcher", "modDir", modDir)

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
		err := modMgr.loadModuleObjects(ctx, modName, modDir, modconf)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	var createdInstances map[string]*serverModule

	log.Info(ctx, "Loading instances", "mod", modName)
	if createdInstances, err = modMgr.loadInstances(ctx, modName, modDir, removedInstances); err != nil {
		return errors.WrapError(ctx, err)
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

	log.Error(ctx, "Readding a watcher", "modName", modName)
	go modMgr.addWatch(ctx, modName, modDir)

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

func (modMgr *moduleManager) loadInstances(ctx core.ServerContext, moduleName, modDir string, removedInstances map[string]*serverModule) (map[string]*serverModule, error) {
	createdInstances := make(map[string]*serverModule)

	for instance, removedMod := range removedInstances {
		if removedMod.isSubInstance(ctx) {
			continue
		}
		instanceConf := removedMod.modConf
		newCtx, _, err := modMgr.setupInstanceContext(ctx, instance, nil, instanceConf, modDir)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		pendingModuleInstances := make(map[string]*pendingModInfo)
		modIns, _, err := modMgr.createModuleInstance(newCtx, instance, moduleName, modDir, nil, instanceConf, pendingModuleInstances)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		createdInstances[instance] = modIns
		log.Error(ctx, "************************", "pending instances", pendingModuleInstances)
		if len(pendingModuleInstances) > 0 {
			instancesCreated, err := modMgr.iterateAndLoadPendingModuleInstances(ctx, pendingModuleInstances)
			if err != nil {
				return nil, errors.WrapError(ctx, err)
			}
			log.Error(ctx, "*********instances created***************", "instancesCreated", instancesCreated)
			for k, v := range instancesCreated {
				createdInstances[k] = v
			}
		}
	}
	log.Info(ctx, "created instances", " created instan", createdInstances)
	//taskManager := ctx.GetServerElement(core.ServerElementTaskManager).(*taskManagerProxy).manager
	chnManager := ctx.GetServerElement(core.ServerElementChannelManager).(*channelManagerProxy).manager
	svcManager := ctx.GetServerElement(core.ServerElementServiceManager).(*serviceManagerProxy).manager
	facManager := ctx.GetServerElement(core.ServerElementFactoryManager).(*factoryManagerProxy).manager
	for modName, modInstance := range createdInstances {
		log.Info(ctx, "Creating factories of module ", "modName", modName, "modInstance", modInstance)
		if err := facManager.createModuleFactories(ctx, modInstance); err != nil {
			return nil, err
		}
		log.Info(ctx, "Creating services of module ", "modName", modName, "modInstance", modInstance)
		if err := svcManager.createModuleServices(ctx, modInstance); err != nil {
			return nil, err
		}
		/*log.Info(ctx, "Creating tasks of module ", "modName", modName, "modInstance", modInstance)
		if err := svcManager.createModuleTasks(ctx, modInstance); err != nil {
			return nil, err
		}*/
		log.Info(ctx, "Creating channels of module ", "modName", modName, "modInstance", modInstance)
		if err := chnManager.createModuleChannels(ctx, modInstance); err != nil {
			return nil, err
		}
	}

	return createdInstances, nil
}

func (modMgr *moduleManager) unloadInstanceLive(ctx core.ServerContext, mod *serverModule, unloadObjs bool, removedInstances map[string]*serverModule) error {
	ctx = ctx.SubContext("Unload live instance " + mod.name)
	var err error
	for ins, parentIns := range modMgr.parentModules {
		if parentIns.name == mod.name {
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
	if removedInstances != nil {
		removedInstances[mod.name] = mod
	}
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
