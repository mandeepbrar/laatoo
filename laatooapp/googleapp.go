// +build appengine
package laatooapp

import (
	"laatoo"
)

func init() {
	configName := "server"
	//create a server with config name
	laatoo.NewServer(configName, laatoo.CONF_SERVERTYPE_GOOGLEAPP)
}
