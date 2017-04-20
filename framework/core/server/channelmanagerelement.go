package server

import (
	"laatoo/framework/core/common"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/server"
)

type channelManagerProxy struct {
	*common.Context
	manager *channelManager
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
	cm := &channelManager{channelStore: make(map[string]server.Channel, 10)}
	cmElemCtx := parentElem.NewCtx("Channel Manager:" + name)
	cmElem := &channelManagerProxy{Context: cmElemCtx.(*common.Context), manager: cm}
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
	cm := &channelManager{channelStore: store}
	cmElemCtx := parent.NewCtx("Channel Manager:" + name)
	cmElem := &channelManagerProxy{Context: cmElemCtx.(*common.Context), manager: cm}
	cm.proxy = cmElem
	return cm, cmElem
}
