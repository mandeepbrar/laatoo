package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/constants"
)

type secretsManager struct {
	name                 string
	secretsMgr           components.SecretsManager
	parentSecretsManager *secretsManager
	proxy                *secretsManagerProxy
	svrContext           core.ServerContext
}

func (sm *secretsManager) Initialize(ctx core.ServerContext, conf config.Config) error {

	svcName, ok := conf.GetString(ctx, constants.CONF_SECRETSVC)

	if ok {
		svcObj, err := ctx.GetService(svcName)
		if err != nil {
			return errors.BadConf(ctx, constants.CONF_SECRETSVC, "error", err)
		}
		sm.secretsMgr, ok = svcObj.(components.SecretsManager)
		if !ok {
			return errors.BadConf(ctx, constants.CONF_SECRETSVC)
		}
	} else {
		secretsMgr, err := defaultSecretsManager(ctx)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		sm.secretsMgr = secretsMgr
	}

	log.Trace(ctx, "Processed Secrets Manager")

	return nil
}

func (sm *secretsManager) get(ctx core.ServerContext, key string) ([]byte, bool) {
	if sm.secretsMgr != nil {
		val, ok := sm.secretsMgr.Get(ctx, key)
		if ok {
			return val, ok
		}
		if sm.parentSecretsManager != nil {
			val, ok := sm.parentSecretsManager.get(ctx, key)
			if ok {
				return val, ok
			}
		}
	}
	return nil, false
}

func (sm *secretsManager) put(ctx core.ServerContext, key string, val []byte) error {
	if sm.secretsMgr != nil {
		err := sm.secretsMgr.Put(ctx, key, val)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return nil
}

func (sm *secretsManager) Start(ctx core.ServerContext) error {
	return nil
}
