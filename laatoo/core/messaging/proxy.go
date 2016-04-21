package messaging

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
	"laatoo/sdk/services"
)

type messagingManagerProxy struct {
	*common.Context
	manager *messagingManager
}

//subscribe to a topic
func (mgr *messagingManagerProxy) SubscribeTopic(ctx core.Context, topic string, handler core.TopicListener) error {
	return mgr.manager.subscribeTopic(ctx, topic, handler)
}

//publish message using
func (mgr *messagingManagerProxy) PublishMessage(ctx core.Context, topic string, message interface{}) error {
	return mgr.manager.publishMessage(ctx, topic, message)
}
