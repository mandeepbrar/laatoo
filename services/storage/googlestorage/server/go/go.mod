module googlestorage

go 1.12

require (
	cloud.google.com/go v0.38.0
	github.com/twinj/uuid v1.0.0
	laatoo/sdk v0.0.0
	storagecommon v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace storagecommon => /modulesrepo/laatoo/services/storage/common
