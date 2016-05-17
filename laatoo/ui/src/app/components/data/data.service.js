(function() {
    'use strict';

    var mod = angular.module('data');

	var Response = {"Success":1, "Unauthorized":2, "InternalError":3, "BadRequest":4, "Failure":5}

    mod.factory('DataService', DataService);

    /** @ngInject */
    function DataService($http) {
		var datasvc={};
		
		datasvc.ExecuteService = function(serviceName, serviceRequest, successMethod, errorMethod) {
			var service = document.Services[serviceName];
			if (service != null && serviceRequest != null) {
				return datasvc.ExecuteServiceObject(service, serviceRequest, successMethod, errorMethod);
			} else {
				errorMethod(buildSvcResponse(Response.InternalError,'Service not found', serviceName));
				return null;
			}
		}
		datasvc.ExecuteServiceObject = function(service, serviceRequest, successMethod, errorMethod) {
			if (service != null && serviceRequest != null) {
				var protocol = getProtocol();
				var req = serviceRequest.GetRequest(protocol);
				if(protocol === 'http') {
					var method = getMethod(service);
					var url = getURL(service, req);
					return datasvc.HttpCall(url, method, req.params, req.data, successMethod, errorMethod);
				}				
			} else {
				errorMethod(buildSvcResponse(Response.InternalError, 'Invalid Request', ""));
				return null;
			}
		};
		
		datasvc.HttpCall = function(url, method, params, data, successMethod, errorMethod) {
			if(method === "" || url === "") {
				errorMethod(buildHttpSvcResponse(Response.InternalError,'Could not build request', url));
				return;				
			}
			console.log(method);
			console.log(url);
			console.log(data);
			return $http({
				method: method,
				url: url,
				params: params,
				data: data
			}).then(
			function successCallback(response) {
				console.log("url"+url);
				console.log(response);
				successMethod(buildHttpSvcResponse(Response.Success, "", response));	
			}, 
			function errorCallback(response) {
				console.log("url"+url);
				console.log(response);
				errorMethod(buildHttpSvcResponse(Response.Failure, "", response));
			});			
			
		};
		
        return datasvc;
    };
	
	function buildHttpSvcResponse(code, msg, res) {
		return buildSvcResponse(code, msg, res.data, res.headers());
	};
	
	function buildSvcResponse(code, msg, data, info) {
		var response = {};
		switch (code) {
			case Response.Success:
				response.code = code;
				response.message = msg;
				response.data = data;
				response.info = info;
			break;
			case Response.Unauthorized:
			break;
			case Response.InternalError:
				response.code = code;
				response.message = msg;
				response.data = data;
				response.info = info;
			break;
			case Response.BadRequest:
			break;
			case Response.Failure:
			break;
		}
		return response;
	};

	function getURL(service, req) {
		var url = service.url;
		if(req.urlparams!=null) {
			for (var param in req.urlparams) {
				url = url.replace(":"+param, req.urlparams[param]);
			}
		}
		if(url.startsWith('http'))  {
			return url;
		} else {
			return document.Application.Backend + url;			
		}
	};
	
	function getMethod(service) {
		if(service.method) {
			return service.method
		}
		return 'GET';
	};
	
	function getProtocol() {
		return 'http';
	};
})();