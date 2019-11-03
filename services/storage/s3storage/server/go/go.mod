module s3storageservice

go 1.12

require (
	github.com/google/btree v1.0.0 // indirect
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/minio/minio-go/v6 v6.0.40
	github.com/twinj/uuid v1.0.0
	laatoo/sdk v0.0.0
	storagecommon v0.0.0
)

replace laatoo/sdk => /laatoo/sdk

replace storagecommon => /modulesrepo/laatoo/services/storage/common
