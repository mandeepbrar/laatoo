package static

import (
	"io"
	"laatoo/libraries/disintegration/imaging"
	"laatoo/sdk/components"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"strconv"
)

const (
	CONF_FILE_OPER          = "operation"
	CONF_FILE_TRANSFORM_STG = "transformedstorage"
	CONF_FILE_STORAGE       = "storage"
	CONF_FILE_DEFAULT       = "default"
	CONF_IMAGE_WIDTH        = "width"
	CONF_IMAGE_HEIGHT       = "height"
)

type FileTransform func(io.Reader, io.Writer) error

type staticFiles struct {
	name                            string
	transformedStorageComponentName string
	transformedFilesStorage         components.StorageComponent
	transformFile                   bool
	transformer                     FileTransform
	storageComponentName            string
	storage                         components.StorageComponent
	defaultImage                    string
	hasDefault                      bool
}

func (svc *staticFiles) Initialize(ctx core.ServerContext, conf config.Config) error {
	stg, ok := conf.GetString(CONF_FILE_STORAGE)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", CONF_FILE_STORAGE)
	}
	svc.storageComponentName = stg
	oper, ok := conf.GetString(CONF_FILE_OPER)
	if ok {
		transformer := svc.getImageTransformationMethod(ctx, conf, oper)
		if transformer == nil {
			return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_CONF, "Conf", CONF_FILE_OPER)
		} else {
			svc.transformFile = true
			svc.transformer = transformer
		}
		transformedStorageComponentName, ok := conf.GetString(CONF_FILE_TRANSFORM_STG)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", CONF_FILE_TRANSFORM_STG)
		}
		svc.transformedStorageComponentName = transformedStorageComponentName
		defaultImage, ok := conf.GetString(CONF_FILE_DEFAULT)
		if ok {
			svc.defaultImage = defaultImage
		}
		svc.hasDefault = ok
	}
	return nil
}

func (svc *staticFiles) Invoke(ctx core.RequestContext) error {
	filename, ok := ctx.GetString(CONF_STATIC_FILEPARAM)
	if ok {
		if !svc.transformFile {
			ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, svc.storage.GetFullPath(ctx, filename), nil))
		} else {
			if svc.transformedFilesStorage.Exists(ctx, filename) {
				ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, svc.transformedFilesStorage.GetFullPath(ctx, filename), nil))
			} else {
				created := svc.createFile(ctx, filename)
				if created {
					ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, svc.transformedFilesStorage.GetFullPath(ctx, filename), nil))
				} else {
					if svc.hasDefault {
						ctx.SetResponse(core.NewServiceResponse(core.StatusServeFile, svc.transformedFilesStorage.GetFullPath(ctx, svc.defaultImage), nil))
					} else {
						ctx.SetResponse(core.StatusNotFoundResponse)
					}
				}
			}
		}
	} else {
		ctx.SetResponse(core.StatusNotFoundResponse)
	}
	return nil
}

func (svc *staticFiles) Start(ctx core.ServerContext) error {
	stgSvc, err := ctx.GetService(svc.storageComponentName)
	if err != nil {
		return err
	}
	svc.storage = stgSvc.(components.StorageComponent)
	if svc.transformFile {
		tstgSvc, err := ctx.GetService(svc.transformedStorageComponentName)
		if err != nil {
			return err
		}
		svc.transformedFilesStorage = tstgSvc.(components.StorageComponent)
	}
	return nil
}

func (svc *staticFiles) getImageTransformationMethod(ctx core.ServerContext, conf config.Config, oper string) FileTransform {
	width := 0
	widthStr, widthok := conf.GetString(CONF_IMAGE_WIDTH)
	if widthok {
		width, _ = strconv.Atoi(widthStr)
	}
	height := 0
	heightStr, heightok := conf.GetString(CONF_IMAGE_HEIGHT)
	if heightok {
		height, _ = strconv.Atoi(heightStr)
	}
	switch oper {
	case "crop":
		{
			return func(reader io.Reader, writer io.Writer) error {
				img, format, err := imaging.Decode(reader)
				log.Logger.Info(ctx, "crop", "format", format)
				if err != nil {
					return err
				}
				log.Logger.Trace(ctx, "cropping image", "format", format)
				dstImage := imaging.CropCenter(img, width, height)
				return imaging.Encode(writer, dstImage, getFormat(format))
			}
		}
	case "fill":
		{
			return func(reader io.Reader, writer io.Writer) error {
				img, format, err := imaging.Decode(reader)
				if err != nil {
					return err
				}
				log.Logger.Trace(ctx, "filling image", "format", format)
				dstImage := imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)
				return imaging.Encode(writer, dstImage, getFormat(format))
			}
		}
	case "fit":
		{
			return func(reader io.Reader, writer io.Writer) error {
				img, format, err := imaging.Decode(reader)
				if err != nil {
					return err
				}
				log.Logger.Trace(ctx, "fitting image", "format", format)
				dstImage := imaging.Fit(img, width, height, imaging.Lanczos)
				return imaging.Encode(writer, dstImage, getFormat(format))
			}
		}
	default:
		{
			return nil
		}
	}
}

func getFormat(format string) imaging.Format {
	switch format {
	case "jpeg":
		return imaging.JPEG
	case "png":
		return imaging.PNG
	case "gif":
		return imaging.GIF
	case "tiff":
		return imaging.PNG
	}
	return imaging.JPEG
}

func (svc *staticFiles) createFile(ctx core.RequestContext, filename string) bool {
	inStr, err := svc.storage.Open(ctx, filename)
	defer inStr.Close()
	if err != nil {
		log.Logger.Trace(ctx, "File does not exist", "sourcefile", filename, "err", err)
		return false
	}

	writer, err := svc.transformedFilesStorage.CreateFile(ctx, filename)
	defer writer.Close()
	if err != nil {
		log.Logger.Trace(ctx, "Error opening source file", "destfile", filename, "err", err)
		return false
	}
	/*destdir, _ := path.Split(destfile)
	os.MkdirAll(destdir, 0755)
	writer, err := os.Create(destfile)
	defer writer.Close()
	if err != nil {
		log.Logger.Info(ctx, "error creating file", "destfile", destfile, "err", err)
		return "", err
	}*/

	err = svc.transformer(inStr, writer)
	if err != nil {
		log.Logger.Trace(ctx, "Error in transformation", "destfile", filename, "err", err)
		return false
	}
	return true
}

/*
func (svc *staticFiles) copyFile(ctx core.Context, fileurl string, writer io.Writer) error {
	_, realpath := svc.parsePath(fileurl)
	rd, err := svc.storage.Open(ctx, realpath)
	defer rd.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, rd)
	return err
}

func (svc *staticFiles) parsePath(url string) (string, string) {
	var prefix string
	if url[0] != '/' && svc.filesUrl[0] == '/' {
		prefix = svc.filesUrl[1:]
	} else {
		prefix = svc.filesUrl
	}
	pathinfolder := strings.TrimPrefix(url, prefix)
	return pathinfolder, fmt.Sprintf("%s%s", svc.directory, pathinfolder)
}*/
