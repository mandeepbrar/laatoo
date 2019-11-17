package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Employee", Object: Employee{}, Metadata: core.NewInfo("","Employee", map[string]interface{}{"descriptor":"{\"name\":\"Employee\",\"inherits\":\"\",\"imports\":[\"bytes\"],\"form\":{\"layout\":[\"Consultants\",\"Contractors\"]},\"fields\":{\"FirstName\":{\"type\":\"string\",\"widget\":{\"props\":{\"label\":\"First Name\"}}},\"LastName\":{\"type\":\"string\",\"widget\":{\"props\":{\"label\":\"Last Name\"}}},\"EmployeeID\":{\"type\":\"string\",\"list\":true,\"widget\":{\"props\":{\"label\":\"Employee Id\"}}},\"Job\":{\"type\":\"storableref\",\"entity\":\"Job\"}}}"})},core.PluginComponent{Name: "Job", Object: Job{}, Metadata: core.NewInfo("","Job", map[string]interface{}{"descriptor":"{\"name\":\"Job\",\"inherits\":\"\",\"imports\":[\"bytes\"],\"form\":{\"overlay\":true,\"layout\":[\"JobID\",\"Title\",\"Description\",\"OrgUnit\"]},\"presave\":true,\"titleField\":\"Account\",\"fields\":{\"JobID\":{\"type\":\"string\"},\"Title\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"OrgUnit\":{\"type\":\"storableref\",\"entity\":\"OrgNode\"}}}"})},core.PluginComponent{Name: "OrgNode", Object: OrgNode{}, Metadata: core.NewInfo("","OrgNode", map[string]interface{}{"descriptor":"{\"name\":\"OrgNode\",\"inherits\":\"\",\"imports\":[\"bytes\"],\"form\":{\"tabs\":{\"General\":[\"Title\",\"Description\",\"Parent\",\"Code\",\"Level\"],\"Node Attributes\":[\"Attr1\",\"Attr2\",\"Attr3\"],\"Node Data\":[\"Data1\",\"Data2\",\"Data3\"]},\"overlay\":true,\"layout\":[\"General\",\"Node Attributes\",\"Node Data\"]},\"titleField\":\"Title\",\"fields\":{\"Title\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Code\":{\"type\":\"string\"},\"Parent\":{\"type\":\"storableref\",\"entity\":\"OrgNode\"},\"Level\":{\"type\":\"string\"},\"Attr1\":{\"type\":\"string\"},\"Attr2\":{\"type\":\"string\"},\"Attr3\":{\"type\":\"string\"},\"Data1\":{\"type\":\"storableref\"},\"Data2\":{\"type\":\"storableref\"},\"Data3\":{\"type\":\"storableref\"}}}"})},
  }
}
