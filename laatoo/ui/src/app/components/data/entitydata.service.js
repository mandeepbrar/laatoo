(function() {
    'use strict';

    var mod = angular.module('data');

    mod.factory('EntityDataService', EntityDataService);

    /** @ngInject */
    function EntityDataService(DataService, RequestBuilderService) {
		var entitydatasvc={}
		entitydatasvc.GetEntity = function(entityName, id, successMethod, errorMethod) {
			var service = {};
			service.method = "GET";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
			var req = RequestBuilderService.DefaultRequest(null, null);
			return DataService.ExecuteServiceObject(service, req, successMethod, errorMethod);
		};
		
		entitydatasvc.SaveEntity = function(entityName, data, successMethod, errorMethod) {
			var service = {};
			service.method = "POST";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase();
			var req = RequestBuilderService.DefaultRequest(null, data);
			return DataService.ExecuteServiceObject(service, req, successMethod, errorMethod);
		};
		
		entitydatasvc.DeleteEntity = function(entityName, id, successMethod, errorMethod) {
			var service = {};
			service.method = "DELETE";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
			var req = RequestBuilderService.DefaultRequest(null, null);
			return DataService.ExecuteServiceObject(service, req, successMethod, errorMethod);
		};

		entitydatasvc.PutEntity = function(entityName, id, data, successMethod, errorMethod) {
			var service = {};
			service.method = "PUT";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
			var req = RequestBuilderService.DefaultRequest(null, data);
			return DataService.ExecuteServiceObject(service, req, successMethod, errorMethod);
		};

				
		entitydatasvc.UpdateEntity = function(entityName, id, fieldmap, successMethod, errorMethod) {
			var service = {};
			service.method = "PUT";
			service.url = document.Application.EntityPrefix+entityName.toLowerCase()+"/"+id;
			var req = RequestBuilderService.DefaultRequest(null, fieldmap);
			return DataService.ExecuteServiceObject(service, req, successMethod, errorMethod);
		};
						
        return entitydatasvc;
    };
})();