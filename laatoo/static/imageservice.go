package laatoostatic

/*
import (
	"disintegration/imaging"
	"fmt"
	"io"
	"laatoocore"
	"laatoosdk/core"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/utils"
	"math/rand"
	"strconv"
	"time"
)

const (
	IMAGE_LOGGING_CONTEXT   = "image_service"
	CONF_IMAGE_SERVICENAME  = "image_service"
	CONF_IMAGE_PUBLICDIR    = "publicdir"
	CONF_DISP_MODES         = "displaymodes"
	CONF_DISP_MODES_OPER    = "operation"
	CONF_DISP_MODES_WIDTH   = "width"
	CONF_DISP_MODES_HEIGHT  = "height"
	CONF_DISP_MODES_DEFAULT = "default"
	CONF_FILE_SVC           = "fileservice"
)

//Environment hosting an application
type ImageService struct {
	fileServiceName string
	fileService     core.Service
	conf            map[string]interface{}
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_IMAGE_SERVICENAME, ImageServiceFactory)
	rand.Seed(time.Now().UTC().UnixNano())
}

//factory method returns the service object to the environment
func ImageServiceFactory(ctx core.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, IMAGE_LOGGING_CONTEXT, "Creating image service")
	svc := &ImageService{}
	svc.conf = conf
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, STATIC_ERROR_MISSING_ROUTER)
	}

	filesvcInt, ok := conf[CONF_FILE_SVC]
	if !ok {
		return nil, errors.ThrowError(ctx, IMAGE_ERROR_MISSING_FILESVC)
	}
	svc.fileServiceName = filesvcInt.(string)

	router := routerInt.(core.Router)

	//get a map of all the pages
	displayModesInt, ok := conf[CONF_DISP_MODES]

	dispModes, ok := displayModesInt.(map[string]interface{})
	if !ok {
		return nil, errors.ThrowError(ctx, IMAGE_ERROR_DISP_MODES_NOT_PROVIDED)
	}

	for name, val := range dispModes {
		//get the config for the dispMode
		dispModeConf := val.(map[string]interface{})
		operationInt, ok := dispModeConf[CONF_DISP_MODES_OPER]
		if ok {
			url := fmt.Sprintf("/%s/*", name)
			oper := operationInt.(string)
			defaultInt, _ := dispModeConf[CONF_DISP_MODES_DEFAULT]
			width := 0
			widthInt, widthok := dispModeConf[CONF_DISP_MODES_WIDTH]
			if widthok {
				width, _ = strconv.Atoi(widthInt.(string))
			}
			height := 0
			heightInt, heightok := dispModeConf[CONF_DISP_MODES_HEIGHT]
			if heightok {
				height, _ = strconv.Atoi(heightInt.(string))
			}
			transformer := svc.getTransformationMethod(ctx, oper, height, width)
			defaultImgArr := svc.getDefaultImageArray(ctx, defaultInt)
			log.Logger.Info(ctx, IMAGE_LOGGING_CONTEXT, "defaultImgArr", "name", name, "defaultImgArr", defaultImgArr)
			router.Get(ctx, url, dispModeConf, func(reqctx core.Context) error {
				log.Logger.Trace(reqctx, IMAGE_LOGGING_CONTEXT, "narray check", "name", name)
				srcpath := reqctx.ParamByIndex(0)
				if len(srcpath) > 0 {
					log.Logger.Trace(reqctx, IMAGE_LOGGING_CONTEXT, "request", "url", url, "reqctx", reqctx, "srcpath", srcpath)
					returl, err := svc.fileService.Execute(reqctx, "TransformFile", map[string]interface{}{"srcpath": srcpath, "destfolder": name, "transformation": transformer})
					if err != nil {
						return svc.handleDefaultImage(reqctx, defaultImgArr)
					}
					log.Logger.Trace(reqctx, IMAGE_LOGGING_CONTEXT, "requestcomplete", "returl", returl)
					return reqctx.Redirect(301, returl.(string))
				}
				return svc.handleDefaultImage(reqctx, defaultImgArr)
			})
			//get the service name to be created for the alias
			log.Logger.Debug(ctx, IMAGE_LOGGING_CONTEXT, "Serving display mode", "Name", name, " URL", url)
		}
	}

	log.Logger.Info(ctx, IMAGE_LOGGING_CONTEXT, "Image service starting")
	return svc, nil
}

//Provides the name of the service
func (svc *ImageService) GetName() string {
	return CONF_IMAGE_SERVICENAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *ImageService) Initialize(ctx core.Context) error {
	fileService, err := ctx.GetService(svc.fileServiceName)
	if err != nil {
		return errors.RethrowError(ctx, IMAGE_ERROR_MISSING_FILESVC, err)
	}
	svc.fileService = fileService
	return nil
}

//The service starts serving when this method is called
func (svc *ImageService) Serve(ctx core.Context) error {
	return nil
}

//Type of service
func (svc *ImageService) GetServiceType() string {
	return core.SERVICE_TYPE_WEB
}

//Execute method
func (svc *ImageService) Execute(ctx core.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}

func (svc *ImageService) getDefaultImageArray(ctx core.Context, defaultImageInt interface{}) *[]string {
	array := &[]string{}
	if defaultImageInt == nil {
		return array
	}
	defaultImage := defaultImageInt.(string)
	imagelen := len(defaultImage)
	if imagelen == 0 {
		return array
	}
	if imagelen > 2 && defaultImage[0] == '<' && defaultImage[imagelen-1] == '>' {
		token := defaultImage[1 : imagelen-2]
		log.Logger.Trace(ctx, IMAGE_LOGGING_CONTEXT, "default images token", "token", token)
		imagesArrInt, ok := svc.conf[token]
		if !ok {
			return array
		}
		imagesArr := imagesArrInt.([]string)
		return &imagesArr
	}
	return &([]string{defaultImage})
}

func (svc *ImageService) getTransformationMethod(ctx core.Context, oper string, height int, width int) utils.FileTransform {
	var transformer utils.FileTransform
	transformer = func(reader io.Reader, writer io.Writer) error {
		img, format, err := imaging.Decode(reader)
		log.Logger.Info(ctx, IMAGE_LOGGING_CONTEXT, "d image", "format", format)
		if err != nil {
			return err
		}
		log.Logger.Trace(ctx, IMAGE_LOGGING_CONTEXT, "default image", "format", format)
		dstImage := imaging.Resize(img, width, height, imaging.Lanczos)
		return imaging.Encode(writer, dstImage, getFormat(format))
	}
	if oper == "crop" {
		transformer = func(reader io.Reader, writer io.Writer) error {
			img, format, err := imaging.Decode(reader)
			log.Logger.Info(ctx, IMAGE_LOGGING_CONTEXT, "crop", "format", format)
			if err != nil {
				return err
			}
			log.Logger.Trace(ctx, IMAGE_LOGGING_CONTEXT, "cropping image", "format", format)
			dstImage := imaging.CropCenter(img, width, height)
			return imaging.Encode(writer, dstImage, getFormat(format))
		}
	}
	if oper == "fill" {
		transformer = func(reader io.Reader, writer io.Writer) error {
			img, format, err := imaging.Decode(reader)
			if err != nil {
				return err
			}
			log.Logger.Trace(ctx, IMAGE_LOGGING_CONTEXT, "filling image", "format", format)
			dstImage := imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)
			return imaging.Encode(writer, dstImage, getFormat(format))
		}
	}
	if oper == "fit" {
		transformer = func(reader io.Reader, writer io.Writer) error {
			img, format, err := imaging.Decode(reader)
			if err != nil {
				return err
			}
			log.Logger.Trace(ctx, IMAGE_LOGGING_CONTEXT, "fitting image", "format", format)
			dstImage := imaging.Fit(img, width, height, imaging.Lanczos)
			return imaging.Encode(writer, dstImage, getFormat(format))
		}
	}
	return transformer
}

func (svc *ImageService) handleDefaultImage(ctx core.Context, arr *[]string) error {
	defaultImageArr := *arr
	length := len(defaultImageArr)
	if length == 0 {
		return ctx.NoContent(404)
	}
	randImage := defaultImageArr[rand.Intn(length)]
	return ctx.Redirect(301, randImage)
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
*/
