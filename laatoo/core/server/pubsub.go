package server

import (
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

var (
	//listeners for pubsub topics
	topicListeners = make(map[string][]core.TopicListener)
)

//subscribe to a topic
func (app *Application) SubscribeTopic(ctx core.Context, topic string, handler core.TopicListener) error {
	listeners, prs := topicListeners[topic]
	if !prs {
		listeners = []core.TopicListener{}
	}
	topicListeners[topic] = append(listeners, handler)
	log.Logger.Trace(ctx, "Subscribed topic", "topicListeners", topicListeners, "topic", topic)
	return nil
}

//publish message using
func (app *Application) PublishMessage(ctx core.Context, topic string, message interface{}) error {
	if app.CommunicationService != nil {
		log.Logger.Trace(ctx, "posting message")
		return app.CommunicationService.Publish(ctx, topic, message)
	}
	return errors.ThrowError(ctx, CORE_ERROR_NOCOMMSVC)
}

func (app *Application) subscribeTopics(ctx core.Context) error {
	if app.CommunicationService != nil {
		topics := make([]string, len(topicListeners))
		i := 0
		for k := range topicListeners {
			topics[i] = k
			i++
		}
		log.Logger.Trace(ctx, "Subscribing topics", "topics", topics)
		app.CommunicationService.Subscribe(ctx, topics, func(ctx core.Context, topic string, message interface{}) {
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
