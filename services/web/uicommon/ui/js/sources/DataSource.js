import { RequestBuilderService } from './RequestBuilder';
import { EntityDataService } from './EntityData';

console.log("uicommon - application services", Application)

class DefaultDataSource {
  constructor() {
    this.ExecuteService = this.ExecuteService.bind(this);
    this.ExecuteServiceObject = this.ExecuteServiceObject.bind(this);
  }
  ExecuteService(serviceName, serviceRequest, config=null) {
    var service = _reg('Services', serviceName);
    if (service != null && serviceRequest != null) {
      return this.ExecuteServiceObject(service, serviceRequest, config);
    } else {
      throw new Error('Service not found ' + serviceName);
    }
  }
  ExecuteServiceObject(service, serviceRequest, config=null) {
    if(!service.protocol) {
      service.protocol = "http"
    }
    if (service != null && serviceRequest != null) {
      var handler =_reg("DataSourceHandlers", service.protocol)
      if(handler == null) {
        console.log("Requested service for handler", service)
        throw new Error('Invalid protocol handler');        
      }
      return handler.ExecuteServiceObject(service, serviceRequest, config)
    }
  }  
}

const DataSource = new DefaultDataSource();
const RequestBuilder = new RequestBuilderService();
const EntityData = new EntityDataService(DataSource, RequestBuilder);
export {
  RequestBuilder,
  DataSource,
  EntityData,
};
