// +build !appengine

package common

import (
	"crypto/tls"
	glctx "golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
)

func GetAppengineContext(ctx *Context) glctx.Context {
	return nil
}

func GetCloudContext(ctx *Context, scope string) glctx.Context {
	return nil
}
func HttpClient(ctx *Context) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func GetOAuthContext(ctx *Context) glctx.Context {
	return oauth2.NoContext
}
