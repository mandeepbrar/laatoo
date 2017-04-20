package entities

import (
	"laatoo/sdk/core"
	"laatoo/sdk/registry"

	"github.com/twinj/uuid"
)

const (
	ENTITY_APPL_NAME = "Application"
)

func init() {
	registry.RegisterObject(ENTITY_APPL_NAME, CreateApplication, CreateApplicationCollection)
}

//Creates object
func CreateApplication(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	app := &Application{}
	app.Id = uuid.NewV4().String()
	app.Deleted = false
	return app, nil
}

//Creates collection
func CreateApplicationCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	collect := make([]Application, length)
	return &collect, nil
}

type Application struct {
	Id          string `json:"Id" bson:"Id"`
	CreatedBy   string `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string `json:"UpdatedBy" bson:"UpdatedBy" `
	UpdatedOn   string `json:"UpdatedOn" bson:"UpdatedOn"`
	Name        string `json:"Name" bson:"Name"`
	Description string `json:"Description" bson:"Description"`
	Status      string `json:"Status" bson:"Status"`
	Deleted     bool   `json:"Deleted" bson:"Deleted"`
}

func (ent *Application) GetId() string {
	return ent.Id
}
func (ent *Application) SetId(id string) {
	ent.Id = id
}
func (ent *Application) GetObjectType() string {
	return ENTITY_APPL_NAME
}

func (ent *Application) GetIdField() string {
	return "Id"
}

func (ent *Application) PreSave(ctx core.RequestContext) error {
	return nil
}
func (article *Application) PostSave(ctx core.RequestContext) error {
	return nil
}
func (article *Application) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (ent *Application) IsNew() bool {
	return ent.CreatedBy == ""
}
func (ent *Application) SetUpdatedOn(val string) {
	ent.UpdatedOn = val
}
func (ent *Application) SetUpdatedBy(val string) {
	ent.UpdatedBy = val
}
func (ent *Application) SetCreatedBy(val string) {
	ent.CreatedBy = val
}
func (ent *Application) GetCreatedBy() string {
	return ent.CreatedBy
}
