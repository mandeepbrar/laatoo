package filter

import (
	"laatoo/sdk/server/core"
)

type AllowedObjectsFilter struct {
	allowedObjects map[string]bool
}

func (af *AllowedObjectsFilter) Allowed(ctx core.ServerContext, objectName string) bool {
	if af.allowedObjects != nil {
		allowed, ok := af.allowedObjects[objectName]
		if ok {
			return allowed
		}
		return false
	}
	return false
}
