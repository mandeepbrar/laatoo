// +build !appengine

package laatoofiles

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/twinj/uuid"
	"io"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"laatoosdk/utils"
	"net/http"
	"os"
	"strings"
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
func FileServiceFactory(ctx *echo.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating file service")
	svc := &FileService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, FILE_ERROR_MISSING_ROUTER)
	}
	filesdirInt, ok := conf[CONF_FILE_FILESDIR]
	if !ok {
		return nil, errors.ThrowError(ctx, FILE_ERROR_MISSING_FILEDIR)
	}
	filesurlInt, ok := conf[CONF_FILE_FILESURL]
	if !ok {
		return nil, errors.ThrowError(ctx, FILE_ERROR_MISSING_FILESURL)
	}

	svc.filesUrl = filesurlInt.(string) + "/"
	router := routerInt.(*echo.Group)
	svc.filesDir = filesdirInt.(string) + "/"
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Got files directory", "Name", filesdirInt)
	router.Post("", svc.processFile)
	return svc, nil
}

//Provides the name of the service
func (svc *FileService) GetName() string {
	return CONF_FILE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *FileService) Initialize(ctx *echo.Context) error {
	return nil
}

//The service starts serving when this method is called
func (svc *FileService) Serve(ctx *echo.Context) error {
	return nil
}

//Type of service
func (svc *FileService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *FileService) Execute(ctx *echo.Context, name string, params map[string]interface{}) (interface{}, error) {
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "here1", "name", name)
	if name == "CopyFile" {
		return nil, svc.copyFile(ctx, params["filename"].(string), params["writer"].(io.Writer))
	}
	if name == "TransformFile" {
		return svc.transformFile(ctx, params["srcpath"].(string), params["destfolder"].(string), params["transformation"].(utils.FileTransform))
	}
	return nil, nil
}

func (svc *FileService) transformFile(ctx *echo.Context, srcpath string, destfolder string, transform utils.FileTransform) (string, error) {
	pathinfolder, realsrcpath := svc.parsePath(srcpath)
	destfile := fmt.Sprintf("%s%s/%s", svc.filesDir, destfolder, pathinfolder)
	request := ctx.Request()
	host := request.Host
	scheme := "https"
	if request.TLS == nil {
		scheme = "http"
	}
	desturl := fmt.Sprintf("%s://%s/%s%s/%s", scheme, host, svc.filesUrl, destfolder, pathinfolder)
	_, err := os.Stat(destfile)
	if os.IsExist(err) {
		return desturl, nil
	}
	log.Logger.Info(ctx, LOGGING_CONTEXT, "file does not exist", "destfile", destfile, "realsrcpath", realsrcpath)

	rd, err := os.Open(realsrcpath)
	defer rd.Close()
	if err != nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "error opening source file", "realsrcpath", realsrcpath, "err", err)
		return "", err
	}
	log.Logger.Info(ctx, LOGGING_CONTEXT, "opened src file", "destfile", destfile, "realsrcpath", realsrcpath)
	os.MkdirAll(fmt.Sprintf("%s%s", svc.filesDir, destfolder), 0755)
	writer, err := os.Create(destfile)
	defer writer.Close()
	if err != nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "error creating file", "destfile", destfile, "err", err)
		return "", err
	}
	log.Logger.Info(ctx, LOGGING_CONTEXT, "transform", "destfile", destfile, "realsrcpath", realsrcpath)

	err = transform(rd, writer)
	if err != nil {
		return "", err
	}
	return desturl, nil
}

func (svc *FileService) copyFile(ctx *echo.Context, fileurl string, writer io.Writer) error {
	_, realpath := svc.parsePath(fileurl)
	rd, err := os.Open(realpath)
	defer rd.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, rd)
	return err
}

func (svc *FileService) parsePath(url string) (string, string) {
	pathinfolder := strings.TrimPrefix(url, svc.filesUrl)
	return pathinfolder, fmt.Sprintf("%s%s", svc.filesDir, pathinfolder)
}

func (svc *FileService) processFile(ctx *echo.Context) error {
	req := ctx.Request()

	err := req.ParseMultipartForm(16 << 20) // Max memory 16 MiB
	if err != nil {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "Error while parsing multipart form", "Error", err)
		return err
	}

	// Read form fields
	//name := c.Form("name")
	//email := c.Form("email")

	// Read files

	files := req.MultipartForm.File["file"]
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Parsed multipart form", "Number of files", len(files))

	url := make([]string, len(files))
	for index, f := range files {
		// Source file
		src, err := f.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		fileName := uuid.NewV4().String()
		fullpath := svc.filesDir + fileName
		// Destination file
		dst, err := os.Create(fullpath)
		if err != nil {
			return err
		}
		defer dst.Close()
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Copying file", "Name", fileName)

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		url[index] = svc.filesUrl + fileName
	}
	return ctx.JSON(http.StatusOK, url)
}
