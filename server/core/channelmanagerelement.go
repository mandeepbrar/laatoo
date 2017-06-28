package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type channelManagerProxy struct {
	manager *channelManager
}

func (proxy *channelManagerProxy) Reference() core.ServerElement {
	return &channelManagerProxy{manager: proxy.manager}
}
func (proxy *channelManagerProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *channelManagerProxy) GetName() string {
	return proxy.manager.name
}
func (proxy *channelManagerProxy) GetType() core.ServerElementType {
	return core.ServerElementChannelManager
}

func (cm *channelManagerProxy) Serve(ctx core.ServerContext, channelName string, svc server.Service, channelConfig config.Config) error {
	channel, ok := cm.manager.channelStore[channelName]
	if ok {
		return channel.Serve(ctx, svc, channelConfig)
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "No such channel", channelName)
	}
}

func newChannelManager(ctx core.ServerContext, name string, parentElem core.ServerElement) (*channelManager, *channelManagerProxy) {
	cm := &channelManager{name: name, channelStore: make(map[string]server.Channel, 10), parent: parentElem}
	cmElem := &channelManagerProxy{manager: cm}
	cm.proxy = cmElem
	return cm, cmElem
}

func childChannelManager(ctx core.ServerContext, name string, parentChannelMgr core.ServerElement, parent core.ServerElement, filters ...server.Filter) (server.ServerElementHandle, core.ServerElement) {
	chanMgrProxy := parentChannelMgr.(*channelManagerProxy)
	chanMgr := chanMgrProxy.manager
	store := make(map[string]server.Channel, len(chanMgr.channelStore))
	for k, v := range chanMgr.channelStore {
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
	cm := &channelManager{name: name, channelStore: store, parent: parent}
	cmElem := &channelManagerProxy{manager: cm}
	cm.proxy = cmElem
	return cm, cmElem
}
