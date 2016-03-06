package data

import (
	"laatoosdk/core"
	"time"
)

type Auditable interface {
	IsNew() bool
	SetUpdatedOn(string)
	SetUpdatedBy(string)
	SetCreatedBy(string)
	GetCreatedBy() string
}

func Audit(ctx core.Context, item interface{}) {
	if item != nil {
		auditable := item.(Auditable)
		if auditable != nil {
			usr := ctx.GetUser()
			if usr != nil {
				id := usr.GetId()
				if auditable.IsNew() {
					auditable.SetCreatedBy(id)
				}
				auditable.SetUpdatedBy(id)
				auditable.SetUpdatedOn(time.Now().Format(time.UnixDate))
			}
		}
	}
}
