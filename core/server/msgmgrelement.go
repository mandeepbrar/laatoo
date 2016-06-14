package server

import (
	"laatoo/core/common"
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type messagingManagerProxy struct {
	*common.Context
	manager *messagingManager
}

//subscribe to a topic
func (mgr *messagingManagerProxy) Subscribe(ctx core.ServerContext, topics []string, handler core.ServiceFunc) error {
	return mgr.manager.subscribeTopic(ctx, topics, handler)
}

//publish message using
func (mgr *messagingManagerProxy) Publish(ctx core.RequestContext, topic string, message interface{}) error {
	return mgr.manager.publishMessage(ctx, topic, message)
}

func newMessagingManager(ctx core.ServerContext, name string, parentElem core.ServerElement, commSvcName string) (*messagingManager, *messagingManagerProxy) {
	msgMgr := &messagingManager{parent: parentElem, topicStore: make(map[string][]core.ServiceFunc, 10), commSvcName: commSvcName}
	msgElemCtx := parentElem.NewCtx(name)
	msgElem := &messagingManagerProxy{Context: msgElemCtx.(*common.Context), manager: msgMgr}
	msgMgr.proxy = msgElem
	return msgMgr, msgElem
}

func childMessagingManager(ctx core.ServerContext, name string, parentMessageMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	msgMgrProxy := parentMessageMgr.(*messagingManagerProxy)
	msgMgr := msgMgrProxy.manager
	store := make(map[string][]core.ServiceFunc, len(msgMgr.topicStore))
	for k, _ := range msgMgr.topicStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			store[k] = []core.ServiceFunc{}
		}
	}
	childmsgMgr := &messagingManager{parent: parent, topicStore: store, commSvcName: msgMgr.commSvcName}
	childmsgMgrElemCtx := parentMessageMgr.NewCtx(name)
	childmsgMgrElem := &messagingManagerProxy{Context: childmsgMgrElemCtx.(*common.Context), manager: childmsgMgr}
	childmsgMgr.proxy = childmsgMgrElem
	return childmsgMgr, childmsgMgrElem
}
