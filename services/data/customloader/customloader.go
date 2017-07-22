package main

import (
	"laatoo/sdk/components/data"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
)

const (
	CONF_CUSTOMLOADER_SVC    = "customloader"
	CONF_CUSTOMLOADER_INJECT = "requiredservices"
	CONF_CUSTOM_LOADER       = "loader"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_CUSTOMLOADER_SVC, Object: customLoaderService{}}}
}

type customLoaderService struct {
	loader         data.CustomLoader
	sourceServices []data.DataComponent
	svcNames       []string
}

func (vs *customLoaderService) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Custom Loader service initialize")
	svcNames, ok := conf.GetStringArray(CONF_CUSTOMLOADER_INJECT)
	if !ok {
		return errors.BadConf(ctx, CONF_CUSTOMLOADER_INJECT)
	}
	vs.svcNames = svcNames
	vs.sourceServices = make([]data.DataComponent, len(svcNames))
	loader, ok := conf.GetString(CONF_CUSTOM_LOADER)
	if ok {
		obj, err := ctx.CreateObject(loader)
		if err != nil {
			return err
		}
		t, ok := obj.(data.CustomLoader)
		if !ok {
			return errors.BadConf(ctx, CONF_CUSTOM_LOADER)
		}
		vs.loader = t
	} else {
		return errors.BadConf(ctx, CONF_CUSTOM_LOADER)
	}
	return nil
}

func (vs *customLoaderService) Start(ctx core.ServerContext) error {
	ctx = ctx.SubContext("Custom Loader start")
	for index, name := range vs.svcNames {
		svc, err := ctx.GetService(name)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		dc, ok := svc.(data.DataComponent)
		if !ok {
			return errors.BadConf(ctx, CONF_CUSTOMLOADER_INJECT, "svc", name)
		}
		vs.sourceServices[index] = dc
	}
	return nil
}

func (vs *customLoaderService) Info() *core.ServiceInfo {
	return &core.ServiceInfo{Description: "Custom data loader service"}
}

func (vs *customLoaderService) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	return vs.loader.LoadData(ctx, vs.sourceServices...)
}
