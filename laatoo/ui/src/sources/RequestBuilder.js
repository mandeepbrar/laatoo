'use strict';

export class RequestBuilderService {
  constructor() {
    this.ParameterSeparatorRequest = this.ParameterSeparatorRequest.bind(this);
    this.DefaultRequest = this.DefaultRequest.bind(this);
    this.URLParamsRequest = this.URLParamsRequest.bind(this);
  }
  ParameterSeparatorRequest(params, data, urlparams) {
    var parameterSeparator = {};
    if(data == null) {
      data = {};
    }
    parameterSeparator.params = params;
    parameterSeparator.data = data;
    parameterSeparator.urlparams = urlparams;
    parameterSeparator.GetRequest = function(protocol) {
      if(protocol == 'http') {
        var http = {};
        http.data = parameterSeparator.data;
        if(parameterSeparator.params == null) {
          http.params = null;
          http.urlparams = null;
          return http;
        }
        var httpUrlParams = {};
        var httpParams = {};
        var count = 0;
        if(parameterSeparator.urlparams != null) {
          for ( var param in parameterSeparator.urlparams) {
            if(param in parameterSeparator.params) {
              httpUrlParams[param] = parameterSeparator.params[param];
              count = count + 1;
            }
          }
        }
        if(count >0) {
          var remaincount = 0;
          for( var param in parameterSeparator.params) {
            if(param in httpUrlParams) {
              continue;
            } else {
              httpParams[param] = parameterSeparator.params[param];
              remaincount = remaincount + 1;
            }
          }
          if(remaincount >0) {
            http.urlparams = httpUrlParams;
            http.params = httpParams;
          } else {
            http.urlparams = httpUrlParams;
            http.params = null;
          }
          return http;
        } else {
          http.urlparams = null;
          http.params = params;
          return http;
        }
      } else {
        var socket = {};
        socket.data = parameterSeparator.data;
        socket.params = params;
        return socket;
      }
    };
    return parameterSeparator;
  }

  DefaultRequest(params, data) {
    var defaultRequest = {};
    if(data == null) {
      data = {};
    }
    defaultRequest.params = params;
    defaultRequest.data = data;
    defaultRequest.GetRequest = function(protocol) {
      var request = {};
      request.data = defaultRequest.data;
      request.params = defaultRequest.params;
      request.urlparams = null;
      return request;
    };
    return defaultRequest
  }

  URLParamsRequest(urlparams, data) {
    var urlparamsRequest = {};
    if(data == null) {
      data = {};
    }
    urlparamsRequest.data = data;
    urlparamsRequest.urlparams = urlparams;
    urlparamsRequest.GetRequest = function(protocol) {
      if(protocol == 'http') {
        var http = {};
        http.data = urlparamsRequest.data;
        http.params = null;
        http.urlparams = urlparamsRequest.urlparams;
        return http;
      } else {
        var socket = {};
        socket.data = urlparamsRequest.data;
        socket.params = urlparamsRequest.urlparams;
        return socket;
      }
    };
    return urlparamsRequest
  }
}
