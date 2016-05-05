package rules

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/rules"
)

type rulesManager struct {
	registeredRules map[string][]rules.Rule
	proxy           *rulesManagerProxy
}

func (rm *rulesManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	ruleNames := conf.AllConfigurations()
	for _, ruleName := range ruleNames {
		ruleConf, err, _ := config.ConfigFileAdapter(conf, ruleName)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		triggerType, ok := ruleConf.GetString(config.CONF_RULE_TRIGGER)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_TRIGGER)
		}
		ruleobj, ok := ruleConf.GetString(config.CONF_RULE_OBJECT)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_OBJECT)
		}
		obj, err := ctx.CreateObject(ruleobj, nil)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		rule, ok := obj.(rules.Rule)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", config.CONF_RULE_OBJECT)
		}
		switch triggerType {
		case config.CONF_RULE_TRIGGER_MSG:
			topic, ok := conf.GetString(config.CONF_RULE_MESSAGETOPIC)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_MESSAGETOPIC)
			}
			ruleMethod := func(rule rules.Rule) core.TopicListener {
				return func(msgctx core.RequestContext, topic string, message interface{}) {
					tr := &rules.Trigger{Event: topic, TriggerType: rules.Message, Data: map[string]interface{}{"message": message}}
					if rule.Condition(msgctx, tr) {
						err := rule.Action(msgctx, tr)
						if err != nil {
							log.Logger.Error(msgctx, err.Error())
						}
					}
				}
			}
			ctx.SubscribeTopic([]string{topic}, ruleMethod(rule))
		case config.CONF_RULE_TRIGGER_EVENT:
			eventType, ok := conf.GetString(config.CONF_RULE_EVENTTYPE)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_EVENTTYPE)
			}
			eventObject, ok := conf.GetString(config.CONF_RULE_EVENTOBJECT)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_RULE_EVENTOBJECT)
			}
			rm.subscribeEvent(ctx, eventType, eventObject, rule)
		default:
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", config.CONF_RULE_TRIGGER)
		}
	}
	return nil
}
func (rm *rulesManager) Start(ctx core.ServerContext) error {
	return nil
}

func (rm *rulesManager) subscribeEvent(ctx core.ServerContext, eventType string, eventObject string, rule rules.Rule) {
	key := fmt.Sprintf("%s#%s", eventType, eventObject)
	regrules, prs := rm.registeredRules[key]
	if !prs {
		regrules = []rules.Rule{}
	}
	regrules = append(regrules, rule)
	rm.registeredRules[key] = regrules
}

func (rm *rulesManager) fireEvent(ctx core.RequestContext, eventType string, eventObject string, data map[string]interface{}) {
	tr := &rules.Trigger{Event: eventType, EventObject: eventObject, TriggerType: rules.Event, Data: data}
	key := fmt.Sprintf("%s#%s", eventType, eventObject)
	regrules, present := rm.registeredRules[key]
	if present {
		for _, rule := range regrules {
			go func(rule rules.Rule) {
				if rule.Condition(ctx, tr) {
					err := rule.Action(ctx, tr)
					if err != nil {
						log.Logger.Error(ctx, err.Error())
					}
				}
			}(rule)
		}
	}
	return
}
