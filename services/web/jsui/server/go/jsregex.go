package main

import (
	"encoding/json"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"regexp"
	"strings"
	"fmt"
)

var (
	jsReplaceRegex    *regexp.Regexp
	jsFMTRegex *regexp.Regexp
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
		jsReplaceRegex, _ = regexp.Compile(`"javascript###replace@@@([a-zA-Z0-9\ _\,\(\)\'\.\(\)\[\]\{\}]+)###"`)
	}
	if jsFMTRegex == nil {
		jsFMTRegex, _ = regexp.Compile(`"javascript###format@@@([^###]*)###"`)
	}

	resStr := input
	found := false
	matches := jsReplaceRegex.FindAllSubmatch([]byte(input), -1)

	if len(matches) != 0 {
		for _, match := range matches {
			resStr = strings.ReplaceAll(resStr, string(match[0]), string(match[1]))
		}
		found = true
	}

	matches = jsFMTRegex.FindAllSubmatch([]byte(input), -1)
	if len(matches) != 0 {
		for _, match := range matches {
			vars := strings.Split(string(match[1]), "@@@")
			fmtStr := vars[0]
			fmtArgs := vars[1:]
			for _, arg := range fmtArgs {
				fmtStr = strings.Replace(fmtStr, "%s", "\"+"+arg+"+\"", 1)
			}
			fmtStr = fmt.Sprintf("\"%s\"", fmtStr)
			resStr = strings.ReplaceAll(resStr, string(match[0]), fmtStr)
		}
		found = true
	}
/*	//arr = jsVarReplaceRegex.FindAllStringIndex(input, -1)
	if len(matches) == 0 {
		return input, false
	} else {
//		return jsVarReplaceRegex.ReplaceAllString(input, `$1`), true
		for _, match := range matches {
			log.Error(ctx, "javascript matches found", "matches", string(match[0]))
		}
		return input, false
	}*/
	return resStr, found
}
