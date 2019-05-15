module dbloginservice

go 1.12

require (
	golang.org/x/crypto v0.0.0-20190510104115-cbcb75029529
	laatoo/sdk v0.0.0
	securitycommon v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace securitycommon => /modulesrepo/laatoo/services/security/common
