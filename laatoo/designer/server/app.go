// +build appengine

package main

import (
	_ "google.golang.org/appengine/remote_api"
	"laatoo/designer/core"
)

func init() {
	core.Start("server.json")
}
