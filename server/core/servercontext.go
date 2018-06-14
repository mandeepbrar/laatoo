package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	slog "laatoo/server/log"
	deflog "log"
)

/*
	Server context movement

	RootCtx -> Server -> Server Create -> Server Init -> Server Start -------- Server context---Context A --- same context or sub context travels
	Server  -> Factory -> Service -> Channel---- Server Context ---- Context A1
	Server Context -> SecurityHandler --- context A2
	ServerContext -> TaskManager - > Tasks --- context A3
	ServerContext -> RulesManager - > Rules --- context A4
	FactoryManager ---- same as above
	Service Manager---- same as above
	Log Manager--- depends on parent... can be changed from factory, service or channel
	....
	ServerContext -> Started Server-> Environment Create -> Environment Init -> environment Start -> AE ---- new context from server..... same context travels in env init
	Environment -> Factory -> Service - > Channel ---- Server Context ----- Context AE1 (for inherited service... it still runs in server )
	Evironment Server Context -> SecurityHandler --- context AE2
	Evironment ServerContext -> TaskManager - > Tasks --- context AE3 --- context does not flow from task manager to task manager because they have parent child relations
	Evironment ServerContext -> RulesManager - > RulesManager --- context AE3 --- context does not flow from task manager to task manager because they have parent child relations
	Evironment ServerContext -> ChannelManager - > Channels --- context AE4
	FactoryManager ---- same as above
	Service Manager---- same as above
	..
	..
	>>>>>>>>>>>>>Factory Manager, Service Manager, Channel MAnager, Rules, Tasks need to pick right context to ensure execution of a service. >>>> Server -> Evironment -> Application-> Factory->Service->Channel
	....
	Started Environment Context -> Create Application Server Context -> Application Init - > Application start ---- AEA
	Server -> Environment -> Application-> Factory -> Service - > Channel -> RequestContext---- Server Context ----- Context AEA1 (for inherited service... it still runs in server )


	New context --- copies an original context---- duplicates the context map, copies the references of server elements.. changes are not refelected back to original context
	Sub context --- name and path change... changes on server elements (excluding logger) are reflected on original context... changes on map are also reflected back...

	e.g. in subcontext... same proxy object as parent is used.... so any changes to proxy in the child are reflected in the parent as well
	in new context... duplicate reference of the same proxy object is used...

*/

type contextElements struct {
	//main server reference
	server server.Server
	//engine to be used in a context
	engine server.Engine
	//object loader to be used in a context
	objectLoader server.ObjectLoader
	//application on which work is being done
	application server.Application
	//environment on which operations are being carried out
	environment server.Environment
	//factory manager applicable in a context
	factoryManager server.FactoryManager
	//service manager applicable in a context
	serviceManager server.ServiceManager
	//security handler applicable in a context
	securityHandler server.SecurityHandler
	//service response handler applicable when a request is being executed
	serviceResponseHandler server.ServiceResponseHandler
	//pubsub
	msgManager server.MessagingManager
	//logger
	logger server.Logger
	//factory for which operation is being done
	factory server.Factory
	//properties
	properties map[string]interface{}
	//service for which an operation/request is being executed
	service        server.Service
	channel        server.Channel
	channelManager server.ChannelManager
	rulesManager   server.RulesManager
	cacheManager   server.CacheManager
	taskManager    server.TaskManager
	moduleManager  server.ModuleManager
	sessionManager server.SessionManager
	module         server.Module

	//open element for client applications
	open1 core.ServerElement
	open2 core.ServerElement
	open3 core.ServerElement
}

//proxy object for the server
//this context is passed during initialization of factories and services
type serverContext struct {
	*common.Context
	level          int
	elements       *contextElements
	childContexts  []*serverContext
	sessionManager *sessionManager
}

//create a new server context
//this is a proxy to the server
func newServerContext() *serverContext {
	ctx := &serverContext{Context: common.NewContext("/"), elements: &contextElements{}, level: 0, childContexts: make([]*serverContext, 0)}
	_, logger := slog.NewLogger(ctx, "Default")
	ctx.setElements(core.ContextMap{core.ServerElementLogger: logger})
	return ctx
}

func (ctx *serverContext) GetService(alias string) (core.Service, error) {
	svcStruct, err := ctx.elements.serviceManager.GetService(ctx, alias)
	if err != nil {
		return nil, err
	}
	return svcStruct.Service(), nil
}

//get a server element applicable to the context by its type
func (ctx *serverContext) GetServerElement(elemType core.ServerElementType) core.ServerElement {
	switch elemType {
	case core.ServerElementServer:
		return ctx.elements.server
	case core.ServerElementEngine:
		return ctx.elements.engine
	case core.ServerElementEnvironment:
		return ctx.elements.environment
	case core.ServerElementLoader:
		return ctx.elements.objectLoader
	case core.ServerElementServiceFactory:
		return ctx.elements.factory
	case core.ServerElementApplication:
		return ctx.elements.application
	case core.ServerElementService:
		return ctx.elements.service
	case core.ServerElementServiceManager:
		return ctx.elements.serviceManager
	case core.ServerElementFactoryManager:
		return ctx.elements.factoryManager
	case core.ServerElementChannel:
		return ctx.elements.channel
	case core.ServerElementChannelManager:
		return ctx.elements.channelManager
	case core.ServerElementSecurityHandler:
		return ctx.elements.securityHandler
	case core.ServerElementMessagingManager:
		return ctx.elements.msgManager
	case core.ServerElementServiceResponseHandler:
		return ctx.elements.serviceResponseHandler
	case core.ServerElementRulesManager:
		return ctx.elements.rulesManager
	case core.ServerElementCacheManager:
		return ctx.elements.cacheManager
	case core.ServerElementTaskManager:
		return ctx.elements.taskManager
	case core.ServerElementModuleManager:
		return ctx.elements.moduleManager
	case core.ServerElementSessionManager:
		return ctx.elements.sessionManager
	case core.ServerElementModule:
		return ctx.elements.module
	case core.ServerElementLogger:
		return ctx.elements.logger
	case core.ServerElementOpen1:
		return ctx.elements.open1
	case core.ServerElementOpen2:
		return ctx.elements.open2
	case core.ServerElementOpen3:
		return ctx.elements.open3
	}
	return nil
}

//create a child context that with the same underlying context
//changes made to context parameters will be visible on the parent
//id of the context is also retained.. this can be used to track flow
func (ctx *serverContext) SubContext(name string) core.ServerContext {
	return ctx.subContext(name)
}

//create a new server context; variables in this context be reflected in parent
//sets a context element
//id of the context is not changed. flow is updated
func (ctx *serverContext) subContext(name string) *serverContext {
	subctx := ctx.SubCtx(name)
	log.Debug(ctx, "Entering new subcontext ", "Elapsed Time ", ctx.GetElapsedTime(), "New Context Name", name)
	return &serverContext{Context: subctx.(*common.Context), elements: ctx.elements, level: ctx.level, sessionManager: ctx.sessionManager, childContexts: ctx.childContexts}
}

//create a new server context from the parent. variables set in this context will not be reflected in parent
//id of the context is changed when new context is created
/*func (ctx *serverContext) NewContext(name string) core.ServerContext {
	return ctx.newContext(name)
}*/

func (ctx *serverContext) newContext(name string) *serverContext {
	newctx := ctx.NewCtx(name)
	log.Debug(ctx, "Entering new server context ", "Elapsed Time ", ctx.GetElapsedTime(), "Name ", name)

	svrCtx := &serverContext{Context: newctx.(*common.Context), elements: &contextElements{}, sessionManager: ctx.sessionManager, level: ctx.level + 1, childContexts: make([]*serverContext, 0)}
	cmap := ctx.getElementsContextMap()
	svrCtx.elements.properties = ctx.elements.properties
	svrCtx.setElementReferences(cmap, true)
	ctx.addChild(svrCtx)
	return svrCtx
}

func (ctx *serverContext) getElementsContextMap() core.ContextMap {
	return core.ContextMap{core.ServerElementServer: ctx.elements.server, core.ServerElementEngine: ctx.elements.engine, core.ServerElementEnvironment: ctx.elements.environment,
		core.ServerElementLoader: ctx.elements.objectLoader, core.ServerElementServiceFactory: ctx.elements.factory, core.ServerElementApplication: ctx.elements.application,
		core.ServerElementService: ctx.elements.service, core.ServerElementServiceManager: ctx.elements.serviceManager, core.ServerElementFactoryManager: ctx.elements.factoryManager,
		core.ServerElementChannel: ctx.elements.channel, core.ServerElementChannelManager: ctx.elements.channelManager, core.ServerElementSecurityHandler: ctx.elements.securityHandler,
		core.ServerElementMessagingManager: ctx.elements.msgManager, core.ServerElementServiceResponseHandler: ctx.elements.serviceResponseHandler, core.ServerElementRulesManager: ctx.elements.rulesManager,
		core.ServerElementCacheManager: ctx.elements.cacheManager, core.ServerElementTaskManager: ctx.elements.taskManager, core.ServerElementLogger: ctx.elements.logger,
		core.ServerElementModuleManager: ctx.elements.moduleManager, core.ServerElementModule: ctx.elements.module,
		core.ServerElementOpen1: ctx.elements.open1, core.ServerElementOpen2: ctx.elements.open2, core.ServerElementOpen3: ctx.elements.open3}
}

func (ctx *serverContext) addChild(child *serverContext) {
	ctx.childContexts = append(ctx.childContexts, child)
}

func (ctx *serverContext) setElements(elements core.ContextMap) {
	ctx.setElementReferences(elements, false)
}

func (ctx *serverContext) setElementReferences(elements core.ContextMap, ref bool) {
	ctxElems := ctx.elements
	for elementToSet, element := range elements {
		if ref && element != nil {
			element = element.Reference()
		}
		switch elementToSet {
		case core.ServerElementServer:
			if element == nil {
				ctxElems.server = nil
			} else {
				ctxElems.server = element.(server.Server)
			}
		case core.ServerElementEngine:
			if element == nil {
				ctxElems.engine = nil
			} else {
				ctxElems.engine = element.(server.Engine)
			}
		case core.ServerElementEnvironment:
			if element == nil {
				ctxElems.environment = nil
			} else {
				ctxElems.environment = element.(server.Environment)
			}
		case core.ServerElementLoader:
			if element == nil {
				ctxElems.objectLoader = nil
			} else {
				ctxElems.objectLoader = element.(server.ObjectLoader)
			}
		case core.ServerElementServiceFactory:
			if element == nil {
				ctxElems.factory = nil
			} else {
				ctxElems.factory = element.(server.Factory)
			}
		case core.ServerElementApplication:
			if element == nil {
				ctxElems.application = nil
			} else {
				ctxElems.application = element.(server.Application)
			}
		case core.ServerElementService:
			if element == nil {
				ctxElems.service = nil
			} else {
				ctxElems.service = element.(server.Service)
			}
		case core.ServerElementChannelManager:
			if element == nil {
				ctxElems.channelManager = nil
			} else {
				ctxElems.channelManager = element.(server.ChannelManager)
			}
		case core.ServerElementChannel:
			if element == nil {
				ctxElems.channel = nil
			} else {
				ctxElems.channel = element.(server.Channel)
			}
		case core.ServerElementServiceManager:
			if element == nil {
				ctxElems.serviceManager = nil
			} else {
				ctxElems.serviceManager = element.(server.ServiceManager)
			}
		case core.ServerElementSessionManager:
			if element == nil {
				ctxElems.sessionManager = nil
				ctx.sessionManager = nil
			} else {
				ctxElems.sessionManager = element.(server.SessionManager)
				ctx.sessionManager = element.(*sessionManagerProxy).manager
			}
		case core.ServerElementFactoryManager:
			if element == nil {
				ctxElems.factoryManager = nil
			} else {
				ctxElems.factoryManager = element.(server.FactoryManager)
			}
		case core.ServerElementServiceResponseHandler:
			if element == nil {
				ctxElems.serviceResponseHandler = nil
			} else {
				ctxElems.serviceResponseHandler = element.(server.ServiceResponseHandler)
			}
		case core.ServerElementSecurityHandler:
			if element == nil {
				ctxElems.securityHandler = nil
			} else {
				ctxElems.securityHandler = element.(server.SecurityHandler)
			}
		case core.ServerElementMessagingManager:
			if element == nil {
				ctxElems.msgManager = nil
			} else {
				ctxElems.msgManager = element.(server.MessagingManager)
			}
		case core.ServerElementRulesManager:
			if element == nil {
				ctxElems.rulesManager = nil
			} else {
				ctxElems.rulesManager = element.(server.RulesManager)
			}
		case core.ServerElementCacheManager:
			if element == nil {
				ctxElems.cacheManager = nil
			} else {
				ctxElems.cacheManager = element.(server.CacheManager)
			}
		case core.ServerElementTaskManager:
			if element == nil {
				ctxElems.taskManager = nil
			} else {
				ctxElems.taskManager = element.(server.TaskManager)
			}
		case core.ServerElementLogger:
			if element == nil {
				ctxElems.logger = nil
			} else {
				ctxElems.logger = element.(server.Logger)
			}
		case core.ServerElementModuleManager:
			if element == nil {
				ctxElems.moduleManager = nil
			} else {
				ctxElems.moduleManager = element.(server.ModuleManager)
			}
		case core.ServerElementModule:
			if element == nil {
				ctxElems.module = nil
			} else {
				ctxElems.module = element.(server.Module)
			}
		case core.ServerElementOpen1:
			ctxElems.open1 = element
		case core.ServerElementOpen2:
			ctxElems.open2 = element
		case core.ServerElementOpen3:
			ctxElems.open3 = element
		}
	}
	for _, c := range ctx.childContexts {
		c.setElements(elements)
	}
}

func (ctx *serverContext) GetServerProperties() map[string]interface{} {
	return ctx.elements.properties
}

func (ctx *serverContext) setServerProperties(props map[string]interface{}) {
	ctx.elements.properties = props
}

//creates a new request with engine context
func (ctx *serverContext) CreateNewRequest(name string, engineCtx interface{}, sessionId string) (core.RequestContext, error) {

	log.Info(ctx, "Creating new request ", "Name", name)
	//a service must be there in the server context if a request is to be created
	if ctx.elements.service == nil {
		return nil, errors.MissingService(ctx, name)
	}
	reqCtx, err := ctx.createNewRequest(name, engineCtx, sessionId)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	reqCtx.logger = ctx.elements.logger
	//svc := ctx.elements.service.(*serviceProxy)

	cacheToUse, ok := ctx.GetString("__cache")
	if ok {
		if ctx.elements.cacheManager != nil {
			cache := ctx.elements.cacheManager.GetCache(ctx, cacheToUse)
			reqCtx.cache = cache
		}
	}

	return reqCtx, nil
}

func (ctx *serverContext) createNewRequest(name string, engineCtx interface{}, sessionId string) (*requestContext, error) {
	//create the request as a child of service context
	//so that the variables set by the service are available while executing a request

	newctx := ctx.NewCtx(name)

	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx, logger: ctx.elements.logger,
		engineContext: engineCtx, sessionId: sessionId, subRequest: false}, nil
}

func (ctx *serverContext) CreateCollection(objectName string, length int) (interface{}, error) {
	return ctx.elements.objectLoader.CreateCollection(ctx, objectName, length)
}

func (ctx *serverContext) CreateObject(objectName string) (interface{}, error) {
	return ctx.elements.objectLoader.CreateObject(ctx, objectName)
}

func (ctx *serverContext) GetObjectCollectionCreator(objectName string) (core.ObjectCollectionCreator, error) {
	return ctx.elements.objectLoader.GetObjectCollectionCreator(ctx, objectName)
}

func (ctx *serverContext) GetObjectCreator(objectName string) (core.ObjectCreator, error) {
	return ctx.elements.objectLoader.GetObjectCreator(ctx, objectName)
}

func (ctx *serverContext) GetObjectMetadata(objectName string) (core.Info, error) {
	return ctx.elements.objectLoader.GetMetaData(ctx, objectName)
}

func (ctx *serverContext) CreateSystemRequest(name string) core.RequestContext {
	reqCtx, err := ctx.createNewRequest(name, nil, "")
	if err != nil {
		log.Error(ctx, "Error while creating system request", "Error", err)
	}
	reqCtx.user = nil
	reqCtx.admin = true
	return reqCtx
}

func (ctx *serverContext) SubscribeTopic(topics []string, lstnr core.MessageListener) error {
	if ctx.elements.msgManager != nil {
		return ctx.elements.msgManager.Subscribe(ctx, topics, lstnr)
	}
	return nil
}

func (ctx *serverContext) CreateConfig() config.Config {
	return make(common.GenericConfig)
}

func (ctx *serverContext) ReadConfigMap(cfg map[string]interface{}) (config.Config, error) {
	res, _ := common.CastToConfig(cfg)
	return res, nil
}

func (ctx *serverContext) ReadConfig(file string, funcs map[string]interface{}) (config.Config, error) {
	return common.NewConfigFromFile(ctx, file, funcs)
}

func (ctx *serverContext) ReadConfigData(data []byte, funcs map[string]interface{}) (config.Config, error) {
	return common.NewConfig(ctx, data, funcs)
}

func (ctx *serverContext) LogTrace(msg string, args ...interface{}) {
	if ctx.elements.logger != nil {
		ctx.elements.logger.Trace(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogDebug(msg string, args ...interface{}) {
	if ctx.elements.logger != nil {
		ctx.elements.logger.Debug(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogInfo(msg string, args ...interface{}) {
	if ctx.elements.logger != nil {
		ctx.elements.logger.Info(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogWarn(msg string, args ...interface{}) {
	if ctx.elements.logger != nil {
		ctx.elements.logger.Warn(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogError(msg string, args ...interface{}) {
	if ctx.elements.logger != nil {
		ctx.elements.logger.Error(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogFatal(msg string, args ...interface{}) {
	if ctx.elements.logger != nil {
		ctx.elements.logger.Fatal(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}
