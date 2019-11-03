package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/log"
)

const (
	singleFileByteLimit = 107374182400 // 1 GB
	chunkSize           = 4096         // 4 KB
)

type abstractserver struct {
	name string

	objectLoader       elements.ObjectLoader
	objectLoaderHandle elements.ServerElementHandle

	channelManager       elements.ChannelManager
	channelManagerHandle elements.ServerElementHandle

	factoryManager       elements.FactoryManager
	factoryManagerHandle elements.ServerElementHandle

	serviceManager       elements.ServiceManager
	serviceManagerHandle elements.ServerElementHandle

	securityHandler       elements.SecurityHandler
	securityHandlerHandle elements.ServerElementHandle

	messagingManager       elements.MessagingManager
	messagingManagerHandle elements.ServerElementHandle

	moduleManager       elements.ModuleManager
	moduleManagerHandle elements.ServerElementHandle

	taskManager       elements.TaskManager
	taskManagerHandle elements.ServerElementHandle

	rulesManager       elements.RulesManager
	rulesManagerHandle elements.ServerElementHandle

	cacheManager       elements.CacheManager
	cacheManagerHandle elements.ServerElementHandle

	secretsManager       elements.SecretsManager
	secretsManagerHandle elements.ServerElementHandle

	sessionManager       elements.SessionManager
	sessionManagerHandle elements.ServerElementHandle

	communicator       elements.Communicator
	communicatorHandle elements.ServerElementHandle

	//engines configured on the abstract server
	engineHandles map[string]elements.ServerElementHandle
	engineConf    map[string]config.Config
	engines       map[string]elements.Engine

	properties map[string]interface{}

	logger       elements.Logger
	loggerHandle elements.ServerElementHandle

	parent *abstractserver

	svrContext *serverContext

	proxy core.ServerElement

	baseDir string
}

func newAbstractServer(svrCtx *serverContext, name string, parent *abstractserver, proxy core.ServerElement, baseDir string) (*abstractserver, error) {
	as := &abstractserver{name: name, parent: parent, proxy: proxy, baseDir: baseDir, svrContext: svrCtx}
	log.Trace(svrCtx, "Base directory set to ", "Name", baseDir)
	svrCtx.Set(config.BASEDIR, baseDir)
	as.engineHandles = make(map[string]elements.ServerElementHandle)
	as.engines = make(map[string]elements.Engine)
	as.engineConf = make(map[string]config.Config)
	as.createNonConfComponents(svrCtx, name, parent, proxy)
	/*	if err := as.installModules(svrCtx, baseDir); err != nil {
		return nil, err
	}*/
	return as, nil
}

/*
func (as *abstractserver) installModules(ctx *serverContext, baseDir string) error {
	modulesDir := path.Join(baseDir, constants.CONF_MODULES_DIR)
	ok, fi, _ := utils.FileExists(modulesDir)
	if ok && fi.IsDir() {
		files, err := ioutil.ReadDir(modulesDir)
		if err != nil {
			return errors.WrapError(ctx, err, "Subdirectory", modulesDir)
		}

		for _, info := range files {
			if !info.IsDir() {
				name := info.Name()
				var extension = filepath.Ext(name)
				if extension == ".zip" {
					modulename := name[0 : len(name)-len(extension)]
					modulefileName := path.Join(modulesDir, name)
					moduleDir := path.Join(modulesDir, modulename)
					ok, fi, _ := utils.FileExists(moduleDir)
					if ok {
						ziptime := info.ModTime()
						dirtime := fi.ModTime()
						diff := ziptime.Sub(dirtime)
						if diff > 0 {
							//remove directory
							err := as.installModule(ctx, baseDir, modulefileName, moduleDir, modulename)
							if err != nil {
								return err
							}
						}
					} else {
						err := as.installModule(ctx, baseDir, modulefileName, moduleDir, modulename)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func (as *abstractserver) installModule(ctx *serverContext, baseDir, moduleFileName, moduleDir, modulename string) error {
	if err := os.RemoveAll(moduleDir); err != nil {
		return errors.WrapError(ctx, err)
	}
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		return errors.WrapError(ctx, err)
	}
	if err := utils.Unzip(moduleFileName, moduleDir); err != nil {
		return errors.WrapError(ctx, err)
	}

	prefix := fmt.Sprintf("__%s__.", modulename)

	servicesPath := path.Join(moduleDir, constants.CONF_SERVICES)
	ok, _, _ := utils.FileExists(servicesPath)
	if ok {
		utils.CopyDir(servicesPath, path.Join(baseDir, modulename, constants.CONF_SERVICES), prefix)
	}

	factoriesPath := path.Join(moduleDir, constants.CONF_FACTORIES)
	ok, _, _ = utils.FileExists(factoriesPath)
	if ok {
		utils.CopyDir(factoriesPath, path.Join(baseDir, modulename, constants.CONF_FACTORIES), prefix)
	}

	objectsPath := path.Join(moduleDir, constants.CONF_OBJECTLDR_OBJECTS)
	ok, _, _ = utils.FileExists(objectsPath)
	if ok {
		utils.CopyDir(objectsPath, path.Join(baseDir, modulename, constants.CONF_OBJECTLDR_OBJECTS), prefix)
	}

	return nil
}
*/
