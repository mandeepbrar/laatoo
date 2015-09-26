package laatoocore

import (
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

var (
	//listeners for pubsub topics
	topicListeners = make(map[string][]service.TopicListener)
)

//subscribe to a topic
func (env *Environment) SubscribeTopic(ctx interface{}, topic string, handler service.TopicListener) {
	listeners, prs := topicListeners[topic]
	if !prs {
		listeners = []service.TopicListener{}
	}
	topicListeners[topic] = append(listeners, handler)
}

//publish message using
func (env *Environment) PublishMessage(ctx interface{}, topic string, message interface{}) error {
	if env.CommunicationService != nil {
		return env.CommunicationService.Publish(ctx, topic, message)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NOCOMMSVC)
}

func (env *Environment) subscribeTopics(ctx interface{}) error {
	if env.CommunicationService != nil {
		topics := make([]string, len(topicListeners))
		i := 0
		for k := range topicListeners {
			topics[i] = k
			i++
		}
		log.Logger.Info(ctx, "core.pubsub", "Subscribing topics", "topics", topics)
		err := env.CommunicationService.Subscribe(ctx, topics, func(ctx interface{}, topic string, message interface{}) {
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
