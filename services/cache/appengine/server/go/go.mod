module appenginecacheservice

go 1.12

require (
	cachecommon v0.0.0
	github.com/ugorji/go v1.1.4 // indirect
	google.golang.org/appengine v1.5.0
	laatoo/sdk v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace cachecommon => /modulesrepo/laatoo/services/cache/common
