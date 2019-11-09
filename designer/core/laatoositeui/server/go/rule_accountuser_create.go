package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/user"

	//"laatoo/sdk/server/components/data"
	"laatoo/sdk/server/components/rules"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"

	/*"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"*/
	"fmt"
)

type AccountUserCreateRule struct {
}

func (rule *AccountUserCreateRule) Initialize(ctx ctx.Context, conf config.Config) error {
	return nil
}

func (rule *AccountUserCreateRule) Condition(ctx core.RequestContext, trigger *rules.Trigger) bool {

	if trigger.Message != nil {
		_, ok := trigger.Message.(*AccountUser)
		if ok {
			return true
		}
	}
	return false
}

func (rule *AccountUserCreateRule) Action(ctx core.RequestContext, trigger *rules.Trigger) error {
	ent, _ := trigger.Message.(*AccountUser)

	usr := user.DefaultUser{}
	usr.Username = ent.User
	usr.Email = ent.Email
	usr.Name = fmt.Sprintf("%s %s", ent.FName, ent.LName)
	usr.Picture = ent.Picture
	usr.SetTenant(ent.GetTenant())
	usr.Realm = ent.Account
	params := map[string]interface{}{"credentials": usr}
	err := ctx.Forward("register", params)
	return err
}
