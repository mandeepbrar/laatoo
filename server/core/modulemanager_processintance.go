package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
)

func (modMgr *moduleManager) loadServices(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.moduleInstances {
		mod := modProxy.(*moduleProxy).mod
		svcCtx := mod.svrContext.SubContext("Load Services")
		log.Debug(svcCtx, "Services to process", "Services", mod.services, "name", mod.name)
		if err := common.ProcessObjects(svcCtx, mod.services, processor); err != nil {
			return errors.WrapError(svcCtx, err)
		}
	}
	return nil
}

/*
processor func(core.ServerContext, config.Config, string) error
*/

func (modMgr *moduleManager) loadFactories(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.moduleInstances {
		mod := modProxy.(*moduleProxy).mod
		facCtx := mod.svrContext.SubContext("Load Factories")
		if err := common.ProcessObjects(facCtx, mod.factories, processor); err != nil {
			return errors.WrapError(facCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadChannels(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.moduleInstances {
		mod := modProxy.(*moduleProxy).mod
		chanCtx := mod.svrContext.SubContext("Load Channels")
		log.Trace(chanCtx, "Channels to process", "channels", mod.channels, "name", mod.name)
		if err := common.ProcessObjects(chanCtx, mod.channels, processor); err != nil {
			return errors.WrapError(chanCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadRules(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.moduleInstances {
		mod := modProxy.(*moduleProxy).mod
		ruleCtx := mod.svrContext.SubContext("Load Rules")
		if err := common.ProcessObjects(ruleCtx, mod.rules, processor); err != nil {
			return errors.WrapError(ruleCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadTasks(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modProxy := range modMgr.moduleInstances {
		mod := modProxy.(*moduleProxy).mod
		taskCtx := mod.svrContext.SubContext("Load Tasks")
		if err := common.ProcessObjects(taskCtx, mod.tasks, processor); err != nil {
			return errors.WrapError(taskCtx, err)
		}
	}
	return nil
}
