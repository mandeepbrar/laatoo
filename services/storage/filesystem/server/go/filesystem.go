package main

import (
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"os"
	"path/filepath"
	common "storagecommon"

	//"net/http"
	"path"
	//"strings"
)

const (
	CONF_FILE_SERVICENAME   = "filesystem"
	CONF_FILE_DEFAULTBUCKET = "defaultbucket"
	CONF_FILESDIR           = "filestoragedir"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_FILE_SERVICENAME, Object: FileSystemSvc{}}}
}

type FileSystemSvc struct {
	core.Service
	filesDir      string
	defaultBucket string
}

func (svc *FileSystemSvc) Invoke(ctx core.RequestContext) error {
	bucket, _ := ctx.GetStringParam("bucket")
	err := common.SaveFiles(ctx, svc, svc.bucket(bucket))
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *FileSystemSvc) bucket(bucketName string) string {
	if bucketName == "" {
		return svc.defaultBucket
	}
	return bucketName
}

func (svc *FileSystemSvc) CreateFile(ctx core.RequestContext, bucket, fileName string, contentType string) (io.WriteCloser, error) {
	fullpath := svc.GetFullPath(ctx, svc.bucket(bucket), fileName)
	destdir, _ := path.Split(fullpath)
	os.MkdirAll(destdir, 0755)
	return os.Create(fullpath)
}

func (svc *FileSystemSvc) Exists(ctx core.RequestContext, bucket, fileName string) bool {
	fullpath := svc.GetFullPath(ctx, bucket, fileName)
	_, err := os.Stat(fullpath)
	if err == nil || os.IsExist(err) {
		return true
	}
	return false
}
func (svc *FileSystemSvc) Open(ctx core.RequestContext, bucket, fileName string) (io.ReadCloser, error) {
	fullpath := svc.GetFullPath(ctx, bucket, fileName)
	return os.Open(fullpath)
}

func (svc *FileSystemSvc) OpenForWrite(ctx core.RequestContext, bucket, fileName string) (io.WriteCloser, error) {
	fullpath := svc.GetFullPath(ctx, bucket, fileName)
	return os.Open(fullpath)
}

func (svc *FileSystemSvc) ServeFile(ctx core.RequestContext, bucket, fileName string) error {
	path := svc.GetFullPath(ctx, bucket, fileName)
	log.Trace(ctx, "Serving file", "filename", path)
	ctx.SetResponse(core.NewServiceResponseWithInfo(core.StatusServeFile, path, nil))
	return nil
}

func (svc *FileSystemSvc) DeleteFiles(ctx core.RequestContext, bucket string, fileName string) (bool, error) {
	return false, nil
}

func (svc *FileSystemSvc) CopyFile(ctx core.RequestContext, bucket, fileName string, dest io.WriteCloser) error {
	err := common.CopyFile(ctx, svc, bucket, fileName, dest)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *FileSystemSvc) GetFullPath(ctx core.RequestContext, bucket, fileName string) string {
	return path.Join(svc.filesDir, fileName)
}

func (svc *FileSystemSvc) SaveFile(ctx core.RequestContext, bucket string, inpStr io.ReadCloser, fileName string, contentType string) (string, error) {
	// Destination file
	dst, err := svc.CreateFile(ctx, bucket, fileName, contentType)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, inpStr); err != nil {
		return "", err
	}
	inpStr.Close()
	return fileName, nil
}

func (svc *FileSystemSvc) ListFiles(ctx core.RequestContext, bucket, pattern string) ([]string, error) {
	return filepath.Glob(path.Join(svc.filesDir, pattern))
}

func (svc *FileSystemSvc) CreateBucket(ctx core.RequestContext, bucket string) error {
	dirToCreate := path.Join(svc.filesDir, bucket)
	return os.Mkdir(dirToCreate, 0744)
}

func (svc *FileSystemSvc) DeleteBucket(ctx core.RequestContext, bucket string) error {
	dirToRemove := path.Join(svc.filesDir, bucket)
	return os.RemoveAll(dirToRemove)
}

func (svc *FileSystemSvc) Initialize(ctx core.ServerContext, conf config.Config) error {
	/*svc.SetDescription(ctx, "File system storage service")
	svc.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
	svc.AddStringConfigurations(ctx, []string{CONF_FILESDIR}, nil)*/

	filesDir, _ := svc.GetStringConfiguration(ctx, CONF_FILESDIR)
	if filepath.IsAbs(filesDir) {
		svc.filesDir = filesDir
	} else {
		baseDir, ok := ctx.GetString(config.MODULEDIR)
		if !ok {
			baseDir, _ = ctx.GetString(config.BASEDIR)
		}
		if filesDir != "" {
			svc.filesDir = path.Join(baseDir, "files", filesDir)
		} else {
			svc.filesDir = path.Join(baseDir, "files")
		}
	}
	svc.defaultBucket, _ = svc.GetStringConfiguration(ctx, CONF_FILE_DEFAULTBUCKET)
	return nil
}
