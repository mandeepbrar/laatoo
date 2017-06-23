FROM golang:latest

RUN go get github.com/dgrijalva/jwt-go github.com/fatih/color google.golang.org/appengine google.golang.org/appengine/datastore github.com/pmylund/go-cache github.com/twinj/uuid github.com/ugorji/go/codec golang.org/x/crypto/bcrypt golang.org/x/net/html

RUN go get github.com/garyburd/redigo/redis google.golang.org/appengine/cloudsql github.com/labstack/echo/middleware google.golang.org/appengine/datastore  google.golang.org/appengine/taskqueue gopkg.in/mgo.v2 github.com/jinzhu/gorm github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql github.com/lib/pq google.golang.org/grpc

RUN go get github.com/tdewolff/minify cloud.google.com/go github.com/prep/beanstalk github.com/rs/cors golang.org/x/image/tiff golang.org/x/image/bmp  github.com/labstack/echo github.com/gin-gonic/gin cloud.google.com/go/storage

RUN go get google.golang.org/appengine/search github.com/blevesearch/bleve/search github.com/blevesearch/bleve github.com/dyatlov/go-opengraph/opengraph

RUN go install github.com/jinzhu/gorm/dialects/mysql github.com/jinzhu/gorm/dialects/postgres google.golang.org/appengine/cloudsql google.golang.org/appengine/datastore golang.org/x/oauth2/facebook/  github.com/tdewolff/minify/html/ cloud.google.com/go/storage google.golang.org/appengine/taskqueue/ github.com/labstack/echo/middleware cloud.google.com/go/storage

ENTRYPOINT /bin/bash
