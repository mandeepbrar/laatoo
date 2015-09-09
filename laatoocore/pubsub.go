package laatoocore

import (
	"laatoosdk/service"
)

var (
	topicListeners = make(map[string][]service.TopicListener)
	publisher      service.PubSub
)

func (env *Environment) SubscribeTopic(topic string, handler service.TopicListener) {
	listeners, prs := topicListeners[topic]
	if !prs {
		listeners = []service.TopicListener{}
	}
	topicListeners[topic] = append(listeners, handler)
}

func (env *Environment) PublishMessage(topic string, message interface{}) error {
	return publisher.PublishMessage(topic, message)
}
