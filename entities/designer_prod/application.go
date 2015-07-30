package designer_prod

import (
	"laatoocore"
)

var (
	ENTITY_NAME = "Application"
)

type Application struct {
	Id        string `json:"Id" bson:"Id"`
	Name      string `json:"Name" bson:"Name"`
	CreatedBy string
	UpdatedBy string
}

func init() {
	laatoocore.RegisterObjectProvider("Application", NewApplication)
}

func NewApplication(conf map[string]interface{}) (interface{}, error) {
	return &Application{}, nil
}

func (ent *Application) RegisterRoutes() {

}
