package main

import (
	"fmt"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	common "storagecommon"

	minio "github.com/minio/minio-go/v6"
)

//"golang.org/x/oauth2"
//"golang.org/x/oauth2/google"
//"net/http"
//"strings"

const (
	CONF_S3STORAGE_SERVICENAME  = "s3storage"
	CONF_S3S_FILESDEFAULTBUCKET = "s3storagedefaultbucket"
	CONF_S3S_PUBLICFILE         = "public"
	CONF_S3S_HOST               = "host"
	CONF_S3S_ENDPOINT           = "endpoint"
	CONF_S3S_SSL                = "ssl"
	CONF_S3S_ACCESSKEY          = "accesskey"
	CONF_S3S_ACCESSID           = "accessid"
	CONF_S3S_LOCATION           = "location"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_S3STORAGE_SERVICENAME, Object: S3StorageSvc{}}}
}

type S3StorageSvc struct {
	core.Service
	endpoint      string
	accessid      string
	accesskey     string
	defaultBucket string
	ssl           bool
	public        bool
	host          string
	location      string
	minioClient   *minio.Client
}

func (svc *S3StorageSvc) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.defaultBucket, _ = svc.GetStringConfiguration(ctx, CONF_S3S_FILESDEFAULTBUCKET)
	svc.public, _ = svc.GetBoolConfiguration(ctx, CONF_S3S_PUBLICFILE)
	svc.endpoint, _ = svc.GetStringConfiguration(ctx, CONF_S3S_ENDPOINT)
	svc.host, _ = svc.GetStringConfiguration(ctx, CONF_S3S_HOST)
	svc.ssl, _ = svc.GetBoolConfiguration(ctx, CONF_S3S_SSL)

	accessid, _ := svc.GetStringConfiguration(ctx, CONF_S3S_ACCESSID)
	accID, ok := svc.GetSecretConfiguration(ctx, accessid)
	if !ok {
		return errors.BadConf(ctx, CONF_S3S_ACCESSID, "Error", "Value not found in secret store", "key", accessid)
	}
	svc.accessid = string(accID)

	accesskey, _ := svc.GetStringConfiguration(ctx, CONF_S3S_ACCESSKEY)
	accKey, ok := svc.GetSecretConfiguration(ctx, accesskey)
	if !ok {
		return errors.BadConf(ctx, CONF_S3S_ACCESSID, "Error", "Value not found in secret store", "key", accesskey)
	}
	svc.accesskey = string(accKey)

	svc.location, _ = svc.GetStringConfiguration(ctx, CONF_S3S_LOCATION)
	var err error
	svc.minioClient, err = svc.createClient(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *S3StorageSvc) Invoke(ctx core.RequestContext) error {
	bucket, _ := ctx.GetStringParam("bucket")
	err := common.SaveFiles(ctx, svc, svc.bucket(bucket))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *S3StorageSvc) createClient(ctx ctx.Context) (*minio.Client, error) {
	// Initialize minio client object.
	return minio.New(svc.endpoint, svc.accessid, svc.accesskey, svc.ssl)
}

func (svc *S3StorageSvc) bucket(bucketName string) string {
	if bucketName == "" {
		return svc.defaultBucket
	}
	return bucketName
}

func (svc *S3StorageSvc) CreateFile(ctx core.RequestContext, bucket, fileName string, contentType string) (io.WriteCloser, error) {
	log.Debug(ctx, "Creating file", "name", fileName)

	pipeReader, pipeWriter := io.Pipe()

	_, err := svc.SaveFile(ctx, bucket, pipeReader, fileName, contentType)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	return pipeWriter, nil
}

func (svc *S3StorageSvc) Exists(ctx core.RequestContext, bucket, fileName string) bool {
	_, err := svc.minioClient.StatObject(svc.bucket(bucket), fileName, minio.StatObjectOptions{})
	if err == nil {
		return true
	}
	return false
}

func (svc *S3StorageSvc) Open(ctx core.RequestContext, bucket, fileName string) (io.ReadCloser, error) {
	object, err := svc.minioClient.GetObjectWithContext(ctx, svc.bucket(bucket), fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	return object, nil
}

func (svc *S3StorageSvc) ServeFile(ctx core.RequestContext, bucket, fileName string) error {
	str, err := svc.Open(ctx, bucket, fileName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	ctx.SetResponse(core.NewServiceResponseWithInfo(core.StatusServeStream, str, nil))
	return nil
}

func (svc *S3StorageSvc) CopyFile(ctx core.RequestContext, bucket, fileName string, dest io.WriteCloser) error {
	err := common.CopyFile(ctx, svc, svc.bucket(bucket), fileName, dest)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *S3StorageSvc) GetFullPath(ctx core.RequestContext, bucket, fileName string) string {
	if svc.public {
		return fmt.Sprintf("https://%/%s/%s", svc.host, svc.bucket(bucket), fileName)
	}
	/*
		return fmt.Sprintf("https://storage.cloud.google.com/%s/%s", svc.bucket, fileName)*/
	return ""
}

func (svc *S3StorageSvc) DeleteFiles(ctx core.RequestContext, bucket string, fileName string) (bool, error) {
	err := svc.minioClient.RemoveObject(svc.bucket(bucket), fileName)
	if err != nil {
		return false, errors.WrapError(ctx, err)
	}
	return true, nil
}

func (svc *S3StorageSvc) ListFiles(ctx core.RequestContext, bucket, pattern string) ([]string, error) {
	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := false
	files := []string{}
	objectCh := svc.minioClient.ListObjectsV2(bucket, pattern, isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			return nil, object.Err
		}
		files = append(files, object.Key)
	}
	return files, nil
}

func (svc *S3StorageSvc) SaveFile(ctx core.RequestContext, bucket string, inpStr io.ReadCloser, fileName string, contentType string) (string, error) {

	bytsWritten, err := svc.minioClient.PutObjectWithContext(ctx, svc.bucket(bucket), fileName, inpStr, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}
	log.Debug(ctx, "Copying complete", "Filename", fileName, "bytes", bytsWritten)
	return fileName, nil
}

func (svc *S3StorageSvc) CreateBucket(ctx core.RequestContext, bucket string) error {
	err := svc.minioClient.MakeBucket(bucket, svc.location)
	if err != nil {
		errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *S3StorageSvc) DeleteBucket(ctx core.RequestContext, bucket string) error {
	err := svc.minioClient.RemoveBucket(bucket)
	if err != nil {
		errors.WrapError(ctx, err)
	}
	return nil
}
