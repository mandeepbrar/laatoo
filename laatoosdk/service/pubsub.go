package service

type TopicListener func(topic string, message interface{})

type PubSub interface {
	Publish(topic string, message interface{}) error
	Subscribe(topics []string, lstnr TopicListener) error
}
