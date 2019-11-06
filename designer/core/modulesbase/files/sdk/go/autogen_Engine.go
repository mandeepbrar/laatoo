package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type Engine_Ref struct {
  Id    string
}

type Engine struct {
	*data.SerializableBase `initialize:"SerializableBase"`
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	EngineType	string `json:"EngineType" bson:"EngineType" datastore:"EngineType"`
	Address	string `json:"Address" bson:"Address" datastore:"Address"`
	Framework	string `json:"Framework" bson:"Framework" datastore:"Framework"`
	SSL	bool `json:"SSL" bson:"SSL" datastore:"SSL"`
	CORS	bool `json:"CORS" bson:"CORS" datastore:"CORS"`
	Path	string `json:"Path" bson:"Path" datastore:"Path"`
	CORSHosts	[]string `json:"CORSHosts" bson:"CORSHosts" datastore:"CORSHosts"`
	QueryParams	[]string `json:"QueryParams" bson:"QueryParams" datastore:"QueryParams"`
}

func (ent *Engine) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Engine",
	}
}


func (ent *Engine) GetName()string {
	return ent.Name
}
func (ent *Engine) SetName(val string) {
	ent.Name=val
}
func (ent *Engine) GetDescription()string {
	return ent.Description
}
func (ent *Engine) SetDescription(val string) {
	ent.Description=val
}
func (ent *Engine) GetEngineType()string {
	return ent.EngineType
}
func (ent *Engine) SetEngineType(val string) {
	ent.EngineType=val
}
func (ent *Engine) GetAddress()string {
	return ent.Address
}
func (ent *Engine) SetAddress(val string) {
	ent.Address=val
}
func (ent *Engine) GetFramework()string {
	return ent.Framework
}
func (ent *Engine) SetFramework(val string) {
	ent.Framework=val
}
func (ent *Engine) GetSSL()bool {
	return ent.SSL
}
func (ent *Engine) SetSSL(val bool) {
	ent.SSL=val
}
func (ent *Engine) GetCORS()bool {
	return ent.CORS
}
func (ent *Engine) SetCORS(val bool) {
	ent.CORS=val
}
func (ent *Engine) GetPath()string {
	return ent.Path
}
func (ent *Engine) SetPath(val string) {
	ent.Path=val
}
func (ent *Engine) GetCORSHosts()[]string {
	return ent.CORSHosts
}
func (ent *Engine) SetCORSHosts(val []string) {
	ent.CORSHosts=val
}
func (ent *Engine) GetQueryParams()[]string {
	return ent.QueryParams
}
func (ent *Engine) SetQueryParams(val []string) {
	ent.QueryParams=val
}
