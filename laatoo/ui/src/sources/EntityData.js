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
    GetEntity(entityName, id) {
			var service = {};
			service.method = "GET";
			service.url = document.Application.EntityPrefix + entityName.toLowerCase()+"/"+id;
			var req = this.RequestBuilder.DefaultRequest(null, null);
			return this.DataSource.ExecuteServiceObject(service, req);
		}

		SaveEntity(entityName, data) {
			var service = {};
			service.method = "POST";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase();
			var req = this.RequestBuilder.DefaultRequest(null, data);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

		DeleteEntity(entityName, id) {
			var service = {};
			service.method = "DELETE";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
			var req = this.RequestBuilder.DefaultRequest(null, null);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

		PutEntity(entityName, id, data) {
			var service = {};
			service.method = "PUT";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
			var req = this.RequestBuilder.DefaultRequest(null, data);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

		UpdateEntity(entityName, id, fieldmap) {
			var service = {};
			service.method = "PUT";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
			var req = this.RequestBuilder.DefaultRequest(null, fieldmap);
			return this.DataSource.ExecuteServiceObject(service, req);
		};

}
