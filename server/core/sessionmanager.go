package core

import (
	"laatoo/sdk/server/components"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"
)

type sessionManager struct {
	name            string
	sessionsCache   components.CacheComponent
	sessionCacheSvc string
	sessionsBucket  string
	proxy           *sessionManagerProxy
}

func (sm *sessionManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	log.Trace(ctx, "Initializing Session Manager")
	var ok bool

	sm.sessionCacheSvc, _ = conf.GetString(ctx, constants.CONF_SESSIONCACHE_SVC)
	sm.sessionsBucket, ok = conf.GetString(ctx, constants.CONF_SESSIONCACHE_BUCKET)
	if !ok {
		sm.sessionsBucket = "__sessions"
	}
	return nil
}

func (sm *sessionManager) Start(ctx core.ServerContext) error {
	if sm.sessionCacheSvc != "" {
		svc, err := ctx.GetService(sm.sessionCacheSvc)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		var ok bool
		sm.sessionsCache, ok = svc.(components.CacheComponent)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", constants.CONF_SESSIONCACHE_SVC)
		}
	}
	return nil
}

func (mgr *sessionManager) newSession(ctx core.ServerContext, id string) (*session, error) {
	sess := newSession(id)
	sess.mgr = mgr
	err := mgr.sessionsCache.PutObject(ctx, mgr.sessionsBucket, sess.id, sess)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return sess, nil
}

func (mgr *sessionManager) getSession(ctx core.ServerContext, sessionId string) (*session, error) {

	obj, ok := mgr.sessionsCache.GetObject(ctx, mgr.sessionsBucket, sessionId, SESSION_OBJ)
	if !ok {
		sess, err := mgr.newSession(ctx, sessionId)
		if err != nil {
			return nil, err
		}
		return sess, nil
	} else {
		return obj.(*session), nil
	}
}

func (mgr *sessionManager) Save(ctx ctx.Context, sess *session) error {
	return mgr.sessionsCache.PutObject(ctx, mgr.sessionsBucket, sess.id, sess)
}

func (mgr *sessionManager) getUserSession(ctx core.ServerContext, userId string) (*session, error) {
	return nil, nil
}

func (mgr *sessionManager) broadcast(ctx core.ServerContext, messageFunc func(core.ServerContext, core.Session) error) error {
	return nil
}
