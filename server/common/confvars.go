package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var confVariables map[string]string

func init() {
	confVariables = make(map[string]string, 0)
	fil := os.Getenv("LAATOO_CONF_VARS")
	if len(fil) == 0 {
		fil = "confvariables.json"
	} else {
		_, err := os.Stat(fil)
		if err != nil {
			fil = "confvariables.json"
		}
	}
	vardata, err := ioutil.ReadFile(fil)
	if err == nil {
		err = json.Unmarshal(vardata, &confVariables)
		if err != nil {
			fmt.Println("Could not read conf variables")
		}
	}
}
