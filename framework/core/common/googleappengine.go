// +build appengine

package common

import (
	"net/http"
	"sync"

	glctx "golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

var (
	gaeMux *laatooGaeMux
)

func init() {
	gaeMux = &laatooGaeMux{ServeMux: http.NewServeMux()}
	http.Handle("/", gaeMux)
}

type laatooGaeMux struct {
	*http.ServeMux
	initializeGae func(*http.Request)
	once          sync.Once
}

func (mux *laatooGaeMux) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	warmupFunc := func() {
		mux.initializeGae(request)
	}
	mux.once.Do(warmupFunc)
	mux.ServeMux.ServeHTTP(w, request)
}

func ConfigureGae(initializer func(*http.Request)) {
	gaeMux.initializeGae = initializer
}

func GaeHandle(pattern string, handler http.Handler) {
	gaeMux.Handle(pattern, handler)
}

func GetAppengineContext(ctx *Context) glctx.Context {
	if ctx == nil || ctx.gaeReq == nil {
		return nil
	}
	return appengine.NewContext(ctx.gaeReq)
}

/*func GetCloudContext(ctx *Context, scope string) glctx.Context {
	appenginectx := GetAppengineContext(ctx)
	if appenginectx == nil {
		log.Print("no request in appengine context")
		return nil
	}
	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.AppEngineTokenSource(appenginectx, scope),
			Base: &urlfetch.Transport{
				Context: appenginectx,
				AllowInvalidServerCertificate: true,
			},
		},
	}
	return cloud.WithContext(appenginectx, appengine.AppID(appenginectx), hc)
}*/

func HttpClient(ctx *Context) *http.Client {
	appenginectx := GetAppengineContext(ctx)
	return &http.Client{
		Transport: &urlfetch.Transport{
			Context: appenginectx,
			AllowInvalidServerCertificate: true,
		},
	}
}

func GetOAuthContext(ctx *Context) glctx.Context {
	return GetAppengineContext(ctx)
}
