package core

import (
	"laatoo/sdk/auth"
	"laatoo/sdk/components"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/server/common"
)

//request context is passed while executing a request
//through all the layers of execution
type requestContext struct {
	*common.Context

	cache components.CacheComponent
	//any context from the engine that received a service request
	engineContext interface{}
	//user who is executing the request
	user auth.User
	//is the user executing a request an admin
	admin bool
	//response data for the request
	responseData *core.ServiceResponse
	parent       *requestContext
	//request body
	//this can be plain bytes, data objects or collections depending upon
	//engine configuration and service expectation
	request interface{}
	//server context that generated this request
	serverContext *serverContext
	//if the request is a subrequest, times are not reported and variables are not cleared
	subRequest bool

	logger components.Logger
}

// context of the engine that received a request
func (ctx *requestContext) EngineRequestContext() interface{} {
	return ctx.engineContext
}

//server context that generated this request
func (ctx *requestContext) ServerContext() core.ServerContext {
	return ctx.serverContext
}
func (ctx *requestContext) GetServerElement(elemType core.ServerElementType) core.ServerElement {
	return ctx.serverContext.GetServerElement(elemType)
}

//subcontext of the request
//retains id and tracks flow along with variables
func (ctx *requestContext) SubContext(name string) core.RequestContext {
	log.Info(ctx, "Entering new request subcontext ", "Name", name, "Elapsed Time ", ctx.GetElapsedTime())
	newctx := ctx.SubCtx(name)
	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx.serverContext, user: ctx.user, admin: ctx.admin, responseData: ctx.responseData, request: ctx.request,
		engineContext: ctx.engineContext, parent: ctx, cache: ctx.cache, subRequest: true}
}

//new context from the request if a part of the request needs to be tracked separately
//as a subflow.
func (ctx *requestContext) NewContext(name string) core.RequestContext {
	log.Info(ctx, "Entering new request context ", "Name", name, "Elapsed Time ", ctx.GetElapsedTime())
	newctx := ctx.NewCtx(name)
	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx.serverContext, user: ctx.user, admin: ctx.admin, responseData: ctx.responseData, request: ctx.request,
		engineContext: ctx.engineContext, parent: ctx, cache: ctx.cache, subRequest: false}
}

func (ctx *requestContext) PutInCache(bucket string, key string, item interface{}) error {
	if ctx.cache != nil {
		return ctx.cache.PutObject(ctx, bucket, key, item)
	} else {
		return nil
	}
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
	if ctx.serverContext.taskManager != nil {
		return ctx.serverContext.taskManager.PushTask(ctx, queue, task)
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
func (ctx *requestContext) SetResponse(responseData *core.ServiceResponse) {
	ctx.responseData = responseData
	if ctx.parent != nil {
		ctx.parent.SetResponse(responseData)
	}
}

//gets response
func (ctx *requestContext) GetResponse() *core.ServiceResponse {
	return ctx.responseData
}

//gets request
func (ctx *requestContext) GetRequest() interface{} {
	return ctx.request
}

//sets the request
func (ctx *requestContext) SetRequest(request interface{}) {
	ctx.request = request
}

func (ctx *requestContext) HasPermission(perm string) bool {
	if ctx.serverContext.securityHandler != nil {
		return ctx.serverContext.securityHandler.HasPermission(ctx, perm)
	}
	return false //ctx.serverContext.HasPermission(ctx, perm)
}

func (ctx *requestContext) SendSynchronousMessage(msgType string, data interface{}) error {
	if ctx.serverContext.rulesManager != nil {
		return ctx.serverContext.rulesManager.SendSynchronousMessage(ctx, msgType, data)
	}
	return nil
}

func (ctx *requestContext) PublishMessage(topic string, message interface{}) {
	if ctx.serverContext.msgManager != nil {
		go func(ctx *requestContext, topic string, message interface{}) {
			err := ctx.serverContext.msgManager.Publish(ctx, topic, message)
			if err != nil {
				log.Error(ctx, err.Error())
			}
		}(ctx, topic, message)
	}
	log.Error(ctx, "Publishing message to non existent manager")
	return
}

//completes a request
func (ctx *requestContext) CompleteRequest() {
	log.Info(ctx, "Completed Request ", "Time taken", ctx.GetElapsedTime())
	//ctx.parentContext = nil
	ctx.engineContext = nil
	//ctx.ParamsStore = nil
	ctx.user = nil
	ctx.admin = false
	ctx.responseData = nil
	ctx.request = nil
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
