package core

type TopicListener func(ctx Context, topic string, message interface{})

type PubSub interface {
	Service
	Publish(ctx Context, topic string, message interface{}) error
	Subscribe(ctx Context, topics []string, lstnr TopicListener) error
}
