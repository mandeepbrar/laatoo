module oauthservice

go 1.12

require (
	github.com/imdario/mergo v0.3.7 // indirect
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a

	laatoo/sdk v0.0.0

	securitycommon v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace securitycommon => /modulesrepo/laatoo/services/security/common
