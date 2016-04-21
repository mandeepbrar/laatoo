package services

import (
	"laatoo/sdk/core"
)

type PubSub interface {
	Publish(ctx core.Context, topic string, message interface{}) error
	Subscribe(ctx core.Context, topics []string, lstnr core.TopicListener) error
}
