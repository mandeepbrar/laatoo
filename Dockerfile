FROM golang

RUN go get github.com/dgrijalva/jwt-go github.com/fatih/color google.golang.org/appengine google.golang.org/appengine/datastore github.com/pmylund/go-cache github.com/twinj/uuid github.com/ugorji/go/codec golang.org/x/crypto/bcrypt golang.org/x/net/html

RUN go get github.com/garyburd/redigo/redis google.golang.org/appengine/cloudsql github.com/labstack/echo/middleware google.golang.org/appengine/datastore  google.golang.org/appengine/taskqueue gopkg.in/mgo.v2 github.com/jinzhu/gorm github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql github.com/lib/pq

RUN go get github.com/tdewolff/minify google.golang.org/cloud github.com/prep/beanstalk github.com/rs/cors golang.org/x/image/tiff golang.org/x/image/bmp  github.com/labstack/echo github.com/gin-gonic/gin

COPY laatoo /go/src/laatoo

ENTRYPOINT /bin/bash
