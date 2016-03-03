package laatoocore

import (
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
)

var (
	//listeners for pubsub topics
	topicListeners = make(map[string][]core.TopicListener)
)

//subscribe to a topic
func (env *Environment) SubscribeTopic(ctx core.Context, topic string, handler core.TopicListener) error {
	log.Logger.Info(ctx, "core.pubsub", "Subscribing topic", "topicListeners", topicListeners, "topic", topic)
	listeners, prs := topicListeners[topic]
	if !prs {
		listeners = []core.TopicListener{}
	}
	topicListeners[topic] = append(listeners, handler)
	log.Logger.Trace(ctx, "core.pubsub", "Subscribed topic", "topicListeners", topicListeners, "topic", topic)
	return nil
}

//publish message using
func (env *Environment) PublishMessage(ctx core.Context, topic string, message interface{}) error {
	if env.CommunicationService != nil {
		log.Logger.Trace(ctx, "core.pubsub", "posting message")
		return env.CommunicationService.Publish(ctx, topic, message)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NOCOMMSVC)
}

func (env *Environment) subscribeTopics(ctx core.Context) error {
	if env.CommunicationService != nil {
		topics := make([]string, len(topicListeners))
		i := 0
		for k := range topicListeners {
			topics[i] = k
			i++
		}
		log.Logger.Trace(ctx, "core.pubsub", "Subscribing topics", "topics", topics)
		env.CommunicationService.Subscribe(ctx, topics, func(ctx core.Context, topic string, message interface{}) {
			lsnrs := topicListeners[topic]
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
