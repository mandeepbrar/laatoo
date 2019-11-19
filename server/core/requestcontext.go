package core

import (
	googleContext "context"
	"laatoo/sdk/server/auth"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"time"
)

//request context is passed while executing a request
//through all the layers of execution
type requestContext struct {
	*common.Context

	cache components.CacheComponent

	sessionId string

	//session to be used for request context
	session core.Session

	//any context from the engine that received a service request
	engineContext interface{}
	//engine to be used in a context
	engine elements.Engine

	//user who is executing the request
	user auth.User
	//is the user executing a request an admin
	admin bool

	parent *requestContext
	//server context that generated this request
	serverContext *serverContext
	//if the request is a subrequest, times are not reported and variables are not cleared
	subRequest bool

	responseData *core.Response

	req *request

	logger components.Logger
}

func (ctx *requestContext) Deadline() (deadline time.Time, ok bool) {
	return ctx.Context.Deadline()
}

func (ctx *requestContext) Done() <-chan struct{} {
	return ctx.Context.Done()
}

func (ctx *requestContext) Err() error {
	return ctx.Context.Err()
}

func (ctx *requestContext) Value(key interface{}) interface{} {
	return ctx.Context.Value(key)
}

func (ctx *requestContext) WithCancel() (ctx.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithCancel(ctx)
	return ctx.WithContext(newgooglectx), cancelFunc
}

func (ctx *requestContext) WithDeadline(timeout time.Time) (ctx.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithDeadline(ctx, timeout)
	return ctx.WithContext(newgooglectx), cancelFunc
}

func (ctx *requestContext) WithTimeout(timeout time.Duration) (ctx.Context, googleContext.CancelFunc) {
	newgooglectx, cancelFunc := googleContext.WithTimeout(ctx, timeout)
	return ctx.WithContext(newgooglectx), cancelFunc
}

func (ctx *requestContext) WithValue(key, val interface{}) ctx.Context {
	newgooglectx := googleContext.WithValue(ctx, key, val)
	return ctx.WithContext(newgooglectx)
}

func (ctx *requestContext) WithContext(parent googleContext.Context) ctx.Context {
	return ctx.subContext(ctx.Name, ctx.Context.WithContext(googleContext.WithValue(parent, "tId", ctx.GetId())).(*common.Context))
}

//subcontext of the request
//retains id and tracks flow along with variables
func (ctx *requestContext) SubContext(name string) core.RequestContext {
	return ctx.subContext(name, ctx.SubCtx(name).(*common.Context))
}

func (ctx *requestContext) subContext(name string, newContext *common.Context) *requestContext {
	log.Info(ctx, "Entering new request subcontext ", "Name", name, "Elapsed Time ", ctx.GetElapsedTime())
	return &requestContext{Context: newContext, serverContext: ctx.serverContext, user: ctx.user, admin: ctx.admin, req: ctx.req, sessionId: ctx.sessionId,
		engine: ctx.engine, engineContext: ctx.engineContext, parent: ctx, cache: ctx.cache, logger: ctx.logger, session: ctx.session, subRequest: true}
}

// context of the engine that received a request
func (ctx *requestContext) EngineRequestContext() interface{} {
	return ctx.engineContext
}

func (ctx *requestContext) EngineRequestParams() map[string]interface{} {
	return ctx.engine.GetRequestParams(ctx)
}

//server context that generated this request
func (ctx *requestContext) ServerContext() core.ServerContext {
	return ctx.serverContext
}
func (ctx *requestContext) GetServerElement(elemType core.ServerElementType) core.ServerElement {
	if elemType == core.ServerElementEngine {
		return ctx.engine
	}
	return ctx.serverContext.GetServerElement(elemType)
}

func (ctx *requestContext) createRequest() *request {
	return &request{Params: make(map[string]core.Param)}
}

//new context from the request if a part of the request needs to be tracked separately
//as a subflow.SESSION_OBJ
/*func (ctx *requestContext) NewContext(name string) core.RequestContext {
	log.Info(ctx, "Entering new request context ", "Name", name, "Elapsed Time ", ctx.GetElapsedTime())
	newctx := ctx.NewCtx(name)
	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx.serverContext, user: ctx.user, admin: ctx.admin, responseData: ctx.responseData, request: ctx.request,
		engineContext: ctx.engineContext, parent: ctx, cache: ctx.cache, subRequest: false}
}*/

func (ctx *requestContext) PutInCache(bucket string, key string, item interface{}) error {
	if ctx.cache != nil {
		return ctx.cache.PutObject(ctx, bucket, key, item)
	} else {
		return nil
	}
}

func (ctx *requestContext) GetCodec(encoding string) (core.Codec, bool) {
	return ctx.serverContext.GetCodec(encoding)
}

func (ctx *requestContext) PutMultiInCache(bucket string, vals map[string]interface{}) error {
	if ctx.cache != nil {
		return ctx.cache.PutObjects(ctx, bucket, vals)
	} else {
		return nil
	}
}
func (ctx *requestContext) IncrementInCache(bucket string, key string) error {
	if ctx.cache != nil {
		return ctx.cache.Increment(ctx, bucket, key)
	} else {
		return nil
	}
}
func (ctx *requestContext) DecrementInCache(bucket string, key string) error {
	if ctx.cache != nil {
		return ctx.cache.Decrement(ctx, bucket, key)
	} else {
		return nil
	}
}

func (ctx *requestContext) PushTask(queue string, task interface{}) error {
	if ctx.serverContext.svrElements.taskManager != nil {
		return ctx.serverContext.svrElements.taskManager.PushTask(ctx, queue, task)
	}
	log.Error(ctx, "No task manager", "queue", queue)
	return nil
}

func (ctx *requestContext) GetFromCache(bucket string, key string) (interface{}, bool) {
	if ctx.cache != nil {
		return ctx.cache.Get(ctx, bucket, key)
	} else {
		return nil, false
	}
}

func (ctx *requestContext) GetMultiFromCache(bucket string, keys []string) map[string]interface{} {
	if ctx.cache != nil {
		return ctx.cache.GetMulti(ctx, bucket, keys)
	}
	return map[string]interface{}{}
}

func (ctx *requestContext) GetObjectFromCache(bucket string, key string, objectType string) (interface{}, bool) {
	if ctx.cache != nil {
		return ctx.cache.GetObject(ctx, bucket, key, objectType)
	} else {
		return nil, false
	}
}

func (ctx *requestContext) GetObjectsFromCache(bucket string, keys []string, objectType string) map[string]interface{} {
	if ctx.cache != nil {
		return ctx.cache.GetObjects(ctx, bucket, keys, objectType)
	}
	return map[string]interface{}{}
}

func (ctx *requestContext) InvalidateCache(bucket string, key string) error {
	if ctx.cache != nil {
		return ctx.cache.Delete(ctx, bucket, key)
	} else {
		return nil
	}
}

//returns user executing a request
func (ctx *requestContext) GetUser() auth.User {
	return ctx.user
}

//if the person is an admin
func (ctx *requestContext) IsAdmin() bool {
	return ctx.admin
}

//sets or gets the response for a request
func (ctx *requestContext) SetResponse(responseData *core.Response) {
	ctx.responseData = responseData
	if ctx.parent != nil {
		ctx.parent.SetResponse(responseData)
	}
}

//gets response
func (ctx *requestContext) GetResponse() *core.Response {
	return ctx.responseData
}

func (ctx *requestContext) CreateCollection(objectName string, length int) (interface{}, error) {
	return ctx.serverContext.CreateCollection(objectName, length)
}
func (ctx *requestContext) CreateObject(objectName string) (interface{}, error) {
	return ctx.serverContext.CreateObject(objectName)
}

func (ctx *requestContext) GetObjectFactory(name string) (core.ObjectFactory, bool) {
	return ctx.serverContext.GetObjectFactory(name)
}

func (ctx *requestContext) GetRegName(object interface{}) string {
	return ctx.serverContext.GetRegName(object)
}

func (ctx *requestContext) GetRequest() core.Request {
	return ctx.req
}

/*func (ctx *requestContext) GetBody() interface{} {
	return ctx.req.GetBody()
}*/

func (ctx *requestContext) GetSession() core.Session {
	if ctx.session == nil {
		var session core.Session
		var err error

		smgr := ctx.serverContext.sessionManager
		if smgr != nil {
			session, err = smgr.getSession(ctx.serverContext, ctx.sessionId)
			if err != nil {
				log.Error(ctx, "Error while retrieving session", "err", err)
			}
		} else {
			session = newSession(ctx.sessionId)
		}
		ctx.session = session
	}
	return ctx.session
}

func (ctx *requestContext) GetParam(name string) (core.Param, bool) {
	return ctx.req.GetParam(ctx, name)
}

func (ctx *requestContext) GetParamValue(name string) (interface{}, bool) {
	return ctx.req.GetParamValue(ctx, name)
}

func (ctx *requestContext) GetParams() map[string]core.Param {
	return ctx.req.GetParams(ctx)
}

func (ctx *requestContext) GetIntParam(name string) (int, bool) {
	return ctx.req.GetIntParam(ctx, name)
}

func (ctx *requestContext) GetStringParam(name string) (string, bool) {
	return ctx.req.GetStringParam(ctx, name)
}

func (ctx *requestContext) GetStringMapParam(name string) (map[string]interface{}, bool) {
	return ctx.req.GetStringMapParam(ctx, name)
}

func (ctx *requestContext) GetStringsMapParam(name string) (map[string]string, bool) {
	return ctx.req.GetStringsMapParam(ctx, name)
}

func (ctx *requestContext) Forward(alias string, vals map[string]interface{}) error {
	svc, err := ctx.serverContext.getService(alias)
	if err != nil {
		return err
	}
	res, err := svc.HandleRequest(ctx.SubContext(alias), vals)
	ctx.SetResponse(res)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *requestContext) CompleteContext() {
	log.Debug(ctx, "Context complete ", "Time elapsed ", ctx.GetElapsedTime())
}

func (ctx *requestContext) SendCommunication(communication interface{}) error {
	comm, ok := communication.(*components.Communication)
	if ok {
		if ctx.serverContext.svrElements.communicator != nil {
			return ctx.serverContext.svrElements.communicator.SendCommunication(ctx, comm)
		}
	}
	return errors.BadConf(ctx, "No communicator service has been configured or bad communication structure")
}

func (ctx *requestContext) ForwardToService(svc core.Service, vals map[string]interface{}) error {
	return ctx.Forward(svc.GetName(), vals)
}

func (ctx *requestContext) HasPermission(perm string) bool {
	if ctx.serverContext.svrElements.securityHandler != nil {
		return ctx.serverContext.svrElements.securityHandler.HasPermission(ctx, perm)
	}
	return false //ctx.serverContext.HasPermission(ctx, perm)
}

func (ctx *requestContext) SendSynchronousMessage(msgType string, data interface{}) error {
	if ctx.serverContext.svrElements.rulesManager != nil {
		return ctx.serverContext.svrElements.rulesManager.SendSynchronousMessage(ctx, msgType, data)
	}
	return nil
}

func (ctx *requestContext) PublishMessage(topic string, message interface{}) {
	if ctx.serverContext.svrElements.msgManager != nil {
		go func(ctx *requestContext, topic string, message interface{}) {
			err := ctx.serverContext.svrElements.msgManager.Publish(ctx, topic, message)
			if err != nil {
				log.Error(ctx, err.Error())
			}
		}(ctx, topic, message)
	}
	log.Error(ctx, "Publishing message to non existent manager")
	return
}

func (ctx *requestContext) setEngine(engine elements.Engine) {
	ctx.engine = engine
}

//completes a request
func (ctx *requestContext) CompleteRequest() {
	log.Info(ctx, "Completed Request ", "Time taken", ctx.GetElapsedTime())
	//ctx.parentContext = nil
	ctx.engineContext = nil
	//ctx.ParamsStore = nil
	ctx.user = nil
	ctx.admin = false
	ctx.serverContext = nil
}

func (ctx *requestContext) LogTrace(msg string, args ...interface{}) {
	ctx.logger.Trace(ctx, msg, args...)
}

func (ctx *requestContext) LogDebug(msg string, args ...interface{}) {
	ctx.logger.Debug(ctx, msg, args...)
}

func (ctx *requestContext) LogInfo(msg string, args ...interface{}) {
	ctx.logger.Info(ctx, msg, args...)
}

func (ctx *requestContext) LogWarn(msg string, args ...interface{}) {
	ctx.logger.Warn(ctx, msg, args...)
}

func (ctx *requestContext) LogError(msg string, args ...interface{}) {
	ctx.logger.Error(ctx, msg, args...)
}

func (ctx *requestContext) LogFatal(msg string, args ...interface{}) {
	ctx.logger.Fatal(ctx, msg, args...)
}
