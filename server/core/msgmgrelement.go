package core

import (
	"laatoo/sdk/core"
	"laatoo/sdk/server"
)

type messagingManagerProxy struct {
	manager *messagingManager
}

func (proxy *messagingManagerProxy) Reference() core.ServerElement {
	return &messagingManagerProxy{manager: proxy.manager}
}
func (proxy *messagingManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *messagingManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *messagingManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementMessagingManager
}

//subscribe to a topic
func (mgr *messagingManagerProxy) Subscribe(ctx core.ServerContext, topics []string, handler core.MessageListener) error {
	return mgr.manager.subscribeTopic(ctx, topics, handler)
}

//publish message using
func (mgr *messagingManagerProxy) Publish(ctx core.RequestContext, topic string, message interface{}) error {
	return mgr.manager.publishMessage(ctx, topic, message)
}

func newMessagingManager(ctx core.ServerContext, name string, commSvcName string) (*messagingManager, *messagingManagerProxy) {
	msgMgr := &messagingManager{name: name, topicStore: make(map[string][]core.MessageListener, 10), commSvcName: commSvcName}
	msgElem := &messagingManagerProxy{manager: msgMgr}
	msgMgr.proxy = msgElem
	return msgMgr, msgElem
}

func childMessagingManager(ctx core.ServerContext, name string, parentMessageMgr core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	msgMgrProxy := parentMessageMgr.(*messagingManagerProxy)
	msgMgr := msgMgrProxy.manager
	store := make(map[string][]core.MessageListener, len(msgMgr.topicStore))
	for k, _ := range msgMgr.topicStore {
		allowed := true
		for _, filter := range filters {
			if !filter.Allowed(ctx, k) {
				allowed = false
				break
			}
		}
		if allowed {
			store[k] = []core.MessageListener{}
		}
	}
	childmsgMgr := &messagingManager{name: name, topicStore: store, commSvcName: msgMgr.commSvcName}
	childmsgMgrElem := &messagingManagerProxy{manager: childmsgMgr}
	childmsgMgr.proxy = childmsgMgrElem
	return childmsgMgr, childmsgMgrElem
}
