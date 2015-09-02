package laatoofiles

import (
	"github.com/labstack/echo"
	"io"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
	"os"
)

const (
	CONF_FILE_SERVICENAME = "file_service"
	CONF_FILE_FILESDIR    = "filesdir"
	CONF_FILE_FILESURL    = "filesurl"
)

type FileService struct {
	filesDir string
	filesUrl string
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_FILE_SERVICENAME, FileServiceFactory)
}

//factory method returns the service object to the environment
func FileServiceFactory(conf map[string]interface{}) (interface{}, error) {
	log.Logger.Infof("Creating file service")
	svc := &FileService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(FILE_ERROR_MISSING_ROUTER)
	}
	filesdirInt, ok := conf[CONF_FILE_FILESDIR]
	if !ok {
		return nil, errors.ThrowError(FILE_ERROR_MISSING_FILEDIR)
	}
	filesurlInt, ok := conf[CONF_FILE_FILESURL]
	if !ok {
		return nil, errors.ThrowError(FILE_ERROR_MISSING_FILESURL)
	}

	svc.filesUrl = filesurlInt.(string)
	router := routerInt.(*echo.Group)
	svc.filesDir = filesdirInt.(string) + "/"
	log.Logger.Infof("Files Directory %s", filesdirInt)
	router.Post("", svc.processFile)
	return svc, nil
}

//Provides the name of the service
func (svc *FileService) GetName() string {
	return CONF_FILE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *FileService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *FileService) Serve() error {
	return nil
}

//Type of service
func (svc *FileService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *FileService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (svc *FileService) processFile(ctx *echo.Context) error {
	req := ctx.Request()

	err := req.ParseMultipartForm(16 << 20) // Max memory 16 MiB
	if err != nil {
		log.Logger.Debugf("Error", err)
		return err
	}

	// Read form fields
	//name := c.Form("name")
	//email := c.Form("email")

	// Read files
	log.Logger.Infof("Request", req)
	log.Logger.Infof("form", req.MultipartForm)

	files := req.MultipartForm.File["file"]
	log.Logger.Infof("files", len(files))

	url := make([]string, len(files))
	for index, f := range files {
		log.Logger.Infof("index", index)
		// Source file
		src, err := f.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		fileName := svc.filesDir + f.Filename
		log.Logger.Debugf("Writing file", fileName)
		// Destination file
		dst, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer dst.Close()
		log.Logger.Debugf("Copying file", fileName)

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		url[index] = svc.filesUrl + "/" + f.Filename
	}
	return ctx.JSON(http.StatusOK, url)
}
