package server

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

type messagingManager struct {
	commSvc    services.PubSub
	parent     core.ServerElement
	proxy      server.MessagingManager
	topicStore map[topic][]core.TopicListener
}

func (msgMgr *messagingManager) msgManager(ctx core.ServerContext, conf config.Config) error {
	msgmgrInitializeCtx := msgMgr.createContext(ctx, "Initialize message manager")
	log.Logger.Trace(msgmgrInitializeCtx, "Create Message Topics")
	err := msgMgr.createTopics(msgmgrInitializeCtx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (msgMgr *messagingManager) Start(ctx core.ServerContext) error {
	return nil
}

func (msgMgr *messagingManager) createTopics(ctx core.ServerContext, conf config.Config) error {
	topicsConf, ok := conf.GetSubConfig(config.CONF_MESSAGE_TOPICS)
	if ok {
		topicNames := topicsConf.AllConfigurations()
		for _, topicName := range topicNames {
			topicConf, err := config.ConfigFileAdapter(topicsConf, topicNames)
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
func (mgr *messagingManager) subscribeTopic(ctx core.Context, topic string, handler core.TopicListener) error {
	listeners, prs := ps.topicListeners[topic]
	if !prs {
		listeners = []core.TopicListener{}
	}
	ps.topicListeners[topic] = append(listeners, handler)
	log.Logger.Trace(ctx, "Subscribed topic", "topic", topic)
	return nil
}

//publish message using
func (mgr *messagingManager) publishMessage(ctx core.Context, topic string, message interface{}) error {
	if ps.commSvc != nil {
		log.Logger.Trace(ctx, "posting message")
		return ps.commSvc.Publish(ctx, topic, message)
	}
	return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Name", "Communication Handler")
}

func (ps *messagingManager) subscribeTopics(ctx core.Context) error {
	if ps.commSvc != nil {
		topics := make([]string, len(ps.topicListeners))
		i := 0
		for k := range ps.topicListeners {
			topics[i] = k
			i++
		}
		log.Logger.Trace(ctx, "Subscribing topics", "topics", topics)
		ps.commSvc.Subscribe(ctx, topics, func(ctx core.Context, topic string, message interface{}) {
			lsnrs := ps.topicListeners[topic]
			for _, val := range lsnrs {
				go val(ctx, topic, message)
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
		core.ContextMap{core.ServerElementMessageManager: msgMgr.proxy}, core.ServerElementMessageManager)
}
