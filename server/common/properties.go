package common

import (
	"io/ioutil"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/imdario/mergo"
)

func MergeProps(obj1, obj2 map[string]interface{}) map[string]interface{} {
	if obj1 == nil {
		return obj2
	}
	if obj2 == nil {
		return obj1
	}
	res := make(map[string]interface{})
	mergo.Merge(&res, obj1)
	mergo.Merge(&res, obj2)
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
		if filepath.Ext(path) == ".yml" {
			cont, err := ioutil.ReadFile(path)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			obj := make(map[string]interface{})
			err = yaml.Unmarshal(cont, &obj)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			cleanMaps(ctx, obj)
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
					properties[k] = MergeProps(defMap, val)
				}
			}
		}
	}
	return properties, nil
}
