export class RequestBuilderService {
  constructor() {
    this.ParameterSeparatorRequest = this.ParameterSeparatorRequest.bind(this);
    this.DefaultRequest = this.DefaultRequest.bind(this);
    this.URLParamsRequest = this.URLParamsRequest.bind(this);
  }
  ParameterSeparatorRequest(params, data, urlparams, headers) {
    var parameterSeparator = {};
    if(data == null) {
      data = {};
    }
    parameterSeparator.params = params;
    parameterSeparator.data = data;
    parameterSeparator.urlparams = urlparams;
    parameterSeparator.headers = headers
    parameterSeparator.GetRequest = function(protocol) {
      if(protocol == 'http') {
        var http = {};
        http.data = parameterSeparator.data;
        http.headers = parameterSeparator.headers;
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
        socket.params = Object.assign({}, params, headers);
        return socket;
      }
    };
    return parameterSeparator;
  }

  DefaultRequest(params, data, headers) {
    var defaultRequest = {};
    if(data == null) {
      data = {};
    }
    defaultRequest.params = params;
    defaultRequest.data = data;
    defaultRequest.headers = headers
    defaultRequest.GetRequest = function(protocol) {
      var request = {};
      request.data = defaultRequest.data;
      request.params = defaultRequest.params;
      request.urlparams = null;
      request.headers = defaultRequest.headers
      return request;
    };
    return defaultRequest
  }

  URLParamsRequest(urlparams, data, headers) {
    var urlparamsRequest = {};
    if(data == null) {
      data = {};
    }
    urlparamsRequest.data = data;
    urlparamsRequest.urlparams = urlparams;
    urlparamsRequest.headers = headers
    urlparamsRequest.GetRequest = function(protocol) {
      if(protocol == 'http') {
        var http = {};
        http.data = urlparamsRequest.data;
        http.params = null;
        http.urlparams = urlparamsRequest.urlparams;
        http.headers = urlparamsRequest.headers
        return http;
      } else {
        var socket = {};
        socket.data = urlparamsRequest.data;
        socket.params = Object.assign({}, urlparamsRequest.urlparams, urlparamsRequest.headers);
        return socket;
      }
    };
    return urlparamsRequest
  }
}
