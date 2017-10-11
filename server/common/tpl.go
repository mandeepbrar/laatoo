package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"laatoo/sdk/ctx"
	"laatoo/sdk/log"
	"text/template"
)

func ProcessTemplate(ctx ctx.Context, cont []byte, funcs map[string]interface{}) ([]byte, error) {
	contextVar := func(args ...string) string {
		val, ok := ctx.GetString(args[0])
		if ok && (len(args) > 1) {
			return fmt.Sprint(args[1], val)
		}
		return val
	}
	funcMap := template.FuncMap{"var": contextVar}
	if funcs != nil {
		for k, v := range funcs {
			funcMap[k] = v.(func(variable string) string)
		}
	}
	temp, err := template.New("temp").Funcs(funcMap).Parse(string(cont))
	if err != nil {
		return nil, err
	}
	result := new(bytes.Buffer)
	anon := struct{}{}
	err = temp.Execute(result, anon)
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

func GetTemplateFileContent(ctx ctx.Context, name string, funcs map[string]interface{}) ([]byte, error) {
	fileData, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	log.Error(ctx, "uprocess Config File", "name", name, "conf", string(fileData))
	return ProcessTemplate(ctx, fileData, funcs)
}
