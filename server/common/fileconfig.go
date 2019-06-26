package common

import (
	"encoding/json"
	"fmt"
	"laatoo/sdk/common/config"
	context "laatoo/sdk/server/ctx"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"

	yaml "gopkg.in/yaml.v2"
)

//creates a new config object from file provided to it
//only works for json configs
func NewConfigFromFile(ctx context.Context, file string, funcs map[string]interface{}) (config.Config, error) {
	fileData, err := utils.GetTemplateFileContent(ctx, file, funcs)
	if err != nil {
		return nil, fmt.Errorf("Error opening config file %s. Error: %s", file, err.Error())
	}
	return newConfig(ctx, fileData, funcs)
}

func NewConfig(ctx context.Context, data []byte, funcs map[string]interface{}) (config.Config, error) {
	fileData, err := utils.ProcessTemplate(ctx, data, funcs)
	if err != nil {
		return nil, fmt.Errorf("Error processing data",  err.Error())
	}
	return newConfig(ctx, fileData, funcs)
}

func newConfig(ctx context.Context, data []byte, funcs map[string]interface{}) (config.Config, error) {
	conf := make(GenericConfig, 50)
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("Error parsing config %s. Error: %s", string(data), err.Error())
	}
	cleanMaps(ctx, conf)
	_, err := json.Marshal(conf)
	if err != nil {
		log.Error(ctx, "marshal error", "err", err)
	}
	log.Trace(ctx, "Conf loaded ", "conf", conf)
	return conf, nil
}

func cleanMaps(ctx context.Context, input map[string]interface{}) {
	for k, v := range input {
		imap, ok := v.(map[interface{}]interface{})
		if ok {
			strmap := convMap(ctx, imap)
			cleanMaps(ctx, strmap)
			input[k] = strmap
		}
		iarr, ok := v.([]interface{})
		if ok {
			input[k] = convArr(ctx, iarr)
		}
	}
}

func convArr(ctx context.Context, arr []interface{}) interface{} {
	if len(arr) == 0 {
		return arr
	}
	_, ok := arr[0].(map[interface{}]interface{})
	if ok {
		resarr := make([]GenericConfig, len(arr))
		for i, item := range arr {
			imap := item.(map[interface{}]interface{})
			resarr[i] = GenericConfig(convMap(ctx, imap))
		}
		return resarr
	}
	_, ok = arr[0].(string)
	if ok {
		resarr := make([]string, len(arr))
		for i, item := range arr {
			str := item.(string)
			resarr[i] = str
		}
		return resarr
	}
	return arr
}

func convMap(ctx context.Context, imap map[interface{}]interface{}) map[string]interface{} {
	strmap := make(map[string]interface{})
	for k, val := range imap {
		strmap[k.(string)] = val
	}
	cleanMaps(ctx, strmap)
	return strmap
}
