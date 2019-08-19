package main

import (
	"laatoo/sdk/server/core"
)


func ObjectsManifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{
    core.PluginComponent{Name: "Application", Object: Application{}, Metadata: core.NewInfo("","Application", map[string]interface{}{"descriptor":"{\"name\":\"Application\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"ServerTemp\",\"EnvironmentTemp\",\"LoggingLevel\",\"LoggingFormat\"],\"Module Instances\":[\"Instances\"],\"Services\":[\"Services\"],\"Channels\":[\"Channels\"],\"Engines\":[\"Engines\"],\"Factories\":[\"Factories\"],\"Rules\":[\"Rules\"],\"Tasks\":[\"Tasks\"],\"Entities\":[\"Entities\"],\"Security\":[\"Security\"]},\"verticaltabs\":true,\"overlay\":true,\"actions\":\"AbstractServer_Actions\",\"submit\":\"Save\",\"layout\":[\"General\",\"Module Instances\",\"Entities\",\"Factories\",\"Services\",\"Engines\",\"Channels\",\"Rules\",\"Tasks\",\"Security\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"ServerTemp\":{\"type\":\"entity\",\"entity\":\"Server\",\"label\":\"Server Template\"},\"EnvironmentTemp\":{\"type\":\"entity\",\"entity\":\"Environment\",\"label\":\"Environment Template\"},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Objects\":{\"type\":\"string\",\"list\":true},\"Instances\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"ModuleInstance\"},\"Services\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Service\"},\"Entities\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Entity\"},\"Factories\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Factory\"},\"Channels\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Channel\"},\"Engines\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Engine\"},\"Rules\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Rule\"},\"Tasks\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Task\"},\"Security\":{\"type\":\"subentity\",\"entity\":\"Security\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"}}}}"})},core.PluginComponent{Name: "Channel", Object: Channel{}, Metadata: core.NewInfo("","Channel", map[string]interface{}{"descriptor":"{\"name\":\"Channel\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\",\"Description\",\"Service\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Service\":{\"type\":\"string\"}}}"})},core.PluginComponent{Name: "Configuration", Object: Configuration{}, Metadata: core.NewInfo("","Configuration", map[string]interface{}{"descriptor":"{\"name\":\"Configuration\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"collection\":\"<nocollection>\",\"form\":{\"layout\":[\"Name\",\"Values\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Values\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"KeyValue\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\",\"props\":{\"mode\":\"dialog\"}}}}}"})},core.PluginComponent{Name: "Deployment", Object: Deployment{}, Metadata: core.NewInfo("","Deployment", map[string]interface{}{"descriptor":"{\"name\":\"Deployment\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"form\":{\"layout\":[\"Name\",\"Description\",\"Solution\",\"Server\",\"Environment\",\"Application\"]},\"titleField\":\"Name\",\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"Environment\":{\"type\":\"entity\",\"entity\":\"Environment\"},\"Server\":{\"type\":\"entity\",\"entity\":\"Server\"},\"Application\":{\"type\":\"entity\",\"entity\":\"Application\"}}}"})},core.PluginComponent{Name: "Engine", Object: Engine{}, Metadata: core.NewInfo("","Engine", map[string]interface{}{"descriptor":"{\"name\":\"Engine\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"EngineType\",\"Address\",\"SSL\"],\"HTML\":[\"Path\",\"CORS\",\"CORSHosts\",\"Framework\",\"QueryParams\"]},\"layout\":[\"General\",\"HTML\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"EngineType\":{\"type\":\"string\"},\"Address\":{\"type\":\"string\"},\"Framework\":{\"type\":\"string\"},\"SSL\":{\"type\":\"bool\"},\"CORS\":{\"type\":\"bool\"},\"Path\":{\"type\":\"string\"},\"CORSHosts\":{\"type\":\"string\",\"list\":true},\"QueryParams\":{\"type\":\"string\",\"list\":true}}}"})},core.PluginComponent{Name: "Entity", Object: Entity{}, Metadata: core.NewInfo("","Entity", map[string]interface{}{"descriptor":"{\"name\":\"Entity\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\"],\"Fields\":[\"Fields\"]},\"layout\":[\"General\",\"Fields\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Fields\":{\"type\":\"subentity\",\"list\":true,\"entity\":\"Field\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\",\"props\":{\"mode\":\"dialog\"}}}}}"})},core.PluginComponent{Name: "Environment", Object: Environment{}, Metadata: core.NewInfo("","Environment", map[string]interface{}{"descriptor":"{\"name\":\"Environment\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"ServerTemp\",\"LoggingLevel\",\"LoggingFormat\"],\"Module Instances\":[\"Instances\"],\"Services\":[\"Services\"],\"Channels\":[\"Channels\"],\"Engines\":[\"Engines\"],\"Factories\":[\"Factories\"],\"Rules\":[\"Rules\"],\"Tasks\":[\"Tasks\"],\"Entities\":[\"Entities\"],\"Security\":[\"Security\"]},\"verticaltabs\":true,\"overlay\":true,\"actions\":\"AbstractServer_Actions\",\"submit\":\"Save\",\"layout\":[\"General\",\"Module Instances\",\"Entities\",\"Factories\",\"Services\",\"Engines\",\"Channels\",\"Rules\",\"Tasks\",\"Security\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"ServerTemp\":{\"type\":\"subentity\",\"entity\":\"Server\",\"label\":\"Server Template\",\"ref\":true},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Objects\":{\"type\":\"string\",\"list\":true},\"Instances\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"ModuleInstance\"},\"Services\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Service\"},\"Entities\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Entity\"},\"Factories\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Factory\"},\"Channels\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Channel\"},\"Engines\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Engine\"},\"Rules\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Rule\"},\"Tasks\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Task\"},\"Security\":{\"type\":\"subentity\",\"entity\":\"Security\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"}}}}"})},core.PluginComponent{Name: "Factory", Object: Factory{}, Metadata: core.NewInfo("","Factory", map[string]interface{}{"descriptor":"{\"name\":\"Factory\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\",\"Description\",\"Object\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Object\":{\"type\":\"string\"}}}"})},core.PluginComponent{Name: "Field", Object: Field{}, Metadata: core.NewInfo("","Field", map[string]interface{}{"descriptor":"{\"name\":\"Field\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"collection\":\"<nocollection>\",\"form\":{\"layout\":[\"Name\",\"Type\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}"})},core.PluginComponent{Name: "ModuleInstance", Object: ModuleInstance{}, Metadata: core.NewInfo("","ModuleInstance", map[string]interface{}{"descriptor":"{\"name\":\"ModuleInstance\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"tabs\":{\"General\":[\"Name\",\"Module\",\"Description\",\"LoggingLevel\",\"LoggingFormat\"],\"Settings\":[\"Settings\"]},\"layout\":[\"General\",\"Settings\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Module\":{\"type\":\"storableref\",\"entity\":\"Module\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\",\"props\":{\"loadData\":true,\"mode\":\"select\",\"dataService\":\"entitlementmodule\"}}},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Description\":{\"type\":\"string\"},\"Settings\":{\"type\":\"stringmap\",\"widget\":{\"name\":\"ModuleConfig\",\"module\":\"modulesrepository\"}}}}"})},core.PluginComponent{Name: "KeyValue", Object: KeyValue{}, Metadata: core.NewInfo("","KeyValue", map[string]interface{}{"descriptor":"{\"name\":\"KeyValue\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"collection\":\"<nocollection>\",\"form\":{\"layout\":[\"Name\",\"Value\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}"})},core.PluginComponent{Name: "Module", Object: Module{}, Metadata: core.NewInfo("","Module", map[string]interface{}{"descriptor":"{\"name\":\"Module\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\",\"Instance\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"bool\"},\"Objects\":{\"type\":\"string\",\"list\":true},\"Params\":{\"type\":\"stringmap\",\"mappedElement\":\"ModuleParam\"},\"ParamsForm\":{\"type\":\"stringmap\"},\"Dependencies\":{\"type\":\"stringmap\"},\"Services\":{\"type\":\"string\",\"list\":true},\"Factories\":{\"type\":\"string\",\"list\":true},\"Channels\":{\"type\":\"string\",\"list\":true},\"Engines\":{\"type\":\"string\",\"list\":true},\"Rules\":{\"type\":\"string\",\"list\":true},\"Tasks\":{\"type\":\"string\",\"list\":true}}}"})},core.PluginComponent{Name: "Rule", Object: Rule{}, Metadata: core.NewInfo("","Rule", map[string]interface{}{"descriptor":"{\"name\":\"Rule\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\",\"Description\",\"Trigger\",\"MessageType\",\"Rule\",\"LoggingLevel\",\"LoggingFormat\",\"Configuration\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Trigger\":{\"type\":\"string\"},\"MessageType\":{\"type\":\"string\",\"label\":\"Message\"},\"Rule\":{\"type\":\"string\"},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Configuration\":{\"type\":\"subentity\",\"entity\":\"Configuration\",\"mode\":\"dialog\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"}}}}"})},core.PluginComponent{Name: "Security", Object: Security{}, Metadata: core.NewInfo("","Security", map[string]interface{}{"descriptor":"{\"name\":\"Security\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Mode\",\"RoleSvc\",\"Publickey\",\"Privatekey\",\"Realm\",\"Supportedrealms\",\"Authservices\",\"Permissions\"]},\"fields\":{\"Mode\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Local\",\"value\":\"local\"},{\"text\":\"Remote\",\"value\":\"remote\"}]}}},\"RoleSvc\":{\"label\":\"Role Data Service\",\"type\":\"string\"},\"Publickey\":{\"type\":\"string\"},\"Privatekey\":{\"type\":\"string\"},\"Realm\":{\"type\":\"string\",\"label\":\"Message\"},\"Supportedrealms\":{\"type\":\"string\",\"label\":\"Supported Realms\",\"list\":true},\"Authservices\":{\"label\":\"Authentication Services\",\"type\":\"string\",\"list\":true},\"Permissions\":{\"label\":\"Permissions\",\"type\":\"string\",\"list\":true}}}"})},core.PluginComponent{Name: "Server", Object: Server{}, Metadata: core.NewInfo("","Server", map[string]interface{}{"descriptor":"{\"name\":\"Server\",\"inherits\":\"\",\"imports\":[],\"multitenant\":true,\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"LoggingLevel\",\"LoggingFormat\"],\"Module Instances\":[\"Instances\"],\"Services\":[\"Services\"],\"Channels\":[\"Channels\"],\"Engines\":[\"Engines\"],\"Factories\":[\"Factories\"],\"Rules\":[\"Rules\"],\"Tasks\":[\"Tasks\"],\"Entities\":[\"Entities\"],\"Security\":[\"Security\"]},\"verticaltabs\":true,\"overlay\":true,\"actions\":\"AbstractServer_Actions\",\"submit\":\"Save\",\"layout\":[\"General\",\"Module Instances\",\"Entities\",\"Factories\",\"Services\",\"Engines\",\"Channels\",\"Rules\",\"Tasks\",\"Security\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Solution\":{\"type\":\"storableref\",\"entity\":\"Solution\"},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Objects\":{\"type\":\"string\",\"list\":true},\"Instances\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"ModuleInstance\"},\"Services\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Service\"},\"Entities\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Entity\"},\"Factories\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Factory\"},\"Channels\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Channel\"},\"Engines\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Engine\"},\"Rules\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Rule\"},\"Tasks\":{\"type\":\"subentity\",\"list\":true,\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"},\"entity\":\"Task\"},\"Security\":{\"type\":\"subentity\",\"entity\":\"Security\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\"}}}}"})},core.PluginComponent{Name: "Service", Object: Service{}, Metadata: core.NewInfo("","Service", map[string]interface{}{"descriptor":"{\"name\":\"Service\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"tabs\":{\"General\":[\"Name\",\"Description\",\"Factory\",\"ServiceMethod\",\"LoggingLevel\",\"LoggingFormat\"],\"Configuration\":[\"Configuration\"]},\"layout\":[\"General\",\"Configuration\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Factory\":{\"type\":\"string\"},\"ServiceMethod\":{\"type\":\"string\",\"label\":\"Service Method\"},\"LoggingLevel\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Error\",\"value\":\"error\"},{\"text\":\"Warn\",\"value\":\"warn\"},{\"text\":\"Info\",\"value\":\"info\"},{\"text\":\"Debug\",\"value\":\"debug\"},{\"text\":\"Trace\",\"value\":\"trace\"}]}}},\"LoggingFormat\":{\"type\":\"string\",\"widget\":{\"name\":\"Select\",\"props\":{\"items\":[{\"text\":\"Json \",\"value\":\"json\"},{\"text\":\"Json Full\",\"value\":\"jsonmax\"},{\"text\":\"Formatted\",\"value\":\"happy\"},{\"text\":\"Formatted Full\",\"value\":\"happymax\"}]}}},\"Configuration\":{\"type\":\"subentity\",\"entity\":\"Configuration\",\"widget\":{\"name\":\"SubEntity\",\"module\":\"subentity\",\"props\":{\"mode\":\"dialog\"}}}}}"})},core.PluginComponent{Name: "Task", Object: Task{}, Metadata: core.NewInfo("","Task", map[string]interface{}{"descriptor":"{\"name\":\"Task\",\"inherits\":\"\",\"imports\":[],\"titleField\":\"Name\",\"form\":{\"layout\":[\"Name\",\"Description\",\"Receiver\",\"Processor\"]},\"fields\":{\"Name\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Receiver\":{\"type\":\"string\"},\"Processor\":{\"type\":\"string\",\"label\":\"Message\"}}}"})},
  }
}
