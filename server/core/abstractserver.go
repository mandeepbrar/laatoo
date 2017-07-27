package core

import (
	"io/ioutil"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"os"
	"path"
	"path/filepath"
)

const (
	singleFileByteLimit = 107374182400 // 1 GB
	chunkSize           = 4096         // 4 KB
)

type abstractserver struct {
	name string

	objectLoader       server.ObjectLoader
	objectLoaderHandle server.ServerElementHandle

	channelManager       server.ChannelManager
	channelManagerHandle server.ServerElementHandle

	factoryManager       server.FactoryManager
	factoryManagerHandle server.ServerElementHandle

	serviceManager       server.ServiceManager
	serviceManagerHandle server.ServerElementHandle

	securityHandler       server.SecurityHandler
	securityHandlerHandle server.ServerElementHandle

	messagingManager       server.MessagingManager
	messagingManagerHandle server.ServerElementHandle

	taskManager       server.TaskManager
	taskManagerHandle server.ServerElementHandle

	rulesManager       server.RulesManager
	rulesManagerHandle server.ServerElementHandle

	cacheManager       server.CacheManager
	cacheManagerHandle server.ServerElementHandle

	//engines configured on the abstract server
	engineHandles map[string]server.ServerElementHandle
	engineConf    map[string]config.Config
	engines       map[string]server.Engine

	logger       server.Logger
	loggerHandle server.ServerElementHandle

	parent *abstractserver

	svrContext *serverContext

	proxy core.ServerElement

	baseDir string
}

func newAbstractServer(svrCtx *serverContext, name string, parent *abstractserver, proxy core.ServerElement, baseDir string) (*abstractserver, error) {
	as := &abstractserver{name: name, parent: parent, proxy: proxy, baseDir: baseDir, svrContext: svrCtx}
	log.Trace(svrCtx, "Base directory set to ", "Name", baseDir)
	svrCtx.Set(constants.CONF_BASE_DIR, baseDir)
	as.engineHandles = make(map[string]server.ServerElementHandle)
	as.engines = make(map[string]server.Engine)
	as.engineConf = make(map[string]config.Config)
	as.createNonConfComponents(svrCtx, name, parent, proxy)
	if err := as.installModules(svrCtx, baseDir); err != nil {
		return nil, err
	}
	return as, nil
}

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
	/*confFile := path.Join(moduleDir, constants.CONF_CONFIG_FILE)
	conf, err := common.NewConfigFromFile(confFile)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/

	servicesPath := path.Join(moduleDir, constants.CONF_SERVICES)
	ok, _, _ := utils.FileExists(servicesPath)
	if ok {
		utils.CopyDir(path.Join(baseDir, constants.CONF_SERVICES), servicesPath)
	}

	factoriesPath := path.Join(moduleDir, constants.CONF_FACTORIES)
	ok, _, _ = utils.FileExists(factoriesPath)
	if ok {
		utils.CopyDir(path.Join(baseDir, constants.CONF_FACTORIES), factoriesPath)
	}

	objectsPath := path.Join(moduleDir, constants.CONF_OBJECTLDR_OBJECTS)
	ok, _, _ = utils.FileExists(objectsPath)
	if ok {
		utils.CopyDir(path.Join(baseDir, constants.CONF_OBJECTLDR_OBJECTS), objectsPath)
	}

	return nil
}
