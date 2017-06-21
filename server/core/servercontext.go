package core

import (
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/server/common"
	deflog "log"
)

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
func (ctx *serverContext) GetServerType() string {
	return ctx.server.(*serverProxy).server.serverType
}

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
	return ctx.newservercontext(ctx.SubCtx(name))
}

//create a new server context; variables in this context be reflected in parent
//sets a context element
//id of the context is not changed. flow is updated
func (ctx *serverContext) SubContextWithElement(name string, primaryElement core.ServerElementType) core.ServerContext {
	retctx := ctx.newservercontext(ctx.SubCtx(name))
	elem := retctx.GetServerElement(primaryElement)
	if elem != nil {
		retctx.element = elem
		retctx.elementType = primaryElement
		return retctx
	}
	return nil
}

//creates a new server context that is duplicate of the parent.
func (ctx *serverContext) newservercontext(context core.Context) *serverContext {
	log.Debug(ctx, "Entering new servercontext ", "Elapsed Time ", ctx.GetElapsedTime(), "New Context Name", context.GetName())
	return &serverContext{Context: context.(*common.Context), server: ctx.server, serviceResponseHandler: ctx.serviceResponseHandler,
		engine: ctx.engine, objectLoader: ctx.objectLoader, application: ctx.application, environment: ctx.environment, securityHandler: ctx.securityHandler,
		factory: ctx.factory, factoryManager: ctx.factoryManager, serviceManager: ctx.serviceManager, service: ctx.service, channel: ctx.channel, msgManager: ctx.msgManager,
		channelManager: ctx.channelManager, rulesManager: ctx.rulesManager, cacheManager: ctx.cacheManager, taskManager: ctx.taskManager, logger: ctx.logger,
		open1: ctx.open1, open2: ctx.open2, open3: ctx.open3, element: ctx.element, elementType: ctx.elementType}
}

func (ctx *serverContext) NewContext(name string) core.ServerContext {
	return ctx.newContext(name)
}
func (ctx *serverContext) newContext(name string) *serverContext {
	return ctx.newservercontext(ctx.NewCtx(name))
}

//create a new server context from the parent. variables set in this context will not be reflected in parent
//id of the context is changed when new context is created
//sets a context element
func (ctx *serverContext) NewContextWithElements(name string, elements core.ContextMap, primaryElement core.ServerElementType) core.ServerContext {
	return ctx.newContextWithElements(name, elements, primaryElement)
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
func (ctx *serverContext) newContextWithElements(name string, elements core.ContextMap, primaryElement core.ServerElementType) *serverContext {
	newctx := ctx.newservercontext(ctx.NewCtx(name))
	newctx.setElements(elements, primaryElement)
	return newctx
}

//creates a new request with engine context
func (ctx *serverContext) CreateNewRequest(name string, engineCtx interface{}) core.RequestContext {
	log.Info(ctx, "Creating new request ", "Name", name)
	//a service must be there in the server context if a request is to be created
	if ctx.service == nil {
		return nil
	}
	reqCtx := ctx.createNewRequest(name, engineCtx, ctx.service)
	cacheToUse, ok := ctx.service.GetString("__cache")
	if ok {
		if ctx.cacheManager != nil {
			cache := ctx.cacheManager.GetCache(ctx, cacheToUse)
			reqCtx.cache = cache
		}
	}
	return reqCtx
}

func (ctx *serverContext) createNewRequest(name string, engineCtx interface{}, parent core.ServerElement) *requestContext {
	//create the request as a child of service context
	//so that the variables set by the service are available while executing a request

	newctx := parent.NewCtx(name)
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
