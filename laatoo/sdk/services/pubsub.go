package services

import (
	"laatoo/sdk/core"
)

type PubSub interface {
	Publish(ctx core.RequestContext, topic string, message interface{}) error
	Subscribe(ctx core.ServerContext, topics []string, lstnr core.TopicListener) error
}
