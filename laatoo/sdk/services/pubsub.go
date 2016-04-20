package services

import (
	"laatoo/sdk/core"
)

type TopicListener func(ctx core.Context, topic string, message interface{})

type PubSub interface {
	Publish(ctx core.Context, topic string, message interface{}) error
	Subscribe(ctx core.Context, topics []string, lstnr TopicListener) error
}
