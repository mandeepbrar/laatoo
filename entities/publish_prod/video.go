package publish_prod

import (
	"github.com/labstack/echo"
	"github.com/twinj/uuid"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	ENTITY_VIDEO_NAME         = "video"
	ENTITY_VIDEO_SERVICE_NAME = "videoservice"
	PERM_VIDEO_VIEW           = "View Video"
	PERM_VIDEO_EDIT           = "Edit Video"
	PERM_VIDEO_CREATE         = "Create Video"
	PERM_VIDEO_DEL            = "Delete Video"
)

type VideoService struct {
	DataStore   data.DataService
	Router      *echo.Group
	dataSvcName string
}

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_VIDEO_SERVICE_NAME, NewVideoService)
	laatoocore.RegisterObjectProvider(ENTITY_VIDEO_NAME, NewVideo)
	laatoocore.RegisterPermissions([]string{PERM_VIDEO_VIEW, PERM_VIDEO_EDIT, PERM_VIDEO_CREATE, PERM_VIDEO_DEL})
}

func NewVideoService(conf map[string]interface{}) (interface{}, error) {

	log.Logger.Infof("Creating entity service", ENTITY_VIDEO_SERVICE_NAME)

	svc := &VideoService{}
	dataSvc, router, err := entities.ParseConfig(conf, svc, ENTITY_VIDEO_NAME)
	if err != nil {
		return nil, errors.RethrowError(entities.ENTITY_ERROR_CONF_INCORRECT, err, ENTITY_VIDEO_SERVICE_NAME)
	}
	svc.Router = router
	svc.dataSvcName = dataSvc
	return svc, nil
}

//Provides the name of the service
func (svc *VideoService) GetName() string {
	return ENTITY_VIDEO_SERVICE_NAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *VideoService) Initialize(ctx service.ServiceContext) error {

	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(entities.ENTITY_ERROR_MISSING_DATASVC, err, ENTITY_VIDEO_SERVICE_NAME)
	}

	svc.DataStore = dataSvc.(data.DataService)

	return nil
}

//The service starts serving when this method is called
func (svc *VideoService) Serve() error {
	return nil
}

//Type of service
func (svc *VideoService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

func (svc *VideoService) GetDataStore() data.DataService {
	return svc.DataStore
}

//Execute method
func (svc *VideoService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

type Video struct {
	Id         string `json:"Id" bson:"Id"`
	CreatedBy  string `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy  string `json:"UpdatedBy" bson:"UpdatedBy"`
	UpdatedOn  string `json:"UpdatedOn" bson:"UpdatedOn"`
	Title      string `json:"Title" bson:"Title"`
	BodyGur    string `json:"BodyGur" bson:"BodyGur"`
	SummaryGur string `json:"SummaryGur" bson:"SummaryGur"`
	BodyEng    string `json:"BodyEng" bson:"BodyEng"`
	SummaryEng string `json:"SummaryEng" bson:"SummaryEng"`
	TitleEng   string `json:"TitleEng" bson:"TitleEng"`
	Video      string `json:"Video" bson:"Video"`
}

func NewVideo(conf map[string]interface{}) (interface{}, error) {
	art := &Video{}
	art.Id = uuid.NewV4().String()
	return art, nil
}

func (ent *Video) GetId() string {
	return ent.Id
}
func (ent *Video) SetId(id string) {
	ent.Id = id
}

func (ent *Video) GetIdField() string {
	return "Id"
}
