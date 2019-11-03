module tokenvalidation

go 1.12

require (
	github.com/google/btree v1.0.0 // indirect
	laatoo/sdk v0.0.0

	securitycommon v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace securitycommon => /modulesrepo/laatoo/services/security/common
