package main

import (
	"io"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"os"

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
	files := *ctx.GetBody().(*map[string]*core.MultipartFile)
	urls := make([]string, len(files))
	i := 0
	for _, fil := range files {
		defer fil.File.Close()
		fileName := uuid.NewV4().String()
		log.Info(ctx, "writing file", "name", fileName, "mimetype", fil.MimeType)
		url, err := svc.SaveFile(ctx, fil.File, fileName, fil.MimeType)
		if err != nil {
			return err
		}
		urls[i] = url
		i++
	}
	log.Info(ctx, "writing file", "urls", urls)
	ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, urls, nil))
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
	ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, path, nil))
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

func (svc *FileSystemSvc) Initialize(ctx core.ServerContext, conf config.Config) error {
	/*svc.SetDescription(ctx, "File system storage service")
	svc.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
	svc.AddStringConfigurations(ctx, []string{CONF_FILESDIR}, nil)*/
	filesDir, ok := svc.GetStringConfiguration(ctx, CONF_FILESDIR)
	if ok {
		svc.filesDir = filesDir
	} else {
		baseDir, ok := ctx.GetString(config.MODULEDIR)
		if !ok {
			baseDir, _ = ctx.GetString(config.BASEDIR)
		}
		svc.filesDir = path.Join(baseDir, "files")
	}
	return nil
}
