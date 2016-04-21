package messaging

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/services"
)

type pubSub struct {
	*common.Context
	commSvc     services.PubSub
	commSvcName string
	//listeners for pubsub topics
	topicListeners map[string][]core.TopicListener
}

func newCommunicationHandler(ctx core.ServerContext, name string, parent core.ServerElement) (server.ServerElementHandle, core.ServerElement) {
	psCtx := parent.NewCtx(name)
	ps := &pubSub{Context: psCtx.(*common.Context), topicListeners: make(map[string][]core.TopicListener)}
	return ps, ps
}

func (ps *pubSub) Initialize(ctx core.ServerContext, conf config.Config) error {
	commSvcName, ok := conf.GetString(config.CONF_COMM_SVC)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "conf", config.CONF_COMM_SVC)
	}
	ps.commSvcName = commSvcName
	return nil
}

func (ps *pubSub) Start(ctx core.ServerContext) error {
	svc, err := ctx.GetService(ps.commSvcName)
	if err != nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_COMM_SVC)
	}
	pubsub, ok := svc.(services.PubSub)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "conf", config.CONF_COMM_SVC)
	}
	ps.commSvc = pubsub

	return nil
}

//subscribe to a topic
func (ps *pubSub) SubscribeTopic(ctx core.Context, topic string, handler core.TopicListener) error {
	listeners, prs := ps.topicListeners[topic]
	if !prs {
		listeners = []core.TopicListener{}
	}
	ps.topicListeners[topic] = append(listeners, handler)
	log.Logger.Trace(ctx, "Subscribed topic", "topic", topic)
	return nil
}

//publish message using
func (ps *pubSub) PublishMessage(ctx core.Context, topic string, message interface{}) error {
	if ps.commSvc != nil {
		log.Logger.Trace(ctx, "posting message")
		return ps.commSvc.Publish(ctx, topic, message)
	}
	return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_SERVICE, "Name", "Communication Handler")
}

func (ps *pubSub) subscribeTopics(ctx core.Context) error {
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

//creates a context specific to environment
func (ps *pubSub) createContext(ctx core.ServerContext, name string) core.ServerContext {
	return ctx.NewContextWithElements(name,
		core.ContextMap{core.ServerElementCommunicationHandler: ps}, core.ServerElementCommunicationHandler)
}
