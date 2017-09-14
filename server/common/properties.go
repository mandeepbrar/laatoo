package common

import (
	"encoding/json"
	"io/ioutil"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/imdario/mergo"
)

func MergeJson(obj1, obj2 map[string]interface{}) map[string]interface{} {
	if obj1 == nil {
		return obj2
	}
	if obj2 == nil {
		return obj1
	}
	res := make(map[string]interface{})
	mergo.MergeWithOverwrite(&res, obj1)
	mergo.MergeWithOverwrite(&res, obj2)
	return res
}

func ReadProperties(ctx core.ServerContext, propsDir string) (map[string]interface{}, error) {
	ok, _, _ := utils.FileExists(propsDir)
	if !ok {
		return nil, nil
	}
	properties := make(map[string]interface{})
	err := filepath.Walk(propsDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".json" {
			cont, err := ioutil.ReadFile(path)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			obj := make(map[string]interface{})
			err = json.Unmarshal(cont, &obj)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			fileName := info.Name()
			locale := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			properties[locale] = obj
		}
		return nil
	})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	defaultProps, ok := properties["default"]
	if ok {
		defMap, ok := defaultProps.(map[string]interface{})
		if ok {
			for k, v := range properties {
				val, ok := v.(map[string]interface{})
				if ok {
					properties[k] = MergeJson(defMap, val)
				}
			}
		}
	}
	return properties, nil
}
