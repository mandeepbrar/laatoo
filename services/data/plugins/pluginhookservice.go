package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

const (
	DATASVC_PLUGINS = "plugins"
)

type pluginHookService struct {
	*data.DataPlugin
	plugins     []data.DataComponent
	baseService *data.DataPlugin
}

func NewPluginHookService(ctx core.ServerContext, name string, conf config.Config) (*pluginHookService, error) {
	pluginNames, ok := conf.GetStringArray(DATASVC_PLUGINS)
	if !ok {
		return nil, errors.MissingArg(ctx, DATASVC_PLUGINS)
	}
	log.Trace(ctx, "Creating plugins ", "name", name, "pluginNames", pluginNames)
	var plugins []data.DataComponent
	var basePlugin data.DataComponent
	baseService := data.NewDataPlugin(ctx)
	basePlugin = baseService
	plugins = append(plugins, basePlugin)
	for _, str := range pluginNames {
		log.Trace(ctx, "Creating data plugin hook", "name", name, "Plugin name", str)
		var svc interface{}
		switch str {
		case DATASVC_CACHE_PLUGIN:
			{
				svc = NewCacheServiceWithBase(ctx, basePlugin)
				basePlugin = svc.(data.DataComponent)
				plugins = append(plugins, basePlugin)
			}
		case DATASVC_JOIN_PLUGIN:
			{
				svc = NewJoinServiceWithBase(ctx, basePlugin)
				basePlugin = svc.(data.DataComponent)
				plugins = append(plugins, basePlugin)
			}
		case DATASVC_CHECKOWNER:
			{
				svc = NewCheckOwnerServiceWithBase(ctx, basePlugin)
				basePlugin = svc.(data.DataComponent)
				plugins = append(plugins, basePlugin)
			}
		}
	}
	return &pluginHookService{DataPlugin: data.NewDataPluginWithBase(ctx, basePlugin), plugins: plugins, baseService: baseService}, nil
}

func (svc *pluginHookService) Initialize(ctx core.ServerContext) error {
	err := svc.baseService.Initialize(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	for _, plugin := range svc.plugins {
		svc := plugin.(core.Service)
		err := svc.Initialize(ctx)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return svc.DataPlugin.Initialize(ctx)
}
func (svc *pluginHookService) Start(ctx core.ServerContext) error {
	err := svc.baseService.Start(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	for _, plugin := range svc.plugins {
		svc := plugin.(core.Service)
		err := svc.Start(ctx)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	return svc.DataPlugin.Start(ctx)
}
