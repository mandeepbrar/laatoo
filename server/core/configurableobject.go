package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/server/common"
	"strings"
)

type configuration struct {
	name         string
	conftype     string
	required     bool
	value        interface{}
	defaultValue interface{}
}

func newConfiguration(name, conftype string, required bool, defaultValue interface{}) core.Configuration {
	return &configuration{name, conftype, required, nil, defaultValue}
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

const (
	CONFIGURATIONS   = "configurations"
	CONFTYPE         = "type"
	CONFDEFAULTVALUE = "default"
	CONFREQ          = "required"
)

func buildConfigurableObject(conf config.Config) *configurableObject {
	co := &configurableObject{objectInfo: buildObjectInfo(conf), configurations: make(map[string]core.Configuration)}
	confs, ok := conf.GetSubConfig(CONFIGURATIONS)
	if ok {
		confNames := confs.AllConfigurations()
		for _, confName := range confNames {
			confDesc, _ := confs.GetSubConfig(confName)
			required, _ := confDesc.GetBool(CONFREQ)
			conftype, _ := confDesc.GetString(CONFTYPE)
			defaultValue, _ := confDesc.Get(CONFDEFAULTVALUE)
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
		val, found = common.LookupContext(ctx, name)
	} else {
		conf := c.(*configuration)
		if conf.value != nil {
			val = conf.value
			valStr, ok := val.(string)
			if ok && strings.HasPrefix(valStr, ":") {
				val, found = common.LookupContext(ctx, valStr[1:])
			}
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
	m, mok := c.(map[string]interface{})
	if mok {
		return config.GenericConfig(m), true
	} else {
		return nil, false
	}
}

func (impl *configurableObject) processInfo(ctx core.ServerContext, conf config.Config) error {
	confs := impl.GetConfigurations()
	for name, configObj := range confs {
		configu := configObj.(*configuration)
		val, ok := conf.Get(name)

		if !ok && configu.required {
			return errors.MissingConf(ctx, name)
		}
		if ok {
			switch configu.conftype {
			case "", config.OBJECTTYPE_STRING:
				val, ok = conf.GetString(name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_STRINGMAP:
				val, ok = conf.GetSubConfig(name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_STRINGARR:
				val, ok = conf.GetStringArray(name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_CONFIG:
				val, ok = conf.GetSubConfig(name)
				if ok {
					configu.value = val
				}
			case config.OBJECTTYPE_BOOL:
				val, ok = conf.GetBool(name)
				if ok {
					configu.value = val
				}
			default:
				configu.value = val
			}

			//check if value provided is a variable name..
			//allow assignment if its a variable name
			if !ok {
				strval, _ := conf.GetString(name)
				if strings.HasPrefix(strval, ":") {
					configu.value = strval
					ok = true
				}
			}

			//configuration was there but wrong type
			if !ok && configu.required {
				return errors.BadConf(ctx, name)
			}
		}
	}

	return nil
}
