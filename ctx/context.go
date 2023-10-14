package ctx

import (
	"context"
	"net/http"
	"time"

	_ "github.com/blang/semver"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/fatih/color"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gorilla/websocket"
	_ "github.com/influxdata/go-syslog/rfc5424"
	_ "github.com/json-iterator/go"
	_ "github.com/labstack/echo"
	_ "github.com/mandeepbrar/go-discover"
	_ "github.com/mholt/archiver"
	_ "github.com/radovskyb/watcher"
	_ "github.com/rs/cors"
	_ "github.com/twinj/uuid"
	_ "github.com/valyala/fastjson"
	_ "goji.io"

	_ "github.com/buraksezer/olric"
	_ "github.com/hashicorp/mdns"
	_ "github.com/json-iterator/go"
	_ "github.com/lesismal/arpc"
	_ "github.com/ugorji/go/codec"
	_ "golang.org/x/oauth2"
	_ "google.golang.org/appengine"
	_ "google.golang.org/grpc"
	_ "gopkg.in/yaml.v3"
)

type Context interface {
	context.Context
	GetId() string
	GetName() string
	GetPath() string
	GetParent() Context
	Get(key string) (interface{}, bool)
	GetCreationTime() time.Time
	GetElapsedTime() time.Duration
	SetGaeReq(req *http.Request)
	Set(key string, value interface{})
	SetVals(vals map[string]interface{})
	GetString(key string) (string, bool)
	GetBool(key string) (bool, bool)
	GetInt(key string) (int, bool)
	GetStringArray(key string) ([]string, bool)
	SubCtx(name string) Context
	NewCtx(name string, newpath bool) Context
	GetAppengineContext() context.Context
	HttpClient() *http.Client
	GetOAuthContext() context.Context
	WithCancel() (Context, context.CancelFunc)
	WithDeadline(timeout time.Time) (Context, context.CancelFunc)
	WithTimeout(timeout time.Duration) (Context, context.CancelFunc)
	WithValue(key, val interface{}) Context
	WithContext(parent context.Context) Context
	CompleteContext()
	Dump()
	LogTrace(msg string, args ...interface{})
	LogDebug(msg string, args ...interface{})
	LogInfo(msg string, args ...interface{})
	LogWarn(msg string, args ...interface{})
	LogError(msg string, args ...interface{})
	LogFatal(msg string, args ...interface{})
}
