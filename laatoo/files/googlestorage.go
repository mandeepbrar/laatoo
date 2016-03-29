// +build appengine

package laatoofiles

import (
	"fmt"
	"github.com/twinj/uuid"
	"google.golang.org/cloud/storage"
	"io"
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
	"net/http"
	"strings"
)

const (
	CONF_GS_SERVICENAME = "googlestorage_file_service"
	CONF_GS_FILESBUCKET = "bucket"
)

type GoogleStorageService struct {
	bucket string
	prefix string
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_GS_SERVICENAME, GoogleStorageServiceFactory)
}

//factory method returns the service object to the environment
func GoogleStorageServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
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
	svc.prefix = fmt.Sprintf("storage.cloud.google.com/%s/", svc.bucket)
	router := routerInt.(core.Router)
	router.Post(ctx, "", conf, svc.processFile)
	return svc, nil
}

//Provides the name of the service
func (svc *GoogleStorageService) GetName() string {
	return CONF_GS_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *GoogleStorageService) Initialize(ctx core.Context) error {
	return nil
}

//The service starts serving when this method is called
func (svc *GoogleStorageService) Serve(ctx core.Context) error {
	return nil
}

//Type of service
func (svc *GoogleStorageService) GetServiceType() string {
	return core.SERVICE_TYPE_WEB
}

//Execute method
func (svc *GoogleStorageService) Execute(ctx core.Context, name string, params map[string]interface{}) (interface{}, error) {
	if name == "CopyFile" {
		return nil, svc.copyFile(ctx, params["filename"].(string), params["writer"].(io.Writer))
	}
	if name == "TransformFile" {
		return svc.transformFile(ctx, params["srcpath"].(string), params["destfolder"].(string), params["transformation"].(utils.FileTransform))
	}
	return nil, nil
}

func (svc *GoogleStorageService) copyFile(ctx core.Context, srcpath string, writer io.Writer) error {
	filepath := svc.parsePath(srcpath)
	log.Logger.Debug(ctx, LOGGING_CONTEXT, "Copying file", filepath)
	cloudCtx := ctx.GetCloudContext(storage.ScopeFullControl)
	client, err := storage.NewClient(cloudCtx)
	if err != nil {
		return err
	}
	reader, err := client.Bucket(svc.bucket).Object(filepath).NewReader(cloudCtx)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, reader)
	return err
}

func (svc *GoogleStorageService) processFile(ctx core.Context) error {
	req := ctx.Request()
	err := req.ParseMultipartForm(16 << 20) // Max memory 16 MiB
	if err != nil {
		log.Logger.Debug(ctx, LOGGING_CONTEXT, "Error while parsing multipart form", "Error", err)
		return err
	}
	cloudCtx := ctx.GetCloudContext(storage.ScopeFullControl)
	client, err := storage.NewClient(cloudCtx)
	if err != nil {
		return err
	}

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

		dst := client.Bucket(svc.bucket).Object(fileName).NewWriter(cloudCtx)
		defer dst.Close()

		dst.ContentType = f.Header.Get("Content-Type")
		dst.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "Copying file", "Name", fileName)

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		if err = dst.Close(); err != nil {
			return err
		}
		returl := fmt.Sprintf("http://%s%s", svc.prefix, fileName)
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "updated object:", "Object", returl)
		url[index] = returl
	}
	return ctx.JSON(http.StatusOK, url)
}

func (svc *GoogleStorageService) transformFile(ctx core.Context, srcpath string, destfolder string, transform utils.FileTransform) (string, error) {
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "Transforming file", "srcpath", srcpath, "destfolder", destfolder)
	appEngineCtx := ctx.GetAppengineContext()
	client, err := storage.NewClient(appEngineCtx)
	if err != nil {
		return "", err
	}
	filepath := svc.parsePath(srcpath)
	log.Logger.Trace(ctx, LOGGING_CONTEXT, "filepath", "filepath", filepath)
	destfile := fmt.Sprintf("%s/%s", destfolder, filepath)
	desturl := fmt.Sprintf("http://%s%s/%s", svc.prefix, destfolder, filepath)

	objattrs, err := client.Bucket(svc.bucket).Object(destfile).Attrs(appEngineCtx)
	if objattrs != nil || err == nil {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "returning dest url ", "desturl", desturl)
		return desturl, nil
	}

	objattrs, err = client.Bucket(svc.bucket).Object(filepath).Attrs(appEngineCtx)
	if err != nil {
		log.Logger.Trace(ctx, LOGGING_CONTEXT, "could not stat object... ", "err", err, "filepath", filepath)
		return "", err
	}

	log.Logger.Trace(ctx, LOGGING_CONTEXT, "file does not exist... ", "destfile", destfile, "filepath", filepath)

	reader, err := client.Bucket(svc.bucket).Object(filepath).NewReader(appEngineCtx)
	defer reader.Close()
	if err != nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "error opening source file", "filepath", filepath, "err", err)
		return "", err
	}

	log.Logger.Trace(ctx, LOGGING_CONTEXT, "opened src file", "destfile", destfile, "filepath", filepath)

	writer := client.Bucket(svc.bucket).Object(destfile).NewWriter(appEngineCtx)
	defer writer.Close()
	writer.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
	writer.ContentType = objattrs.ContentType

	log.Logger.Trace(ctx, LOGGING_CONTEXT, "transform", "destfile", destfile, "filepath", filepath)

	err = transform(reader, writer)
	if err != nil {
		log.Logger.Info(ctx, LOGGING_CONTEXT, "Error in transformation", "destfile", destfile, "err", err)
		return "", err
	}
	return desturl, nil
}

func (svc *GoogleStorageService) parsePath(url string) string {
	indexOfPrefix := strings.Index(url, svc.prefix)
	if indexOfPrefix < 0 {
		return url
	}
	return url[indexOfPrefix+len(svc.prefix):]
}
