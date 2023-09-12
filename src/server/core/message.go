package core

import "laatoo/sdk/server/auth"

type MessageListener func(ctx RequestContext, message *Message, info map[string]interface{}) error

type Message struct {
	Data   interface{}
	Tenant auth.TenantInfo
	User   auth.User
}
