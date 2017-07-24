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
	CONF_FILE_SERVICENAME     = "filesystem"
	CONF_FILES_SERVICEFACTORY = "filesystemfactory"
	CONF_FILESDIR             = "filesdir"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_FILE_SERVICENAME, Object: FileSystemSvc{}},
		core.PluginComponent{Name: CONF_FILES_SERVICEFACTORY, ObjectCreator: core.NewFactory(func() interface{} { return &FileSystemSvc{} })}}
}

type FileSystemSvc struct {
	core.Service
	filesDir string
}

func (svc *FileSystemSvc) Invoke(ctx core.RequestContext, req core.Request) (*core.Response, error) {
	log.Info(ctx, "writing file")
	files := *req.GetBody().(*map[string]*core.MultipartFile)
	urls := make([]string, len(files))
	i := 0
	for _, fil := range files {
		defer fil.File.Close()
		fileName := uuid.NewV4().String()
		log.Info(ctx, "writing file", "name", fileName, "mimetype", fil.MimeType)
		url, err := svc.SaveFile(ctx, fil.File, fileName, fil.MimeType)
		if err != nil {
			return nil, err
		}
		urls[i] = url
		i++
	}
	log.Info(ctx, "writing file", "urls", urls)
	return core.NewServiceResponse(core.StatusSuccess, urls, nil), nil
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

func (svc *FileSystemSvc) ServeFile(ctx core.RequestContext, fileName string) (*core.Response, error) {
	return core.NewServiceResponse(core.StatusServeFile, svc.GetFullPath(ctx, fileName), nil), nil
}

func (svc *FileSystemSvc) GetFullPath(ctx core.RequestContext, fileName string) string {
	log.Error(ctx, "Full Path:***********", "path", svc.filesDir+fileName)
	return svc.filesDir + fileName
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

func (svc *FileSystemSvc) Initialize(ctx core.ServerContext) error {
	svc.SetDescription("File system storage service")
	svc.SetRequestType(config.CONF_OBJECT_STRINGMAP, false, false)
	svc.AddStringConfigurations([]string{CONF_FILESDIR}, nil)
	return nil
}

func (svc *FileSystemSvc) Start(ctx core.ServerContext) error {
	filesDir, _ := svc.GetConfiguration(CONF_FILESDIR)
	svc.filesDir = filesDir.(string)

	return nil
}
