package laatoocore

import (
	"github.com/labstack/echo"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

var (
	//listeners for pubsub topics
	topicListeners = make(map[string][]service.TopicListener)
)

//subscribe to a topic
func (env *Environment) SubscribeTopic(ctx *echo.Context, topic string, handler service.TopicListener) {
	log.Logger.Info(ctx, "core.pubsub", "Subscribing topic", "topicListeners", topicListeners, "topic", topic)
	listeners, prs := topicListeners[topic]
	if !prs {
		listeners = []service.TopicListener{}
	}
	topicListeners[topic] = append(listeners, handler)
	log.Logger.Trace(ctx, "core.pubsub", "Subscribed topic", "topicListeners", topicListeners, "topic", topic)
}

//publish message using
func (env *Environment) PublishMessage(ctx *echo.Context, topic string, message interface{}) error {
	if env.CommunicationService != nil {
		log.Logger.Trace(ctx, "core.pubsub", "posting message")
		return env.CommunicationService.Publish(ctx, topic, message)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NOCOMMSVC)
}

func (env *Environment) subscribeTopics(ctx *echo.Context) error {
	if env.CommunicationService != nil {
		topics := make([]string, len(topicListeners))
		i := 0
		for k := range topicListeners {
			topics[i] = k
			i++
		}
		log.Logger.Trace(ctx, "core.pubsub", "Subscribing topics", "topics", topics)
		err := env.CommunicationService.Subscribe(ctx, topics, func(ctx *echo.Context, topic string, message interface{}) {
			lsnrs := topicListeners[topic]
			for _, val := range lsnrs {
				go val(ctx, topic, message)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}
