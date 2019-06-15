package main

import (
	"encoding/json"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"regexp"
)

var (
	jsReplaceRegex    *regexp.Regexp
	jsVarReplaceRegex *regexp.Regexp
)

func processJS(ctx core.ServerContext, input string) string {
	val, found := processJSRegex(ctx, input)
	if !found {
		val, e := json.Marshal(input)
		if e != nil {
			log.Error(ctx, "Error in marshalling string", "string", input, "error", e)
		}
		return string(val)
	} else {
		return val
	}
}

func processJSRegex(ctx core.ServerContext, input string) (string, bool) {
	if jsReplaceRegex == nil {
		jsReplaceRegex, _ = regexp.Compile(`javascript#@#([a-zA-Z0-9\ _\,\(\)\'\.\(\)\[\]\{\}]+)#@#`)
	}
	if jsVarReplaceRegex == nil {
		jsVarReplaceRegex, _ = regexp.Compile(`javascript###([a-zA-Z0-9\ _\,\(\)\'\.\(\)\[\]\{\}]+)###`)
	}

	arr := jsReplaceRegex.FindAllStringIndex(input, -1)
	if len(arr) != 0 {
		return `"` + jsReplaceRegex.ReplaceAllString(input, `"+$1+"`) + `"`, true
	}

	arr = jsVarReplaceRegex.FindAllStringIndex(input, -1)
	if len(arr) == 0 {
		return input, false
	} else {
		return jsVarReplaceRegex.ReplaceAllString(input, `$1`), true
	}
}
