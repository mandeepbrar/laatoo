package modulesbase


import (
  
  "laatoo/sdk/server/components/data"
)

type Task_Ref struct {
  Id    string
}

type Task struct {
  Id    string `json:"Id" bson:"Id" datastore:"Id"`
  
	Name	string `json:"Name" bson:"Name" datastore:"Name"`
	Description	string `json:"Description" bson:"Description" datastore:"Description"`
	Receiver	string `json:"Receiver" bson:"Receiver" datastore:"Receiver"`
	Processor	string `json:"Processor" bson:"Processor" datastore:"Processor"`
}

func (ent *Task) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Name",
		Type:            "Task",
	}
}


func (ent *Task) GetName()string {
	return ent.Name
}
func (ent *Task) SetName(val string) {
	ent.Name=val
}
func (ent *Task) GetDescription()string {
	return ent.Description
}
func (ent *Task) SetDescription(val string) {
	ent.Description=val
}
func (ent *Task) GetReceiver()string {
	return ent.Receiver
}
func (ent *Task) SetReceiver(val string) {
	ent.Receiver=val
}
func (ent *Task) GetProcessor()string {
	return ent.Processor
}
func (ent *Task) SetProcessor(val string) {
	ent.Processor=val
}
