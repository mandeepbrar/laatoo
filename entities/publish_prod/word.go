package publish_prod

import (
	"github.com/labstack/echo"
	"laatoocore"
	"laatoosdk/data"
	"laatoosdk/entities"
	"laatoosdk/errors"
	"laatoosdk/log"
	"laatoosdk/service"
)

const (
	ENTITY_WORD_NAME         = "word"
	ENTITY_WORD_SERVICE_NAME = "wordservice"
)

type WordService struct {
	DataStore   data.DataService
	Router      *echo.Group
	dataSvcName string
}

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_WORD_SERVICE_NAME, NewWordService)
	laatoocore.RegisterObjectProvider(ENTITY_WORD_NAME, NewWord)
}

func NewWordService(conf map[string]interface{}) (interface{}, error) {

	log.Logger.Infof("Creating entity service", ENTITY_WORD_SERVICE_NAME)

	svc := &WordService{}
	dataSvc, router, err := entities.ParseConfig(conf, svc, ENTITY_WORD_NAME)
	if err != nil {
		return nil, errors.RethrowError(entities.ENTITY_ERROR_CONF_INCORRECT, err, ENTITY_WORD_SERVICE_NAME)
	}
	svc.Router = router
	svc.dataSvcName = dataSvc

	return svc, nil
}

//Provides the name of the service
func (svc *WordService) GetName() string {
	return ENTITY_WORD_SERVICE_NAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *WordService) Initialize(ctx service.ServiceContext) error {

	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(entities.ENTITY_ERROR_MISSING_DATASVC, err, ENTITY_WORD_SERVICE_NAME)
	}

	svc.DataStore = dataSvc.(data.DataService)

	return nil
}

//The service starts serving when this method is called
func (svc *WordService) Serve() error {
	return nil
}

//Type of service
func (svc *WordService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

func (svc *WordService) GetDataStore() data.DataService {
	return svc.DataStore
}

//Execute method
func (svc *WordService) Execute(name string, params map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

type Word struct {
	WordGur  string `json:"WordGur" bson:"WordGur"`
	WordEng  string `json:"WordEng" bson:"WordEng"`
	Reviewed bool   `json:"Reviewed" bson:"Reviewed"`
}

func NewWord(conf map[string]interface{}) (interface{}, error) {
	word := &Word{Reviewed: false}
	return word, nil
}

func (ent *Word) GetId() string {
	return ent.WordGur
}
func (ent *Word) SetId(id string) {
	ent.WordGur = id
}

func (ent *Word) GetIdField() string {
	return "WordGur"
}
