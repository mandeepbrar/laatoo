// +build appengine

package main

import (
	"laatoocore"
)

func init() {
	start("server", laatoocore.CONF_SERVERTYPE_GOOGLEAPP)
}
