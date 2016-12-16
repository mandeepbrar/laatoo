'use strict';

export class EntityDataService {
    constructor(DataSource, RequestBuilder) {
        this.DataSource = DataSource;
        this.RequestBuilder = RequestBuilder;
        this.GetEntity = this.GetEntity.bind(this);
        this.SaveEntity = this.SaveEntity.bind(this);
        this.DeleteEntity = this.DeleteEntity.bind(this);
        this.PutEntity = this.PutEntity.bind(this);
        this.UpdateEntity = this.UpdateEntity.bind(this);
    }
    GetEntity(entityName, id, headers, svc) {
      if(svc) {
        var req = this.RequestBuilder.URLParamsRequest({":id": id}, null, headers);
  			return this.DataSource.ExecuteService(svc, req);
      } else {
        var service = {};
  			service.method = "GET";
        service.url = document.Application.EntityPrefix + entityName.toLowerCase()+"/"+id;
        var req = this.RequestBuilder.DefaultRequest(null, null, headers);
  			return this.DataSource.ExecuteServiceObject(service, req);
      }
		}

		SaveEntity(entityName, data, headers, svc) {
      var req = this.RequestBuilder.DefaultRequest(null, data, headers);
      if(svc) {
  			return this.DataSource.ExecuteService(svc, req);
      } else {
        var service = {};
  			service.method = "POST";
        service.url = document.Application.EntityPrefix+entityName.toLowerCase();
        return this.DataSource.ExecuteServiceObject(service, req);
      }
		};

		DeleteEntity(entityName, id, headers, svc) {
      if(svc) {
        var req = this.RequestBuilder.URLParamsRequest({":id": id}, null, headers);
  			return this.DataSource.ExecuteService(svc, req);
      } else {
        var service = {};
  			service.method = "DELETE";
        service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
        var req = this.RequestBuilder.DefaultRequest(null, null, headers);
  			return this.DataSource.ExecuteServiceObject(service, req);
      }
		};

		PutEntity(entityName, id, data, headers, svc) {
      if(svc) {
        var req = this.RequestBuilder.URLParamsRequest({":id": id}, null, headers);
  			return this.DataSource.ExecuteService(svc, req);
      } else {
        var service = {};
  			service.method = "PUT";
        service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
        var req = this.RequestBuilder.DefaultRequest(null, data, headers);
  			return this.DataSource.ExecuteServiceObject(service, req);
      }
		};

		UpdateEntity(entityName, id, fieldmap, headers, svc) {
      if(svc) {
        var req = this.RequestBuilder.URLParamsRequest({":id": id}, null, headers);
  			return this.DataSource.ExecuteService(svc, req);
      } else {
        var service = {};
  			service.method = "PUT";
        service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
        var req = this.RequestBuilder.DefaultRequest(null, fieldmap, headers);
  			return this.DataSource.ExecuteServiceObject(service, req);
      }
		};

}
