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
	ENTITY_ARTICLE_NAME         = "article"
	ENTITY_ARTICLE_SERVICE_NAME = "articleservice"
)

type ArticleService struct {
	DataStore   data.DataService
	Router      *echo.Group
	dataSvcName string
}

func init() {
	laatoocore.RegisterObjectProvider(ENTITY_ARTICLE_SERVICE_NAME, NewArticleService)
	laatoocore.RegisterObjectProvider(ENTITY_ARTICLE_NAME, NewArticle)
}

func NewArticleService(conf map[string]interface{}) (interface{}, error) {

	log.Logger.Infof("Creating entity service", ENTITY_ARTICLE_SERVICE_NAME)

	svc := &ArticleService{}
	dataSvc, router, err := entities.ParseConfig(conf, svc, ENTITY_ARTICLE_NAME, NewArticle)
	if err != nil {
		return nil, errors.RethrowError(entities.ENTITY_ERROR_CONF_INCORRECT, err, ENTITY_ARTICLE_SERVICE_NAME)
	}
	svc.Router = router
	svc.dataSvcName = dataSvc
	return svc, nil
}

//Provides the name of the service
func (svc *ArticleService) GetName() string {
	return ENTITY_ARTICLE_SERVICE_NAME
}

//Initialize the service. Consumer of a service passes the data
func (svc *ArticleService) Initialize(ctx service.ServiceContext) error {

	dataSvc, err := ctx.GetService(svc.dataSvcName)
	if err != nil {
		return errors.RethrowError(entities.ENTITY_ERROR_MISSING_DATASVC, err, ENTITY_ARTICLE_SERVICE_NAME)
	}

	svc.DataStore = dataSvc.(data.DataService)

	return nil
}

//The service starts serving when this method is called
func (svc *ArticleService) Serve() error {
	return nil
}

//Type of service
func (svc *ArticleService) GetServiceType() string {
	return service.SERVICE_TYPE_WEB
}

func (svc *ArticleService) GetDataStore() data.DataService {
	return svc.DataStore
}

type Article struct {
	Id        string `json:"Id" bson:"Id"`
	CreatedBy string
	UpdatedBy string
	Title     string `json:"Title" bson:"Title"`
}

func NewArticle(conf map[string]interface{}) (interface{}, error) {
	return &Article{}, nil
}

func (ent *Article) GetId() string {
	return ent.Id
}
func (ent *Article) SetId(id string) {
	ent.Id = id
}

func (ent *Article) GetIdField() string {
	return "Id"
}
