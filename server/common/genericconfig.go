package common

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
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
			log.Trace(ctx, "Config", "configurationName", configurationName, "val", str)
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

func (conf GenericConfig) ToMap() map[string]interface{} {
	return map[string]interface{}(conf)
}

func (conf GenericConfig) GetRoot(ctx ctx.Context) (string, config.Config, bool) {
	confNames := conf.AllConfigurations(ctx)
	if len(confNames) == 1 {
		rootElem := confNames[0]
		rootConf, _ := conf.GetSubConfig(ctx, rootElem)
		return rootElem, rootConf, true
	}
	return "", nil, false
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
		strarr, sok := val.([]string)
		if sok {
			return strarr, true
		}

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
		retVal, cok := val.([]config.Config)
		if cok {
			return retVal, true
		}
		confArr, cok := val.([]GenericConfig)
		if cok {
			retVal = make([]config.Config, len(confArr))
			for index, val := range confArr {
				retVal[index] = val
			}
			return retVal, true
		}
		cArr, cok := val.([]interface{})
		if !cok {
			return nil, false
		}
		retVal = make([]config.Config, len(cArr))
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

func (conf GenericConfig) GetStringMap(ctx ctx.Context, configurationName string) (map[string]interface{}, bool) {
	val, found := conf[configurationName]
	if found {
		cf, ok := val.(map[string]interface{})
		if ok {
			return cf, ok
		}
	}
	return nil, false
}

func (conf GenericConfig) GetStringsMap(ctx ctx.Context, configurationName string) (map[string]string, bool) {
	val, found := conf[configurationName]
	if found {
		cf, ok := val.(map[string]interface{})
		if ok {
			sm := make(map[string]string)
			for key, val := range cf {
				strval, ok := val.(string)
				if !ok {
					return nil, false
				}
				sm[key] = strval
			}
			return sm, true
		} else {
			res, ok := val.(map[string]string)
			if ok {
				return res, ok
			}
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
