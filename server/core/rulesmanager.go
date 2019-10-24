package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components/rules"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

type rulesManager struct {
	name            string
	registeredRules map[string]map[string]rules.Rule
	rulesStore      map[string]rules.Rule
	proxy           *rulesManagerProxy
	svrContext      core.ServerContext
}

func (rm *rulesManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	ruleMgrCtx := ctx.SubContext("Rules Manager")
	rulesConf, err, ok := common.ConfigFileAdapter(ruleMgrCtx, conf, constants.CONF_RULES)
	if err != nil {
		return errors.WrapError(ruleMgrCtx, err)
	}

	modManager := ctx.GetServerElement(core.ServerElementModuleManager).(*moduleManagerProxy).modMgr

	if err = modManager.loadRules(ctx, rm.processRuleConf); err != nil {
		return err
	}

	if ok {
		log.Debug(ruleMgrCtx, "Initializing rules manager")
		ruleNames := rulesConf.AllConfigurations(ruleMgrCtx)
		for _, ruleName := range ruleNames {
			ruleCtx := ruleMgrCtx.SubContext("Creating rule" + ruleName)
			log.Debug(ruleCtx, "Creating rule", "Name", ruleName)
			ruleConf, err, _ := common.ConfigFileAdapter(ruleCtx, rulesConf, ruleName)
			if err != nil {
				return errors.WrapError(ruleCtx, err)
			}
			if err := rm.processRuleConf(ruleCtx, ruleConf, ruleName); err != nil {
				return err
			}
		}
	}

	baseDir, _ := ctx.GetString(config.BASEDIR)
	return rm.processRulesFromFolder(ruleMgrCtx, baseDir)
}

func (rm *rulesManager) Start(ctx core.ServerContext) error {
	return nil
}

func (rm *rulesManager) processRulesFromFolder(ctx core.ServerContext, folder string) error {
	objs, err := rm.loadRulesFromDirectory(ctx, folder)
	if err != nil {
		return err
	}

	if err = common.ProcessObjects(ctx, objs, rm.processRuleConf); err != nil {
		return err
	}
	return nil
}

func (rm *rulesManager) loadRulesFromDirectory(ctx core.ServerContext, folder string) (map[string]config.Config, error) {
	return common.ProcessDirectoryFiles(ctx, folder, constants.CONF_RULES, true)
}

func (rm *rulesManager) processRuleConf(ruleCtx core.ServerContext, ruleConf config.Config, ruleName string) error {
	triggerType, ok := ruleConf.GetString(ruleCtx, constants.CONF_RULE_TRIGGER)
	if !ok {
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", constants.CONF_RULE_TRIGGER)
	}
	ruleobj, ok := ruleConf.GetString(ruleCtx, constants.CONF_RULE_OBJECT)
	if !ok {
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", constants.CONF_RULE_OBJECT)
	}
	obj, err := ruleCtx.CreateObject(ruleobj)
	if err != nil {
		return errors.WrapError(ruleCtx, err)
	}
	init := obj.(core.Initializable)
	err = init.Initialize(ruleCtx, ruleConf)
	if err != nil {
		return errors.WrapError(ruleCtx, err)
	}
	rule, ok := obj.(rules.Rule)
	if !ok {
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_BAD_CONF, "Conf", constants.CONF_RULE_OBJECT)
	}
	rm.rulesStore[ruleName] = rule
	switch triggerType {
	case constants.CONF_RULE_TRIGGER_ASYNC:
		msgType, ok := ruleConf.GetString(ruleCtx, constants.CONF_RULE_MSGTYPE)
		if !ok {
			return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", constants.CONF_RULE_MSGTYPE)
		}
		lsnr := func(rule rules.Rule, msgType string) core.MessageListener {
			return func(msgctx core.RequestContext, data interface{}, info map[string]interface{}) error {
				go func() {
					tr := &rules.Trigger{MessageType: msgType, TriggerType: rules.AsynchronousMessage, Message: data}
					if rule.Condition(msgctx, tr) {
						err := rule.Action(msgctx, tr)
						if err != nil {
							log.Error(msgctx, err.Error())
						}
					}

				}()
				return nil
			}
		}(rule, msgType)
		ruleCtx.SubscribeTopic([]string{msgType}, lsnr, ruleName)
	case constants.CONF_RULE_TRIGGER_SYNC:
		msgType, ok := ruleConf.GetString(ruleCtx, constants.CONF_RULE_MSGTYPE)
		if !ok {
			return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", constants.CONF_RULE_MSGTYPE)
		}
		rm.subscribeSynchronousMessage(ruleCtx, msgType, rule, ruleName)
	default:
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_BAD_CONF, "Conf", constants.CONF_RULE_TRIGGER)
	}
	return nil
}

func (rm *rulesManager) subscribeSynchronousMessage(ctx core.ServerContext, msgType string, rule rules.Rule, ruleName string) {
	regrules, prs := rm.registeredRules[msgType]
	if !prs {
		regrules = make(map[string]rules.Rule)
	}
	regrules[ruleName] = rule
	rm.registeredRules[msgType] = regrules
}

func (rm *rulesManager) sendSynchronousMessage(ctx core.RequestContext, msgType string, data interface{}) error {
	log.Error(ctx, "sending synchronous message", "msgtype", msgType)
	tr := &rules.Trigger{MessageType: msgType, TriggerType: rules.SynchronousMessage, Message: data}
	regrules, present := rm.registeredRules[msgType]
	if present {
		for ruleName, rule := range regrules {
			log.Error(ctx, "Executing rule. Checking condition", "name", ruleName, "rule", rule)
			if rule.Condition(ctx, tr) {
				log.Debug(ctx, "Executing rule", "message", msgType)
				err := rule.Action(ctx, tr)
				log.Debug(ctx, "Executed rule", "message", msgType)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (rm *rulesManager) unloadModuleRules(ctx core.ServerContext, mod *serverModule) error {
	ctx = ctx.SubContext("unload rules")
	if err := common.ProcessObjects(ctx, mod.rules, rm.unloadRule); err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (rm *rulesManager) unloadRule(ctx core.ServerContext, ruleConf config.Config, ruleName string) error {
	ctx = ctx.SubContext("unload rule" + ruleName)
	unloadCtx := ctx.SubContext("Unload rule")
	triggerType, _ := ruleConf.GetString(unloadCtx, constants.CONF_RULE_TRIGGER)
	msgType, _ := ruleConf.GetString(ctx, constants.CONF_RULE_MSGTYPE)
	switch triggerType {
	case constants.CONF_RULE_TRIGGER_ASYNC:
		msgManager := ctx.GetServerElement(core.ServerElementMessagingManager).(*messagingManagerProxy).manager
		msgManager.unsubscribeTopic(ctx, []string{msgType}, ruleName)
	case constants.CONF_RULE_TRIGGER_SYNC:
		rule, ok := rm.rulesStore[ruleName]
		if ok {
			regrules, prs := rm.registeredRules[msgType]
			if prs {
				for idx, existingrule := range regrules {
					if existingrule == rule {
						delete(regrules, idx)
					}
				}
			}
		}
	}
	return nil
}
