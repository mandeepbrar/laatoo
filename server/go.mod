module laatoo/server

go 1.12

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/fatih/color v1.7.0
	github.com/frankban/quicktest v1.4.1 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/influxdata/go-syslog v1.0.1
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/pierrec/lz4 v2.2.6+incompatible // indirect
	github.com/radovskyb/watcher v1.0.7
	github.com/rs/cors v1.7.0
	github.com/twinj/uuid v1.0.0
	github.com/ugorji/go/codec v1.1.7
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	goji.io v2.0.2+incompatible
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/appengine v1.6.2
	gopkg.in/yaml.v2 v2.2.2
	laatoo/sdk v0.0.0
)

replace laatoo/sdk => /laatoo/sdk
