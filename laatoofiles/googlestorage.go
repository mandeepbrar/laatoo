// +build appengine

package laatoofiles

import (
	"github.com/labstack/echo"
	"google.golang.org/cloud/storage"
	"io"
	"laatoocore"
	"laatoosdk/context"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"net/http"
)

const (
	CONF_GS_SERVICENAME = "googlestorage_file_service"
	CONF_GS_FILESBUCKET = "bucket"
)

type GoogleStorageService struct {
	bucket string
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_GS_SERVICENAME, GoogleStorageServiceFactory)
}

//factory method returns the service object to the environment
func GoogleStorageServiceFactory(ctx interface{}, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, LOGGING_CONTEXT, "Creating google storage file service")
	svc := &GoogleStorageService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, FILE_ERROR_MISSING_ROUTER)
	}
	bucketInt, ok := conf[CONF_GS_FILESBUCKET]
	if !ok {
		return nil, errors.ThrowError(ctx, FILE_ERROR_MISSING_BUCKET)
	}
	svc.bucket = bucketInt.(string)
	router := routerInt.(*echo.Group)
	router.Post("", svc.processFile)
	return svc, nil
}

//Provides the name of the service
func (svc *GoogleStorageService) GetName() string {
	return CONF_GS_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *GoogleStorageService) Initialize(ctx service.ServiceContext) error {
	return nil
}

//The service starts serving when this method is called
func (svc *GoogleStorageService) Serve(ctx interface{}) error {
	return nil
}

//Type of service
func (svc *GoogleStorageService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *GoogleStorageService) Execute(ctx interface{}, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}

func (svc *GoogleStorageService) processFile(ctx *echo.Context) error {
	req := ctx.Request()
	err := req.ParseMultipartForm(16 << 20) // Max memory 16 MiB
	if err != nil {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "Error while parsing multipart form", "Error", err)
		return err
	}
	cloudCtx := context.GetCloudContext(ctx, storage.ScopeFullControl)
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

		dst := storage.NewWriter(cloudCtx, svc.bucket, f.Filename)

		dst.ContentType = f.Header.Get("Content-Type")
		dst.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Copying file", "Name", f.Filename)

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		if err = dst.Close(); err != nil {
			return err
		}
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "updated object:", "Object", dst.Object().MediaLink)
		url[index] = dst.Object().MediaLink
	}
	return ctx.JSON(http.StatusOK, url)
}
