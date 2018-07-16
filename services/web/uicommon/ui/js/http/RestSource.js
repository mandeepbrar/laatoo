import axios from 'axios';

class RestDataSource {
  constructor() {
    this.ExecuteServiceObject = this.ExecuteServiceObject.bind(this);
    this.HttpCall = this.HttpCall.bind(this);
    this.buildHttpSvcResponse = this.buildHttpSvcResponse.bind(this);  
  }

  ExecuteServiceObject(service, serviceRequest, config) {
    var method = this.getMethod(service);
    var req = serviceRequest.GetRequest("http");
    var url = this.getURL(service, req);
    return this.HttpCall(url, method, req.params, req.data, req.headers, config);
  }

  HttpCall(url, method, params, data, headers, config=null) {
    let service = this;
    var promise = new Promise(
      function (resolve, reject) {
        if (method === "" || url === "") {
          reject(service.buildHttpSvcResponse(Response.InternalError, 'Could not build request', url));
          return;
        }
        let successCallback = function(response) {
          if (response.status < 300) {
            let res = service.buildHttpSvcResponse(Response.Success, "", response);
            resolve(res);
          } else {
            reject(service.buildHttpSvcResponse(Response.Failure, "", response));
          }
        };
        let errorCallback = function(response) {
          reject(service.buildHttpSvcResponse(Response.Failure, "", response));
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

  buildHttpSvcResponse(code, msg, res) {
    if(res instanceof Error) {
      return this.buildSvcResponse(code, msg, res, {});
    }
    return this.buildSvcResponse(code, msg, res.data, res.headers, res.status);
  }

  buildSvcResponse(code, msg, data, info, statuscode) {
    var response = {};
    switch (code) {
      default:
        response.code = code;
        response.message = msg;
        response.data = data;
        response.info = info;
        response.statuscode = statuscode;
    }
    console.log(response);
    return response;
  }

  getURL(service, req) {
    var url = service.url;
    if (req.urlparams != null) {
      for (var param in req.urlparams) {
        url = url.replace(":" + param, req.urlparams[param]);
      }
    }
    if (url.startsWith('http')) {
      return url;
    } else {
      return Application.Backend + url;
    }
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