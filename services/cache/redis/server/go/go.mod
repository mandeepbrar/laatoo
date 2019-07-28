module redisservice

go 1.12

require (
	cachecommon v0.0.0
	github.com/garyburd/redigo v1.6.0
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/ugorji/go v1.1.4 // indirect

	laatoo/sdk v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace cachecommon => /modulesrepo/laatoo/services/cache/common
