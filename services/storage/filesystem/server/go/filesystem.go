package main

import (
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/log"
	"os"
	"path/filepath"

	"github.com/twinj/uuid"
	//"net/http"
	"path"
	//"strings"
)

const (
	CONF_FILE_SERVICENAME = "filesystem"
	CONF_FILESDIR         = "filestoragedir"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_FILE_SERVICENAME, Object: FileSystemSvc{}}}
}

type FileSystemSvc struct {
	core.Service
	filesDir string
}

func (svc *FileSystemSvc) Invoke(ctx core.RequestContext) error {
	log.Info(ctx, "writing file")
	val, _ := ctx.GetParamValue("Data")
	log.Info(ctx, "File System Service", "Data", val)
	files := val.(map[string]*core.MultipartFile)
	urls := map[string]string{}
	i := 0
	for filNam, fil := range files {
		defer fil.File.Close()
		fileName := uuid.NewV4().String()
		log.Info(ctx, "writing file", "name", fileName, "mimetype", fil.MimeType)
		url, err := svc.SaveFile(ctx, fil.File, fileName, fil.MimeType)
		if err != nil {
			return err
		}
		urls[filNam] = url
		i++
	}
	log.Info(ctx, "writing file", "urls", urls)
	ctx.SetResponse(core.SuccessResponse(urls))
	return nil
}
func (svc *FileSystemSvc) CreateFile(ctx core.RequestContext, fileName string, contentType string) (io.WriteCloser, error) {
	fullpath := svc.GetFullPath(ctx, fileName)
	destdir, _ := path.Split(fullpath)
	os.MkdirAll(destdir, 0755)
	return os.Create(fullpath)
}

func (svc *FileSystemSvc) Exists(ctx core.RequestContext, fileName string) bool {
	fullpath := svc.GetFullPath(ctx, fileName)
	_, err := os.Stat(fullpath)
	if err == nil || os.IsExist(err) {
		return true
	}
	return false
}
func (svc *FileSystemSvc) Open(ctx core.RequestContext, fileName string) (io.ReadCloser, error) {
	fullpath := svc.GetFullPath(ctx, fileName)
	return os.Open(fullpath)
}

func (svc *FileSystemSvc) ServeFile(ctx core.RequestContext, fileName string) error {
	path := svc.GetFullPath(ctx, fileName)
	log.Trace(ctx, "Serving file", "filename", path)
	ctx.SetResponse(core.NewServiceResponseWithInfo(core.StatusServeFile, path, nil))
	return nil
}

func (svc *FileSystemSvc) CopyFile(ctx core.RequestContext, fileName string, dest io.WriteCloser) error {
	// Source file
	src, err := svc.Open(ctx, fileName)
	if err != nil {
		return err
	}
	defer src.Close()

	if _, err = io.Copy(dest, src); err != nil {
		return err
	}
	return nil
}

func (svc *FileSystemSvc) GetFullPath(ctx core.RequestContext, fileName string) string {
	return path.Join(svc.filesDir, fileName)
}

func (svc *FileSystemSvc) SaveFile(ctx core.RequestContext, inpStr io.ReadCloser, fileName string, contentType string) (string, error) {
	// Destination file
	dst, err := svc.CreateFile(ctx, fileName, contentType)
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

func (svc *FileSystemSvc) ListFiles(ctx core.RequestContext, pattern string) ([]string, error) {
	return filepath.Glob(path.Join(svc.filesDir, pattern))
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
	return nil
}
