module dbloginservice

go 1.12

require (
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/google/btree v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472
	laatoo/sdk v0.0.0
	securitycommon v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace securitycommon => /modulesrepo/laatoo/services/security/common
