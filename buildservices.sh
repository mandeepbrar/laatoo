#!/bin/bash

rm -fr dist/lib/
mkdir -p dist/lib/laatoo/server


#cache services
go build -buildmode=plugin -o dist/lib/laatoo/server/cache/memorycache.so services/cache/memory/memorycacheservice.go
go build -buildmode=plugin -o dist/lib/laatoo/server/cache/appenginecache.so services/cache/appengine/appenginecache.go
go build -buildmode=plugin -o dist/lib/laatoo/server/cache/rediscache.so services/cache/redis/rediscacheservice.go

#data services
go build -buildmode=plugin -o dist/lib/laatoo/server/data/dataadapter.so services/data/dataadapter/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/data/mongo.so services/data/mongo/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/data/sql.so services/data/sql/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/data/gae.so services/data/gae/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/data/plugins.so services/data/plugins/*.go


#pubsub services
go build -buildmode=plugin -o dist/lib/laatoo/server/pubsub/redis.so services/pubsub/redis/*.go

#search services
go build -buildmode=plugin -o dist/lib/laatoo/server/search/bleve.so services/search/bleve/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/search/google.so services/search/google/*.go

#task services
go build -buildmode=plugin -o dist/lib/laatoo/server/tasks/beanstalk.so services/tasks/beanstalk/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/tasks/gae.so services/tasks/gae/*.go

#storage
go build -buildmode=plugin -o dist/lib/laatoo/server/storage/googlestorage.so services/storage/googlestorage/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/storage/filestorage.so services/storage/filesystem/*.go

#static content
go build -buildmode=plugin -o dist/lib/laatoo/server/static/staticcontent.so services/static/*.go

#Security
go build -buildmode=plugin -o dist/lib/laatoo/server/security/misc.so services/security/misc/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/security/db.so services/security/db/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/security/keyauth.so services/security/keyauth/*.go
go build -buildmode=plugin -o dist/lib/laatoo/server/security/oauth.so services/security/oauth/*.go


#middleware
go build -buildmode=plugin -o dist/lib/laatoo/server/middleware/basicmiddleware.so middleware/*.go
