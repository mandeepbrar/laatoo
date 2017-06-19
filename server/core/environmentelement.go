package core

import (
	"laatoo/server/common"
)

type environmentProxy struct {
	*common.Context
	env *environment
}
