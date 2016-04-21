package server

import (
	"laatoo/sdk/core"
)

type MessagingManager interface {
	core.ServerElement
	Publish(ctx core.Context, topic string, message interface{}) error
	Subscribe(ctx core.Context, topics []string, lstnr core.TopicListener) error
}
