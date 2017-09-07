package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type configuration struct {
	name         string
	conftype     string
	required     bool
	value        interface{}
	defaultValue interface{}
}

func newConfiguration(name, conftype string, required bool, defaultValue interface{}) *configuration {
	return &configuration{name, conftype, required, nil, defaultValue}
}
func (conf *configuration) clone() *configuration {
	c := newConfiguration(conf.name, conf.conftype, conf.required, conf.defaultValue)
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
	*objectInfo
	configurations map[string]core.Configuration
}

func newConfigurableObject(description, objectType string) *configurableObject {
	return &configurableObject{objectInfo: newObjectInfo(description, objectType), configurations: make(map[string]core.Configuration)}
}
func (impl *configurableObject) clone() *configurableObject {
	inf := &configurableObject{objectInfo: impl.objectInfo.clone()}
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
)

func buildConfigurableObject(ctx core.ServerContext, conf config.Config) *configurableObject {
	co := &configurableObject{objectInfo: buildObjectInfo(ctx, conf), configurations: make(map[string]core.Configuration)}
	confs, ok := conf.GetSubConfig(ctx, CONFIGURATIONS)
	if ok {
		confNames := confs.AllConfigurations(ctx)
		for _, confName := range confNames {
			confDesc, _ := confs.GetSubConfig(ctx, confName)
			required, _ := confDesc.GetBool(ctx, CONFREQ)
			conftype, _ := confDesc.GetString(ctx, CONFTYPE)
			defaultValue, _ := confDesc.Get(ctx, CONFDEFAULTVALUE)
			co.configurations[confName] = newConfiguration(confName, conftype, required, defaultValue)
		}
	}
	return co
}

func (impl *configurableObject) setConfigurations(confs []core.Configuration) {
	if confs != nil {
		for _, c := range confs {
			impl.configurations[c.GetName()] = c
		}
	}
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
		impl.configurations[name] = newConfiguration(name, config.OBJECTTYPE_STRING, required, defaultValue)
	}
}

func (impl *configurableObject) AddStringConfiguration(ctx core.ServerContext, name string) {
	impl.AddConfigurations(ctx, map[string]string{name: config.OBJECTTYPE_STRING})
}

func (impl *configurableObject) AddConfigurations(ctx core.ServerContext, configs map[string]string) {
	for name, typ := range configs {
		impl.configurations[name] = newConfiguration(name, typ, true, nil)
	}
}

func (impl *configurableObject) AddOptionalConfigurations(ctx core.ServerContext, configs map[string]string, defaultValues map[string]interface{}) {
	for name, typ := range configs {
		var defaultValue interface{}
		if defaultValues != nil {
			defaultValue = defaultValues[name]
		}
		impl.configurations[name] = newConfiguration(name, typ, false, defaultValue)
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

func (impl *configurableObject) GetBoolConfiguration(ctx core.ServerContext, name string) (bool, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return false, ok
	}
	return c.(bool), ok
}

func (impl *configurableObject) GetMapConfiguration(ctx core.ServerContext, name string) (config.Config, bool) {
	c, ok := impl.GetConfiguration(ctx, name)
	if !ok && c == nil {
		return nil, ok
	}
	conf, cok := c.(config.Config)
	if cok {
		return conf, cok
	}
	m, mok := c.(map[string]interface{})
	if mok {
		return config.GenericConfig(m), true
	} else {
		return nil, false
	}
}

func (impl *configurableObject) processInfo(ctx core.ServerContext, conf config.Config) error {
	log.Trace(ctx, "Processing Configurations")
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
			case config.OBJECTTYPE_STRINGMAP:
				val, ok = conf.GetSubConfig(ctx, name)
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
