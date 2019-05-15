module redisservice

go 1.12

require (
	github.com/garyburd/redigo v1.6.0

	laatoo/sdk v0.0.0
	cachecommon v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace cachecommon => /modulesrepo/laatoo/services/cache/common
