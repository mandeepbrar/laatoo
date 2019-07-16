package main

import (
	"fmt"
	"io"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"cloud.google.com/go/storage"
	"github.com/twinj/uuid"
)

//"golang.org/x/oauth2"
//"golang.org/x/oauth2/google"
//"net/http"
//"strings"

const (
	CONF_GOOGLESTORAGE_SERVICENAME = "googlestorage"
	CONF_GS_FILESBUCKET            = "googlestoragebucket"
	CONF_GS_PUBLICFILE             = "public"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_GOOGLESTORAGE_SERVICENAME, Object: GoogleStorageSvc{}}}
}

type GoogleStorageSvc struct {
	core.Service
	bucket string
	public bool
}

func (svc *GoogleStorageSvc) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.bucket, _ = svc.GetStringConfiguration(ctx, CONF_GS_FILESBUCKET)
	svc.public, _ = svc.GetBoolConfiguration(ctx, CONF_GS_PUBLICFILE)
	/*
		svc.SetDescription(ctx, "Google storage service")
		svc.SetRequestType(ctx, config.CONF_OBJECT_STRINGMAP, false, false)
		svc.AddStringConfigurations(ctx, []string{CONF_GS_FILESBUCKET}, nil)
		svc.AddOptionalConfigurations(ctx, map[string]string{CONF_GS_PUBLICFILE: config.CONF_OBJECT_BOOL}, map[string]interface{}{CONF_GS_PUBLICFILE: false})*/
	return nil
}

func (svc *GoogleStorageSvc) Invoke(ctx core.RequestContext) error {
	val, _ := ctx.GetParamValue("Data")
	files := *val.(*map[string]*core.MultipartFile)
	urls := map[string]string{}
	i := 0
	for _, fil := range files {
		defer fil.File.Close()
		fileName := uuid.NewV4().String()
		log.Debug(ctx, "writing file", "name", fileName, "MimeType", fil.MimeType)
		url, err := svc.SaveFile(ctx, fil.File, fileName, fil.MimeType)
		if err != nil {
			log.Debug(ctx, "Error while invoking upload", "err", err)
			return errors.WrapError(ctx, err)
		}
		urls[fileName] = url
		i++
	}
	ctx.SetResponse(core.SuccessResponse(urls))
	return nil
}

func (svc *GoogleStorageSvc) CreateFile(ctx core.RequestContext, fileName string, contentType string) (io.WriteCloser, error) {
	log.Debug(ctx, "Creating file", "name", fileName, "bucket", svc.bucket)

	appengineCtx := ctx.GetAppengineContext()
	client, err := storage.NewClient(appengineCtx)
	if err != nil {
		log.Debug(ctx, "Error while creating file", "err", err)
		return nil, errors.WrapError(ctx, err)
	}

	dst := client.Bucket(svc.bucket).Object(fileName).NewWriter(appengineCtx)
	if contentType != "" {
		dst.ContentType = contentType
	}

	dst.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	return dst, nil
}

func (svc *GoogleStorageSvc) Exists(ctx core.RequestContext, fileName string) bool {
	appengineCtx := ctx.GetAppengineContext()
	client, err := storage.NewClient(appengineCtx)
	if err != nil {
		log.Debug(ctx, "Error while creating file", "err", err)
		return false
	}
	defer client.Close()

	_, err = client.Bucket(svc.bucket).Object(fileName).Attrs(appengineCtx)
	if err == nil {
		return true
	}
	return false
}
func (svc *GoogleStorageSvc) Open(ctx core.RequestContext, fileName string) (io.ReadCloser, error) {
	appengineCtx := ctx.GetAppengineContext()
	client, err := storage.NewClient(appengineCtx)
	if err != nil {
		log.Debug(ctx, "Error while opening", "err", err)
		return nil, errors.WrapError(ctx, err)
	}
	return client.Bucket(svc.bucket).Object(fileName).NewReader(appengineCtx)
}

func (svc *GoogleStorageSvc) ServeFile(ctx core.RequestContext, fileName string) error {
	ctx.SetResponse(core.NewServiceResponseWithInfo(core.StatusRedirect, svc.GetFullPath(ctx, fileName), nil))
	return nil
}

func (svc *GoogleStorageSvc) CopyFile(ctx core.RequestContext, fileName string, dest io.WriteCloser) error {
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

func (svc *GoogleStorageSvc) GetFullPath(ctx core.RequestContext, fileName string) string {
	if svc.public {
		return fmt.Sprintf("https://storage.googleapis.com/%s/%s", svc.bucket, fileName)
	}
	return fmt.Sprintf("https://storage.cloud.google.com/%s/%s", svc.bucket, fileName)
}

func (svc *GoogleStorageSvc) ListFiles(ctx core.RequestContext, pattern string) ([]string, error) {
	/******TODO**********/
	//log.Debug(ctx, "Creating file", "name", fileName, "bucket", svc.bucket)
	/*
		appengineCtx := ctx.GetAppengineContext()
		client, err := storage.NewClient(appengineCtx)
		if err != nil {
			log.Debug(ctx, "Error while creating file", "err", err)
			return nil, errors.WrapError(ctx, err)
		}

		dst := client.Bucket(svc.bucket).Object(fileName).NewWriter(appengineCtx)
		if contentType != "" {
			dst.ContentType = contentType
		}
		dst.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
		return dst, nil*/
	return nil, nil

}

func (svc *GoogleStorageSvc) SaveFile(ctx core.RequestContext, inpStr io.ReadCloser, fileName string, contentType string) (string, error) {
	log.Debug(ctx, "Saving file", "name", fileName)
	// Destination file
	dst, err := svc.CreateFile(ctx, fileName, contentType)
	if err != nil {
		log.Debug(ctx, "Error while opening", "err", err)
		return "", errors.WrapError(ctx, err)
	}
	defer dst.Close()

	numbytes, err := io.Copy(dst, inpStr)

	if err != nil {
		log.Debug(ctx, "Error while saving", "err", err)
		return "", errors.WrapError(ctx, err)
	}
	dst.Close()
	inpStr.Close()
	log.Debug(ctx, "Copying complete", "Filename", fileName, "bucket", svc.bucket, "bytes", numbytes)
	return fileName, nil
}

/*
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



func (svc *GoogleStorageService) copyFile(ctx core.Context, srcpath string, writer io.Writer) error {
	filepath := svc.parsePath(srcpath)
	log.Debug(ctx, LOGGING_CONTEXT, "Copying file", filepath)
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
		log.Debug(ctx, LOGGING_CONTEXT, "Error while parsing multipart form", "Error", err)
		return err
	}
	cloudCtx := ctx.GetCloudContext(storage.ScopeFullControl)
	client, err := storage.NewClient(cloudCtx)
	if err != nil {
		return err
	}

	// Read files
	files := req.MultipartForm.File["file"]
	log.Debug(ctx, LOGGING_CONTEXT, "Parsed multipart form", "Number of files", len(files))

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
		log.Trace(ctx, LOGGING_CONTEXT, "Copying file", "Name", fileName)

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		if err = dst.Close(); err != nil {
			return err
		}
		returl := fmt.Sprintf("http://%s%s", svc.prefix, fileName)
		log.Trace(ctx, LOGGING_CONTEXT, "updated object:", "Object", returl)
		url[index] = returl
	}
	return ctx.JSON(http.StatusOK, url)
}

func (svc *GoogleStorageService) transformFile(ctx core.Context, srcpath string, destfolder string, transform utils.FileTransform) (string, error) {
	log.Trace(ctx, LOGGING_CONTEXT, "Transforming file", "srcpath", srcpath, "destfolder", destfolder)
	appEngineCtx := ctx.GetAppengineContext()
	client, err := storage.NewClient(appEngineCtx)
	if err != nil {
		return "", err
	}
	filepath := svc.parsePath(srcpath)
	log.Trace(ctx, LOGGING_CONTEXT, "filepath", "filepath", filepath)
	destfile := fmt.Sprintf("%s/%s", destfolder, filepath)
	desturl := fmt.Sprintf("http://%s%s/%s", svc.prefix, destfolder, filepath)

	objattrs, err := client.Bucket(svc.bucket).Object(destfile).Attrs(appEngineCtx)
	if objattrs != nil || err == nil {
		log.Trace(ctx, LOGGING_CONTEXT, "returning dest url ", "desturl", desturl)
		return desturl, nil
	}

	objattrs, err = client.Bucket(svc.bucket).Object(filepath).Attrs(appEngineCtx)
	if err != nil {
		log.Trace(ctx, LOGGING_CONTEXT, "could not stat object... ", "err", err, "filepath", filepath)
		return "", err
	}

	log.Trace(ctx, LOGGING_CONTEXT, "file does not exist... ", "destfile", destfile, "filepath", filepath)

	reader, err := client.Bucket(svc.bucket).Object(filepath).NewReader(appEngineCtx)
	defer reader.Close()
	if err != nil {
		log.Info(ctx, LOGGING_CONTEXT, "error opening source file", "filepath", filepath, "err", err)
		return "", err
	}

	log.Trace(ctx, LOGGING_CONTEXT, "opened src file", "destfile", destfile, "filepath", filepath)

	writer := client.Bucket(svc.bucket).Object(destfile).NewWriter(appEngineCtx)
	defer writer.Close()
	writer.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
	writer.ContentType = objattrs.ContentType

	log.Trace(ctx, LOGGING_CONTEXT, "transform", "destfile", destfile, "filepath", filepath)

	err = transform(reader, writer)
	if err != nil {
		log.Info(ctx, LOGGING_CONTEXT, "Error in transformation", "destfile", destfile, "err", err)
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
}*/
