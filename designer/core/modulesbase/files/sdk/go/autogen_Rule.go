package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

/*
type Rule_Ref struct {
  Id    string
}*/

type Rule struct {
	data.Storable 
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Trigger	string `json:"Trigger" bson:"Trigger" datastore:"Trigger"`
	MessageType	string `json:"MessageType" bson:"MessageType" datastore:"MessageType"`
	Rule	string `json:"Rule" bson:"Rule" datastore:"Rule"`
	LoggingLevel	string `json:"LoggingLevel" bson:"LoggingLevel" datastore:"LoggingLevel"`
	LoggingFormat	string `json:"LoggingFormat" bson:"LoggingFormat" datastore:"LoggingFormat"`
	Settings	map[string]interface{} `json:"Settings" bson:"Settings" datastore:"Settings"`
}

func (ent *Rule) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Rule",
	}
}


func (ent *Rule) GetName()string {
	return ent.Name
}
func (ent *Rule) SetName(val string) {
	ent.Name=val
}
func (ent *Rule) GetDescription()string {
	return ent.Description
}
func (ent *Rule) SetDescription(val string) {
	ent.Description=val
}
func (ent *Rule) GetTrigger()string {
	return ent.Trigger
}
func (ent *Rule) SetTrigger(val string) {
	ent.Trigger=val
}
func (ent *Rule) GetMessageType()string {
	return ent.MessageType
}
func (ent *Rule) SetMessageType(val string) {
	ent.MessageType=val
}
func (ent *Rule) GetRule()string {
	return ent.Rule
}
func (ent *Rule) SetRule(val string) {
	ent.Rule=val
}
func (ent *Rule) GetLoggingLevel()string {
	return ent.LoggingLevel
}
func (ent *Rule) SetLoggingLevel(val string) {
	ent.LoggingLevel=val
}
func (ent *Rule) GetLoggingFormat()string {
	return ent.LoggingFormat
}
func (ent *Rule) SetLoggingFormat(val string) {
	ent.LoggingFormat=val
}
func (ent *Rule) GetSettings()map[string]interface{} {
	return ent.Settings
}
func (ent *Rule) SetSettings(val map[string]interface{}) {
	ent.Settings=val
}
