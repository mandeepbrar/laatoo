package service

type TopicListener func(topic string, message interface{})

type PubSub interface {
	PublishMessage(topic string, message interface{}) error
	Subscribe(topic string, lstnr TopicListener)
}
