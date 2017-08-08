package core

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"log"
	"strings"
)

type configuration struct {
	name         string
	conftype     string
	required     bool
	value        interface{}
	defaultValue interface{}
}

type configurableObject struct {
	configurations map[string]interface{}
}

func newConfigurableObject() *configurableObject {
	return &configurableObject{configurations: make(map[string]interface{})}
}

func (impl *configurableObject) GetConfigurations() map[string]interface{} {
	return impl.configurations
}
func (impl *configurableObject) AddStringConfigurations(ctx core.ServerContext, configs []string, defaultValues []string) {
	if defaultValues != nil && len(configs) != len(defaultValues) {
		log.Fatal("Length of configurations not equal to length of default values", ctx.GetName())
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
		impl.configurations[name] = &configuration{name, config.CONF_OBJECT_STRING, required, nil, defaultValue}
	}
}

func (impl *configurableObject) AddStringConfiguration(ctx core.ServerContext, name string) {
	impl.AddConfigurations(ctx, map[string]string{name: config.CONF_OBJECT_STRING})
}

func (impl *configurableObject) AddConfigurations(ctx core.ServerContext, configs map[string]string) {
	for name, typ := range configs {
		impl.configurations[name] = &configuration{name, typ, true, nil, nil}
	}
}

func (impl *configurableObject) AddOptionalConfigurations(ctx core.ServerContext, configs map[string]string, defaultValues map[string]interface{}) {
	for name, typ := range configs {
		var defaultValue interface{}
		if defaultValues != nil {
			defaultValue = defaultValues[name]
		}
		impl.configurations[name] = &configuration{name, typ, false, nil, defaultValue}
	}
}

func (impl *configurableObject) lookupContext(ctx core.ServerContext, name string) (interface{}, bool) {
	val, found := ctx.Get(name)
	if !found {
		val, found = ctx.GetVariable(name)
		if found {
			return val, found
		} else {
			return nil, false
		}
	} else {
		return val, found
	}
}

func (impl *configurableObject) GetConfiguration(ctx core.ServerContext, name string) (interface{}, bool) {
	var val interface{}
	c, found := impl.configurations[name]
	if !found {
		val, found = impl.lookupContext(ctx, name)
	} else {
		conf := c.(*configuration)
		if conf.value != nil {
			val = conf.value
			valStr, ok := val.(string)
			if ok && strings.HasPrefix(valStr, ":") {
				val, found = impl.lookupContext(ctx, valStr[1:])
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
			case "", config.CONF_OBJECT_STRING:
				val, ok = conf.GetString(name)
				if ok {
					configu.value = val
				}
			case config.CONF_OBJECT_STRINGMAP:
				val, ok = conf.GetSubConfig(name)
				if ok {
					configu.value = val
				}
			case config.CONF_OBJECT_STRINGARR:
				val, ok = conf.GetStringArray(name)
				if ok {
					configu.value = val
				}
			case config.CONF_OBJECT_CONFIG:
				val, ok = conf.GetSubConfig(name)
				if ok {
					configu.value = val
				}
			case config.CONF_OBJECT_BOOL:
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
