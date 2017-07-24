package main

import (
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	CHECK_CREATION_USER = "CheckCreationUser"
)

type CheckCreationUser struct {
	core.Service
}

func (svc *CheckCreationUser) Initialize(ctx core.ServerContext) error {
	svc.SetDescription("Creation user check middleware service. Checks if the id that created the entity is the same as the logged in user")
	return nil
}

func (svc *CheckCreationUser) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
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
