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
    GetEntity(entityName, id, headers, svcUrl) {
			var service = {};
			service.method = "GET";
      if(svcUrl) {
        service.url = svcUrl +"/"+id;
      } else {
        service.url = document.Application.EntityPrefix + entityName.toLowerCase()+"/"+id;
      }
			var req = this.RequestBuilder.DefaultRequest(null, null, headers);
			return this.DataSource.ExecuteServiceObject(service, req);
		}

		SaveEntity(entityName, data, headers, svcUrl) {
			var service = {};
			service.method = "POST";
      if(svcUrl) {
        service.url = svcUrl
      } else {
        service.url = document.Application.EntityPrefix+entityName.toLowerCase();
      }
			var req = this.RequestBuilder.DefaultRequest(null, data, headers);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

		DeleteEntity(entityName, id, headers, svcUrl) {
			var service = {};
			service.method = "DELETE";
      if(svcUrl) {
        service.url = svcUrl +"/"+id;
      } else {
        service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
      }
			var req = this.RequestBuilder.DefaultRequest(null, null, headers);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

		PutEntity(entityName, id, data, headers, svcUrl) {
			var service = {};
			service.method = "PUT";
      if(svcUrl) {
        service.url = svcUrl +"/"+id;
      } else {
        service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
      }
			var req = this.RequestBuilder.DefaultRequest(null, data, headers);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

		UpdateEntity(entityName, id, fieldmap, headers, svcUrl) {
			var service = {};
			service.method = "PUT";
      if(svcUrl) {
        service.url = svcUrl +"/"+id;
      } else {
        service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
      }
			var req = this.RequestBuilder.DefaultRequest(null, fieldmap, headers);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

}
