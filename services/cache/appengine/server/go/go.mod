module appenginecacheservice

go 1.12

require (
	cachecommon v0.0.0
	github.com/ugorji/go v1.1.4 // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	google.golang.org/appengine v1.5.0
	laatoo/sdk v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace cachecommon => /modulesrepo/laatoo/services/cache/common
