// +build !appengine

package core

import (
	"crypto/tls"
	glctx "golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
)

func GetAppengineContext(ctx RequestContext) glctx.Context {
	return nil
}

func GetCloudContext(ctx RequestContext, scope string) glctx.Context {
	return nil
}
func HttpClient(ctx RequestContext) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func GetOAuthContext(ctx Context) glctx.Context {
	return oauth2.NoContext
}
