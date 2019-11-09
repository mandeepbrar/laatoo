package main

import (
  
  "laatoo/sdk/server/components/data"
)

type Employee_Ref struct {
  Id    string
  Title string
}

type Employee struct {
	data.Storable `laatoo:"auditable, softdelete"`
  
	FirstName	string `json:"FirstName" bson:"FirstName" datastore:"FirstName"`
	LastName	string `json:"LastName" bson:"LastName" datastore:"LastName"`
	EmployeeID	[]string `json:"EmployeeID" bson:"EmployeeID" datastore:"EmployeeID"`
	Job	Job_Ref `json:"Job" bson:"Job" datastore:"Job"`
}

func (ent *Employee) Config() *data.StorableConfig {
	return &data.StorableConfig{
		IdField:         "Id",
    LabelField:      "Title",
		Type:            "Employee",
		PreSave:         false,
		PostSave:        false,
		PostLoad:        false,
		Auditable:       true,
		Multitenant:     false,
		Collection:      "Employee",
		Cacheable:       false,
	}
}
