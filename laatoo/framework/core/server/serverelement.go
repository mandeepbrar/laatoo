package server

import (
	"laatoo/framework/core/common"
	//	"laatoo/sdk/config"
)

type serverProxy struct {
	*common.Context
	server *serverObject
}
