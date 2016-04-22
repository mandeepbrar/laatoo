package server

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/services"
)

type messagingManager struct {
	commSvcName string
	commSvc     services.PubSub
	parent      core.ServerElement
	proxy       server.MessagingManager
	topicStore  map[string][]core.TopicListener
}

func (msgMgr *messagingManager) Initialize(ctx core.ServerContext, conf config.Config) error {
	msgmgrInitializeCtx := msgMgr.createContext(ctx, "Initialize message manager")
	log.Logger.Trace(msgmgrInitializeCtx, "Create Message Topics")
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
	pubsub, ok := commSvc.(services.PubSub)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF)
	}
	msgMgr.commSvc = pubsub
	return msgMgr.subscribeTopics(ctx)
}

func (msgMgr *messagingManager) createTopics(ctx core.ServerContext, conf config.Config) error {
	topicsConf, ok := conf.GetSubConfig(config.CONF_MESSAGE_TOPICS)
	if ok {
		topicNames := topicsConf.AllConfigurations()
		for _, topicName := range topicNames {
			topicConf, err, _ := config.ConfigFileAdapter(topicsConf, topicName)
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
	msgMgr.topicStore[name] = []core.TopicListener{}
	return nil
}

//subscribe to a topic
func (mgr *messagingManager) subscribeTopic(ctx core.ServerContext, topics []string, handler core.TopicListener) error {
	for _, topic := range topics {
		listeners, prs := mgr.topicStore[topic]
		if !prs {
			log.Logger.Error(ctx, "Topic not allowed for Subscription", "Topic", topic)
			return nil
		}
		mgr.topicStore[topic] = append(listeners, handler)
		log.Logger.Trace(ctx, "Subscribed topic", "topic", topic)
	}
	return nil
}

//publish message using
func (mgr *messagingManager) publishMessage(ctx core.RequestContext, topic string, message interface{}) error {
	_, ok := mgr.topicStore[topic]
	if !ok {
		log.Logger.Error(ctx, "Topic not allowed for Publishing", "Topic", topic)
		return nil
	}
	if mgr.commSvc != nil {
		log.Logger.Trace(ctx, "posting message")
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
		log.Logger.Trace(ctx, "Subscribing topics", "topics", topics)
		mgr.commSvc.Subscribe(ctx, topics, func(reqctx core.RequestContext, topic string, message interface{}) {
			lsnrs := mgr.topicStore[topic]
			for _, val := range lsnrs {
				go val(reqctx, topic, message)
			}
		})
		/*if err != nil {
			return err
		}*/
	}
	return nil
}

//creates a context specific to factory manager
func (msgMgr *messagingManager) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementMessagingManager: msgMgr.proxy}, core.ServerElementMessagingManager)
}
