module laatoo/server

go 1.12

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/fatih/color v1.7.0
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/imdario/mergo v0.3.7
	github.com/influxdata/go-syslog v1.0.1
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.8 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/radovskyb/watcher v1.0.6
	github.com/rs/cors v1.6.0
	github.com/twinj/uuid v1.0.0
	github.com/ugorji/go v1.1.4
	github.com/valyala/fasttemplate v1.0.1 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	goji.io v2.0.2+incompatible
	golang.org/x/net v0.0.0-20190509222800-a4d6f7feada5
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/text v0.3.0
	google.golang.org/appengine v1.5.0
	gopkg.in/yaml.v2 v2.2.2
	laatoo/sdk v0.0.0
)

replace laatoo/sdk => /laatoo/sdk
