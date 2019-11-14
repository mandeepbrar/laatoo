package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

/*
type Service_Ref struct {
  Id    string
}*/

type Service struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Factory	string `json:"Factory" bson:"Factory" datastore:"Factory"`
	ServiceMethod	string `json:"ServiceMethod" bson:"ServiceMethod" datastore:"ServiceMethod"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *Service) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Service",
	}
}


func (ent *Service) GetName()string {
	return ent.Name
}
func (ent *Service) SetName(val string) {
	ent.Name=val
}
func (ent *Service) GetDescription()string {
	return ent.Description
}
func (ent *Service) SetDescription(val string) {
	ent.Description=val
}
func (ent *Service) GetFactory()string {
	return ent.Factory
}
func (ent *Service) SetFactory(val string) {
	ent.Factory=val
}
func (ent *Service) GetServiceMethod()string {
	return ent.ServiceMethod
}
func (ent *Service) SetServiceMethod(val string) {
	ent.ServiceMethod=val
}
func (ent *Service) GetLoggingLevel()string {
	return ent.LoggingLevel
}
func (ent *Service) SetLoggingLevel(val string) {
	ent.LoggingLevel=val
}
func (ent *Service) GetLoggingFormat()string {
	return ent.LoggingFormat
}
func (ent *Service) SetLoggingFormat(val string) {
	ent.LoggingFormat=val
}
func (ent *Service) GetSettings()map[string]interface{} {
	return ent.Settings
}
func (ent *Service) SetSettings(val map[string]interface{}) {
	ent.Settings=val
}
