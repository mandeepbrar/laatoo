package entities

import (
	"laatoo/sdk/core"

	"github.com/twinj/uuid"
)

const (
	ENTITY_SOLUTION_NAME = "Solution"
)

//Creates object
func CreateSolution(ctx core.Context, args core.MethodArgs) (interface{}, error) {
	sol := &Solution{}
	sol.Id = uuid.NewV4().String()
	sol.Deleted = false
	return sol, nil
}

//Creates collection
func CreateSolutionCollection(ctx core.Context, length int, args core.MethodArgs) (interface{}, error) {
	collect := make([]Solution, length)
	return &collect, nil
}

type Solution struct {
	Id          string `json:"Id" bson:"Id"`
	CreatedBy   string `json:"CreatedBy" bson:"CreatedBy"`
	UpdatedBy   string `json:"UpdatedBy" bson:"UpdatedBy" `
	UpdatedOn   string `json:"UpdatedOn" bson:"UpdatedOn"`
	Name        string `json:"Name" bson:"Name"`
	Description string `json:"Description" bson:"Description"`
	Status      string `json:"Status" bson:"Status"`
	Deleted     bool   `json:"Deleted" bson:"Deleted"`
}

func (ent *Solution) GetId() string {
	return ent.Id
}
func (ent *Solution) SetId(id string) {
	ent.Id = id
}
func (ent *Solution) GetObjectType() string {
	return ENTITY_APPL_NAME
}

func (ent *Solution) GetIdField() string {
	return "Id"
}

func (ent *Solution) PreSave(ctx core.RequestContext) error {
	return nil
}
func (article *Solution) PostSave(ctx core.RequestContext) error {
	return nil
}
func (article *Solution) PostLoad(ctx core.RequestContext) error {
	return nil
}
func (ent *Solution) IsNew() bool {
	return ent.CreatedBy == ""
}
func (ent *Solution) SetUpdatedOn(val string) {
	ent.UpdatedOn = val
}
func (ent *Solution) SetUpdatedBy(val string) {
	ent.UpdatedBy = val
}
func (ent *Solution) SetCreatedBy(val string) {
	ent.CreatedBy = val
}
func (ent *Solution) GetCreatedBy() string {
	return ent.CreatedBy
}
