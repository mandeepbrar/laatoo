package main

import (
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CHECK_USER_OWN_ACCOUNT = "OwnUserAccountEnforce"
)

type OwnUserAccountEnforce struct {
	core.Service
}

func (svc *OwnUserAccountEnforce) Initialize(ctx core.ServerContext) error {
	svc.SetDescription("Own User check. Check if the user is performing operations on his own account")
	return nil
}

func Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	id, ok := ctx.GetString("id")
	if !ok {
		return core.StatusUnauthorizedResponse, nil
	}
	usr := ctx.GetUser()
	if usr == nil {
		log.Trace(ctx, "Entity not accessible by anonymous user", "user", id)
		return core.StatusUnauthorizedResponse, nil
	}
	if id != usr.GetId() {
		log.Trace(ctx, "Entity accessible only by owner", "entity", id, "user", usr.GetId())
		return core.StatusUnauthorizedResponse, nil
	}
	log.Trace(ctx, "User allowed", "entity", id, "user", usr.GetId())
	return nil, nil
}
