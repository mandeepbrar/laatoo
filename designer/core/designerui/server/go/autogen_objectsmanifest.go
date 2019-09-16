package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Application", Object: Application{}, Metadata: core.NewInfo("","Application", map[string]interface{}{"descriptor":"{\"name\":\"Application\",\"inherits\":\"\",\"imports\":[\"laatoo/sdk/modules/modulesbase\"],\"multitenant\":true,\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"ServerTemp\",\"EnvironmentTemp\",\"LoggingLevel\",\"LoggingFormat\"],\"Module Instances\":[\"Instances\"],\"Services\":[\"Services\"],\"Channels\":[\"Channels\"],\"Engines\":[\"Engines\"],\"Factories\":[\"Factories\"],\"Rules\":[\"Rules\"],\"Tasks\":[\"Tasks\"],\"Entities\":[\"Entities\"],\"Security\":[\"Security\"]},\"verticaltabs\":true,\"overlay\":true,\"actions\":\"AbstractServer_Actions\",\"submit\":\"Save\",\"layout\":[\"General\",\"Module Instances\",\"Entities\",\"Factories\",\"Services\",\"Engines\",\"Channels\",\"Rules\",\"Tasks\",\"Security\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"ServerTemp\":{\"type\":\"entity\",\"entity\":\"Server\",\"label\":\"Server Template\"},\"EnvironmentTemp\":{\"type\":\"entity\",\"entity\":\"Environment\",\"label\":\"Environment Template\"},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Objects\":{\"type\":\"string\",\"list\":true},\"Instances\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.ModuleInstance\"},\"Services\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Service\"},\"Entities\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Entity\"},\"Factories\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Factory\"},\"Channels\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Channel\"},\"Engines\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Engine\"},\"Rules\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Rule\"},\"Tasks\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Task\"},\"Security\":{\"type\":\"subentity\",\"entity\":\"Security\"}}}"})},core.PluginComponent{Name: "Deployment", Object: Deployment{}, Metadata: core.NewInfo("","Deployment", map[string]interface{}{"descriptor":"{\"name\":\"Deployment\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"collection\":\"<nocollection>\",\"form\":{\"layout\":[\"Name\",\"Description\",\"Solution\",\"Server\",\"Environment\",\"Application\"]},\"titleField\":\"Name\",\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"Environment\":{\"type\":\"entity\",\"entity\":\"Environment\"},\"Server\":{\"type\":\"entity\",\"entity\":\"Server\"},\"Application\":{\"type\":\"entity\",\"entity\":\"Application\"}}}"})},core.PluginComponent{Name: "Environment", Object: Environment{}, Metadata: core.NewInfo("","Environment", map[string]interface{}{"descriptor":"{\"name\":\"Environment\",\"inherits\":\"\",\"imports\":[\"laatoo/sdk/modules/modulesbase\"],\"multitenant\":true,\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"ServerTemp\",\"LoggingLevel\",\"LoggingFormat\"],\"Module Instances\":[\"Instances\"],\"Services\":[\"Services\"],\"Channels\":[\"Channels\"],\"Engines\":[\"Engines\"],\"Factories\":[\"Factories\"],\"Rules\":[\"Rules\"],\"Tasks\":[\"Tasks\"],\"Entities\":[\"Entities\"],\"Security\":[\"Security\"]},\"verticaltabs\":true,\"overlay\":true,\"actions\":\"AbstractServer_Actions\",\"submit\":\"Save\",\"layout\":[\"General\",\"Module Instances\",\"Entities\",\"Factories\",\"Services\",\"Engines\",\"Channels\",\"Rules\",\"Tasks\",\"Security\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"ServerTemp\":{\"type\":\"subentity\",\"entity\":\"Server\",\"label\":\"Server Template\",\"ref\":true},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Objects\":{\"type\":\"string\",\"list\":true},\"Instances\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.ModuleInstance\"},\"Services\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Service\"},\"Entities\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Entity\"},\"Factories\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Factory\"},\"Channels\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Channel\"},\"Engines\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Engine\"},\"Rules\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Rule\"},\"Tasks\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Task\"},\"Security\":{\"type\":\"subentity\",\"entity\":\"Security\"}}}"})},core.PluginComponent{Name: "Security", Object: Security{}, Metadata: core.NewInfo("","Security", map[string]interface{}{"descriptor":"{\"name\":\"Security\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"collection\":\"<nocollection>\",\"form\":{\"layout\":[\"Mode\",\"RoleSvc\",\"Publickey\",\"Privatekey\",\"Realm\",\"Supportedrealms\",\"Authservices\",\"Permissions\"]},\"fields\":{\"Mode\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Local\",\"value\":\"local\"},{\"text\":\"Remote\",\"value\":\"remote\"}]}}},\"RoleSvc\":{\"label\":\"Role Data Service\",\"type\":\"string\"},\"Publickey\":{\"type\":\"string\"},\"Privatekey\":{\"type\":\"string\"},\"Realm\":{\"type\":\"string\",\"label\":\"Message\"},\"Supportedrealms\":{\"type\":\"string\",\"label\":\"Supported Realms\",\"widget\":{\"props\":{\"mode\":\"summary\"}},\"list\":true},\"Authservices\":{\"label\":\"Authentication Services\",\"type\":\"string\",\"widget\":{\"props\":{\"mode\":\"summary\"}},\"list\":true},\"Permissions\":{\"label\":\"Permissions\",\"type\":\"string\",\"list\":true}}}"})},core.PluginComponent{Name: "Server", Object: Server{}, Metadata: core.NewInfo("","Server", map[string]interface{}{"descriptor":"{\"name\":\"Server\",\"inherits\":\"\",\"imports\":[\"laatoo/sdk/modules/modulesbase\"],\"multitenant\":true,\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"LoggingLevel\",\"LoggingFormat\"],\"Module Instances\":[\"Instances\"],\"Services\":[\"Services\"],\"Channels\":[\"Channels\"],\"Engines\":[\"Engines\"],\"Factories\":[\"Factories\"],\"Rules\":[\"Rules\"],\"Tasks\":[\"Tasks\"],\"Entities\":[\"Entities\"],\"Security\":[\"Security\"]},\"verticaltabs\":true,\"overlay\":true,\"actions\":\"AbstractServer_Actions\",\"submit\":\"Save\",\"layout\":[\"General\",\"Module Instances\",\"Entities\",\"Factories\",\"Services\",\"Engines\",\"Channels\",\"Rules\",\"Tasks\",\"Security\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Objects\":{\"type\":\"string\",\"list\":true},\"Instances\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.ModuleInstance\"},\"Services\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Service\"},\"Entities\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Entity\"},\"Factories\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Factory\"},\"Channels\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Channel\"},\"Engines\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Engine\"},\"Rules\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Rule\"},\"Tasks\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"modulesbase.Task\"},\"Security\":{\"type\":\"subentity\",\"entity\":\"Security\"}}}"})},
  }
}
