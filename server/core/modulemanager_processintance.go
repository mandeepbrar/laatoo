package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
)

func (modMgr *moduleManager) loadServices(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modInstance := range modMgr.moduleInstances {
		svcCtx := modInstance.svrContext.SubContext("Load Services")
		log.Debug(svcCtx, "Services to process", "Services", modInstance.services, "name", modInstance.name)
		if err := common.ProcessObjects(svcCtx, modInstance.services, processor); err != nil {
			return errors.WrapError(svcCtx, err)
		}
	}
	return nil
}

/*
processor func(core.ServerContext, config.Config, string) error
*/

func (modMgr *moduleManager) loadFactories(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modInstance := range modMgr.moduleInstances {
		facCtx := modInstance.svrContext.SubContext("Load Factories")
		if err := common.ProcessObjects(facCtx, modInstance.factories, processor); err != nil {
			return errors.WrapError(facCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadChannels(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modInstance := range modMgr.moduleInstances {
		chanCtx := modInstance.svrContext.SubContext("Load Channels")
		log.Trace(chanCtx, "Channels to process", "channels", modInstance.channels, "name", modInstance.name)
		if err := common.ProcessObjects(chanCtx, modInstance.channels, processor); err != nil {
			return errors.WrapError(chanCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadRules(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modInstance := range modMgr.moduleInstances {
		ruleCtx := modInstance.svrContext.SubContext("Load Rules")
		if err := common.ProcessObjects(ruleCtx, modInstance.rules, processor); err != nil {
			return errors.WrapError(ruleCtx, err)
		}
	}
	return nil
}

func (modMgr *moduleManager) loadTasks(ctx core.ServerContext, processor func(core.ServerContext, config.Config, string) error) error {
	for _, modInstance := range modMgr.moduleInstances {
		taskCtx := modInstance.svrContext.SubContext("Load Tasks")
		if err := common.ProcessObjects(taskCtx, modInstance.tasks, processor); err != nil {
			return errors.WrapError(taskCtx, err)
		}
	}
	return nil
}
