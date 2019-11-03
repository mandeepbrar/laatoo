package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

type configuration struct {
	name         string
	conftype     string
	required     bool
	value        interface{}
	defaultValue interface{}
	variable     string
	conf         config.Config
}

func newConfiguration(name, conftype string, required bool, defaultValue interface{}, variable string) *configuration {
	return &configuration{name, conftype, required, nil, defaultValue, variable, nil}
}
func (conf *configuration) clone() *configuration {
	c := newConfiguration(conf.name, conf.conftype, conf.required, conf.defaultValue, conf.variable)
	c.value = conf.value
	return c
}
func (conf *configuration) GetName() string {
	return conf.name
}
func (conf *configuration) IsRequired() bool {
	return conf.required
}
func (conf *configuration) GetDefaultValue() interface{} {
	return conf.defaultValue
}
func (conf *configuration) GetValue() interface{} {
	return conf.value
}
func (conf *configuration) GetType() string {
	return conf.conftype
}

type configurableObject struct {
	core.Info
	name           string
	configurations map[string]core.Configuration
}

func newConfigurableObject(name, description, objectType string) *configurableObject {
	return &configurableObject{name: name, Info: newObjectInfo(description, objectType), configurations: make(map[string]core.Configuration)}
}
func (impl *configurableObject) clone() *configurableObject {
	inf := &configurableObject{name: impl.name, Info: impl.Info}
	inf.configurations = make(map[string]core.Configuration, len(impl.configurations))
	for k, v := range impl.configurations {
		inf.configurations[k] = v.(*configuration).clone()
	}
	return inf
}

const (
	CONFIGURATIONS   = "configurations"
	CONFTYPE         = "type"
	CONFDEFAULTVALUE = "default"
	CONFREQ          = "required"
	CONFVARTOSET     = "variable"
)

func buildConfigurableObject(ctx core.ServerContext, name string, conf config.Config) *configurableObject {
	co := &configurableObject{name: name, Info: buildObjectInfo(ctx, conf), configurations: make(map[string]core.Configuration)}
	confs, ok := conf.GetSubConfig(ctx, CONFIGURATIONS)
	if ok {
		confNames := confs.AllConfigurations(ctx)
		for _, confName := range confNames {
			confDesc, ok := confs.GetSubConfig(ctx, confName)
			if ok {
				required, _ := confDesc.GetBool(ctx, CONFREQ)
				conftype, _ := confDesc.GetString(ctx, CONFTYPE)
				defaultValue, _ := confDesc.Get(ctx, CONFDEFAULTVALUE)
				variable, _ := confDesc.GetString(ctx, CONFVARTOSET)
				configObj := newConfiguration(confName, conftype, required, defaultValue, variable)
				configObj.conf = confDesc
				co.configurations[confName] = configObj
				log.Error(ctx, "configurable object", "confName", confName, "conf", co.configurations[confName])
			}
		}
	}
	return co
}

func (impl *configurableObject) GetName() string {
	return impl.name
}

func (impl *configurableObject) setConfigurations(confs []core.Configuration) {
	if confs != nil {
		for _, c := range confs {
			impl.configurations[c.GetName()] = c
		}
	}
}

func (impl *configurableObject) getConfigurationsToInject() map[string]interface{} {
	confsToInject := make(map[string]interface{})
	if impl.configurations != nil {
		for _, c := range impl.configurations {
			conf := c.(*configuration)
			val := conf.GetValue()
			if conf.variable != "" && val != nil {
				confsToInject[conf.variable] = val
			}
		}
	}
	return confsToInject
}

func (impl *configurableObject) GetConfigurations() map[string]core.Configuration {
	return impl.configurations
}
func (impl *configurableObject) AddStringConfigurations(ctx core.ServerContext, configs []string, defaultValues []string) {
	if defaultValues != nil && len(configs) != len(defaultValues) {
		log.Error(ctx, "Length of configurations not equal to length of default values", ctx.GetName())
		return
	}
	required := true
	if defaultValues != nil {
		required = false
	}
	for index, name := range configs {
		defaultValue := ""
		if defaultValues != nil {
			defaultValue = defaultValues[index]
		}
		impl.configurations[name] = newConfiguration(name, config.OBJECTTYPE_STRING, required, defaultValue, "")
	}
}

func (impl *configurableObject) AddStringConfiguration(ctx core.ServerContext, name string) {
	impl.AddConfigurations(ctx, map[string]string{name: config.OBJECTTYPE_STRING})
}

func (impl *configurableObject) AddConfigurations(ctx core.ServerContext, configs map[string]string) {
	for name, typ := range configs {
		impl.configurations[name] = newConfiguration(name, typ, true, nil, "")
	}
}

func (impl *configurableObject) AddOptionalConfigurations(ctx core.ServerContext, configs map[string]string, defaultValues map[string]interface{}) {
	for name, typ := range configs {
		var defaultValue interface{}
		if defaultValues != nil {
			defaultValue = defaultValues[name]
		}
		impl.configurations[name] = newConfiguration(name, typ, false, defaultValue, "")
	}
}

func (impl *configurableObject) GetConfiguration(ctx core.ServerContext, name string) (interface{}, bool) {
	var val interface{}
	c, found := impl.configurations[name]
	if !found {
		return nil, false
	} else {
		conf := c.(*configuration)
		if conf.value != nil {
			val = conf.value
		} else {
			val = conf.defaultValue
			found = false
		}
	}
	return val, found
}

func (impl *configurableObject) GetStringConfiguration(ctx core.ServerContext, name string) (string, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return "", ok
	}
	return c.(string), ok
}

func (impl *configurableObject) GetStringArrayConfiguration(ctx core.ServerContext, name string) ([]string, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return nil, false
	}
	val, ok := c.([]string)
	if !ok {
		return nil, false
	}
	return val, true
}

func (impl *configurableObject) GetStringsMapConfiguration(ctx core.ServerContext, name string) (map[string]string, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return nil, false
	}
	val, ok := c.(map[string]string)
	if !ok {
		return nil, false
	}
	return val, true
}

func (impl *configurableObject) GetBoolConfiguration(ctx core.ServerContext, name string) (bool, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return false, false
	}
	return c.(bool), ok
}

func (impl *configurableObject) GetMapConfiguration(ctx core.ServerContext, name string) (config.Config, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return nil, false
	}
	conf, cok := c.(config.Config)
	if cok {
		return conf, cok
	}
	m, mok := c.(map[string]interface{})
	if mok {
		conf := ctx.CreateConfig()
		conf.SetVals(ctx, m)
		return conf, true
	} else {
		return nil, false
	}
}

func (impl *configurableObject) GetSecretConfiguration(ctx core.ServerContext, name string) ([]byte, bool) {
	secretsManager := ctx.GetServerElement(core.ServerElementSecretsManager).(elements.SecretsManager)
	if secretsManager != nil {
		return secretsManager.Get(ctx, name)
	}
	log.Warn(ctx, "Secrets manager not found")
	return nil, false
}

func (impl *configurableObject) processInfo(ctx core.ServerContext, conf config.Config) error {
	confs := impl.GetConfigurations()
	for name, configObj := range confs {
		configu := configObj.(*configuration)
		val, ok := conf.Get(ctx, name)
		if !ok && configu.required {
			return errors.MissingConf(ctx, name)
		}
		if ok {
			switch configu.conftype {
			case "", config.OBJECTTYPE_STRING:
				val, ok = conf.GetString(ctx, name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_STRINGSMAP:
				val, ok = conf.GetStringsMap(ctx, name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_STRINGMAP:
				val, ok = conf.GetStringMap(ctx, name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_STRINGARR:
				val, ok = conf.GetStringArray(ctx, name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_CONFIG:
				val, ok = conf.GetSubConfig(ctx, name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_BOOL:
				val, ok = conf.GetBool(ctx, name)
				if ok {
					configu.value = val
				}
			default:
				configu.value = val
			}

			//configuration was there but wrong type
			if !ok && configu.required {
				return errors.BadConf(ctx, name)
			}
		}
	}

	return nil
}
