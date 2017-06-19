package core

import (
	"laatoo/server/common"
)

type serverProxy struct {
	*common.Context
	server *serverObject
}
