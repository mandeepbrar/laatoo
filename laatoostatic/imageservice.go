package laatoostatic

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"laatoocore"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
	"laatoosdk/utils"
)

const (
	IMAGE_LOGGING_CONTEXT  = "image_service"
	CONF_IMAGE_SERVICENAME = "image_service"
	CONF_IMAGE_PUBLICDIR   = "publicdir"
	CONF_DISP_MODES        = "displaymodes"
	CONF_DISP_MODES_OPER   = "operation"
	CONF_FILE_SVC          = "fileservice"
)

//Environment hosting an application
type ImageService struct {
	fileServiceName string
	fileService     service.Service
}

//Initialize service, register provider with laatoo
func init() {
	laatoocore.RegisterObjectProvider(CONF_IMAGE_SERVICENAME, ImageServiceFactory)
}

//factory method returns the service object to the environment
func ImageServiceFactory(ctx *echo.Context, conf map[string]interface{}) (interface{}, error) {
	log.Logger.Info(ctx, IMAGE_LOGGING_CONTEXT, "Creating image service")
	svc := &ImageService{}
	routerInt, ok := conf[laatoocore.CONF_ENV_ROUTER]
	if !ok {
		return nil, errors.ThrowError(ctx, STATIC_ERROR_MISSING_ROUTER)
	}

	filesvcInt, ok := conf[CONF_FILE_SVC]
	if !ok {
		return nil, errors.ThrowError(ctx, IMAGE_ERROR_MISSING_FILESVC)
	}
	svc.fileServiceName = filesvcInt.(string)

	router := routerInt.(*echo.Group)

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
			url := fmt.Sprintf("/%s/:srcpath", name)
			oper := operationInt.(string)
			var transformer utils.FileTransform
			transformer = func(reader io.Reader, writer io.Writer) error {
				_, err := io.Copy(writer, reader)
				return err
			}
			if oper == "crop" {
				transformer = func(reader io.Reader, writer io.Writer) error {
					_, err := io.Copy(writer, reader)
					return err
				}
			}
			if oper == "resize_and_crop" {
				transformer = func(reader io.Reader, writer io.Writer) error {
					_, err := io.Copy(writer, reader)
					return err
				}
			}
			router.Get(url, func(reqctx *echo.Context) error {
				log.Logger.Trace(reqctx, IMAGE_LOGGING_CONTEXT, "request", url)
				srcpath := reqctx.P(0)
				returl, err := svc.fileService.Execute(reqctx, "TransformFile", map[string]interface{}{"srcpath": srcpath, "destfolder": name, "transformation": transformer})
				if err != nil {
					return reqctx.NoContent(404)
				}
				log.Logger.Trace(reqctx, IMAGE_LOGGING_CONTEXT, "requestcomplete", "returl", returl)
				return reqctx.Redirect(303, returl.(string))
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
func (svc *ImageService) Initialize(ctx *echo.Context) error {
	svcenv := ctx.Get(laatoocore.CONF_ENV_CONTEXT).(service.Environment)
	fileService, err := svcenv.GetService(ctx, svc.fileServiceName)
	if err != nil {
		return errors.RethrowError(ctx, IMAGE_ERROR_MISSING_FILESVC, err)
	}
	svc.fileService = fileService
	return nil
}

//The service starts serving when this method is called
func (svc *ImageService) Serve(ctx *echo.Context) error {
	return nil
}

//Type of service
func (svc *ImageService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

//Execute method
func (svc *ImageService) Execute(ctx *echo.Context, name string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
