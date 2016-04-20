package filter

import (
	"laatoo/sdk/core"
)

type RestrictedObjectsFilter struct {
	restrictedObjects map[string]bool
}

func (rf *RestrictedObjectsFilter) Allowed(ctx core.ServerContext, objectName string) bool {
	if rf.restrictedObjects != nil {
		restricted, ok := rf.restrictedObjects[objectName]
		if ok {
			return !restricted
		}
		return true
	}
	return true
}
