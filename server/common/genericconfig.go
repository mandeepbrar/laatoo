package common

import (
	"laatoo/sdk/config"
	"laatoo/sdk/ctx"
	"laatoo/sdk/utils"
)

type GenericConfig map[string]interface{}

func fillVariables(ctx ctx.Context, val interface{}) interface{} {
	/*expr, ok := val.(string)
	if !ok {
		return val
	}
	cont, err := utils.ProcessTemplate(ctx, []byte(expr), nil)
	if err != nil {
		return val
	}
	return string(cont)*/
	return val
}

//Get string configuration value
func (conf GenericConfig) GetString(ctx ctx.Context, configurationName string) (string, bool) {
	val, found := conf[configurationName]
	if found {
		str, ok := fillVariables(ctx, val).(string)
		if ok {
			return str, true
		}
		return "", false
	}
	return "", false
}

func (conf GenericConfig) Clone() config.Config {
	res := make(GenericConfig, len(conf))
	for k, v := range conf {
		mapV, ok := v.(GenericConfig)
		if ok {
			res[k] = mapV.Clone().(GenericConfig)
		} else {
			res[k] = v
		}

	}
	return res
}

//Get string configuration value
func (conf GenericConfig) GetBool(ctx ctx.Context, configurationName string) (bool, bool) {
	val, found := conf[configurationName]
	if found {
		b, ok := val.(bool)
		if ok {
			return b, true
		}
		val = fillVariables(ctx, val)
		b, ok = val.(bool)
		if ok {
			return b, true
		}
	}
	return false, false
}

//Get string configuration value
func (conf GenericConfig) Get(ctx ctx.Context, configurationName string) (interface{}, bool) {
	val, cok := conf[configurationName]
	if cok {
		return fillVariables(ctx, val), true
	}
	return nil, false
}

func (conf GenericConfig) GetStringArray(ctx ctx.Context, configurationName string) ([]string, bool) {
	val, found := conf[configurationName]
	if found {
		arr, cok := val.([]interface{})
		if !cok {
			return nil, false
		}
		retVal := make([]string, len(arr))
		var ok bool
		for index, val := range arr {
			retVal[index], ok = fillVariables(ctx, val).(string)
			if !ok {
				return nil, false
			}
		}
		return retVal, true
	}
	return nil, false
}

func (conf GenericConfig) GetConfigArray(ctx ctx.Context, configurationName string) ([]config.Config, bool) {
	val, found := conf[configurationName]
	if found {
		confArr, cok := val.([]GenericConfig)
		if cok {
			retVal := make([]config.Config, len(confArr))
			for index, val := range confArr {
				retVal[index] = val
			}
			return retVal, true
		}
		cArr, cok := val.([]interface{})
		if !cok {
			return nil, false
		}
		retVal := make([]config.Config, len(cArr))
		for index, val := range cArr {
			var gc GenericConfig
			gc, ok := val.(map[string]interface{})
			if !ok {
				return nil, false
			}
			retVal[index] = gc
		}
		return retVal, true
	}
	return nil, false
}

func (conf GenericConfig) AllConfigurations(ctx ctx.Context) []string {
	return utils.MapKeys(conf)
}

func (conf GenericConfig) checkConfig(ctx ctx.Context, val interface{}) (config.Config, bool) {
	var gc GenericConfig
	cf, ok := val.(map[string]interface{})
	if ok {
		gc = cf
		return gc, true
	} else {
		c, ok := val.(config.Config)
		if ok {
			return c, true
		} else {
			return nil, false
		}
	}
}

func (conf GenericConfig) GetSubConfig(ctx ctx.Context, configurationName string) (config.Config, bool) {
	val, found := conf[configurationName]
	if found {
		c, ok := conf.checkConfig(ctx, val)
		if ok {
			return c, true
		} else {
			/*			lookupVal := fillVariables(ctx, val)
						if lookupVal != val {
							c, ok := conf.checkConfig(ctx, lookupVal)
							if ok {
								return c, true
							}
						}*/
		}
	}
	return nil, false
}

//Set string configuration value
func (conf GenericConfig) SetString(ctx ctx.Context, configurationName string, configurationValue string) {
	conf.Set(ctx, configurationName, configurationValue)
}

func (conf GenericConfig) Set(ctx ctx.Context, configurationName string, configurationValue interface{}) {
	conf[configurationName] = configurationValue
}

func (conf GenericConfig) SetVals(ctx ctx.Context, vals map[string]interface{}) {
	if vals != nil {
		for k, v := range vals {
			conf.Set(ctx, k, v)
		}
	}
}
