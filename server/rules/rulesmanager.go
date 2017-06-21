package rules

import (
	"laatoo/sdk/components/rules"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/common"
)

type rulesManager struct {
	registeredRules map[string][]rules.Rule
	proxy           *rulesManagerProxy
}

func (rm *rulesManager) Initialize(ctx core.ServerContext, conf config.Config) error {

	rulesConf, err, ok := common.ConfigFileAdapter(ctx, conf, config.CONF_RULES)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if ok {
		ruleMgrCtx := rm.createContext(ctx, "Rules Manager")
		log.Logger.Debug(ruleMgrCtx, "Initializing rules manager")
		ruleNames := rulesConf.AllConfigurations()
		for _, ruleName := range ruleNames {
			ruleCtx := ruleMgrCtx.SubContext("Creating rule" + ruleName)
			log.Logger.Debug(ruleCtx, "Creating rule", "Name", ruleName)
			ruleConf, err, _ := common.ConfigFileAdapter(ctx, rulesConf, ruleName)
			if err != nil {
				return errors.WrapError(ruleCtx, err)
			}
			if err := rm.processRuleConf(ruleCtx, ruleConf, ruleName); err != nil {
				return err
			}
		}
	}

	if err := common.ProcessDirectoryFiles(ctx, config.CONF_RULES, rm.processRuleConf); err != nil {
		return err
	}

	return nil
}
func (rm *rulesManager) Start(ctx core.ServerContext) error {
	return nil
}

func (rm *rulesManager) processRuleConf(ruleCtx core.ServerContext, ruleConf config.Config, ruleName string) error {
	triggerType, ok := ruleConf.GetString(config.CONF_RULE_TRIGGER)
	if !ok {
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_TRIGGER)
	}
	ruleobj, ok := ruleConf.GetString(config.CONF_RULE_OBJECT)
	if !ok {
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_OBJECT)
	}
	obj, err := ruleCtx.CreateObject(ruleobj)
	if err != nil {
		return errors.WrapError(ruleCtx, err)
	}
	init := obj.(core.Initializable)
	err = init.Init(ruleCtx, map[string]interface{}{"conf": ruleConf})
	if err != nil {
		return errors.WrapError(ruleCtx, err)
	}
	rule, ok := obj.(rules.Rule)
	if !ok {
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_BAD_CONF, "Conf", config.CONF_RULE_OBJECT)
	}
	switch triggerType {
	case config.CONF_RULE_TRIGGER_ASYNC:
		msgType, ok := ruleConf.GetString(config.CONF_RULE_MSGTYPE)
		if !ok {
			return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_MSGTYPE)
		}
		ruleMethod := func(rule rules.Rule, msgType string) core.ServiceFunc {
			return func(msgctx core.RequestContext) error {
				go func() {
					tr := &rules.Trigger{MessageType: msgType, TriggerType: rules.AsynchronousMessage, Message: msgctx.GetRequest()}
					if rule.Condition(msgctx, tr) {
						err := rule.Action(msgctx, tr)
						if err != nil {
							log.Logger.Error(msgctx, err.Error())
						}
					}

				}()
				return nil
			}
		}
		ruleCtx.SubscribeTopic([]string{msgType}, ruleMethod(rule, msgType))
	case config.CONF_RULE_TRIGGER_SYNC:
		msgType, ok := ruleConf.GetString(config.CONF_RULE_MSGTYPE)
		if !ok {
			return errors.ThrowError(ruleCtx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_MSGTYPE)
		}
		rm.subscribeSynchronousMessage(ruleCtx, msgType, rule)
	default:
		return errors.ThrowError(ruleCtx, errors.CORE_ERROR_BAD_CONF, "Conf", config.CONF_RULE_TRIGGER)
	}
	return nil
}

func (rm *rulesManager) subscribeSynchronousMessage(ctx core.ServerContext, msgType string, rule rules.Rule) {
	regrules, prs := rm.registeredRules[msgType]
	if !prs {
		regrules = []rules.Rule{}
	}
	regrules = append(regrules, rule)
	rm.registeredRules[msgType] = regrules
}

func (rm *rulesManager) sendSynchronousMessage(ctx core.RequestContext, msgType string, data interface{}) error {
	tr := &rules.Trigger{MessageType: msgType, TriggerType: rules.SynchronousMessage, Message: data}
	regrules, present := rm.registeredRules[msgType]
	if present {
		for _, rule := range regrules {
			log.Logger.Error(ctx, "Executing rule")
			if rule.Condition(ctx, tr) {
				log.Logger.Error(ctx, "Executing rule", "message", msgType)
				err := rule.Action(ctx, tr)
				log.Logger.Error(ctx, "Executed rule", "message", msgType)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (rm *rulesManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementRulesManager: rm.proxy}, core.ServerElementRulesManager)
}
