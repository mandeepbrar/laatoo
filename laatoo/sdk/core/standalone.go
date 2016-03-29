// +build !appengine

package core

import (
	"crypto/tls"
	glctx "golang.org/x/net/context"
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
