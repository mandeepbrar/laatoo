package core

import (
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
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

*/

//proxy object for the server
//this context is passed during initialization of factories and services
type serverContext struct {
	*common.Context
	element     core.ServerElement
	elementType core.ServerElementType
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
	//service for which an operation/request is being executed
	service        server.Service
	channel        server.Channel
	channelManager server.ChannelManager
	rulesManager   server.RulesManager
	cacheManager   server.CacheManager
	taskManager    server.TaskManager
	//open element for client applications
	open1 core.ServerElement
	open2 core.ServerElement
	open3 core.ServerElement
}

//create a new server context
//this is a proxy to the server
func newServerContext() *serverContext {
	return &serverContext{Context: common.NewContext("/")}
}

//returns the server type i.e. standalone or google app engine
/*func (ctx *serverContext) GetServerType() string {
	return ctx.server.(*serverProxy).server.serverType
}*/

//returns the primary element
func (ctx *serverContext) GetElement() core.ServerElement {
	return ctx.element
}

func (ctx *serverContext) GetElementType() core.ServerElementType {
	return ctx.elementType
}

func (ctx *serverContext) GetService(alias string) (core.Service, error) {
	svcStruct, err := ctx.serviceManager.GetService(ctx, alias)
	if err != nil {
		return nil, err
	}
	return svcStruct.Service(), nil
}

//get a server element applicable to the context by its type
func (ctx *serverContext) GetServerElement(elemType core.ServerElementType) core.ServerElement {
	switch elemType {
	case core.ServerElementServer:
		return ctx.server
	case core.ServerElementEngine:
		return ctx.engine
	case core.ServerElementEnvironment:
		return ctx.environment
	case core.ServerElementLoader:
		return ctx.objectLoader
	case core.ServerElementServiceFactory:
		return ctx.factory
	case core.ServerElementApplication:
		return ctx.application
	case core.ServerElementService:
		return ctx.service
	case core.ServerElementServiceManager:
		return ctx.serviceManager
	case core.ServerElementFactoryManager:
		return ctx.factoryManager
	case core.ServerElementChannel:
		return ctx.channel
	case core.ServerElementChannelManager:
		return ctx.channelManager
	case core.ServerElementSecurityHandler:
		return ctx.securityHandler
	case core.ServerElementMessagingManager:
		return ctx.msgManager
	case core.ServerElementServiceResponseHandler:
		return ctx.serviceResponseHandler
	case core.ServerElementRulesManager:
		return ctx.rulesManager
	case core.ServerElementCacheManager:
		return ctx.cacheManager
	case core.ServerElementTaskManager:
		return ctx.taskManager
	case core.ServerElementLogger:
		return ctx.logger
	case core.ServerElementOpen1:
		return ctx.open1
	case core.ServerElementOpen2:
		return ctx.open2
	case core.ServerElementOpen3:
		return ctx.open3
	}
	return nil
}

//create a child context that with the same underlying context
//changes made to context parameters will be visible on the parent
//id of the context is also retained.. this can be used to track flow
func (ctx *serverContext) SubContext(name string) core.ServerContext {
	return ctx.subContext(name)
}

func (ctx *serverContext) subContext(name string) *serverContext {
	subctx := ctx.SubCtx(name)
	log.Debug(ctx, "Entering new subcontext ", "Elapsed Time ", ctx.GetElapsedTime(), "New Context Name", name)
	return &serverContext{Context: subctx.(*common.Context), server: ctx.server, serviceResponseHandler: ctx.serviceResponseHandler,
		engine: ctx.engine, objectLoader: ctx.objectLoader, application: ctx.application, environment: ctx.environment, securityHandler: ctx.securityHandler,
		factory: ctx.factory, factoryManager: ctx.factoryManager, serviceManager: ctx.serviceManager, service: ctx.service, channel: ctx.channel, msgManager: ctx.msgManager,
		channelManager: ctx.channelManager, rulesManager: ctx.rulesManager, cacheManager: ctx.cacheManager, taskManager: ctx.taskManager, logger: ctx.logger,
		open1: ctx.open1, open2: ctx.open2, open3: ctx.open3, element: ctx.element, elementType: ctx.elementType}
}

//create a new server context; variables in this context be reflected in parent
//sets a context element
//id of the context is not changed. flow is updated
func (ctx *serverContext) SubContextWithElement(name string, primaryElement core.ServerElementType) core.ServerContext {
	retctx := ctx.subContext(name)
	elem := retctx.GetServerElement(primaryElement)
	if elem != nil {
		retctx.element = elem
		retctx.elementType = primaryElement
		return retctx
	}
	return nil
}

func (ctx *serverContext) NewContext(name string) core.ServerContext {
	return ctx.newContext(name)
}
func (ctx *serverContext) newContext(name string) *serverContext {
	newctx := ctx.NewCtx(name)
	log.Debug(ctx, "Entering new server context ", "Elapsed Time ", ctx.GetElapsedTime(), "New Context Name", name)
	svrCtx := &serverContext{Context: newctx.(*common.Context), server: ctx.checkNil(ctx.server), serviceResponseHandler: ctx.checkNil(ctx.serviceResponseHandler).(server.ServiceResponseHandler),
		engine: ctx.checkNil(ctx.engine).(server.Engine), objectLoader: ctx.checkNil(ctx.objectLoader).(server.ObjectLoader), application: ctx.checkNil(ctx.application).(server.Application),
		environment: ctx.checkNil(ctx.environment).(server.Environment), securityHandler: ctx.checkNil(ctx.securityHandler).(server.SecurityHandler), factory: ctx.checkNil(ctx.factory).(server.Factory),
		factoryManager: ctx.checkNil(ctx.factoryManager).(server.FactoryManager), serviceManager: ctx.checkNil(ctx.serviceManager).(server.ServiceManager), service: ctx.checkNil(ctx.service).(server.Service),
		channel: ctx.checkNil(ctx.channel).(server.Channel), msgManager: ctx.checkNil(ctx.msgManager).(server.MessagingManager), channelManager: ctx.checkNil(ctx.channelManager).(server.ChannelManager),
		rulesManager: ctx.checkNil(ctx.rulesManager).(server.RulesManager), cacheManager: ctx.checkNil(ctx.cacheManager).(server.CacheManager), taskManager: ctx.checkNil(ctx.taskManager).(server.TaskManager),
		logger: ctx.checkNil(ctx.logger).(server.Logger), open1: ctx.open1, open2: ctx.open2, open3: ctx.open3}
	elem := svrCtx.GetServerElement(ctx.elementType)
	svrCtx.element = elem
	svrCtx.elementType = ctx.elementType
	return svrCtx
}

//create a new server context from the parent. variables set in this context will not be reflected in parent
//id of the context is changed when new context is created
//sets a context element
func (ctx *serverContext) NewContextWithElements(name string, elements core.ContextMap, primaryElement core.ServerElementType) core.ServerContext {
	return ctx.newContextWithElements(name, elements, primaryElement)
}

func (ctx *serverContext) newContextWithElements(name string, elements core.ContextMap, primaryElement core.ServerElementType) *serverContext {
	newctx := ctx.newContext(name)
	newctx.setElements(elements, primaryElement)
	return newctx
}

func (ctx *serverContext) checkNil(elem core.ServerElement) core.ServerElement {
	if elem != nil {
		return elem.Reference()
	} else {
		return nil
	}
}

func (ctx *serverContext) setElements(elements core.ContextMap, primaryElement core.ServerElementType) {
	for elementToSet, element := range elements {
		switch elementToSet {
		case core.ServerElementServer:
			if element == nil {
				ctx.server = nil
			} else {
				ctx.server = element.(server.Server)
			}
		case core.ServerElementEngine:
			if element == nil {
				ctx.engine = nil
			} else {
				ctx.engine = element.(server.Engine)
			}
		case core.ServerElementEnvironment:
			if element == nil {
				ctx.environment = nil
			} else {
				ctx.environment = element.(server.Environment)
			}
		case core.ServerElementLoader:
			if element == nil {
				ctx.objectLoader = nil
			} else {
				ctx.objectLoader = element.(server.ObjectLoader)
			}
		case core.ServerElementServiceFactory:
			if element == nil {
				ctx.factory = nil
			} else {
				ctx.factory = element.(server.Factory)
			}
		case core.ServerElementApplication:
			if element == nil {
				ctx.application = nil
			} else {
				ctx.application = element.(server.Application)
			}
		case core.ServerElementService:
			if element == nil {
				ctx.service = nil
			} else {
				ctx.service = element.(server.Service)
			}
		case core.ServerElementChannelManager:
			if element == nil {
				ctx.channelManager = nil
			} else {
				ctx.channelManager = element.(server.ChannelManager)
			}
		case core.ServerElementChannel:
			if element == nil {
				ctx.channel = nil
			} else {
				ctx.channel = element.(server.Channel)
			}
		case core.ServerElementServiceManager:
			if element == nil {
				ctx.serviceManager = nil
			} else {
				ctx.serviceManager = element.(server.ServiceManager)
			}
		case core.ServerElementFactoryManager:
			if element == nil {
				ctx.factoryManager = nil
			} else {
				ctx.factoryManager = element.(server.FactoryManager)
			}
		case core.ServerElementServiceResponseHandler:
			if element == nil {
				ctx.serviceResponseHandler = nil
			} else {
				ctx.serviceResponseHandler = element.(server.ServiceResponseHandler)
			}
		case core.ServerElementSecurityHandler:
			if element == nil {
				ctx.securityHandler = nil
			} else {
				ctx.securityHandler = element.(server.SecurityHandler)
			}
		case core.ServerElementMessagingManager:
			if element == nil {
				ctx.msgManager = nil
			} else {
				ctx.msgManager = element.(server.MessagingManager)
			}
		case core.ServerElementRulesManager:
			if element == nil {
				ctx.rulesManager = nil
			} else {
				ctx.rulesManager = element.(server.RulesManager)
			}
		case core.ServerElementCacheManager:
			if element == nil {
				ctx.cacheManager = nil
			} else {
				ctx.cacheManager = element.(server.CacheManager)
			}
		case core.ServerElementTaskManager:
			if element == nil {
				ctx.taskManager = nil
			} else {
				ctx.taskManager = element.(server.TaskManager)
			}
		case core.ServerElementLogger:
			if element == nil {
				ctx.logger = nil
			} else {
				ctx.logger = element.(server.Logger)
			}
		case core.ServerElementOpen1:
			ctx.open1 = element
		case core.ServerElementOpen2:
			ctx.open2 = element
		case core.ServerElementOpen3:
			ctx.open3 = element
		}
	}
	elem := ctx.GetServerElement(primaryElement)
	if elem != nil {
		ctx.element = elem
		ctx.elementType = primaryElement
	}
}

//creates a new request with engine context
func (ctx *serverContext) CreateNewRequest(name string, params interface{}) core.RequestContext {

	rparams := params.(*common.RequestContextParams)
	log.Info(ctx, "Creating new request ", "Name", name)
	//a service must be there in the server context if a request is to be created
	if ctx.service == nil {
		return nil
	}
	reqCtx := ctx.createNewRequest(name, rparams.EngineContext, ctx.service)

	if rparams.Logger != nil {
		reqCtx.logger = rparams.Logger
	}

	if rparams.Cache != nil {
		reqCtx.cache = rparams.Cache
	}

	return reqCtx
}

func (ctx *serverContext) createNewRequest(name string, engineCtx interface{}, parent core.ServerElement) *requestContext {
	//create the request as a child of service context
	//so that the variables set by the service are available while executing a request

	newctx := ctx.NewCtx(name)
	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx, logger: ctx.logger,
		engineContext: engineCtx, subRequest: false}
}

func (ctx *serverContext) CreateCollection(objectName string, length int) (interface{}, error) {
	return ctx.objectLoader.CreateCollection(ctx, objectName, length)
}

func (ctx *serverContext) CreateObject(objectName string) (interface{}, error) {
	return ctx.objectLoader.CreateObject(ctx, objectName)
}

func (ctx *serverContext) GetObjectCollectionCreator(objectName string) (core.ObjectCollectionCreator, error) {
	return ctx.objectLoader.GetObjectCollectionCreator(ctx, objectName)
}

func (ctx *serverContext) GetObjectCreator(objectName string) (core.ObjectCreator, error) {
	return ctx.objectLoader.GetObjectCreator(ctx, objectName)
}

func (ctx *serverContext) GetMethod(methodName string) (core.ServiceFunc, error) {
	return ctx.objectLoader.GetMethod(ctx, methodName)
}

func (ctx *serverContext) CreateSystemRequest(name string) core.RequestContext {
	reqCtx := ctx.createNewRequest(name, nil, ctx.element)
	reqCtx.user = nil
	reqCtx.admin = true
	return reqCtx
}

func (ctx *serverContext) SubscribeTopic(topics []string, lstnr core.ServiceFunc) error {
	if ctx.msgManager != nil {
		return ctx.msgManager.Subscribe(ctx, topics, lstnr)
	}
	return nil
}
func (ctx *serverContext) LogTrace(msg string, args ...interface{}) {
	if ctx.logger != nil {
		ctx.logger.Trace(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogDebug(msg string, args ...interface{}) {
	if ctx.logger != nil {
		ctx.logger.Debug(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogInfo(msg string, args ...interface{}) {
	if ctx.logger != nil {
		ctx.logger.Info(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogWarn(msg string, args ...interface{}) {
	if ctx.logger != nil {
		ctx.logger.Warn(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogError(msg string, args ...interface{}) {
	if ctx.logger != nil {
		ctx.logger.Error(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogFatal(msg string, args ...interface{}) {
	if ctx.logger != nil {
		ctx.logger.Fatal(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}
