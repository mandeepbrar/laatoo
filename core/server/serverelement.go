package server

import (
	"laatoo/core/common"
	//	"laatoo/sdk/config"
)

type serverProxy struct {
	*common.Context
	server *serverObject
}
