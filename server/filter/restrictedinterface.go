package filter

import (
	"laatoo/sdk/server/core"
	"reflect"
)

type RestrictedInterfaceFilter struct {
	restrictedInterfaces map[reflect.Type]bool
}

func (rf *RestrictedInterfaceFilter) Allowed(ctx core.ServerContext, objectName string) bool {
	/*	obj, err := loader.CreateObject(ctx, objectName, nil)
		if err != nil {
			return false
		}
		objType := reflect.TypeOf(obj)
		for interfaceType, restricted := range rf.restrictedInterfaces {
			if objType.Implements(interfaceType) {
				return restricted
			}
		}*/
	return true
}
