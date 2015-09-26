package service

type TopicListener func(ctx interface{}, topic string, message interface{})

type PubSub interface {
	Publish(ctx interface{}, topic string, message interface{}) error
	Subscribe(ctx interface{}, topics []string, lstnr TopicListener) error
}
