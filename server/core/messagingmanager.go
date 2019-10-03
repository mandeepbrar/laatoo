package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/server/common"
	"laatoo/server/constants"
)

type messagingManager struct {
	name        string
	commSvcName string
	commSvc     components.PubSubComponent
	proxy       elements.MessagingManager
	topicStore  map[string][]*messagingListenerHolder
	svrContext  core.ServerContext
}

type messagingListenerHolder struct {
	name string
	lsnr core.MessageListener
}

func (msgMgr *messagingManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	msgmgrInitializeCtx := ctx.SubContext("Initialize message manager")
	log.Trace(msgmgrInitializeCtx, "Create Message Topics")
	err := msgMgr.createTopics(msgmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (msgMgr *messagingManager) Start(ctx core.ServerContext) error {
	commSvc, err := ctx.GetService(msgMgr.commSvcName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	pubsub, ok := commSvc.(components.PubSubComponent)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF)
	}
	msgMgr.commSvc = pubsub
	return msgMgr.subscribeTopics(ctx)
}

func (msgMgr *messagingManager) createTopics(ctx core.ServerContext, conf config.Config) error {
	topicsConf, ok := conf.GetSubConfig(ctx, constants.CONF_MESSAGE_TOPICS)
	if ok {
		topicNames := topicsConf.AllConfigurations(ctx)
		for _, topicName := range topicNames {
			topicConf, err, _ := common.ConfigFileAdapter(ctx, topicsConf, topicName)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			err = msgMgr.createTopic(ctx, topicName, topicConf)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
		}
	}
	return nil
}

func (msgMgr *messagingManager) createTopic(ctx core.ServerContext, name string, conf config.Config) error {
	msgMgr.topicStore[name] = []*messagingListenerHolder{}
	return nil
}

//subscribe to a topic
func (mgr *messagingManager) subscribeTopic(ctx core.ServerContext, topics []string, handler core.MessageListener, lsnrId string) error {
	for _, topic := range topics {
		listeners, prs := mgr.topicStore[topic]
		if !prs {
			log.Error(ctx, "Topic not allowed for Subscription", "Topic", topic)
			return nil
		}
		lsnrHolder := &messagingListenerHolder{lsnr: handler, name: lsnrId}
		mgr.topicStore[topic] = append(listeners, lsnrHolder)
		log.Trace(ctx, "Subscribed topic", "topic", topic)
	}
	return nil
}

//unsubscribe to a topic
func (mgr *messagingManager) unsubscribeTopic(ctx core.ServerContext, topics []string, lsnrId string) error {
	for _, topic := range topics {
		listeners, prs := mgr.topicStore[topic]
		if !prs {
			for idx, lsnr := range listeners {
				if lsnr.name == lsnrId {
					listeners[idx] = nil
				}
			}
		}
	}
	log.Trace(ctx, "Unsubscribed topics", "topics", topics)
	return nil
}

//publish message using
func (mgr *messagingManager) publishMessage(ctx core.RequestContext, topic string, message interface{}) error {
	_, ok := mgr.topicStore[topic]
	if !ok {
		log.Error(ctx, "Topic not allowed for Publishing", "Topic", topic)
		return nil
	}
	if mgr.commSvc != nil {
		log.Trace(ctx, "posting message")
		return mgr.commSvc.Publish(ctx, topic, message)
	}
	return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Name", "Messaging Manager")
}

func (mgr *messagingManager) subscribeTopics(ctx core.ServerContext) error {
	if mgr.commSvc != nil {
		topics := make([]string, len(mgr.topicStore))
		i := 0
		for k := range mgr.topicStore {
			topics[i] = k
			i++
		}
		log.Trace(ctx, "Subscribing topics", "topics", topics)
		mgr.commSvc.Subscribe(ctx, topics, func(reqctx core.RequestContext, data interface{}, params map[string]interface{}) error {
			topic, ok := reqctx.GetString("messagetype")
			if ok {
				lsnrs := mgr.topicStore[topic]
				for _, val := range lsnrs {
					go val.lsnr(reqctx, data, params)
				}
			}
			return nil
		})
		/*if err != nil {
			return err
		}*/
	}
	return nil
}
