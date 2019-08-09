package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"strings"
)

type ConfigFormsBuilder struct {
	core.Service
	objectConfig map[string]map[string]core.Configuration
}

func (svc *ConfigFormsBuilder) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.objectConfig = make(map[string]map[string]core.Configuration)
	return nil
}

func (svc *ConfigFormsBuilder) Load(ctx core.ServerContext, modInfo *components.ModInfo) error {
	modName := modInfo.ModName
	_, exists := svc.objectConfig[modName]
	if !exists {
		svc.objectConfig[modName] = modInfo.Configurations
	}
	return nil
}

func (svc *ConfigFormsBuilder) Loaded(ctx core.ServerContext) error {
	return nil
}

func (svc *ConfigFormsBuilder) UILoad(ctx core.ServerContext) map[string]config.Config {
	return nil
}
func (svc *ConfigFormsBuilder) LoadingComplete(ctx core.ServerContext) map[string]config.Config {
	reg := ctx.CreateConfig()
	forms := ctx.CreateConfig()
	reg.Set(ctx, "Forms", forms)
	for modName, objConf := range svc.objectConfig {
		formConf := svc.createConfForm(ctx, modName, objConf)
		forms.Set(ctx, "module_config_"+strings.ToLower(modName), formConf)
	}

	return map[string]config.Config{"registry": reg}
}
func (svc *ConfigFormsBuilder) Unloaded(ctx core.ServerContext, insName, modName string) error {
	return nil
}
func (svc *ConfigFormsBuilder) Unloading(ctx core.ServerContext, insName, modName string) error {
	return nil
}
func (svc *ConfigFormsBuilder) createConfForm(ctx core.ServerContext, modName string, configs map[string]core.Configuration) config.Config {
	configForm := ctx.CreateConfig()
	configFormInfo := ctx.CreateConfig()
	//configFormInfo.Set(ctx, "entity", entity.object)
	configFormInfo.Set(ctx, "className", fmt.Sprint(" configform ", "module_config_"+strings.ToLower(modName)))
	configFormInfo.Set(ctx, "info", configFormInfo)

	configFormFields := ctx.CreateConfig()

	for fname, fConf := range configs {
		fieldConf := ctx.CreateConfig()
		fieldConf.Set(ctx, "type", fConf.GetType())
		fieldConf.Set(ctx, "required", fConf.IsRequired())
		fieldConf.Set(ctx, "default", fConf.GetDefaultValue())
		fieldConf.Set(ctx, "className", " configformfield "+fname)
		configFormFields.Set(ctx, fname, fieldConf)
	}
	configForm.Set(ctx, "fields", configFormFields)
	return configForm
}
