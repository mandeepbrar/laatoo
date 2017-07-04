package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/constants"
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

func newAbstractServer(svrCtx *serverContext, name string, parent *abstractserver, proxy core.ServerElement, baseDir string) *abstractserver {
	as := &abstractserver{name: name, parent: parent, proxy: proxy, baseDir: baseDir, svrContext: svrCtx}
	log.Trace(svrCtx, "Base directory set to ", "Name", baseDir)
	svrCtx.Set(constants.CONF_BASE_DIR, baseDir)
	as.engineHandles = make(map[string]server.ServerElementHandle)
	as.engines = make(map[string]server.Engine)
	as.engineConf = make(map[string]config.Config)
	as.createNonConfComponents(svrCtx, name, parent, proxy)
	return as
}
