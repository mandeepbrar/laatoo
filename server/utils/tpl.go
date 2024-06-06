package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"text/template"

	genutils "laatoo.io/sdk/utils"

	"laatoo.io/sdk/ctx"
	"laatoo.io/sdk/server/log"
)

func ProcessTemplate(ctx ctx.Context, cont []byte, funcs map[string]interface{}) ([]byte, error) {
	contextVar := func(args ...string) string {
		val, ok := ctx.Get(args[0])
		if ok {
			strval, ok := val.(string)
			if ok {
				if len(args) > 1 && strings.TrimSpace(strval) != "" {
					return fmt.Sprint(args[1], val)
				}
				return strval
			}
			if val == nil {
				return ""
			}
			retval, err := json.Marshal(val)
			if err != nil {
				log.Error(ctx, "Error in conf", "Err", err)
			}
			return string(retval)
		}
		return ""
	}

	defaultVar := func(args ...string) string {
		_, ok := ctx.Get(args[0])
		if !ok {
			return contextVar(args[1])
		} else {
			return contextVar(args[0])
		}
	}

	is := func(variable string) bool {
		val, _ := ctx.GetBool(variable)
		return val
	}

	exists := func(variable string) bool {
		_, ok := ctx.Get(variable)
		return ok
	}

	contains := func(variable string, val string) bool {
		vals, ok := ctx.GetStringArray(variable)
		if ok {
			return genutils.StrContains(vals, val) >= 0
		}
		return false
	}

	equals := func(variable string, val string) bool {
		valToCompare, ok := ctx.Get(variable)
		if ok {
			return fmt.Sprint("%v", valToCompare) == val
		}
		return false
	}

	json := func(variable string) string {
		varval, ok := ctx.Get(variable)
		if ok {
			val, _ := json.Marshal(varval)
			if val != nil {
				return string(val)
			}
		}
		return ""
	}

	jsReplace := func(args ...string) string {
		return fmt.Sprintf("javascript###replace@@@%s###", args[0])
	}

	jsFormat := func(args ...string) string {
		vars := ""
		if len(args) > 1 {
			vars = strings.Join(args[1:], "@@@")
		}
		return fmt.Sprintf("javascript###format@@@%s@@@%s###", args[0], vars)
	}

	funcMap := template.FuncMap{"var": contextVar, "is": is, "jsreplace": jsReplace, "jsformat": jsFormat, "default": defaultVar, "equals": equals, "upper": strings.ToUpper, "lower": strings.ToLower, "title": strings.Title, "exists": exists, "contains": contains, "json": json}
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

	c := result.String()
	re1 := regexp.MustCompile(`\[\[(.*)\]\]`)
	c = re1.ReplaceAllString(c, "<![CDATA[$1]]>")

	return []byte(c), nil
}

func GetTemplateFileContent(ctx ctx.Context, name string, funcs map[string]interface{}) ([]byte, error) {
	fileData, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	cont, err := ProcessTemplate(ctx, fileData, funcs)
	return cont, err
}
