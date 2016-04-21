package messaging

import (
	"laatoo/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

func NewMessagingManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*messagingManager, *messagingManagerProxy) {
	msgMgr := &messagingManager{parent: parentElem, topicStore: make(map[topic]bool, 10)}
	msgElemCtx := parentElem.NewCtx(name)
	msgElem := &messagingManagerProxy{Context: msgElemCtx.(*common.Context), manager: msgMgr}
	msgMgr.proxy = msgElem
	return msgMgr, msgElem
}

func childChannelManager(ctx core.ServerContext, name string, parentMessageMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	msgMgrProxy := parentMessageMgr.(*channelManagerProxy)
	msgMgr := msgMgrProxy.manager
	store := make(map[topic]bool, len(msgMgr.topicStore))
	for k, v := range msgMgr.topicStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			store[k] = v
		}
	}
	msgMgr := &messagingManager{parent: parent, topicStore: store}
	msgMgrElemCtx := parentMessageMgr.NewCtx(name)
	msgMgrElem := &messagingManagerProxy{Context: msgMgrElemCtx.(*common.Context), manager: msgMgr}
	msgMgr.proxy = msgMgrElem
	return msgMgr, msgMgrElem
}
