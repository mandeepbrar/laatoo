#!/bin/bash

#cache services
go build -buildmode=plugin -o dist/lib/cache/memorycache.so services/cache/memory/memorycacheservice.go
go build -buildmode=plugin -o dist/lib/cache/appenginecache.so services/cache/appengine/appenginecache.go
go build -buildmode=plugin -o dist/lib/cache/rediscache.so services/cache/redis/rediscacheservice.go

#data services
go build -buildmode=plugin -o dist/lib/data/dataadapter.so services/data/dataadapter/dataadapter.go
go build -buildmode=plugin -o dist/lib/data/mongo.so services/data/mongo/*.go
go build -buildmode=plugin -o dist/lib/data/sql.so services/data/sql/*.go
go build -buildmode=plugin -o dist/lib/data/gae.so services/data/gae/*.go
go build -buildmode=plugin -o dist/lib/data/plugins.so services/data/plugins/*.go
go build -buildmode=plugin -o dist/lib/data/customloader.so services/data/customloader/*.go


#pubsub services
go build -buildmode=plugin -o dist/lib/pubsub/redis.so services/pubsub/redis/*.go

#search services
go build -buildmode=plugin -o dist/lib/search/bleve.so services/search/bleve/*.go
go build -buildmode=plugin -o dist/lib/search/google.so services/search/google/*.go

#task services
go build -buildmode=plugin -o dist/lib/tasks/beanstalk.so services/tasks/beanstalk/*.go
go build -buildmode=plugin -o dist/lib/tasks/gae.so services/tasks/gae/*.go

#storage
go build -buildmode=plugin -o dist/lib/storage/googlestorage.so services/storage/googlestorage/*.go
go build -buildmode=plugin -o dist/lib/storage/filestorage.so services/storage/filesystem/*.go

#static content
go build -buildmode=plugin -o dist/lib/static/staticcontent.so services/static/*.go

#Security
go build -buildmode=plugin -o dist/lib/security/misc.so services/security/misc/*.go
go build -buildmode=plugin -o dist/lib/security/db.so services/security/db/*.go
go build -buildmode=plugin -o dist/lib/security/keyauth.so services/security/keyauth/*.go
go build -buildmode=plugin -o dist/lib/security/oauth.so services/security/oauth/*.go


#middleware
go build -buildmode=plugin -o dist/lib/middleware/basicmiddleware.so middleware/*.go
