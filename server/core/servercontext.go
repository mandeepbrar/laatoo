package core

import (
	googleContext "context"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/codecs"
	"laatoo/server/common"
	slog "laatoo/server/log"
	deflog "log"
	"time"
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
	server elements.Server
	//object loader to be used in a context
	objectLoader elements.ObjectLoader
	//application on which work is being done
	application elements.Application
	//environment on which operations are being carried out
	environment elements.Environment
	//factory manager applicable in a context
	factoryManager elements.FactoryManager
	//service manager applicable in a context
	serviceManager elements.ServiceManager
	//security handler applicable in a context
	securityHandler elements.SecurityHandler
	//service response handler applicable when a request is being executed
	serviceResponseHandler elements.ServiceResponseHandler
	//pubsub
	msgManager elements.MessagingManager
	//logger
	logger elements.Logger
	//factory for which operation is being done
	factory elements.Factory
	//properties
	properties map[string]interface{}
	//service for which an operation/request is being executed
	service        elements.Service
	channel        elements.Channel
	channelManager elements.ChannelManager
	rulesManager   elements.RulesManager
	cacheManager   elements.CacheManager
	taskManager    elements.TaskManager
	moduleManager  elements.ModuleManager
	sessionManager elements.SessionManager
	module         elements.Module
	communicator   elements.Communicator
	secretsManager elements.SecretsManager

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
	svrElements    *contextElements
	childContexts  []*serverContext
	sessionManager *sessionManager
}

//create a new server context
//this is a proxy to the server
func newServerContext() *serverContext {
	ctx := &serverContext{Context: common.NewContext("/"), svrElements: &contextElements{}, level: 0, childContexts: make([]*serverContext, 0)}
	_, logger := slog.NewLogger(ctx, "Default")
	ctx.setElements(core.ContextMap{core.ServerElementLogger: logger})
	return ctx
}

func (ctx *serverContext) Deadline() (deadline time.Time, ok bool) {
	return ctx.Context.Deadline()
}

func (ctx *serverContext) Done() <-chan struct{} {
	return ctx.Context.Done()
}

func (ctx *serverContext) Err() error {
	return ctx.Context.Err()
}

func (ctx *serverContext) Value(key interface{}) interface{} {
	return ctx.Context.Value(key)
}

func (ctx *serverContext) WithCancel() (ctx.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithCancel(ctx)
	return ctx.WithContext(newgooglectx), cancelFunc
}

func (ctx *serverContext) WithDeadline(timeout time.Time) (ctx.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithDeadline(ctx, timeout)
	return ctx.WithContext(newgooglectx), cancelFunc
}

func (ctx *serverContext) WithTimeout(timeout time.Duration) (ctx.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithTimeout(ctx, timeout)
	return ctx.WithContext(newgooglectx), cancelFunc
}

func (ctx *serverContext) WithValue(key, val interface{}) ctx.Context {
	newgooglectx := googleContext.WithValue(ctx, key, val)
	return ctx.WithContext(newgooglectx)
}

func (ctx *serverContext) WithContext(parent googleContext.Context) ctx.Context {
	return ctx.subContext(ctx.Name, ctx.Context.WithContext(googleContext.WithValue(parent, "tId", ctx.GetId())).(*common.Context))
}

func (ctx *serverContext) GetService(alias string) (core.Service, error) {
	svcStruct, err := ctx.svrElements.serviceManager.GetService(ctx, alias)
	if err != nil {
		return nil, err
	}
	return svcStruct.Service(), nil
}

func (ctx *serverContext) getService(alias string) (elements.Service, error) {
	return ctx.svrElements.serviceManager.(*serviceManagerProxy).manager.getService(ctx, alias)
}

func (ctx *serverContext) GetServiceContext(serviceName string) (core.ServerContext, error) {
	return ctx.svrElements.serviceManager.GetServiceContext(ctx, serviceName)
}

//get a server element applicable to the context by its type
func (ctx *serverContext) GetServerElement(elemType core.ServerElementType) core.ServerElement {
	switch elemType {
	case core.ServerElementServer:
		return ctx.svrElements.server
	case core.ServerElementEnvironment:
		return ctx.svrElements.environment
	case core.ServerElementLoader:
		return ctx.svrElements.objectLoader
	case core.ServerElementServiceFactory:
		return ctx.svrElements.factory
	case core.ServerElementApplication:
		return ctx.svrElements.application
	case core.ServerElementService:
		return ctx.svrElements.service
	case core.ServerElementServiceManager:
		return ctx.svrElements.serviceManager
	case core.ServerElementFactoryManager:
		return ctx.svrElements.factoryManager
	case core.ServerElementChannel:
		return ctx.svrElements.channel
	case core.ServerElementChannelManager:
		return ctx.svrElements.channelManager
	case core.ServerElementSecurityHandler:
		return ctx.svrElements.securityHandler
	case core.ServerElementMessagingManager:
		return ctx.svrElements.msgManager
	case core.ServerElementServiceResponseHandler:
		return ctx.svrElements.serviceResponseHandler
	case core.ServerElementRulesManager:
		return ctx.svrElements.rulesManager
	case core.ServerElementCacheManager:
		return ctx.svrElements.cacheManager
	case core.ServerElementTaskManager:
		return ctx.svrElements.taskManager
	case core.ServerElementModuleManager:
		return ctx.svrElements.moduleManager
	case core.ServerElementSessionManager:
		return ctx.svrElements.sessionManager
	case core.ServerElementModule:
		return ctx.svrElements.module
	case core.ServerElementLogger:
		return ctx.svrElements.logger
	case core.ServerElementCommunicator:
		return ctx.svrElements.communicator
	case core.ServerElementSecretsManager:
		return ctx.svrElements.secretsManager
	case core.ServerElementOpen1:
		return ctx.svrElements.open1
	case core.ServerElementOpen2:
		return ctx.svrElements.open2
	case core.ServerElementOpen3:
		return ctx.svrElements.open3
	}
	return nil
}

//create a child context that with the same underlying context
//changes made to context parameters will be visible on the parent
//id of the context is also retained.. this can be used to track flow
func (ctx *serverContext) SubContext(name string) core.ServerContext {
	return ctx.subContext(name, ctx.SubCtx(name).(*common.Context))
}

//create a new server context; variables in this context be reflected in parent
//sets a context element
//id of the context is not changed. flow is updated
func (ctx *serverContext) subContext(name string, parent *common.Context) *serverContext {
	log.Debug(ctx, "Entering new subcontext ", "Elapsed Time ", ctx.GetElapsedTime(), "New Context Name", name)
	return &serverContext{Context: parent, svrElements: ctx.svrElements, level: ctx.level, sessionManager: ctx.sessionManager, childContexts: ctx.childContexts}
}

//create a new server context from the parent. variables set in this context will not be reflected in parent
//id of the context is changed when new context is created
/*func (ctx *serverContext) NewContext(name string) core.ServerContext {
	return ctx.newContext(name)
}*/

func (ctx *serverContext) newContext(name string) *serverContext {
	newctx := ctx.NewCtx(name)
	log.Debug(ctx, "Entering new server context ", "Elapsed Time ", ctx.GetElapsedTime(), "Name ", name)

	svrCtx := &serverContext{Context: newctx.(*common.Context), svrElements: &contextElements{}, sessionManager: ctx.sessionManager, level: ctx.level + 1, childContexts: make([]*serverContext, 0)}
	cmap := ctx.getElementsContextMap()
	svrCtx.svrElements.properties = ctx.svrElements.properties
	svrCtx.setElementReferences(cmap, true)
	ctx.addChild(svrCtx)
	return svrCtx
}

func (ctx *serverContext) getElementsContextMap() core.ContextMap {
	return core.ContextMap{core.ServerElementServer: ctx.svrElements.server, core.ServerElementEnvironment: ctx.svrElements.environment,
		core.ServerElementLoader: ctx.svrElements.objectLoader, core.ServerElementServiceFactory: ctx.svrElements.factory, core.ServerElementApplication: ctx.svrElements.application,
		core.ServerElementService: ctx.svrElements.service, core.ServerElementServiceManager: ctx.svrElements.serviceManager, core.ServerElementFactoryManager: ctx.svrElements.factoryManager,
		core.ServerElementChannel: ctx.svrElements.channel, core.ServerElementChannelManager: ctx.svrElements.channelManager, core.ServerElementSecurityHandler: ctx.svrElements.securityHandler,
		core.ServerElementMessagingManager: ctx.svrElements.msgManager, core.ServerElementServiceResponseHandler: ctx.svrElements.serviceResponseHandler, core.ServerElementRulesManager: ctx.svrElements.rulesManager,
		core.ServerElementCacheManager: ctx.svrElements.cacheManager, core.ServerElementTaskManager: ctx.svrElements.taskManager, core.ServerElementLogger: ctx.svrElements.logger,
		core.ServerElementModuleManager: ctx.svrElements.moduleManager, core.ServerElementModule: ctx.svrElements.module, core.ServerElementCommunicator: ctx.svrElements.communicator,
		core.ServerElementSecretsManager: ctx.svrElements.secretsManager, core.ServerElementOpen1: ctx.svrElements.open1, core.ServerElementOpen2: ctx.svrElements.open2, core.ServerElementOpen3: ctx.svrElements.open3}
}

func (ctx *serverContext) addChild(child *serverContext) {
	ctx.childContexts = append(ctx.childContexts, child)
}

func (ctx *serverContext) setElements(elements core.ContextMap) {
	ctx.setElementReferences(elements, false)
}

func (ctx *serverContext) setElementReferences(svrelements core.ContextMap, ref bool) {
	ctxElems := ctx.svrElements
	for elementToSet, element := range svrelements {
		if ref && element != nil {
			element = element.Reference()
		}
		switch elementToSet {
		case core.ServerElementServer:
			if element == nil {
				ctxElems.server = nil
			} else {
				ctxElems.server = element.(elements.Server)
			}
		case core.ServerElementEnvironment:
			if element == nil {
				ctxElems.environment = nil
			} else {
				ctxElems.environment = element.(elements.Environment)
			}
		case core.ServerElementLoader:
			if element == nil {
				ctxElems.objectLoader = nil
			} else {
				ctxElems.objectLoader = element.(elements.ObjectLoader)
			}
		case core.ServerElementServiceFactory:
			if element == nil {
				ctxElems.factory = nil
			} else {
				ctxElems.factory = element.(elements.Factory)
			}
		case core.ServerElementApplication:
			if element == nil {
				ctxElems.application = nil
			} else {
				ctxElems.application = element.(elements.Application)
			}
		case core.ServerElementService:
			if element == nil {
				ctxElems.service = nil
			} else {
				ctxElems.service = element.(elements.Service)
			}
		case core.ServerElementChannelManager:
			if element == nil {
				ctxElems.channelManager = nil
			} else {
				ctxElems.channelManager = element.(elements.ChannelManager)
			}
		case core.ServerElementChannel:
			if element == nil {
				ctxElems.channel = nil
			} else {
				ctxElems.channel = element.(elements.Channel)
			}
		case core.ServerElementServiceManager:
			if element == nil {
				ctxElems.serviceManager = nil
			} else {
				ctxElems.serviceManager = element.(elements.ServiceManager)
			}
		case core.ServerElementSessionManager:
			if element == nil {
				ctxElems.sessionManager = nil
				ctx.sessionManager = nil
			} else {
				ctxElems.sessionManager = element.(elements.SessionManager)
				ctx.sessionManager = element.(*sessionManagerProxy).manager
			}
		case core.ServerElementFactoryManager:
			if element == nil {
				ctxElems.factoryManager = nil
			} else {
				ctxElems.factoryManager = element.(elements.FactoryManager)
			}
		case core.ServerElementServiceResponseHandler:
			if element == nil {
				ctxElems.serviceResponseHandler = nil
			} else {
				ctxElems.serviceResponseHandler = element.(elements.ServiceResponseHandler)
			}
		case core.ServerElementSecurityHandler:
			if element == nil {
				ctxElems.securityHandler = nil
			} else {
				ctxElems.securityHandler = element.(elements.SecurityHandler)
			}
		case core.ServerElementMessagingManager:
			if element == nil {
				ctxElems.msgManager = nil
			} else {
				ctxElems.msgManager = element.(elements.MessagingManager)
			}
		case core.ServerElementRulesManager:
			if element == nil {
				ctxElems.rulesManager = nil
			} else {
				ctxElems.rulesManager = element.(elements.RulesManager)
			}
		case core.ServerElementCacheManager:
			if element == nil {
				ctxElems.cacheManager = nil
			} else {
				ctxElems.cacheManager = element.(elements.CacheManager)
			}
		case core.ServerElementTaskManager:
			if element == nil {
				ctxElems.taskManager = nil
			} else {
				ctxElems.taskManager = element.(elements.TaskManager)
			}
		case core.ServerElementLogger:
			if element == nil {
				ctxElems.logger = nil
			} else {
				ctxElems.logger = element.(elements.Logger)
			}
		case core.ServerElementModuleManager:
			if element == nil {
				ctxElems.moduleManager = nil
			} else {
				ctxElems.moduleManager = element.(elements.ModuleManager)
			}
		case core.ServerElementModule:
			if element == nil {
				ctxElems.module = nil
			} else {
				ctxElems.module = element.(elements.Module)
			}
		case core.ServerElementCommunicator:
			if element == nil {
				ctxElems.communicator = nil
			} else {
				ctxElems.communicator = element.(elements.Communicator)
			}
		case core.ServerElementSecretsManager:
			if element == nil {
				ctxElems.secretsManager = nil
			} else {
				ctxElems.secretsManager = element.(elements.SecretsManager)
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
		c.setElements(svrelements)
	}
}

func (ctx *serverContext) GetServerProperties() map[string]interface{} {
	return ctx.svrElements.properties
}

func (ctx *serverContext) GetCodec(encoding string) (core.Codec, bool) {
	return codecs.GetCodec(encoding)
}

func (ctx *serverContext) setServerProperties(props map[string]interface{}) {
	ctx.svrElements.properties = props
}

//creates a new request with engine context
func (ctx *serverContext) CreateNewRequest(name string, engine interface{}, engineCtx interface{}, sessionId string) (core.RequestContext, error) {

	log.Info(ctx, "Creating new request ", "Name", name)
	//a service must be there in the server context if a request is to be created
	if ctx.svrElements.service == nil {
		return nil, errors.MissingService(ctx, name)
	}
	var eng elements.Engine
	if engine != nil {
		eng = engine.(elements.Engine)
	}
	reqCtx, err := ctx.createNewRequest(name, eng, engineCtx, sessionId)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	reqCtx.logger = ctx.svrElements.logger
	//svc := ctx.svrElements.service.(*serviceProxy)

	cacheToUse, ok := ctx.GetString("__cache")
	if ok {
		if ctx.svrElements.cacheManager != nil {
			cache := ctx.svrElements.cacheManager.GetCache(ctx, cacheToUse)
			reqCtx.cache = cache
		}
	}

	return reqCtx, nil
}

func (ctx *serverContext) createNewRequest(name string, engine elements.Engine, engineCtx interface{}, sessionId string) (*requestContext, error) {
	//create the request as a child of service context
	//so that the variables set by the service are available while executing a request

	newctx := ctx.NewCtx(name)

	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx, logger: ctx.svrElements.logger, engine: engine,
		engineContext: engineCtx, sessionId: sessionId, subRequest: false}, nil
}

func (ctx *serverContext) CreateCollection(objectName string, length int) (interface{}, error) {
	return ctx.svrElements.objectLoader.CreateCollection(ctx, objectName, length)
}

func (ctx *serverContext) CreateObject(objectName string) (interface{}, error) {
	return ctx.svrElements.objectLoader.CreateObject(ctx, objectName)
}

func (ctx *serverContext) GetObjectMetadata(objectName string) (core.Info, error) {
	return ctx.svrElements.objectLoader.GetMetaData(ctx, objectName)
}

func (ctx *serverContext) CreateSystemRequest(name string) core.RequestContext {
	reqCtx, err := ctx.createNewRequest(name, nil, nil, "")
	if err != nil {
		log.Error(ctx, "Error while creating system request", "Error", err)
	}
	reqCtx.user = nil
	reqCtx.admin = true
	return reqCtx
}

func (ctx *serverContext) CompleteContext() {
	log.Debug(ctx, "Context complete ", "Time elapsed ", ctx.GetElapsedTime())
}

func (ctx *serverContext) GetObjectFactory(name string) (core.ObjectFactory, bool) {
	return ctx.svrElements.objectLoader.GetObjectFactory(ctx, name)
}

func (ctx *serverContext) GetRegName(object interface{}) string {
	return ctx.svrElements.objectLoader.GetRegName(ctx, object)
}

func (ctx *serverContext) SubscribeTopic(topics []string, lstnr core.MessageListener, lsnrID string) error {
	if ctx.svrElements.msgManager != nil {
		return ctx.svrElements.msgManager.Subscribe(ctx, topics, lstnr, lsnrID)
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
	if ctx.svrElements.logger != nil {
		ctx.svrElements.logger.Trace(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogDebug(msg string, args ...interface{}) {
	if ctx.svrElements.logger != nil {
		ctx.svrElements.logger.Debug(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogInfo(msg string, args ...interface{}) {
	if ctx.svrElements.logger != nil {
		ctx.svrElements.logger.Info(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogWarn(msg string, args ...interface{}) {
	if ctx.svrElements.logger != nil {
		ctx.svrElements.logger.Warn(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogError(msg string, args ...interface{}) {
	if ctx.svrElements.logger != nil {
		ctx.svrElements.logger.Error(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}

func (ctx *serverContext) LogFatal(msg string, args ...interface{}) {
	if ctx.svrElements.logger != nil {
		ctx.svrElements.logger.Fatal(ctx, msg, args...)
	} else {
		deflog.Println(msg)
	}
}
