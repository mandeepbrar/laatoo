package laatoocore

import (
	"laatoosdk/log"
	"laatoosdk/service"
)

var (
	topicListeners = make(map[string][]service.TopicListener)
)

func (env *Environment) SubscribeTopic(topic string, handler service.TopicListener) {
	listeners, prs := topicListeners[topic]
	if !prs {
		listeners = []service.TopicListener{}
	}
	topicListeners[topic] = append(listeners, handler)
}

func (env *Environment) PublishMessage(topic string, message interface{}) error {
	return env.pubSub.Publish(topic, message)
}

func (env *Environment) subscribeTopics() error {
	if env.pubSub != nil {
		topics := make([]string, len(topicListeners))
		i := 0
		for k := range topicListeners {
			topics[i] = k
			i++
		}
		log.Logger.Infof("topics ", topics)
		err := env.pubSub.Subscribe(topics, func(topic string, message interface{}) {
			lsnrs := topicListeners[topic]
			for _, val := range lsnrs {
				go val(topic, message)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}
