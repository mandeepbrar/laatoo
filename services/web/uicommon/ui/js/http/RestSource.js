import axios from 'axios';
import {formatUrl} from '../utils';

class RestDataSource {
  constructor() {
    this.ExecuteServiceObject = this.ExecuteServiceObject.bind(this);
    this.HttpCall = this.HttpCall.bind(this);
    this.buildHttpSvcResponse = this.buildHttpSvcResponse.bind(this); 
  }

  ExecuteServiceObject(service, serviceRequest, config) {
    var method = this.getMethod(service);
    var req = serviceRequest.GetRequest("http");
    let urlparams = Object.assign({}, req.urlparams, service.urlparams)
    let url = formatUrl(service.url, urlparams)
    if (!url.startsWith('http')) {
      url = Application.Backend + url;
    }
    let data = req.data
    if(service.postArgs && method == 'POST') {
      data = Object.assign({}, req.data, service.postArgs);
    }
    console.log("executing service", service, url, data, req)
    return this.HttpCall(service, url, method, req.params, data, req.headers, config);
  }

  HttpCall(service, url, method, params, data, headers, config=null) {
    let requestor = this;
    var promise = new Promise(
      function (resolve, reject) {
        if (method === "" || url === "") {
          reject(requestor.buildHttpSvcResponse(service, Response.InternalError, 'Could not build request', url));
          return;
        }
        let successCallback = function(response) {
          if (response.status < 300) {
            let res = requestor.buildHttpSvcResponse(service, Response.Success, "", response);
            resolve(res);
          } else {
            reject(requestor.buildHttpSvcResponse(service, Response.Failure, "", response));
          }
        };
        let errorCallback = function(response) {
          reject(requestor.buildHttpSvcResponse(service, Response.Failure, "", response));
        };
        if(method == 'DELETE' || method == 'GET') {
          data = null;
        }
        if(!headers) {
          headers = {}
        }
        headers[Application.Security.AuthToken] = Storage.auth;
        let req = {
          method: method,
          url: url,
          data: data,
          headers: headers,
          params: params,
          responseType: 'json'
        };
        if(config) {
          req = Object.assign({}, req, config)
        }
        console.log("Request.. ",req);

        axios(req).then(successCallback, errorCallback);
      });
    return promise;
  }

  createFullUrl(url, params) {
    if (params != null && Object.keys(params).length != 0) {
      return url + "?" + Object.keys(data).map(function(key) {
        return [key, data[key]].map(encodeURIComponent).join("=");
      }).join("&");
    }
    return url
  }

  buildHttpSvcResponse(service, code, msg, res) {
    if(res instanceof Error) {
      return this.buildSvcResponse(service, code, msg, res, {});
    }
    return this.buildSvcResponse(service, code, msg, res.data, res.headers, res.status);
  }

  buildSvcResponse(service, code, msg, data, info, statuscode) {
    var response = {};
    if(data) {
      if(service.responseField) {
        if(Array.isArray(data)) {
          data = data.map((item)=>{
            return item[service.responseField]
          })
        }
        else {
          data = data[service.responseField]
        }
      }
      if(service.transformResponse) {
        let method = _reg('Methods', service.transformResponse)
        if(method) {
          data = method(data)
        }
      }
    }

    switch (code) {
      default:
        response.code = code;
        response.message = msg;
        response.data = data;
        response.info = info;
        response.statuscode = statuscode;
    }
    console.log("service response", service, response)  
    return response;
  }

  getMethod(service) {
    if (service.method) {
      return service.method
    }
    return 'GET';
  }
}

export {
  RestDataSource
}