package entities

import (
	"laatoo/sdk/core"

	"github.com/twinj/uuid"
)

const (
	ENTITY_ENV_NAME = "Environment"
)

//Creates object
func CreateEnvironment(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	svr := &Environment{}
	svr.Id = uuid.NewV4().String()
	svr.Deleted = false
	return svr, nil
}

//Creates collection
func CreateEnvironmentCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	collect := make([]Environment, length)
	return &collect, nil
}

type Environment struct {
	Id          string `json:"Id" bson:"Id"`
	CreatedBy   string `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string `json:"UpdatedBy" bson:"UpdatedBy" `
	UpdatedOn   string `json:"UpdatedOn" bson:"UpdatedOn"`
	Name        string `json:"Name" bson:"Name"`
	Description string `json:"Description" bson:"Description"`
	Status      string `json:"Status" bson:"Status"`
	Deleted     bool   `json:"Deleted" bson:"Deleted"`
}

func (ent *Environment) GetId() string {
	return ent.Id
}
func (ent *Environment) SetId(id string) {
	ent.Id = id
}
func (ent *Environment) GetObjectType() string {
	return ENTITY_ENV_NAME
}

func (ent *Environment) GetIdField() string {
	return "Id"
}

func (ent *Environment) PreSave(ctx core.RequestContext) error {
	return nil
}
func (en *Environment) PostSave(ctx core.RequestContext) error {
	return nil
}
func (ent *Environment) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (ent *Environment) IsNew() bool {
	return ent.CreatedBy == ""
}
func (ent *Environment) SetUpdatedOn(val string) {
	ent.UpdatedOn = val
}
func (ent *Environment) SetUpdatedBy(val string) {
	ent.UpdatedBy = val
}
func (ent *Environment) SetCreatedBy(val string) {
	ent.CreatedBy = val
}
func (ent *Environment) GetCreatedBy() string {
	return ent.CreatedBy
}
