package main

import (
	"image"
	"io"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	CONF_FILE_OPER             = "operation"
	CONF_FILE_TRANSFORM_STG    = "transformedstorage"
	CONF_FILE_STORAGE          = "storage"
	CONF_FILE_DEFAULT          = "default"
	CONF_IMAGE_WIDTH           = "width"
	CONF_IMAGE_HEIGHT          = "height"
	CONF_STATICFILE_FILEBUCKET = "bucket"
)

type FileTransform func(io.Reader, io.Writer) error

type StaticFiles struct {
	core.Service
	name                            string
	transformedStorageComponentName string
	transformedFilesStorage         components.StorageComponent
	transformFile                   bool
	transformer                     FileTransform
	storageComponentName            string
	storage                         components.StorageComponent
	defaultImage                    string
	hasDefault                      bool
	bucket                          string
}

/*
func (svc *StaticFiles) Initialize(ctx core.ServerContext) error {
	svc.SetDescription(ctx, "Static files service")
	svc.AddStringConfigurations(ctx, []string{CONF_FILE_STORAGE}, nil)
	svc.AddStringConfigurations(ctx, []string{CONF_FILE_OPER, CONF_FILE_TRANSFORM_STG, CONF_FILE_DEFAULT, CONF_IMAGE_WIDTH, CONF_IMAGE_HEIGHT}, []string{"", "", "", "0", "0"})
	svc.AddParam(ctx, CONF_STATIC_FILEPARAM, config.OBJECTTYPE_STRING, false)

	return nil
}*/

func (svc *StaticFiles) Invoke(ctx core.RequestContext) error {
	fn, ok := ctx.GetParam(CONF_STATIC_FILEPARAM)
	log.Trace(ctx, "Received request for file", "filename", fn)
	if ok {
		filename := strings.TrimLeft(fn.GetValue().(string), "/")
		if !svc.transformFile {
			return svc.storage.ServeFile(ctx, svc.bucket, filename)
		} else {
			if svc.transformedFilesStorage.Exists(ctx, svc.bucket, filename) {
				return svc.transformedFilesStorage.ServeFile(ctx, svc.bucket, filename)
			} else {
				created := svc.createFile(ctx, filename)
				if created {
					return svc.transformedFilesStorage.ServeFile(ctx, svc.bucket, filename)
				} else {
					if svc.hasDefault {
						return svc.transformedFilesStorage.ServeFile(ctx, svc.bucket, svc.defaultImage)
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

func (svc *StaticFiles) Start(ctx core.ServerContext) error {
	svc.bucket, _ = svc.GetStringConfiguration(ctx, CONF_STATICFILE_FILEBUCKET)
	stg, _ := svc.GetConfiguration(ctx, CONF_FILE_STORAGE)
	svc.storageComponentName = stg.(string)

	oper, ok := svc.GetConfiguration(ctx, CONF_FILE_OPER)
	if ok {
		transformer := svc.getImageTransformationMethod(ctx, oper.(string))
		if transformer == nil {
			return errors.BadConf(ctx, CONF_FILE_OPER)
		} else {
			svc.transformFile = true
			svc.transformer = transformer
		}

		transformedStorageComponentName, ok := svc.GetConfiguration(ctx, CONF_FILE_TRANSFORM_STG)
		if !ok {
			return errors.MissingConf(ctx, CONF_FILE_TRANSFORM_STG)
		}
		svc.transformedStorageComponentName = transformedStorageComponentName.(string)

		defaultImage, ok := svc.GetConfiguration(ctx, CONF_FILE_DEFAULT)
		if ok {
			svc.defaultImage = defaultImage.(string)
		}
		svc.hasDefault = ok
	}

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

func (svc *StaticFiles) getImageTransformationMethod(ctx core.ServerContext, oper string) FileTransform {
	widthStr, _ := svc.GetConfiguration(ctx, CONF_IMAGE_WIDTH)
	width, _ := strconv.Atoi(widthStr.(string))

	heightStr, _ := svc.GetConfiguration(ctx, CONF_IMAGE_HEIGHT)
	height, _ := strconv.Atoi(heightStr.(string))

	switch oper {
	case "crop":
		{
			return func(reader io.Reader, writer io.Writer) error {
				img, format, err := decode(reader)
				log.Info(ctx, "crop", "format", format)
				if err != nil {
					return err
				}
				log.Trace(ctx, "cropping image", "format", format)
				dstImage := imaging.CropCenter(img, width, height)
				return imaging.Encode(writer, dstImage, getFormat(format))
			}
		}
	case "fill":
		{
			return func(reader io.Reader, writer io.Writer) error {
				img, format, err := decode(reader)
				if err != nil {
					return err
				}
				log.Trace(ctx, "filling image", "format", format)
				dstImage := imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)
				return imaging.Encode(writer, dstImage, getFormat(format))
			}
		}
	case "fit":
		{
			return func(reader io.Reader, writer io.Writer) error {
				img, format, err := decode(reader)
				if err != nil {
					return err
				}
				log.Trace(ctx, "fitting image", "format", format)
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

func decode(reader io.Reader) (image.Image, string, error) {
	//img, format, err := imaging.Decode(reader)
	return image.Decode(reader)
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

func (svc *StaticFiles) createFile(ctx core.RequestContext, filename string) bool {
	log.Trace(ctx, "Opening file", "filename", filename)
	inStr, err := svc.storage.Open(ctx, svc.bucket, filename)
	if err != nil {
		log.Trace(ctx, "File does not exist", "sourcefile", filename, "err", err)
		return false
	}
	defer inStr.Close()

	writer, err := svc.transformedFilesStorage.CreateFile(ctx, svc.bucket, filename, "")
	if err != nil {
		log.Trace(ctx, "Error opening source file", "destfile", filename, "err", err)
		return false
	}
	defer writer.Close()

	err = svc.transformer(inStr, writer)
	if err != nil {
		log.Trace(ctx, "Error in transformation", "destfile", filename, "err", err)
		return false
	}
	return true
}
