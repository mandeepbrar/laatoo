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
        this.EntityPrefix = "/"
    }
    SetPrefix = (pre) => {
      this.EntityPrefix = pre
    }
    GetEntity(entityName, id, headers, svc) {
      if(svc) {
        var req = this.RequestBuilder.URLParamsRequest({":id": id}, null, headers);
  			return this.DataSource.ExecuteService(svc, req);
      } else {
        var service = {};
  			service.method = "GET";
        service.url = this.EntityPrefix + entityName.toLowerCase()+"/"+id;
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
        service.url = this.EntityPrefix+entityName.toLowerCase();
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
        service.url = this.EntityPrefix+entityName.toLowerCase()+"/"+id;
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
        service.url = this.EntityPrefix+entityName.toLowerCase()+"/"+id;
        var req = this.RequestBuilder.DefaultRequest(null, data, headers);
  			return this.DataSource.ExecuteServiceObject(service, req);
      }
		};

    ListEntities(entityName, criteria, headers, svc) {
      if(svc) {
        var req = this.RequestBuilder.URLParamsRequest(criteria, null, headers);
  			return this.DataSource.ExecuteService(svc, req);
      } else {
        var service = {};
  			service.method = "POST";
        service.url = this.EntityPrefix+entityName.toLowerCase()+"/view";
        var req = this.RequestBuilder.DefaultRequest(criteria, null, headers);
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
        service.url = this.EntityPrefix+entityName.toLowerCase()+"/"+id;
        var req = this.RequestBuilder.DefaultRequest(null, fieldmap, headers);
  			return this.DataSource.ExecuteServiceObject(service, req);
      }
		};

}
