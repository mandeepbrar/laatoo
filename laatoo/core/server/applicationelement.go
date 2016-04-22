package server

import (
	"laatoo/core/common"
	//	"laatoo/sdk/config"
)

type applicationProxy struct {
	*common.Context
	app *application
}
