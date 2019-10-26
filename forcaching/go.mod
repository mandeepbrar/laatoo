module forcaching

require (

	cloud.google.com/go v0.38.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/blevesearch/bleve v0.7.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.0
	github.com/dsnet/compress v0.0.1
	github.com/fatih/color v1.7.0
	github.com/frankban/quicktest v1.4.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/mock v1.3.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1
	github.com/google/btree v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20190502144155-8358a9778bd1 // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/imdario/mergo v0.3.7
	github.com/influxdata/go-syslog v1.0.1
	github.com/jinzhu/gorm v1.9.8
	github.com/klauspost/compress v1.5.0 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/kr/pretty v0.1.0
	github.com/kr/pty v1.1.1
	github.com/kr/text v0.1.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0
	github.com/lib/pq v1.1.1
	github.com/mattn/go-colorable v0.1.2
	github.com/mattn/go-isatty v0.0.9
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/nwaples/rardecode v1.0.0
	github.com/olivere/elastic v6.2.17+incompatible
	github.com/pierrec/lz4 v2.2.6+incompatible
	github.com/pmylund/go-cache v2.1.0+incompatible
	github.com/prep/beanstalk v1.2.2
	github.com/radovskyb/watcher v1.0.7
	github.com/rs/cors v1.7.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/tdewolff/minify v2.3.6+incompatible
	github.com/tdewolff/parse v2.3.4+incompatible
	github.com/twinj/uuid v1.0.0
	github.com/ugorji/go v1.1.7
	github.com/ugorji/go/codec v1.1.7
	github.com/ulikunitz/xz v0.5.6
	github.com/valyala/bytebufferpool v1.0.0
	github.com/valyala/fasttemplate v1.0.1
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.uber.org/cadence v0.9.3
	go.uber.org/yarpc v1.41.0
	go.uber.org/zap v1.10.0
	goji.io v2.0.2+incompatible
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472
	golang.org/x/exp v0.0.0-20190510132918-efd6b22b2522 // indirect
	golang.org/x/image v0.0.0-20190507092727-e4e5bf290fec
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422 // indirect
	golang.org/x/mobile v0.0.0-20190509164839-32b2708ab171 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	golang.org/x/tools v0.0.0-20190606124116-d0a3d012864b
	google.golang.org/api v0.5.0 // indirect
	google.golang.org/appengine v1.6.2
	google.golang.org/genproto v0.0.0-20190508193815-b515fa19cec8 // indirect
	google.golang.org/grpc v1.20.1
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/yaml.v2 v2.2.2
	honnef.co/go/tools v0.0.0-20190418001031-e561f6794a2a // indirect
)

//go mod download -json gopkg.in/mgo.v2/bson gopkg.in/mgo.v2 google.golang.org/appengine/taskqueue google.golang.org/appengine/datastore google.golang.org/appengine/search  google.golang.org/appengine/cloudsql golang.org/x/oauth2/google golang.org/x/oauth2/facebook golang.org/x/oauth2 golang.org/x/net/context golang.org/x/image/tiff golang.org/x/image/bmp github.com/ugorji/go/codec golang.org/x/crypto/bcrypt github.com/twinj/uuid github.com/tdewolff/minify github.com/rs/cors github.com/prep/beanstalk github.com/pmylund/go-cache github.com/labstack/echo github.com/labstack/echo/middleware github.com/jinzhu/gorm github.com/jinzhu/gorm/dialects/mysql github.com/gin-gonic/gin github.com/garyburd/redigo/redis github.com/fatih/color github.com/dgrijalva/jwt-go github.com/dyatlov/go-opengraph/opengraph github.com/blevesearch/bleve/search github.com/blevesearch/bleve cloud.google.com/go/storage github.com/lib/pq google.golang.org/appengine golang.org/x/net/html github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql github.com/lib/pq  google.golang.org/grpc cloud.google.com/go  github.com/jinzhu/gorm/dialects/postgres github.com/tdewolff/minify/html/ github.com/golang/protobuf/protoc-gen-go github.com/blang/semver github.com/mholt/archiver github.com/imdario/mergo gopkg.in/yaml.v2 github.com/fsnotify/fsnotify github.com/gorilla/websocket github.com/zpencerq/godux github.com/olivere/elastic github.com/influxdata/go-syslog github.com/crewjam/rfc5424 github.com/radovskyb/watcher
//go mod download -json go.uber.org/cadence go.uber.org/zap go.uber.org/yarpc

go 1.13
