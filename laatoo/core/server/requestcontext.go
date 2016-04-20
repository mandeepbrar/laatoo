package server

import (
	"laatoo/core/common"
	"laatoo/sdk/auth"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"time"
)

//request context is passed while executing a request
//through all the layers of execution
type requestContext struct {
	*common.Context
	//any context from the engine that received a service request
	engineContext interface{}
	//user who is executing the request
	user auth.User
	//is the user executing a request an admin
	admin bool
	//response data for the request
	responseData *core.ServiceResponse
	//request body
	//this can be plain bytes, data objects or collections depending upon
	//engine configuration and service expectation
	requestBody interface{}
	//server context that generated this request
	serverContext *serverContext
	//time at which the request was created
	createTime time.Time
	//if the request is a subrequest, times are not reported and variables are not cleared
	subRequest bool
}

// context of the engine that received a request
func (ctx *requestContext) EngineRequestContext() interface{} {
	return ctx.engineContext
}

//server context that generated this request
func (ctx *requestContext) ServerContext() core.ServerContext {
	return ctx.serverContext
}

//subcontext of the request
//retains id and tracks flow along with variables
func (ctx *requestContext) SubContext(name string) core.RequestContext {
	newctx := ctx.SubCtx(name)
	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx.serverContext, user: ctx.user, admin: ctx.admin, responseData: ctx.responseData, requestBody: ctx.requestBody,
		engineContext: ctx.engineContext, createTime: time.Now(), subRequest: true}
}

//new context from the request if a part of the request needs to be tracked separately
//as a subflow.
func (ctx *requestContext) NewContext(name string) core.RequestContext {
	newctx := ctx.NewCtx(name)
	return &requestContext{Context: newctx.(*common.Context), serverContext: ctx.serverContext, user: ctx.user, admin: ctx.admin, responseData: ctx.responseData, requestBody: ctx.requestBody,
		engineContext: ctx.engineContext, createTime: time.Now(), subRequest: false}
}

func (ctx *requestContext) PublishMessage(topic string, message interface{}) error {
	return nil //ctx.serverContext.PublishMessage(topic, message)
}

func (ctx *requestContext) PutInCache(key string, item interface{}) error {
	return nil //ctx.serverContext.PutInCache(key, item)
}

func (ctx *requestContext) GetFromCache(key string, val interface{}) bool {
	return false //ctx.serverContext.GetFromCache(key, val)
}

func (ctx *requestContext) GetMultiFromCache(keys []string, val map[string]interface{}) bool {
	return false //ctx.serverContext.GetMultiFromCache(keys, val)
}

func (ctx *requestContext) InvalidateCache(key string) error {
	return nil //ctx.serverContext.DeleteFromCache(key)
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
}

//gets response
func (ctx *requestContext) GetResponse() *core.ServiceResponse {
	return ctx.responseData
}

//gets request
func (ctx *requestContext) GetRequest() interface{} {
	return ctx.requestBody
}

//sets the request
func (ctx *requestContext) SetRequest(requestBody interface{}) {
	ctx.requestBody = requestBody
}

func (ctx *requestContext) HasPermission(perm string) bool {
	return false //ctx.serverContext.HasPermission(ctx, perm)
}
func (ctx *requestContext) GetRolePermissions(role []string) ([]string, bool) {
	return nil, false //ctx.serverContext.GetRolePermissions(role)
}

//completes a request
func (ctx *requestContext) CompleteRequest() {
	if !ctx.subRequest {
		completionTime := time.Now()
		elapsedTime := completionTime.Sub(ctx.createTime)
		log.Logger.Info(ctx, "Request Complete", "Time taken", elapsedTime)
	}
	//ctx.parentContext = nil
	ctx.engineContext = nil
	//ctx.ParamsStore = nil
	ctx.user = nil
	ctx.admin = false
	ctx.responseData = nil
	ctx.requestBody = nil
	ctx.serverContext = nil
}
