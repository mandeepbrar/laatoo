package entities

import (
	"laatoo/sdk/core"

	"github.com/twinj/uuid"
)

const (
	ENTITY_SERVER_NAME = "Server"
)

//Creates object
func CreateServer(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	svr := &Server{}
	svr.Id = uuid.NewV4().String()
	svr.Deleted = false
	return svr, nil
}

//Creates collection
func CreateServerCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	collect := make([]Server, length)
	return &collect, nil
}

type Server struct {
	Id          string `json:"Id" bson:"Id"`
	CreatedBy   string `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string `json:"UpdatedBy" bson:"UpdatedBy" `
	UpdatedOn   string `json:"UpdatedOn" bson:"UpdatedOn"`
	Name        string `json:"Name" bson:"Name"`
	Description string `json:"Description" bson:"Description"`
	Status      string `json:"Status" bson:"Status"`
	Deleted     bool   `json:"Deleted" bson:"Deleted"`
}

func (ent *Server) GetId() string {
	return ent.Id
}
func (ent *Server) SetId(id string) {
	ent.Id = id
}
func (ent *Server) GetObjectType() string {
	return ENTITY_SERVER_NAME
}

func (ent *Server) GetIdField() string {
	return "Id"
}

func (ent *Server) PreSave(ctx core.RequestContext) error {
	return nil
}
func (en *Server) PostSave(ctx core.RequestContext) error {
	return nil
}
func (ent *Server) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (ent *Server) IsNew() bool {
	return ent.CreatedBy == ""
}
func (ent *Server) SetUpdatedOn(val string) {
	ent.UpdatedOn = val
}
func (ent *Server) SetUpdatedBy(val string) {
	ent.UpdatedBy = val
}
func (ent *Server) SetCreatedBy(val string) {
	ent.CreatedBy = val
}
func (ent *Server) GetCreatedBy() string {
	return ent.CreatedBy
}
