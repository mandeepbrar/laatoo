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
	ENTITY_MEHFIL_NAME         = "mehfil"
	ENTITY_MEHFIL_SERVICE_NAME = "mehfilservice"
)

type MehfilService struct {
	DataStore   data.DataService
	Router      *echo.Group
	dataSvcName string
}

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_MEHFIL_SERVICE_NAME, NewMehfilService)
	laatoocore.RegisterObjectProvider(ENTITY_MEHFIL_NAME, NewMehfil)
}

func NewMehfilService(conf map[string]interface{}) (interface{}, error) {

	log.Logger.Infof("Creating entity service", ENTITY_MEHFIL_SERVICE_NAME)

	svc := &MehfilService{}
	dataSvc, router, err := entities.ParseConfig(conf, svc, ENTITY_MEHFIL_NAME)
	if err != nil {
		return nil, errors.RethrowError(entities.ENTITY_ERROR_CONF_INCORRECT, err, ENTITY_MEHFIL_SERVICE_NAME)
	}
	svc.Router = router
	svc.dataSvcName = dataSvc
	return svc, nil
}

//Provides the name of the service
func (svc *MehfilService) GetName() string {
	return ENTITY_MEHFIL_SERVICE_NAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *MehfilService) Initialize(ctx service.ServiceContext) error {

	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(entities.ENTITY_ERROR_MISSING_DATASVC, err, ENTITY_MEHFIL_SERVICE_NAME)
	}

	svc.DataStore = dataSvc.(data.DataService)

	return nil
}

//The service starts serving when this method is called
func (svc *MehfilService) Serve() error {
	return nil
}

//Type of service
func (svc *MehfilService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

func (svc *MehfilService) GetDataStore() data.DataService {
	return svc.DataStore
}

//Execute method
func (svc *MehfilService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

type Mehfil struct {
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
	Image      string `json:"Image" bson:"Image"`
}

func NewMehfil(conf map[string]interface{}) (interface{}, error) {
	art := &Mehfil{}
	art.Id = uuid.NewV4().String()
	return art, nil
}

func (ent *Mehfil) GetId() string {
	return ent.Id
}
func (ent *Mehfil) SetId(id string) {
	ent.Id = id
}

func (ent *Mehfil) GetIdField() string {
	return "Id"
}
