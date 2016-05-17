package core

import (
	glctx "golang.org/x/net/context"
	"net/http"
)

type Context interface {
	GetId() string
	GetName() string
	GetParent() Context
	Get(key string) (interface{}, bool)
	SetGaeReq(req *http.Request)
	Set(key string, value interface{})
	GetString(key string) (string, bool)
	GetInt(key string) (int, bool)
	GetStringArray(key string) ([]string, bool)
	SubCtx(name string) Context
	NewCtx(flow string) Context
	GetAppengineContext() glctx.Context
	GetCloudContext(scope string) glctx.Context
	HttpClient() *http.Client
	GetOAuthContext() glctx.Context
}
