(function webpackUniversalModuleDefinition(root, factory) {
	if(typeof exports === 'object' && typeof module === 'object')
		module.exports = factory(require("react"), require("redux-saga"), require("react-redux"), require("redux"), require("babel-polyfill"), require("md5"));
	else if(typeof define === 'function' && define.amd)
		define(["react", "redux-saga", "react-redux", "redux", "babel-polyfill", "md5"], factory);
	else if(typeof exports === 'object')
		exports["laatoocommon"] = factory(require("react"), require("redux-saga"), require("react-redux"), require("redux"), require("babel-polyfill"), require("md5"));
	else
		root["laatoocommon"] = factory(root["react"], root["redux-saga"], root["react-redux"], root["redux"], root["babel-polyfill"], root["md5"]);
})(this, function(__WEBPACK_EXTERNAL_MODULE_6__, __WEBPACK_EXTERNAL_MODULE_7__, __WEBPACK_EXTERNAL_MODULE_10__, __WEBPACK_EXTERNAL_MODULE_22__, __WEBPACK_EXTERNAL_MODULE_64__, __WEBPACK_EXTERNAL_MODULE_65__) {
return /******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;
/******/
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			exports: {},
/******/ 			id: moduleId,
/******/ 			loaded: false
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "./assets/";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var _Globals = __webpack_require__(4);
	
	var _DataSource = __webpack_require__(5);
	
	var _reducers = __webpack_require__(47);
	
	var _ActionNames = __webpack_require__(2);
	
	var _utils = __webpack_require__(3);
	
	var _View = __webpack_require__(19);
	
	var _Entity = __webpack_require__(18);
	
	var _reduxSaga = __webpack_require__(7);
	
	var _reduxSaga2 = _interopRequireDefault(_reduxSaga);
	
	var _sagas = __webpack_require__(52);
	
	var _GroupLoad = __webpack_require__(42);
	
	var _View2 = __webpack_require__(45);
	
	var _ViewData = __webpack_require__(17);
	
	var _LoginComponent = __webpack_require__(43);
	
	__webpack_require__(64);
	
	var _gurmukhikeymap = __webpack_require__(55);
	
	var _gurmukhikeymap2 = _interopRequireDefault(_gurmukhikeymap);
	
	var _colors = __webpack_require__(41);
	
	var _colors2 = _interopRequireDefault(_colors);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	function _toConsumableArray(arr) { if (Array.isArray(arr)) { for (var i = 0, arr2 = Array(arr.length); i < arr.length; i++) { arr2[i] = arr[i]; } return arr2; } else { return Array.from(arr); } }
	
	var redux = __webpack_require__(22);
	
	
	/*
	
	*/
	
	function createStore(reducers, initialState, middleware, sagas, enhancers) {
	  var sagaMiddleware = (0, _reduxSaga2.default)();
	  enhancers = redux.compose.apply(redux, [redux.applyMiddleware.apply(redux, [sagaMiddleware].concat(_toConsumableArray(middleware)))].concat(_toConsumableArray(enhancers)));
	  if (!reducers) {
	    reducers = {};
	  }
	  // mount it on the Store
	  var store = redux.createStore(redux.combineReducers(reducers), initialState, enhancers);
	
	  // then run the saga
	  (0, _sagas.runSagas)(sagaMiddleware, sagas);
	  return store;
	}
	
	console.log("color from laatoo ", _colors2.default);
	var moduleExports = {
	  Storage: _Globals.Storage,
	  Application: _Globals.Application,
	  Window: _Globals.Window,
	  GurmukhiKeymap: _gurmukhikeymap2.default,
	  RequestBuilder: _DataSource.RequestBuilder,
	  DataSource: _DataSource.DataSource,
	  Response: _DataSource.Response,
	  EntityData: _DataSource.EntityData,
	  Reducers: _reducers.Reducers,
	  ViewReducer: _View.ViewReducer,
	  Colors: _colors2.default,
	  View: _View2.View,
	  ViewData: _ViewData.ViewData,
	  EntityReducer: _Entity.EntityReducer,
	  LoginComponent: _LoginComponent.LoginComponent,
	  ActionNames: _ActionNames.ActionNames,
	  formatUrl: _utils.formatUrl,
	  createStore: createStore,
	  createAction: _utils.createAction,
	  LaatooError: _utils.LaatooError,
	  hasPermission: _utils.hasPermission,
	  GroupLoad: _GroupLoad.GroupLoad,
	  Sagas: _sagas.Sagas
	};
	
	module.exports = moduleExports;

/***/ }),
/* 1 */
/***/ (function(module, exports, __webpack_require__) {

	/* WEBPACK VAR INJECTION */(function(Buffer) {'use strict';
	
	var bind = __webpack_require__(16);
	
	/*global toString:true*/
	
	// utils is a library of generic helper functions non-specific to axios
	
	var toString = Object.prototype.toString;
	
	/**
	 * Determine if a value is an Array
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is an Array, otherwise false
	 */
	function isArray(val) {
	  return toString.call(val) === '[object Array]';
	}
	
	/**
	 * Determine if a value is a Node Buffer
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a Node Buffer, otherwise false
	 */
	function isBuffer(val) {
	  return ((typeof Buffer !== 'undefined') && (Buffer.isBuffer) && (Buffer.isBuffer(val)));
	}
	
	/**
	 * Determine if a value is an ArrayBuffer
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is an ArrayBuffer, otherwise false
	 */
	function isArrayBuffer(val) {
	  return toString.call(val) === '[object ArrayBuffer]';
	}
	
	/**
	 * Determine if a value is a FormData
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is an FormData, otherwise false
	 */
	function isFormData(val) {
	  return (typeof FormData !== 'undefined') && (val instanceof FormData);
	}
	
	/**
	 * Determine if a value is a view on an ArrayBuffer
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a view on an ArrayBuffer, otherwise false
	 */
	function isArrayBufferView(val) {
	  var result;
	  if ((typeof ArrayBuffer !== 'undefined') && (ArrayBuffer.isView)) {
	    result = ArrayBuffer.isView(val);
	  } else {
	    result = (val) && (val.buffer) && (val.buffer instanceof ArrayBuffer);
	  }
	  return result;
	}
	
	/**
	 * Determine if a value is a String
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a String, otherwise false
	 */
	function isString(val) {
	  return typeof val === 'string';
	}
	
	/**
	 * Determine if a value is a Number
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a Number, otherwise false
	 */
	function isNumber(val) {
	  return typeof val === 'number';
	}
	
	/**
	 * Determine if a value is undefined
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if the value is undefined, otherwise false
	 */
	function isUndefined(val) {
	  return typeof val === 'undefined';
	}
	
	/**
	 * Determine if a value is an Object
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is an Object, otherwise false
	 */
	function isObject(val) {
	  return val !== null && typeof val === 'object';
	}
	
	/**
	 * Determine if a value is a Date
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a Date, otherwise false
	 */
	function isDate(val) {
	  return toString.call(val) === '[object Date]';
	}
	
	/**
	 * Determine if a value is a File
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a File, otherwise false
	 */
	function isFile(val) {
	  return toString.call(val) === '[object File]';
	}
	
	/**
	 * Determine if a value is a Blob
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a Blob, otherwise false
	 */
	function isBlob(val) {
	  return toString.call(val) === '[object Blob]';
	}
	
	/**
	 * Determine if a value is a Function
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a Function, otherwise false
	 */
	function isFunction(val) {
	  return toString.call(val) === '[object Function]';
	}
	
	/**
	 * Determine if a value is a Stream
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a Stream, otherwise false
	 */
	function isStream(val) {
	  return isObject(val) && isFunction(val.pipe);
	}
	
	/**
	 * Determine if a value is a URLSearchParams object
	 *
	 * @param {Object} val The value to test
	 * @returns {boolean} True if value is a URLSearchParams object, otherwise false
	 */
	function isURLSearchParams(val) {
	  return typeof URLSearchParams !== 'undefined' && val instanceof URLSearchParams;
	}
	
	/**
	 * Trim excess whitespace off the beginning and end of a string
	 *
	 * @param {String} str The String to trim
	 * @returns {String} The String freed of excess whitespace
	 */
	function trim(str) {
	  return str.replace(/^\s*/, '').replace(/\s*$/, '');
	}
	
	/**
	 * Determine if we're running in a standard browser environment
	 *
	 * This allows axios to run in a web worker, and react-native.
	 * Both environments support XMLHttpRequest, but not fully standard globals.
	 *
	 * web workers:
	 *  typeof window -> undefined
	 *  typeof document -> undefined
	 *
	 * react-native:
	 *  navigator.product -> 'ReactNative'
	 */
	function isStandardBrowserEnv() {
	  if (typeof navigator !== 'undefined' && navigator.product === 'ReactNative') {
	    return false;
	  }
	  return (
	    typeof window !== 'undefined' &&
	    typeof document !== 'undefined'
	  );
	}
	
	/**
	 * Iterate over an Array or an Object invoking a function for each item.
	 *
	 * If `obj` is an Array callback will be called passing
	 * the value, index, and complete array for each item.
	 *
	 * If 'obj' is an Object callback will be called passing
	 * the value, key, and complete object for each property.
	 *
	 * @param {Object|Array} obj The object to iterate
	 * @param {Function} fn The callback to invoke for each item
	 */
	function forEach(obj, fn) {
	  // Don't bother if no value provided
	  if (obj === null || typeof obj === 'undefined') {
	    return;
	  }
	
	  // Force an array if not already something iterable
	  if (typeof obj !== 'object' && !isArray(obj)) {
	    /*eslint no-param-reassign:0*/
	    obj = [obj];
	  }
	
	  if (isArray(obj)) {
	    // Iterate over array values
	    for (var i = 0, l = obj.length; i < l; i++) {
	      fn.call(null, obj[i], i, obj);
	    }
	  } else {
	    // Iterate over object keys
	    for (var key in obj) {
	      if (Object.prototype.hasOwnProperty.call(obj, key)) {
	        fn.call(null, obj[key], key, obj);
	      }
	    }
	  }
	}
	
	/**
	 * Accepts varargs expecting each argument to be an object, then
	 * immutably merges the properties of each object and returns result.
	 *
	 * When multiple objects contain the same key the later object in
	 * the arguments list will take precedence.
	 *
	 * Example:
	 *
	 * ```js
	 * var result = merge({foo: 123}, {foo: 456});
	 * console.log(result.foo); // outputs 456
	 * ```
	 *
	 * @param {Object} obj1 Object to merge
	 * @returns {Object} Result of all merge properties
	 */
	function merge(/* obj1, obj2, obj3, ... */) {
	  var result = {};
	  function assignValue(val, key) {
	    if (typeof result[key] === 'object' && typeof val === 'object') {
	      result[key] = merge(result[key], val);
	    } else {
	      result[key] = val;
	    }
	  }
	
	  for (var i = 0, l = arguments.length; i < l; i++) {
	    forEach(arguments[i], assignValue);
	  }
	  return result;
	}
	
	/**
	 * Extends object a by mutably adding to it the properties of object b.
	 *
	 * @param {Object} a The object to be extended
	 * @param {Object} b The object to copy properties from
	 * @param {Object} thisArg The object to bind function to
	 * @return {Object} The resulting value of object a
	 */
	function extend(a, b, thisArg) {
	  forEach(b, function assignValue(val, key) {
	    if (thisArg && typeof val === 'function') {
	      a[key] = bind(val, thisArg);
	    } else {
	      a[key] = val;
	    }
	  });
	  return a;
	}
	
	module.exports = {
	  isArray: isArray,
	  isArrayBuffer: isArrayBuffer,
	  isBuffer: isBuffer,
	  isFormData: isFormData,
	  isArrayBufferView: isArrayBufferView,
	  isString: isString,
	  isNumber: isNumber,
	  isObject: isObject,
	  isUndefined: isUndefined,
	  isDate: isDate,
	  isFile: isFile,
	  isBlob: isBlob,
	  isFunction: isFunction,
	  isStream: isStream,
	  isURLSearchParams: isURLSearchParams,
	  isStandardBrowserEnv: isStandardBrowserEnv,
	  forEach: forEach,
	  merge: merge,
	  extend: extend,
	  trim: trim
	};
	
	/* WEBPACK VAR INJECTION */}.call(exports, __webpack_require__(57).Buffer))

/***/ }),
/* 2 */
/***/ (function(module, exports) {

	"use strict";
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	/* Populated by react-webpack-redux:action */
	var ActionNames = exports.ActionNames = {
	
	  LOGIN: "LOGIN",
	  LOGGING_IN: "LOGGING_IN",
	  LOGIN_SUCCESS: "LOGIN_SUCCESS",
	  LOGIN_FAILURE: "LOGIN_FAILURE",
	  LOGOUT: "LOGOUT",
	
	  CONTAINER_REFRESH: "CONTAINER_REFRESH",
	
	  VIEW_FETCH: "VIEW_FETCH",
	  VIEW_FETCHING: "VIEW_FETCHING",
	  VIEW_FETCH_SUCCESS: "VIEW_FETCH_SUCCESS",
	  VIEW_FETCH_FAILED: "VIEW_FETCH_FAILED",
	  VIEW_ITEM_REMOVE: "VIEW_ITEM_REMOVE",
	  VIEW_ITEM_RELOAD: "VIEW_ITEM_RELOAD",
	
	  ENTITY_GET: "ENTITY_GET",
	  ENTITY_GETTING: "ENTITY_GETTING",
	  ENTITY_GET_SUCCESS: "ENTITY_GET_SUCCESS",
	  ENTITY_GET_FAILED: "ENTITY_GET_FAILED",
	
	  ENTITY_SAVE: "ENTITY_SAVE",
	  ENTITY_SAVING: "ENTITY_SAVING",
	  ENTITY_SAVE_SUCCESS: "ENTITY_SAVE_SUCCESS",
	  ENTITY_SAVE_FAILURE: "ENTITY_SAVE_FAILURE",
	
	  ENTITY_PUT: "ENTITY_PUT",
	  ENTITY_PUTTING: "ENTITY_PUTTING",
	  ENTITY_PUT_SUCCESS: "ENTITY_UPDATE_SUCCESS",
	  ENTITY_PUT_FAILURE: "ENTITY_PUT_FAILURE",
	
	  ENTITY_UPDATE: "ENTITY_UPDATE",
	  ENTITY_UPDATING: "ENTITY_UPDATING",
	  ENTITY_UPDATE_SUCCESS: "ENTITY_UPDATE_SUCCESS",
	  ENTITY_UPDATE_FAILURE: "ENTITY_UPDATE_FAILURE",
	
	  GROUP_LOAD: "GROUP_LOAD",
	
	  PAGE_CHANGE: "@@reduxdirector/LOCATION_CHANGE",
	
	  ENTITY_DELETE: "ENTITY_DELETE",
	  ENTITY_DELETING: "ENTITY_DELETING",
	  ENTITY_DELETE_SUCCESS: "ENTITY_DELETE_SUCCESS",
	  ENTITY_DELETE_FAILURE: "ENTITY_DELETE_FAILURE"
	};

/***/ }),
/* 3 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.LaatooError = undefined;
	exports.createAction = createAction;
	exports.formatUrl = formatUrl;
	exports.hasPermission = hasPermission;
	
	var _Globals = __webpack_require__(4);
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }
	
	function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }
	
	var LaatooError = exports.LaatooError = function (_Error) {
	  _inherits(LaatooError, _Error);
	
	  function LaatooError(type, rootError, args) {
	    _classCallCheck(this, LaatooError);
	
	    var _this = _possibleConstructorReturn(this, (LaatooError.__proto__ || Object.getPrototypeOf(LaatooError)).call(this, type));
	
	    _this.name = _this.constructor.name;
	    _this.message = type;
	    if (typeof Error.captureStackTrace === 'function') {
	      Error.captureStackTrace(_this, _this.constructor);
	    } else {
	      _this.stack = new Error(type).stack;
	    }
	    _this.type = type;
	    _this.rootError = rootError;
	    _this.args = args;
	    return _this;
	  }
	
	  return LaatooError;
	}(Error);
	
	function createAction(type, payload, meta) {
	  var error = payload instanceof Error;
	  console.log("created action", type, payload, meta, error);
	  return {
	    type: type,
	    payload: payload,
	    meta: meta,
	    error: error
	  };
	}
	
	function formatUrl(url, params) {
	  var newurl = url;
	  if (params) {
	    for (var key in params) {
	      var val = params[key];
	      newurl = newurl.replace(new RegExp(":" + key, "g"), val);
	    }
	  }
	  return newurl;
	}
	
	function hasPermission(permission) {
	  var hasPermission = true;
	  if (permission && permission != "") {
	    var permissions = localStorage.permissions;
	    if (permissions) {
	      if (permissions.indexOf(permission) < 0) {
	        hasPermission = false;
	      }
	    }
	  }
	  return hasPermission;
	}

/***/ }),
/* 4 */
/***/ (function(module, exports) {

	'use strict';
	
	var native = typeof document == 'undefined';
	var storage = native ? {} : localStorage;
	var application = native ? {} : document.InitConfig;
	var wind = native ? {} : window;
	application.native = native;
	
	module.exports = {
	  Storage: storage,
	  Application: application,
	  Window: wind
	};

/***/ }),
/* 5 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.Response = exports.EntityData = exports.DataSource = exports.RequestBuilder = undefined;
	
	var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
	
	var _RequestBuilder = __webpack_require__(54);
	
	var _EntityData = __webpack_require__(53);
	
	var _axios = __webpack_require__(23);
	
	var _axios2 = _interopRequireDefault(_axios);
	
	var _Globals = __webpack_require__(4);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	var Response = {
	  Success: "Success",
	  Unauthorized: "Unauthorized",
	  InternalError: "InternalError",
	  BadRequest: "BadRequest",
	  Failure: "Failure"
	};
	console.log("application services", _Globals.Application);
	
	var DefaultDataSource = function () {
	  function DefaultDataSource() {
	    _classCallCheck(this, DefaultDataSource);
	
	    this.ExecuteService = this.ExecuteService.bind(this);
	    this.ExecuteServiceObject = this.ExecuteServiceObject.bind(this);
	    this.HttpCall = this.HttpCall.bind(this);
	    this.buildHttpSvcResponse = this.buildHttpSvcResponse.bind(this);
	  }
	
	  _createClass(DefaultDataSource, [{
	    key: 'ExecuteService',
	    value: function ExecuteService(serviceName, serviceRequest) {
	      var config = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : null;
	
	      console.log("application services", _Globals.Application);
	      var service = _Globals.Application.Services[serviceName];
	      if (service != null && serviceRequest != null) {
	        return this.ExecuteServiceObject(service, serviceRequest, config);
	      } else {
	        throw new Error('Service not found' + serviceName);
	      }
	    }
	  }, {
	    key: 'ExecuteServiceObject',
	    value: function ExecuteServiceObject(service, serviceRequest) {
	      var config = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : null;
	
	      if (service != null && serviceRequest != null) {
	        var protocol = this.getProtocol();
	        var req = serviceRequest.GetRequest(protocol);
	        if (protocol === 'http') {
	          var method = this.getMethod(service);
	          var url = this.getURL(service, req);
	          return this.HttpCall(url, method, req.params, req.data, req.headers, config);
	        }
	      } else {
	        throw new Error('Invalid Request');
	      }
	    }
	  }, {
	    key: 'HttpCall',
	    value: function HttpCall(url, method, params, data, headers) {
	      var config = arguments.length > 5 && arguments[5] !== undefined ? arguments[5] : null;
	
	      var service = this;
	      var promise = new Promise(function (resolve, reject) {
	        if (method === "" || url === "") {
	          reject(service.buildHttpSvcResponse(Response.InternalError, 'Could not build request', url));
	          return;
	        }
	        var successCallback = function successCallback(response) {
	          if (response.status < 300) {
	            var res = service.buildHttpSvcResponse(Response.Success, "", response);
	            resolve(res);
	          } else {
	            reject(service.buildHttpSvcResponse(Response.Failure, "", response));
	          }
	        };
	        var errorCallback = function errorCallback(response) {
	          reject(service.buildHttpSvcResponse(Response.Failure, "", response));
	        };
	        if (method == 'DELETE' || method == 'GET') {
	          data = null;
	        }
	        if (!headers) {
	          headers = {};
	        }
	        headers[_Globals.Application.Security.AuthToken] = _Globals.Storage.auth;
	        var req = {
	          method: method,
	          url: url,
	          data: data,
	          headers: headers,
	          params: params,
	          responseType: 'json'
	        };
	        if (config) {
	          req = Object.assign({}, req, config);
	        }
	        console.log("Request.. ", req);
	        (0, _axios2.default)(req).then(successCallback, errorCallback);
	      });
	      return promise;
	    }
	  }, {
	    key: 'createFullUrl',
	    value: function createFullUrl(url, params) {
	      if (params != null && Object.keys(params).length != 0) {
	        return url + "?" + Object.keys(data).map(function (key) {
	          return [key, data[key]].map(encodeURIComponent).join("=");
	        }).join("&");
	      }
	      return url;
	    }
	  }, {
	    key: 'buildHttpSvcResponse',
	    value: function buildHttpSvcResponse(code, msg, res) {
	      if (res instanceof Error) {
	        return this.buildSvcResponse(code, msg, res, {});
	      }
	      return this.buildSvcResponse(code, msg, res.data, res.headers, res.status);
	    }
	  }, {
	    key: 'buildSvcResponse',
	    value: function buildSvcResponse(code, msg, data, info, statuscode) {
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
	  }, {
	    key: 'getURL',
	    value: function getURL(service, req) {
	      var url = service.url;
	      if (req.urlparams != null) {
	        for (var param in req.urlparams) {
	          url = url.replace(":" + param, req.urlparams[param]);
	        }
	      }
	      if (url.startsWith('http')) {
	        return url;
	      } else {
	        return _Globals.Application.Backend + url;
	      }
	    }
	  }, {
	    key: 'getMethod',
	    value: function getMethod(service) {
	      if (service.method) {
	        return service.method;
	      }
	      return 'GET';
	    }
	  }, {
	    key: 'getProtocol',
	    value: function getProtocol() {
	      return 'http';
	    }
	  }]);
	
	  return DefaultDataSource;
	}();
	
	var DataSource = new DefaultDataSource();
	var RequestBuilder = new _RequestBuilder.RequestBuilderService();
	var EntityData = new _EntityData.EntityDataService(DataSource, RequestBuilder);
	exports.RequestBuilder = RequestBuilder;
	exports.DataSource = DataSource;
	exports.EntityData = EntityData;
	exports.Response = Response;

/***/ }),
/* 6 */
/***/ (function(module, exports) {

	module.exports = __WEBPACK_EXTERNAL_MODULE_6__;

/***/ }),
/* 7 */
/***/ (function(module, exports) {

	module.exports = __WEBPACK_EXTERNAL_MODULE_7__;

/***/ }),
/* 8 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	exports.__esModule = true;
	
	var _io = __webpack_require__(21);
	
	Object.defineProperty(exports, 'take', {
		enumerable: true,
		get: function get() {
			return _io.take;
		}
	});
	Object.defineProperty(exports, 'takem', {
		enumerable: true,
		get: function get() {
			return _io.takem;
		}
	});
	Object.defineProperty(exports, 'put', {
		enumerable: true,
		get: function get() {
			return _io.put;
		}
	});
	Object.defineProperty(exports, 'all', {
		enumerable: true,
		get: function get() {
			return _io.all;
		}
	});
	Object.defineProperty(exports, 'race', {
		enumerable: true,
		get: function get() {
			return _io.race;
		}
	});
	Object.defineProperty(exports, 'call', {
		enumerable: true,
		get: function get() {
			return _io.call;
		}
	});
	Object.defineProperty(exports, 'apply', {
		enumerable: true,
		get: function get() {
			return _io.apply;
		}
	});
	Object.defineProperty(exports, 'cps', {
		enumerable: true,
		get: function get() {
			return _io.cps;
		}
	});
	Object.defineProperty(exports, 'fork', {
		enumerable: true,
		get: function get() {
			return _io.fork;
		}
	});
	Object.defineProperty(exports, 'spawn', {
		enumerable: true,
		get: function get() {
			return _io.spawn;
		}
	});
	Object.defineProperty(exports, 'join', {
		enumerable: true,
		get: function get() {
			return _io.join;
		}
	});
	Object.defineProperty(exports, 'cancel', {
		enumerable: true,
		get: function get() {
			return _io.cancel;
		}
	});
	Object.defineProperty(exports, 'select', {
		enumerable: true,
		get: function get() {
			return _io.select;
		}
	});
	Object.defineProperty(exports, 'actionChannel', {
		enumerable: true,
		get: function get() {
			return _io.actionChannel;
		}
	});
	Object.defineProperty(exports, 'cancelled', {
		enumerable: true,
		get: function get() {
			return _io.cancelled;
		}
	});
	Object.defineProperty(exports, 'flush', {
		enumerable: true,
		get: function get() {
			return _io.flush;
		}
	});
	Object.defineProperty(exports, 'getContext', {
		enumerable: true,
		get: function get() {
			return _io.getContext;
		}
	});
	Object.defineProperty(exports, 'setContext', {
		enumerable: true,
		get: function get() {
			return _io.setContext;
		}
	});
	Object.defineProperty(exports, 'takeEvery', {
		enumerable: true,
		get: function get() {
			return _io.takeEvery;
		}
	});
	Object.defineProperty(exports, 'takeLatest', {
		enumerable: true,
		get: function get() {
			return _io.takeLatest;
		}
	});
	Object.defineProperty(exports, 'throttle', {
		enumerable: true,
		get: function get() {
			return _io.throttle;
		}
	});

/***/ }),
/* 9 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	exports.__esModule = true;
	
	var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };
	
	var _typeof = typeof Symbol === "function" && typeof Symbol.iterator === "symbol" ? function (obj) { return typeof obj; } : function (obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; };
	
	exports.check = check;
	exports.hasOwn = hasOwn;
	exports.remove = remove;
	exports.deferred = deferred;
	exports.arrayOfDeffered = arrayOfDeffered;
	exports.delay = delay;
	exports.createMockTask = createMockTask;
	exports.autoInc = autoInc;
	exports.makeIterator = makeIterator;
	exports.log = log;
	exports.deprecate = deprecate;
	var sym = exports.sym = function sym(id) {
	  return '@@redux-saga/' + id;
	};
	
	var TASK = exports.TASK = sym('TASK');
	var HELPER = exports.HELPER = sym('HELPER');
	var MATCH = exports.MATCH = sym('MATCH');
	var CANCEL = exports.CANCEL = sym('CANCEL_PROMISE');
	var SAGA_ACTION = exports.SAGA_ACTION = sym('SAGA_ACTION');
	var SELF_CANCELLATION = exports.SELF_CANCELLATION = sym('SELF_CANCELLATION');
	var konst = exports.konst = function konst(v) {
	  return function () {
	    return v;
	  };
	};
	var kTrue = exports.kTrue = konst(true);
	var kFalse = exports.kFalse = konst(false);
	var noop = exports.noop = function noop() {};
	var ident = exports.ident = function ident(v) {
	  return v;
	};
	
	function check(value, predicate, error) {
	  if (!predicate(value)) {
	    log('error', 'uncaught at check', error);
	    throw new Error(error);
	  }
	}
	
	var hasOwnProperty = Object.prototype.hasOwnProperty;
	function hasOwn(object, property) {
	  return is.notUndef(object) && hasOwnProperty.call(object, property);
	}
	
	var is = exports.is = {
	  undef: function undef(v) {
	    return v === null || v === undefined;
	  },
	  notUndef: function notUndef(v) {
	    return v !== null && v !== undefined;
	  },
	  func: function func(f) {
	    return typeof f === 'function';
	  },
	  number: function number(n) {
	    return typeof n === 'number';
	  },
	  string: function string(s) {
	    return typeof s === 'string';
	  },
	  array: Array.isArray,
	  object: function object(obj) {
	    return obj && !is.array(obj) && (typeof obj === 'undefined' ? 'undefined' : _typeof(obj)) === 'object';
	  },
	  promise: function promise(p) {
	    return p && is.func(p.then);
	  },
	  iterator: function iterator(it) {
	    return it && is.func(it.next) && is.func(it.throw);
	  },
	  iterable: function iterable(it) {
	    return it && is.func(Symbol) ? is.func(it[Symbol.iterator]) : is.array(it);
	  },
	  task: function task(t) {
	    return t && t[TASK];
	  },
	  observable: function observable(ob) {
	    return ob && is.func(ob.subscribe);
	  },
	  buffer: function buffer(buf) {
	    return buf && is.func(buf.isEmpty) && is.func(buf.take) && is.func(buf.put);
	  },
	  pattern: function pattern(pat) {
	    return pat && (is.string(pat) || (typeof pat === 'undefined' ? 'undefined' : _typeof(pat)) === 'symbol' || is.func(pat) || is.array(pat));
	  },
	  channel: function channel(ch) {
	    return ch && is.func(ch.take) && is.func(ch.close);
	  },
	  helper: function helper(it) {
	    return it && it[HELPER];
	  },
	  stringableFunc: function stringableFunc(f) {
	    return is.func(f) && hasOwn(f, 'toString');
	  }
	};
	
	var object = exports.object = {
	  assign: function assign(target, source) {
	    for (var i in source) {
	      if (hasOwn(source, i)) {
	        target[i] = source[i];
	      }
	    }
	  }
	};
	
	function remove(array, item) {
	  var index = array.indexOf(item);
	  if (index >= 0) {
	    array.splice(index, 1);
	  }
	}
	
	var array = exports.array = {
	  'from': function from(obj) {
	    var arr = Array(obj.length);
	    for (var i in obj) {
	      if (hasOwn(obj, i)) {
	        arr[i] = obj[i];
	      }
	    }
	    return arr;
	  }
	};
	
	function deferred() {
	  var props = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
	
	  var def = _extends({}, props);
	  var promise = new Promise(function (resolve, reject) {
	    def.resolve = resolve;
	    def.reject = reject;
	  });
	  def.promise = promise;
	  return def;
	}
	
	function arrayOfDeffered(length) {
	  var arr = [];
	  for (var i = 0; i < length; i++) {
	    arr.push(deferred());
	  }
	  return arr;
	}
	
	function delay(ms) {
	  var val = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : true;
	
	  var timeoutId = void 0;
	  var promise = new Promise(function (resolve) {
	    timeoutId = setTimeout(function () {
	      return resolve(val);
	    }, ms);
	  });
	
	  promise[CANCEL] = function () {
	    return clearTimeout(timeoutId);
	  };
	
	  return promise;
	}
	
	function createMockTask() {
	  var _ref;
	
	  var running = true;
	  var _result = void 0,
	      _error = void 0;
	
	  return _ref = {}, _ref[TASK] = true, _ref.isRunning = function isRunning() {
	    return running;
	  }, _ref.result = function result() {
	    return _result;
	  }, _ref.error = function error() {
	    return _error;
	  }, _ref.setRunning = function setRunning(b) {
	    return running = b;
	  }, _ref.setResult = function setResult(r) {
	    return _result = r;
	  }, _ref.setError = function setError(e) {
	    return _error = e;
	  }, _ref;
	}
	
	function autoInc() {
	  var seed = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 0;
	
	  return function () {
	    return ++seed;
	  };
	}
	
	var uid = exports.uid = autoInc();
	
	var kThrow = function kThrow(err) {
	  throw err;
	};
	var kReturn = function kReturn(value) {
	  return { value: value, done: true };
	};
	function makeIterator(next) {
	  var thro = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : kThrow;
	  var name = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : '';
	  var isHelper = arguments[3];
	
	  var iterator = { name: name, next: next, throw: thro, return: kReturn };
	
	  if (isHelper) {
	    iterator[HELPER] = true;
	  }
	  if (typeof Symbol !== 'undefined') {
	    iterator[Symbol.iterator] = function () {
	      return iterator;
	    };
	  }
	  return iterator;
	}
	
	/**
	  Print error in a useful way whether in a browser environment
	  (with expandable error stack traces), or in a node.js environment
	  (text-only log output)
	 **/
	function log(level, message) {
	  var error = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : '';
	
	  /*eslint-disable no-console*/
	  if (typeof window === 'undefined') {
	    console.log('redux-saga ' + level + ': ' + message + '\n' + (error && error.stack || error));
	  } else {
	    console[level](message, error);
	  }
	}
	
	function deprecate(fn, deprecationWarning) {
	  return function () {
	    if (false) log('warn', deprecationWarning);
	    return fn.apply(undefined, arguments);
	  };
	}
	
	var updateIncentive = exports.updateIncentive = function updateIncentive(deprecated, preferred) {
	  return deprecated + ' has been deprecated in favor of ' + preferred + ', please update your code';
	};
	
	var internalErr = exports.internalErr = function internalErr(err) {
	  return new Error('\n  redux-saga: Error checking hooks detected an inconsistent state. This is likely a bug\n  in redux-saga code and not yours. Thanks for reporting this in the project\'s github repo.\n  Error: ' + err + '\n');
	};
	
	var createSetContextWarning = exports.createSetContextWarning = function createSetContextWarning(ctx, props) {
	  return (ctx ? ctx + '.' : '') + 'setContext(props): argument ' + props + ' is not a plain object';
	};
	
	var wrapSagaDispatch = exports.wrapSagaDispatch = function wrapSagaDispatch(dispatch) {
	  return function (action) {
	    return dispatch(Object.defineProperty(action, SAGA_ACTION, { value: true }));
	  };
	};
	
	var cloneableGenerator = exports.cloneableGenerator = function cloneableGenerator(generatorFunc) {
	  return function () {
	    for (var _len = arguments.length, args = Array(_len), _key = 0; _key < _len; _key++) {
	      args[_key] = arguments[_key];
	    }
	
	    var history = [];
	    var gen = generatorFunc.apply(undefined, args);
	    return {
	      next: function next(arg) {
	        history.push(arg);
	        return gen.next(arg);
	      },
	      clone: function clone() {
	        var clonedGen = cloneableGenerator(generatorFunc).apply(undefined, args);
	        history.forEach(function (arg) {
	          return clonedGen.next(arg);
	        });
	        return clonedGen;
	      },
	      return: function _return(value) {
	        return gen.return(value);
	      },
	      throw: function _throw(exception) {
	        return gen.throw(exception);
	      }
	    };
	  };
	};

/***/ }),
/* 10 */
/***/ (function(module, exports) {

	module.exports = __WEBPACK_EXTERNAL_MODULE_10__;

/***/ }),
/* 11 */
/***/ (function(module, exports, __webpack_require__) {

	/* WEBPACK VAR INJECTION */(function(process) {'use strict';
	
	var utils = __webpack_require__(1);
	var normalizeHeaderName = __webpack_require__(38);
	
	var DEFAULT_CONTENT_TYPE = {
	  'Content-Type': 'application/x-www-form-urlencoded'
	};
	
	function setContentTypeIfUnset(headers, value) {
	  if (!utils.isUndefined(headers) && utils.isUndefined(headers['Content-Type'])) {
	    headers['Content-Type'] = value;
	  }
	}
	
	function getDefaultAdapter() {
	  var adapter;
	  if (typeof XMLHttpRequest !== 'undefined') {
	    // For browsers use XHR adapter
	    adapter = __webpack_require__(12);
	  } else if (typeof process !== 'undefined') {
	    // For node use HTTP adapter
	    adapter = __webpack_require__(12);
	  }
	  return adapter;
	}
	
	var defaults = {
	  adapter: getDefaultAdapter(),
	
	  transformRequest: [function transformRequest(data, headers) {
	    normalizeHeaderName(headers, 'Content-Type');
	    if (utils.isFormData(data) ||
	      utils.isArrayBuffer(data) ||
	      utils.isBuffer(data) ||
	      utils.isStream(data) ||
	      utils.isFile(data) ||
	      utils.isBlob(data)
	    ) {
	      return data;
	    }
	    if (utils.isArrayBufferView(data)) {
	      return data.buffer;
	    }
	    if (utils.isURLSearchParams(data)) {
	      setContentTypeIfUnset(headers, 'application/x-www-form-urlencoded;charset=utf-8');
	      return data.toString();
	    }
	    if (utils.isObject(data)) {
	      setContentTypeIfUnset(headers, 'application/json;charset=utf-8');
	      return JSON.stringify(data);
	    }
	    return data;
	  }],
	
	  transformResponse: [function transformResponse(data) {
	    /*eslint no-param-reassign:0*/
	    if (typeof data === 'string') {
	      try {
	        data = JSON.parse(data);
	      } catch (e) { /* Ignore */ }
	    }
	    return data;
	  }],
	
	  timeout: 0,
	
	  xsrfCookieName: 'XSRF-TOKEN',
	  xsrfHeaderName: 'X-XSRF-TOKEN',
	
	  maxContentLength: -1,
	
	  validateStatus: function validateStatus(status) {
	    return status >= 200 && status < 300;
	  }
	};
	
	defaults.headers = {
	  common: {
	    'Accept': 'application/json, text/plain, */*'
	  }
	};
	
	utils.forEach(['delete', 'get', 'head'], function forEachMethodNoData(method) {
	  defaults.headers[method] = {};
	});
	
	utils.forEach(['post', 'put', 'patch'], function forEachMethodWithData(method) {
	  defaults.headers[method] = utils.merge(DEFAULT_CONTENT_TYPE);
	});
	
	module.exports = defaults;
	
	/* WEBPACK VAR INJECTION */}.call(exports, __webpack_require__(60)))

/***/ }),
/* 12 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	var settle = __webpack_require__(30);
	var buildURL = __webpack_require__(33);
	var parseHeaders = __webpack_require__(39);
	var isURLSameOrigin = __webpack_require__(37);
	var createError = __webpack_require__(15);
	var btoa = (typeof window !== 'undefined' && window.btoa && window.btoa.bind(window)) || __webpack_require__(32);
	
	module.exports = function xhrAdapter(config) {
	  return new Promise(function dispatchXhrRequest(resolve, reject) {
	    var requestData = config.data;
	    var requestHeaders = config.headers;
	
	    if (utils.isFormData(requestData)) {
	      delete requestHeaders['Content-Type']; // Let the browser set it
	    }
	
	    var request = new XMLHttpRequest();
	    var loadEvent = 'onreadystatechange';
	    var xDomain = false;
	
	    // For IE 8/9 CORS support
	    // Only supports POST and GET calls and doesn't returns the response headers.
	    // DON'T do this for testing b/c XMLHttpRequest is mocked, not XDomainRequest.
	    if (("production") !== 'test' &&
	        typeof window !== 'undefined' &&
	        window.XDomainRequest && !('withCredentials' in request) &&
	        !isURLSameOrigin(config.url)) {
	      request = new window.XDomainRequest();
	      loadEvent = 'onload';
	      xDomain = true;
	      request.onprogress = function handleProgress() {};
	      request.ontimeout = function handleTimeout() {};
	    }
	
	    // HTTP basic authentication
	    if (config.auth) {
	      var username = config.auth.username || '';
	      var password = config.auth.password || '';
	      requestHeaders.Authorization = 'Basic ' + btoa(username + ':' + password);
	    }
	
	    request.open(config.method.toUpperCase(), buildURL(config.url, config.params, config.paramsSerializer), true);
	
	    // Set the request timeout in MS
	    request.timeout = config.timeout;
	
	    // Listen for ready state
	    request[loadEvent] = function handleLoad() {
	      if (!request || (request.readyState !== 4 && !xDomain)) {
	        return;
	      }
	
	      // The request errored out and we didn't get a response, this will be
	      // handled by onerror instead
	      // With one exception: request that using file: protocol, most browsers
	      // will return status as 0 even though it's a successful request
	      if (request.status === 0 && !(request.responseURL && request.responseURL.indexOf('file:') === 0)) {
	        return;
	      }
	
	      // Prepare the response
	      var responseHeaders = 'getAllResponseHeaders' in request ? parseHeaders(request.getAllResponseHeaders()) : null;
	      var responseData = !config.responseType || config.responseType === 'text' ? request.responseText : request.response;
	      var response = {
	        data: responseData,
	        // IE sends 1223 instead of 204 (https://github.com/mzabriskie/axios/issues/201)
	        status: request.status === 1223 ? 204 : request.status,
	        statusText: request.status === 1223 ? 'No Content' : request.statusText,
	        headers: responseHeaders,
	        config: config,
	        request: request
	      };
	
	      settle(resolve, reject, response);
	
	      // Clean up request
	      request = null;
	    };
	
	    // Handle low level network errors
	    request.onerror = function handleError() {
	      // Real errors are hidden from us by the browser
	      // onerror should only fire if it's a network error
	      reject(createError('Network Error', config));
	
	      // Clean up request
	      request = null;
	    };
	
	    // Handle timeout
	    request.ontimeout = function handleTimeout() {
	      reject(createError('timeout of ' + config.timeout + 'ms exceeded', config, 'ECONNABORTED'));
	
	      // Clean up request
	      request = null;
	    };
	
	    // Add xsrf header
	    // This is only done if running in a standard browser environment.
	    // Specifically not if we're in a web worker, or react-native.
	    if (utils.isStandardBrowserEnv()) {
	      var cookies = __webpack_require__(35);
	
	      // Add xsrf header
	      var xsrfValue = (config.withCredentials || isURLSameOrigin(config.url)) && config.xsrfCookieName ?
	          cookies.read(config.xsrfCookieName) :
	          undefined;
	
	      if (xsrfValue) {
	        requestHeaders[config.xsrfHeaderName] = xsrfValue;
	      }
	    }
	
	    // Add headers to the request
	    if ('setRequestHeader' in request) {
	      utils.forEach(requestHeaders, function setRequestHeader(val, key) {
	        if (typeof requestData === 'undefined' && key.toLowerCase() === 'content-type') {
	          // Remove Content-Type if data is undefined
	          delete requestHeaders[key];
	        } else {
	          // Otherwise add header to the request
	          request.setRequestHeader(key, val);
	        }
	      });
	    }
	
	    // Add withCredentials to request if needed
	    if (config.withCredentials) {
	      request.withCredentials = true;
	    }
	
	    // Add responseType to request if needed
	    if (config.responseType) {
	      try {
	        request.responseType = config.responseType;
	      } catch (e) {
	        // Expected DOMException thrown by browsers not compatible XMLHttpRequest Level 2.
	        // But, this can be suppressed for 'json' type as it can be parsed by default 'transformResponse' function.
	        if (config.responseType !== 'json') {
	          throw e;
	        }
	      }
	    }
	
	    // Handle progress if needed
	    if (typeof config.onDownloadProgress === 'function') {
	      request.addEventListener('progress', config.onDownloadProgress);
	    }
	
	    // Not all browsers support upload events
	    if (typeof config.onUploadProgress === 'function' && request.upload) {
	      request.upload.addEventListener('progress', config.onUploadProgress);
	    }
	
	    if (config.cancelToken) {
	      // Handle cancellation
	      config.cancelToken.promise.then(function onCanceled(cancel) {
	        if (!request) {
	          return;
	        }
	
	        request.abort();
	        reject(cancel);
	        // Clean up request
	        request = null;
	      });
	    }
	
	    if (requestData === undefined) {
	      requestData = null;
	    }
	
	    // Send the request
	    request.send(requestData);
	  });
	};


/***/ }),
/* 13 */
/***/ (function(module, exports) {

	'use strict';
	
	/**
	 * A `Cancel` is an object that is thrown when an operation is canceled.
	 *
	 * @class
	 * @param {string=} message The message.
	 */
	function Cancel(message) {
	  this.message = message;
	}
	
	Cancel.prototype.toString = function toString() {
	  return 'Cancel' + (this.message ? ': ' + this.message : '');
	};
	
	Cancel.prototype.__CANCEL__ = true;
	
	module.exports = Cancel;


/***/ }),
/* 14 */
/***/ (function(module, exports) {

	'use strict';
	
	module.exports = function isCancel(value) {
	  return !!(value && value.__CANCEL__);
	};


/***/ }),
/* 15 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var enhanceError = __webpack_require__(29);
	
	/**
	 * Create an Error with the specified message, config, error code, and response.
	 *
	 * @param {string} message The error message.
	 * @param {Object} config The config.
	 * @param {string} [code] The error code (for example, 'ECONNABORTED').
	 @ @param {Object} [response] The response.
	 * @returns {Error} The created error.
	 */
	module.exports = function createError(message, config, code, response) {
	  var error = new Error(message);
	  return enhanceError(error, config, code, response);
	};


/***/ }),
/* 16 */
/***/ (function(module, exports) {

	'use strict';
	
	module.exports = function bind(fn, thisArg) {
	  return function wrap() {
	    var args = new Array(arguments.length);
	    for (var i = 0; i < args.length; i++) {
	      args[i] = arguments[i];
	    }
	    return fn.apply(thisArg, args);
	  };
	};


/***/ }),
/* 17 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.ViewData = undefined;
	
	var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
	
	var _react = __webpack_require__(6);
	
	var _react2 = _interopRequireDefault(_react);
	
	var _reactRedux = __webpack_require__(10);
	
	var _ActionNames = __webpack_require__(2);
	
	var _utils = __webpack_require__(3);
	
	var _DataSource = __webpack_require__(5);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }
	
	function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }
	
	var ViewDataComp = function (_React$Component) {
	  _inherits(ViewDataComp, _React$Component);
	
	  function ViewDataComp(props) {
	    _classCallCheck(this, ViewDataComp);
	
	    var _this = _possibleConstructorReturn(this, (ViewDataComp.__proto__ || Object.getPrototypeOf(ViewDataComp)).call(this, props));
	
	    _this.setPage = _this.setPage.bind(_this);
	    _this.selectedItems = _this.selectedItems.bind(_this);
	    _this.itemStatus = _this.itemStatus.bind(_this);
	    _this.viewrefs = _this.viewrefs.bind(_this);
	    _this.itemCount = _this.itemCount.bind(_this);
	    _this.reload = _this.reload.bind(_this);
	    _this.setFilter = _this.setFilter.bind(_this);
	    _this.loadMore = _this.loadMore.bind(_this);
	    _this.canLoadMore = _this.canLoadMore.bind(_this);
	    _this.getView = _this.getView.bind(_this);
	    _this.methods = { reload: _this.reload, canLoadMore: _this.canLoadMore, loadMore: _this.loadMore, setFilter: _this.setFilter, itemCount: _this.itemCount, viewrefs: _this.viewrefs, itemStatus: _this.itemStatus, selectedItems: _this.selectedItems, setPage: _this.setPage };
	    _this.addMethod = _this.addMethod.bind(_this);
	    _this.state = { lastLoadTime: -1 };
	    _this.numItems = 0;
	    return _this;
	  }
	
	  _createClass(ViewDataComp, [{
	    key: 'componentWillMount',
	    value: function componentWillMount() {
	      this.filter = this.props.defaultFilter;
	    }
	  }, {
	    key: 'componentDidMount',
	    value: function componentDidMount() {
	      console.log("mounting view data", this.props);
	      if (this.props.load && !this.props.externalLoad) {
	        this.props.loadView(this.props.currentPage, this.filter);
	      }
	    }
	  }, {
	    key: 'componentWillReceiveProps',
	    value: function componentWillReceiveProps(nextprops) {
	      if (nextprops.load) {
	        nextprops.loadView(nextprops.currentPage, this.filter);
	      }
	    }
	  }, {
	    key: 'shouldComponentUpdate',
	    value: function shouldComponentUpdate(nextProps, nextState) {
	      if (!nextProps.forceUpdate && this.lastRenderTime) {
	        if (nextProps.lastUpdateTime) {
	          if (this.lastRenderTime >= nextProps.lastUpdateTime) {
	            return false;
	          }
	        } else {
	          return false;
	        }
	      }
	      return true;
	    }
	  }, {
	    key: 'addMethod',
	    value: function addMethod(name, method) {
	      this.methods[name] = method;
	    }
	  }, {
	    key: 'reload',
	    value: function reload() {
	      this.props.loadView(this.props.currentPage, this.filter);
	    }
	  }, {
	    key: 'canLoadMore',
	    value: function canLoadMore() {
	      return this.props.currentPage < this.props.totalPages;
	    }
	  }, {
	    key: 'viewrefs',
	    value: function viewrefs() {
	      return this.refs;
	    }
	  }, {
	    key: 'itemCount',
	    value: function itemCount() {
	      return this.numItems;
	    }
	  }, {
	    key: 'selectedItems',
	    value: function selectedItems() {
	      var selectedItems = [];
	      for (var i = 0; i < this.numItems; i++) {
	        var refName = "item" + i;
	        var item = this.refs[refName];
	        if (item.selected) {
	          selectedItems.push(item.id);
	        }
	      }
	      return selectedItems;
	    }
	  }, {
	    key: 'itemStatus',
	    value: function itemStatus() {
	      var items = {};
	      for (var i = 0; i < this.numItems; i++) {
	        var refName = "item" + i;
	        var item = this.refs[refName];
	        items[item.id] = item.selected;
	      }
	      return items;
	    }
	  }, {
	    key: 'setPage',
	    value: function setPage(newPage) {
	      this.props.loadView(newPage, this.filter);
	    }
	  }, {
	    key: 'setFilter',
	    value: function setFilter(filter) {
	      this.filter = filter;
	      this.props.loadView(1, this.filter);
	    }
	  }, {
	    key: 'getView',
	    value: function getView(items, currentPage, totalPages) {
	      if (this.props.getView) {
	        return this.props.getView(this, items, currentPage, totalPages);
	      }
	      return null;
	    }
	  }, {
	    key: 'loadMore',
	    value: function loadMore() {
	      if (this.props.currentPage >= this.props.totalPages) {
	        return false;
	      } else {
	        if (this.props.currentPage) {
	          this.props.loadIncrementally(this.props.currentPage + 1, this.filter);
	          return true;
	        }
	      }
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      this.lastRenderTime = this.props.lastUpdateTime;
	      var view = this.getView(this.props.items, this.props.currentPage, this.props.totalPages);
	      this.items = this.props.items;
	      return view;
	    }
	  }]);
	
	  return ViewDataComp;
	}(_react2.default.Component);
	
	var mapStateToProps = function mapStateToProps(state, ownProps) {
	  var props = {
	    reducer: ownProps.reducer,
	    paginate: ownProps.paginate,
	    pageSize: ownProps.pageSize,
	    defaultFilter: ownProps.defaultFilter,
	    externalLoad: ownProps.externalLoad,
	    urlParams: ownProps.urlParams,
	    postArgs: ownProps.postArgs,
	    getView: ownProps.getView,
	    incrementalLoad: ownProps.incrementalLoad,
	    currentPage: ownProps.currentPage,
	    className: ownProps.className,
	    style: ownProps.style,
	    totalPages: 1,
	    load: false,
	    items: null
	  };
	  var view = null;
	  if (!ownProps.globalReducer) {
	    if (state.router && state.router.routeStore) {
	      view = state.router.routeStore[ownProps.reducer];
	    }
	  } else {
	    view = state[ownProps.reducer];
	  }
	  if (view) {
	    if (view.status == "Loaded") {
	      props.items = view.data;
	      props.currentPage = view.currentPage;
	      props.totalPages = view.totalPages;
	      props.lastUpdateTime = view.lastUpdateTime;
	      props.latestPageData = view.latestPageData;
	      return props;
	    }
	    if (view.status == "NotLoaded") {
	      props.load = true;
	      return props;
	    }
	  }
	  return props;
	};
	
	function loadData(dispatch, ownProps, pagenum, filter, incrementalLoad) {
	  if (!pagenum) {
	    pagenum = 1;
	  }
	  var queryParams = {};
	  if (ownProps.paginate) {
	    queryParams.pagesize = ownProps.pageSize;
	    queryParams.pagenum = pagenum;
	  }
	  var postArgs = Object.assign({}, ownProps.postArgs, filter);
	  var payload = { queryParams: queryParams, postArgs: postArgs };
	  var meta = { serviceName: ownProps.viewService, reducer: ownProps.reducer, incrementalLoad: incrementalLoad };
	  dispatch((0, _utils.createAction)(_ActionNames.ActionNames.VIEW_FETCH, payload, meta));
	}
	
	var mapDispatchToProps = function mapDispatchToProps(dispatch, ownProps) {
	  return {
	    loadView: function loadView(pagenum, filter) {
	      loadData(dispatch, ownProps, pagenum, filter, false);
	    },
	    loadIncrementally: function loadIncrementally(pagenum, filter) {
	      loadData(dispatch, ownProps, pagenum, filter, true);
	    }
	  };
	};
	
	var ViewData = (0, _reactRedux.connect)(mapStateToProps, mapDispatchToProps, null, { withRef: true })(ViewDataComp);
	
	exports.ViewData = ViewData;

/***/ }),
/* 18 */
/***/ (function(module, exports, __webpack_require__) {

	"use strict";
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.EntityReducer = undefined;
	
	var _ActionNames = __webpack_require__(2);
	
	function EntityReducer(reducerName) {
	  var initialState = {
	    status: "NotLoaded",
	    entityId: "",
	    entityName: "",
	    data: {}
	  };
	
	  return function (state, action) {
	    //return state if this is not the correct copy of reducer
	    if (!action || !action.meta || !action.meta.reducer || reducerName != action.meta.reducer) {
	      if (!state) {
	        return initialState;
	      }
	      if (action.type == _ActionNames.ActionNames.LOGOUT) {
	        return initialState;
	      }
	      return state;
	    }
	    if (action.type) {
	      switch (action.type) {
	        case _ActionNames.ActionNames.ENTITY_GETTING:
	          return Object.assign({}, state, {
	            status: "Loading",
	            entityName: action.payload.entityName,
	            entityId: action.payload.entityId
	          });
	
	        case _ActionNames.ActionNames.ENTITY_GET_SUCCESS:
	          var st = Object.assign({}, state, {
	            status: "Loaded",
	            lastUpdateTime: new Date().getTime(),
	            data: action.payload.data
	          });
	          return st;
	        case _ActionNames.ActionNames.ENTITY_GET_FAILED:
	          {
	            return Object.assign({}, state, {
	              status: "LoadingFailed",
	              data: null
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_SAVING:
	          {
	            return Object.assign({}, state, {
	              status: "Saving",
	              entityName: action.payload.entityName
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_SAVE_SUCCESS:
	          {
	            return Object.assign({}, state, {
	              status: "Saved"
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_SAVE_FAILURE:
	          {
	            return Object.assign({}, state, {
	              status: "SavingFailed"
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_UPDATING:
	          {
	            return Object.assign({}, state, {
	              status: "Updating",
	              entityName: action.payload.entityName,
	              entityId: action.payload.entityId
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_UPDATE_SUCCESS:
	          {
	            return Object.assign({}, state, {
	              status: "Updated"
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_UPDATE_FAILURE:
	          {
	            return Object.assign({}, state, {
	              status: "UpdateFailed"
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_PUTTING:
	          {
	            return Object.assign({}, state, {
	              status: "Updating",
	              entityName: action.payload.entityName,
	              entityId: action.payload.entityId
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_PUT_SUCCESS:
	          {
	            return Object.assign({}, state, {
	              status: "Updated"
	            });
	          }
	
	        case _ActionNames.ActionNames.ENTITY_PUT_FAILURE:
	          {
	            return Object.assign({}, state, {
	              status: "UpdateFailed"
	            });
	          }
	
	        default:
	          if (!state) {
	            return initialState;
	          }
	          return state;
	      }
	    }
	  };
	}
	
	exports.EntityReducer = EntityReducer;

/***/ }),
/* 19 */
/***/ (function(module, exports, __webpack_require__) {

	"use strict";
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.ViewReducer = undefined;
	
	var _ActionNames = __webpack_require__(2);
	
	function ViewReducer(reducerName) {
	  var initialState = {
	    status: "NotLoaded",
	    data: {},
	    currentPage: 1,
	    totalPages: 1,
	    pagesize: -1
	  };
	
	  return function (state, action) {
	    //return state if this is not the correct copy of reducer
	    if (!action || !action.meta || !action.meta.reducer || reducerName != action.meta.reducer) {
	      if (!state) {
	        return initialState;
	      }
	      if (action.type == _ActionNames.ActionNames.LOGOUT) {
	        return initialState;
	      }
	      return state;
	    }
	    if (action.type) {
	      switch (action.type) {
	        case _ActionNames.ActionNames.VIEW_FETCHING:
	          var pagenum = 1;
	          var pagesize = -1;
	          if (action.payload.queryParams && action.payload.queryParams.pagenum) {
	            pagenum = action.payload.queryParams.pagenum;
	            pagesize = action.payload.queryParams.pagesize;
	          }
	          return Object.assign({}, state, {
	            status: "Fetching",
	            currentPage: pagenum,
	            pagesize: pagesize
	          });
	
	        case _ActionNames.ActionNames.VIEW_FETCH_SUCCESS:
	          var totalPages = 1;
	          if (action.meta.info && action.meta.info.totalrecords) {
	            var totalrecords = action.meta.info.totalrecords;
	            if (totalrecords > 0 && state.pagesize > 0) {
	              totalPages = Math.ceil(totalrecords / state.pagesize);
	            }
	          }
	          var newData = null;
	          var data = state.data;
	          if (data && action.meta.incrementalLoad) {
	            if (action.payload) {
	              if (Array.isArray(action.payload)) {
	                newData = data.concat(action.payload);
	              } else {
	                newData = Object.assign(data, action.payload);
	              }
	            }
	          } else {
	            newData = action.payload;
	          }
	          return Object.assign({}, state, {
	            status: "Loaded",
	            data: newData,
	            lastUpdateTime: new Date().getTime(),
	            totalPages: totalPages
	          });
	
	        case _ActionNames.ActionNames.VIEW_FETCH_FAILED:
	          {
	            return Object.assign({}, initialState, {
	              status: "LoadingFailed"
	            });
	          }
	
	        case _ActionNames.ActionNames.VIEW_ITEM_RELOAD:
	          {
	            var index = action.meta.Index;
	            if (index == null) {
	              return state;
	            }
	            var _data = state.data;
	            var _newData = null;
	            //remove by index for arrays and keys for map
	            if (Array.isArray(_data)) {
	              var ind = -1;
	              if (typeof index == "string") {
	                ind = parseInt(index);
	              } else {
	                ind = index;
	              }
	              _newData = _data.slice(0);
	              console.log("newData ", _newData, " index", ind, "  payload ", action.payload);
	              _newData[ind] = action.payload;
	            } else {
	              _newData = Object.assign({}, _data);
	              _newData[index] = action.payload;
	            }
	            return Object.assign({}, state, {
	              data: _newData,
	              lastUpdateTime: new Date().getTime()
	            });
	          }
	
	        case _ActionNames.ActionNames.VIEW_ITEM_REMOVE:
	          {
	            var _index = action.payload.Index;
	            if (_index == null) {
	              return state;
	            }
	            var _data2 = state.data;
	            var _newData2 = null;
	            //remove by index for arrays and keys for map
	            if (Array.isArray(_data2)) {
	              var _ind = -1;
	              if (typeof _index == "string") {
	                _ind = parseInt(_index);
	              } else {
	                _ind = _index;
	              }
	              _newData2 = _data2.slice(0);
	              _newData2.splice(_ind, 1);
	            } else {
	              _newData2 = Object.assign({}, _data2);
	              delete _newData2[_index];
	              if (_newData2 == null) {
	                _newData2 = {};
	              }
	            }
	            return Object.assign({}, state, {
	              data: _newData2,
	              lastUpdateTime: new Date().getTime()
	            });
	          }
	
	        default:
	          if (!state) {
	            return initialState;
	          }
	          return state;
	      }
	    }
	  };
	}
	
	exports.ViewReducer = ViewReducer;

/***/ }),
/* 20 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	exports.__esModule = true;
	exports.buffers = exports.BUFFER_OVERFLOW = undefined;
	
	var _utils = __webpack_require__(9);
	
	var BUFFER_OVERFLOW = exports.BUFFER_OVERFLOW = 'Channel\'s Buffer overflow!';
	
	var ON_OVERFLOW_THROW = 1;
	var ON_OVERFLOW_DROP = 2;
	var ON_OVERFLOW_SLIDE = 3;
	var ON_OVERFLOW_EXPAND = 4;
	
	var zeroBuffer = { isEmpty: _utils.kTrue, put: _utils.noop, take: _utils.noop };
	
	function ringBuffer() {
	  var limit = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 10;
	  var overflowAction = arguments[1];
	
	  var arr = new Array(limit);
	  var length = 0;
	  var pushIndex = 0;
	  var popIndex = 0;
	
	  var push = function push(it) {
	    arr[pushIndex] = it;
	    pushIndex = (pushIndex + 1) % limit;
	    length++;
	  };
	
	  var take = function take() {
	    if (length != 0) {
	      var it = arr[popIndex];
	      arr[popIndex] = null;
	      length--;
	      popIndex = (popIndex + 1) % limit;
	      return it;
	    }
	  };
	
	  var flush = function flush() {
	    var items = [];
	    while (length) {
	      items.push(take());
	    }
	    return items;
	  };
	
	  return {
	    isEmpty: function isEmpty() {
	      return length == 0;
	    },
	    put: function put(it) {
	      if (length < limit) {
	        push(it);
	      } else {
	        var doubledLimit = void 0;
	        switch (overflowAction) {
	          case ON_OVERFLOW_THROW:
	            throw new Error(BUFFER_OVERFLOW);
	          case ON_OVERFLOW_SLIDE:
	            arr[pushIndex] = it;
	            pushIndex = (pushIndex + 1) % limit;
	            popIndex = pushIndex;
	            break;
	          case ON_OVERFLOW_EXPAND:
	            doubledLimit = 2 * limit;
	
	            arr = flush();
	
	            length = arr.length;
	            pushIndex = arr.length;
	            popIndex = 0;
	
	            arr.length = doubledLimit;
	            limit = doubledLimit;
	
	            push(it);
	            break;
	          default:
	          // DROP
	        }
	      }
	    },
	    take: take, flush: flush
	  };
	}
	
	var buffers = exports.buffers = {
	  none: function none() {
	    return zeroBuffer;
	  },
	  fixed: function fixed(limit) {
	    return ringBuffer(limit, ON_OVERFLOW_THROW);
	  },
	  dropping: function dropping(limit) {
	    return ringBuffer(limit, ON_OVERFLOW_DROP);
	  },
	  sliding: function sliding(limit) {
	    return ringBuffer(limit, ON_OVERFLOW_SLIDE);
	  },
	  expanding: function expanding(initialSize) {
	    return ringBuffer(initialSize, ON_OVERFLOW_EXPAND);
	  }
	};

/***/ }),
/* 21 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	exports.__esModule = true;
	exports.asEffect = exports.takem = undefined;
	exports.take = take;
	exports.put = put;
	exports.all = all;
	exports.race = race;
	exports.call = call;
	exports.apply = apply;
	exports.cps = cps;
	exports.fork = fork;
	exports.spawn = spawn;
	exports.join = join;
	exports.cancel = cancel;
	exports.select = select;
	exports.actionChannel = actionChannel;
	exports.cancelled = cancelled;
	exports.flush = flush;
	exports.getContext = getContext;
	exports.setContext = setContext;
	exports.takeEvery = takeEvery;
	exports.takeLatest = takeLatest;
	exports.throttle = throttle;
	
	var _utils = __webpack_require__(9);
	
	var _sagaHelpers = __webpack_require__(62);
	
	var IO = (0, _utils.sym)('IO');
	var TAKE = 'TAKE';
	var PUT = 'PUT';
	var ALL = 'ALL';
	var RACE = 'RACE';
	var CALL = 'CALL';
	var CPS = 'CPS';
	var FORK = 'FORK';
	var JOIN = 'JOIN';
	var CANCEL = 'CANCEL';
	var SELECT = 'SELECT';
	var ACTION_CHANNEL = 'ACTION_CHANNEL';
	var CANCELLED = 'CANCELLED';
	var FLUSH = 'FLUSH';
	var GET_CONTEXT = 'GET_CONTEXT';
	var SET_CONTEXT = 'SET_CONTEXT';
	
	var TEST_HINT = '\n(HINT: if you are getting this errors in tests, consider using createMockTask from redux-saga/utils)';
	
	var effect = function effect(type, payload) {
	  var _ref;
	
	  return _ref = {}, _ref[IO] = true, _ref[type] = payload, _ref;
	};
	
	function take() {
	  var patternOrChannel = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : '*';
	
	  if (arguments.length) {
	    (0, _utils.check)(arguments[0], _utils.is.notUndef, 'take(patternOrChannel): patternOrChannel is undefined');
	  }
	  if (_utils.is.pattern(patternOrChannel)) {
	    return effect(TAKE, { pattern: patternOrChannel });
	  }
	  if (_utils.is.channel(patternOrChannel)) {
	    return effect(TAKE, { channel: patternOrChannel });
	  }
	  throw new Error('take(patternOrChannel): argument ' + String(patternOrChannel) + ' is not valid channel or a valid pattern');
	}
	
	take.maybe = function () {
	  var eff = take.apply(undefined, arguments);
	  eff[TAKE].maybe = true;
	  return eff;
	};
	
	var takem = exports.takem = (0, _utils.deprecate)(take.maybe, (0, _utils.updateIncentive)('takem', 'take.maybe'));
	
	function put(channel, action) {
	  if (arguments.length > 1) {
	    (0, _utils.check)(channel, _utils.is.notUndef, 'put(channel, action): argument channel is undefined');
	    (0, _utils.check)(channel, _utils.is.channel, 'put(channel, action): argument ' + channel + ' is not a valid channel');
	    (0, _utils.check)(action, _utils.is.notUndef, 'put(channel, action): argument action is undefined');
	  } else {
	    (0, _utils.check)(channel, _utils.is.notUndef, 'put(action): argument action is undefined');
	    action = channel;
	    channel = null;
	  }
	  return effect(PUT, { channel: channel, action: action });
	}
	
	put.resolve = function () {
	  var eff = put.apply(undefined, arguments);
	  eff[PUT].resolve = true;
	  return eff;
	};
	
	put.sync = (0, _utils.deprecate)(put.resolve, (0, _utils.updateIncentive)('put.sync', 'put.resolve'));
	
	function all(effects) {
	  return effect(ALL, effects);
	}
	
	function race(effects) {
	  return effect(RACE, effects);
	}
	
	function getFnCallDesc(meth, fn, args) {
	  (0, _utils.check)(fn, _utils.is.notUndef, meth + ': argument fn is undefined');
	
	  var context = null;
	  if (_utils.is.array(fn)) {
	    var _fn = fn;
	    context = _fn[0];
	    fn = _fn[1];
	  } else if (fn.fn) {
	    var _fn2 = fn;
	    context = _fn2.context;
	    fn = _fn2.fn;
	  }
	  if (context && _utils.is.string(fn) && _utils.is.func(context[fn])) {
	    fn = context[fn];
	  }
	  (0, _utils.check)(fn, _utils.is.func, meth + ': argument ' + fn + ' is not a function');
	
	  return { context: context, fn: fn, args: args };
	}
	
	function call(fn) {
	  for (var _len = arguments.length, args = Array(_len > 1 ? _len - 1 : 0), _key = 1; _key < _len; _key++) {
	    args[_key - 1] = arguments[_key];
	  }
	
	  return effect(CALL, getFnCallDesc('call', fn, args));
	}
	
	function apply(context, fn) {
	  var args = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : [];
	
	  return effect(CALL, getFnCallDesc('apply', { context: context, fn: fn }, args));
	}
	
	function cps(fn) {
	  for (var _len2 = arguments.length, args = Array(_len2 > 1 ? _len2 - 1 : 0), _key2 = 1; _key2 < _len2; _key2++) {
	    args[_key2 - 1] = arguments[_key2];
	  }
	
	  return effect(CPS, getFnCallDesc('cps', fn, args));
	}
	
	function fork(fn) {
	  for (var _len3 = arguments.length, args = Array(_len3 > 1 ? _len3 - 1 : 0), _key3 = 1; _key3 < _len3; _key3++) {
	    args[_key3 - 1] = arguments[_key3];
	  }
	
	  return effect(FORK, getFnCallDesc('fork', fn, args));
	}
	
	function spawn(fn) {
	  for (var _len4 = arguments.length, args = Array(_len4 > 1 ? _len4 - 1 : 0), _key4 = 1; _key4 < _len4; _key4++) {
	    args[_key4 - 1] = arguments[_key4];
	  }
	
	  var eff = fork.apply(undefined, [fn].concat(args));
	  eff[FORK].detached = true;
	  return eff;
	}
	
	function join() {
	  for (var _len5 = arguments.length, tasks = Array(_len5), _key5 = 0; _key5 < _len5; _key5++) {
	    tasks[_key5] = arguments[_key5];
	  }
	
	  if (tasks.length > 1) {
	    return all(tasks.map(function (t) {
	      return join(t);
	    }));
	  }
	  var task = tasks[0];
	  (0, _utils.check)(task, _utils.is.notUndef, 'join(task): argument task is undefined');
	  (0, _utils.check)(task, _utils.is.task, 'join(task): argument ' + task + ' is not a valid Task object ' + TEST_HINT);
	  return effect(JOIN, task);
	}
	
	function cancel() {
	  for (var _len6 = arguments.length, tasks = Array(_len6), _key6 = 0; _key6 < _len6; _key6++) {
	    tasks[_key6] = arguments[_key6];
	  }
	
	  if (tasks.length > 1) {
	    return all(tasks.map(function (t) {
	      return cancel(t);
	    }));
	  }
	  var task = tasks[0];
	  if (tasks.length === 1) {
	    (0, _utils.check)(task, _utils.is.notUndef, 'cancel(task): argument task is undefined');
	    (0, _utils.check)(task, _utils.is.task, 'cancel(task): argument ' + task + ' is not a valid Task object ' + TEST_HINT);
	  }
	  return effect(CANCEL, task || _utils.SELF_CANCELLATION);
	}
	
	function select(selector) {
	  for (var _len7 = arguments.length, args = Array(_len7 > 1 ? _len7 - 1 : 0), _key7 = 1; _key7 < _len7; _key7++) {
	    args[_key7 - 1] = arguments[_key7];
	  }
	
	  if (arguments.length === 0) {
	    selector = _utils.ident;
	  } else {
	    (0, _utils.check)(selector, _utils.is.notUndef, 'select(selector,[...]): argument selector is undefined');
	    (0, _utils.check)(selector, _utils.is.func, 'select(selector,[...]): argument ' + selector + ' is not a function');
	  }
	  return effect(SELECT, { selector: selector, args: args });
	}
	
	/**
	  channel(pattern, [buffer])    => creates an event channel for store actions
	**/
	function actionChannel(pattern, buffer) {
	  (0, _utils.check)(pattern, _utils.is.notUndef, 'actionChannel(pattern,...): argument pattern is undefined');
	  if (arguments.length > 1) {
	    (0, _utils.check)(buffer, _utils.is.notUndef, 'actionChannel(pattern, buffer): argument buffer is undefined');
	    (0, _utils.check)(buffer, _utils.is.buffer, 'actionChannel(pattern, buffer): argument ' + buffer + ' is not a valid buffer');
	  }
	  return effect(ACTION_CHANNEL, { pattern: pattern, buffer: buffer });
	}
	
	function cancelled() {
	  return effect(CANCELLED, {});
	}
	
	function flush(channel) {
	  (0, _utils.check)(channel, _utils.is.channel, 'flush(channel): argument ' + channel + ' is not valid channel');
	  return effect(FLUSH, channel);
	}
	
	function getContext(prop) {
	  (0, _utils.check)(prop, _utils.is.string, 'getContext(prop): argument ' + prop + ' is not a string');
	  return effect(GET_CONTEXT, prop);
	}
	
	function setContext(props) {
	  (0, _utils.check)(props, _utils.is.object, (0, _utils.createSetContextWarning)(null, props));
	  return effect(SET_CONTEXT, props);
	}
	
	function takeEvery(patternOrChannel, worker) {
	  for (var _len8 = arguments.length, args = Array(_len8 > 2 ? _len8 - 2 : 0), _key8 = 2; _key8 < _len8; _key8++) {
	    args[_key8 - 2] = arguments[_key8];
	  }
	
	  return fork.apply(undefined, [_sagaHelpers.takeEveryHelper, patternOrChannel, worker].concat(args));
	}
	
	function takeLatest(patternOrChannel, worker) {
	  for (var _len9 = arguments.length, args = Array(_len9 > 2 ? _len9 - 2 : 0), _key9 = 2; _key9 < _len9; _key9++) {
	    args[_key9 - 2] = arguments[_key9];
	  }
	
	  return fork.apply(undefined, [_sagaHelpers.takeLatestHelper, patternOrChannel, worker].concat(args));
	}
	
	function throttle(ms, pattern, worker) {
	  for (var _len10 = arguments.length, args = Array(_len10 > 3 ? _len10 - 3 : 0), _key10 = 3; _key10 < _len10; _key10++) {
	    args[_key10 - 3] = arguments[_key10];
	  }
	
	  return fork.apply(undefined, [_sagaHelpers.throttleHelper, ms, pattern, worker].concat(args));
	}
	
	var createAsEffectType = function createAsEffectType(type) {
	  return function (effect) {
	    return effect && effect[IO] && effect[type];
	  };
	};
	
	var asEffect = exports.asEffect = {
	  take: createAsEffectType(TAKE),
	  put: createAsEffectType(PUT),
	  all: createAsEffectType(ALL),
	  race: createAsEffectType(RACE),
	  call: createAsEffectType(CALL),
	  cps: createAsEffectType(CPS),
	  fork: createAsEffectType(FORK),
	  join: createAsEffectType(JOIN),
	  cancel: createAsEffectType(CANCEL),
	  select: createAsEffectType(SELECT),
	  actionChannel: createAsEffectType(ACTION_CHANNEL),
	  cancelled: createAsEffectType(CANCELLED),
	  flush: createAsEffectType(FLUSH),
	  getContext: createAsEffectType(GET_CONTEXT),
	  setContext: createAsEffectType(SET_CONTEXT)
	};

/***/ }),
/* 22 */
/***/ (function(module, exports) {

	module.exports = __WEBPACK_EXTERNAL_MODULE_22__;

/***/ }),
/* 23 */
/***/ (function(module, exports, __webpack_require__) {

	module.exports = __webpack_require__(24);

/***/ }),
/* 24 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	var bind = __webpack_require__(16);
	var Axios = __webpack_require__(26);
	var defaults = __webpack_require__(11);
	
	/**
	 * Create an instance of Axios
	 *
	 * @param {Object} defaultConfig The default config for the instance
	 * @return {Axios} A new instance of Axios
	 */
	function createInstance(defaultConfig) {
	  var context = new Axios(defaultConfig);
	  var instance = bind(Axios.prototype.request, context);
	
	  // Copy axios.prototype to instance
	  utils.extend(instance, Axios.prototype, context);
	
	  // Copy context to instance
	  utils.extend(instance, context);
	
	  return instance;
	}
	
	// Create the default instance to be exported
	var axios = createInstance(defaults);
	
	// Expose Axios class to allow class inheritance
	axios.Axios = Axios;
	
	// Factory for creating new instances
	axios.create = function create(instanceConfig) {
	  return createInstance(utils.merge(defaults, instanceConfig));
	};
	
	// Expose Cancel & CancelToken
	axios.Cancel = __webpack_require__(13);
	axios.CancelToken = __webpack_require__(25);
	axios.isCancel = __webpack_require__(14);
	
	// Expose all/spread
	axios.all = function all(promises) {
	  return Promise.all(promises);
	};
	axios.spread = __webpack_require__(40);
	
	module.exports = axios;
	
	// Allow use of default import syntax in TypeScript
	module.exports.default = axios;


/***/ }),
/* 25 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var Cancel = __webpack_require__(13);
	
	/**
	 * A `CancelToken` is an object that can be used to request cancellation of an operation.
	 *
	 * @class
	 * @param {Function} executor The executor function.
	 */
	function CancelToken(executor) {
	  if (typeof executor !== 'function') {
	    throw new TypeError('executor must be a function.');
	  }
	
	  var resolvePromise;
	  this.promise = new Promise(function promiseExecutor(resolve) {
	    resolvePromise = resolve;
	  });
	
	  var token = this;
	  executor(function cancel(message) {
	    if (token.reason) {
	      // Cancellation has already been requested
	      return;
	    }
	
	    token.reason = new Cancel(message);
	    resolvePromise(token.reason);
	  });
	}
	
	/**
	 * Throws a `Cancel` if cancellation has been requested.
	 */
	CancelToken.prototype.throwIfRequested = function throwIfRequested() {
	  if (this.reason) {
	    throw this.reason;
	  }
	};
	
	/**
	 * Returns an object that contains a new `CancelToken` and a function that, when called,
	 * cancels the `CancelToken`.
	 */
	CancelToken.source = function source() {
	  var cancel;
	  var token = new CancelToken(function executor(c) {
	    cancel = c;
	  });
	  return {
	    token: token,
	    cancel: cancel
	  };
	};
	
	module.exports = CancelToken;


/***/ }),
/* 26 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var defaults = __webpack_require__(11);
	var utils = __webpack_require__(1);
	var InterceptorManager = __webpack_require__(27);
	var dispatchRequest = __webpack_require__(28);
	var isAbsoluteURL = __webpack_require__(36);
	var combineURLs = __webpack_require__(34);
	
	/**
	 * Create a new instance of Axios
	 *
	 * @param {Object} instanceConfig The default config for the instance
	 */
	function Axios(instanceConfig) {
	  this.defaults = instanceConfig;
	  this.interceptors = {
	    request: new InterceptorManager(),
	    response: new InterceptorManager()
	  };
	}
	
	/**
	 * Dispatch a request
	 *
	 * @param {Object} config The config specific for this request (merged with this.defaults)
	 */
	Axios.prototype.request = function request(config) {
	  /*eslint no-param-reassign:0*/
	  // Allow for axios('example/url'[, config]) a la fetch API
	  if (typeof config === 'string') {
	    config = utils.merge({
	      url: arguments[0]
	    }, arguments[1]);
	  }
	
	  config = utils.merge(defaults, this.defaults, { method: 'get' }, config);
	
	  // Support baseURL config
	  if (config.baseURL && !isAbsoluteURL(config.url)) {
	    config.url = combineURLs(config.baseURL, config.url);
	  }
	
	  // Hook up interceptors middleware
	  var chain = [dispatchRequest, undefined];
	  var promise = Promise.resolve(config);
	
	  this.interceptors.request.forEach(function unshiftRequestInterceptors(interceptor) {
	    chain.unshift(interceptor.fulfilled, interceptor.rejected);
	  });
	
	  this.interceptors.response.forEach(function pushResponseInterceptors(interceptor) {
	    chain.push(interceptor.fulfilled, interceptor.rejected);
	  });
	
	  while (chain.length) {
	    promise = promise.then(chain.shift(), chain.shift());
	  }
	
	  return promise;
	};
	
	// Provide aliases for supported request methods
	utils.forEach(['delete', 'get', 'head', 'options'], function forEachMethodNoData(method) {
	  /*eslint func-names:0*/
	  Axios.prototype[method] = function(url, config) {
	    return this.request(utils.merge(config || {}, {
	      method: method,
	      url: url
	    }));
	  };
	});
	
	utils.forEach(['post', 'put', 'patch'], function forEachMethodWithData(method) {
	  /*eslint func-names:0*/
	  Axios.prototype[method] = function(url, data, config) {
	    return this.request(utils.merge(config || {}, {
	      method: method,
	      url: url,
	      data: data
	    }));
	  };
	});
	
	module.exports = Axios;


/***/ }),
/* 27 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	
	function InterceptorManager() {
	  this.handlers = [];
	}
	
	/**
	 * Add a new interceptor to the stack
	 *
	 * @param {Function} fulfilled The function to handle `then` for a `Promise`
	 * @param {Function} rejected The function to handle `reject` for a `Promise`
	 *
	 * @return {Number} An ID used to remove interceptor later
	 */
	InterceptorManager.prototype.use = function use(fulfilled, rejected) {
	  this.handlers.push({
	    fulfilled: fulfilled,
	    rejected: rejected
	  });
	  return this.handlers.length - 1;
	};
	
	/**
	 * Remove an interceptor from the stack
	 *
	 * @param {Number} id The ID that was returned by `use`
	 */
	InterceptorManager.prototype.eject = function eject(id) {
	  if (this.handlers[id]) {
	    this.handlers[id] = null;
	  }
	};
	
	/**
	 * Iterate over all the registered interceptors
	 *
	 * This method is particularly useful for skipping over any
	 * interceptors that may have become `null` calling `eject`.
	 *
	 * @param {Function} fn The function to call for each interceptor
	 */
	InterceptorManager.prototype.forEach = function forEach(fn) {
	  utils.forEach(this.handlers, function forEachHandler(h) {
	    if (h !== null) {
	      fn(h);
	    }
	  });
	};
	
	module.exports = InterceptorManager;


/***/ }),
/* 28 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	var transformData = __webpack_require__(31);
	var isCancel = __webpack_require__(14);
	var defaults = __webpack_require__(11);
	
	/**
	 * Throws a `Cancel` if cancellation has been requested.
	 */
	function throwIfCancellationRequested(config) {
	  if (config.cancelToken) {
	    config.cancelToken.throwIfRequested();
	  }
	}
	
	/**
	 * Dispatch a request to the server using the configured adapter.
	 *
	 * @param {object} config The config that is to be used for the request
	 * @returns {Promise} The Promise to be fulfilled
	 */
	module.exports = function dispatchRequest(config) {
	  throwIfCancellationRequested(config);
	
	  // Ensure headers exist
	  config.headers = config.headers || {};
	
	  // Transform request data
	  config.data = transformData(
	    config.data,
	    config.headers,
	    config.transformRequest
	  );
	
	  // Flatten headers
	  config.headers = utils.merge(
	    config.headers.common || {},
	    config.headers[config.method] || {},
	    config.headers || {}
	  );
	
	  utils.forEach(
	    ['delete', 'get', 'head', 'post', 'put', 'patch', 'common'],
	    function cleanHeaderConfig(method) {
	      delete config.headers[method];
	    }
	  );
	
	  var adapter = config.adapter || defaults.adapter;
	
	  return adapter(config).then(function onAdapterResolution(response) {
	    throwIfCancellationRequested(config);
	
	    // Transform response data
	    response.data = transformData(
	      response.data,
	      response.headers,
	      config.transformResponse
	    );
	
	    return response;
	  }, function onAdapterRejection(reason) {
	    if (!isCancel(reason)) {
	      throwIfCancellationRequested(config);
	
	      // Transform response data
	      if (reason && reason.response) {
	        reason.response.data = transformData(
	          reason.response.data,
	          reason.response.headers,
	          config.transformResponse
	        );
	      }
	    }
	
	    return Promise.reject(reason);
	  });
	};


/***/ }),
/* 29 */
/***/ (function(module, exports) {

	'use strict';
	
	/**
	 * Update an Error with the specified config, error code, and response.
	 *
	 * @param {Error} error The error to update.
	 * @param {Object} config The config.
	 * @param {string} [code] The error code (for example, 'ECONNABORTED').
	 @ @param {Object} [response] The response.
	 * @returns {Error} The error.
	 */
	module.exports = function enhanceError(error, config, code, response) {
	  error.config = config;
	  if (code) {
	    error.code = code;
	  }
	  error.response = response;
	  return error;
	};


/***/ }),
/* 30 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var createError = __webpack_require__(15);
	
	/**
	 * Resolve or reject a Promise based on response status.
	 *
	 * @param {Function} resolve A function that resolves the promise.
	 * @param {Function} reject A function that rejects the promise.
	 * @param {object} response The response.
	 */
	module.exports = function settle(resolve, reject, response) {
	  var validateStatus = response.config.validateStatus;
	  // Note: status is not exposed by XDomainRequest
	  if (!response.status || !validateStatus || validateStatus(response.status)) {
	    resolve(response);
	  } else {
	    reject(createError(
	      'Request failed with status code ' + response.status,
	      response.config,
	      null,
	      response
	    ));
	  }
	};


/***/ }),
/* 31 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	
	/**
	 * Transform the data for a request or a response
	 *
	 * @param {Object|String} data The data to be transformed
	 * @param {Array} headers The headers for the request or response
	 * @param {Array|Function} fns A single function or Array of functions
	 * @returns {*} The resulting transformed data
	 */
	module.exports = function transformData(data, headers, fns) {
	  /*eslint no-param-reassign:0*/
	  utils.forEach(fns, function transform(fn) {
	    data = fn(data, headers);
	  });
	
	  return data;
	};


/***/ }),
/* 32 */
/***/ (function(module, exports) {

	'use strict';
	
	// btoa polyfill for IE<10 courtesy https://github.com/davidchambers/Base64.js
	
	var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
	
	function E() {
	  this.message = 'String contains an invalid character';
	}
	E.prototype = new Error;
	E.prototype.code = 5;
	E.prototype.name = 'InvalidCharacterError';
	
	function btoa(input) {
	  var str = String(input);
	  var output = '';
	  for (
	    // initialize result and counter
	    var block, charCode, idx = 0, map = chars;
	    // if the next str index does not exist:
	    //   change the mapping table to "="
	    //   check if d has no fractional digits
	    str.charAt(idx | 0) || (map = '=', idx % 1);
	    // "8 - idx % 1 * 8" generates the sequence 2, 4, 6, 8
	    output += map.charAt(63 & block >> 8 - idx % 1 * 8)
	  ) {
	    charCode = str.charCodeAt(idx += 3 / 4);
	    if (charCode > 0xFF) {
	      throw new E();
	    }
	    block = block << 8 | charCode;
	  }
	  return output;
	}
	
	module.exports = btoa;


/***/ }),
/* 33 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	
	function encode(val) {
	  return encodeURIComponent(val).
	    replace(/%40/gi, '@').
	    replace(/%3A/gi, ':').
	    replace(/%24/g, '$').
	    replace(/%2C/gi, ',').
	    replace(/%20/g, '+').
	    replace(/%5B/gi, '[').
	    replace(/%5D/gi, ']');
	}
	
	/**
	 * Build a URL by appending params to the end
	 *
	 * @param {string} url The base of the url (e.g., http://www.google.com)
	 * @param {object} [params] The params to be appended
	 * @returns {string} The formatted url
	 */
	module.exports = function buildURL(url, params, paramsSerializer) {
	  /*eslint no-param-reassign:0*/
	  if (!params) {
	    return url;
	  }
	
	  var serializedParams;
	  if (paramsSerializer) {
	    serializedParams = paramsSerializer(params);
	  } else if (utils.isURLSearchParams(params)) {
	    serializedParams = params.toString();
	  } else {
	    var parts = [];
	
	    utils.forEach(params, function serialize(val, key) {
	      if (val === null || typeof val === 'undefined') {
	        return;
	      }
	
	      if (utils.isArray(val)) {
	        key = key + '[]';
	      }
	
	      if (!utils.isArray(val)) {
	        val = [val];
	      }
	
	      utils.forEach(val, function parseValue(v) {
	        if (utils.isDate(v)) {
	          v = v.toISOString();
	        } else if (utils.isObject(v)) {
	          v = JSON.stringify(v);
	        }
	        parts.push(encode(key) + '=' + encode(v));
	      });
	    });
	
	    serializedParams = parts.join('&');
	  }
	
	  if (serializedParams) {
	    url += (url.indexOf('?') === -1 ? '?' : '&') + serializedParams;
	  }
	
	  return url;
	};


/***/ }),
/* 34 */
/***/ (function(module, exports) {

	'use strict';
	
	/**
	 * Creates a new URL by combining the specified URLs
	 *
	 * @param {string} baseURL The base URL
	 * @param {string} relativeURL The relative URL
	 * @returns {string} The combined URL
	 */
	module.exports = function combineURLs(baseURL, relativeURL) {
	  return relativeURL
	    ? baseURL.replace(/\/+$/, '') + '/' + relativeURL.replace(/^\/+/, '')
	    : baseURL;
	};


/***/ }),
/* 35 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	
	module.exports = (
	  utils.isStandardBrowserEnv() ?
	
	  // Standard browser envs support document.cookie
	  (function standardBrowserEnv() {
	    return {
	      write: function write(name, value, expires, path, domain, secure) {
	        var cookie = [];
	        cookie.push(name + '=' + encodeURIComponent(value));
	
	        if (utils.isNumber(expires)) {
	          cookie.push('expires=' + new Date(expires).toGMTString());
	        }
	
	        if (utils.isString(path)) {
	          cookie.push('path=' + path);
	        }
	
	        if (utils.isString(domain)) {
	          cookie.push('domain=' + domain);
	        }
	
	        if (secure === true) {
	          cookie.push('secure');
	        }
	
	        document.cookie = cookie.join('; ');
	      },
	
	      read: function read(name) {
	        var match = document.cookie.match(new RegExp('(^|;\\s*)(' + name + ')=([^;]*)'));
	        return (match ? decodeURIComponent(match[3]) : null);
	      },
	
	      remove: function remove(name) {
	        this.write(name, '', Date.now() - 86400000);
	      }
	    };
	  })() :
	
	  // Non standard browser env (web workers, react-native) lack needed support.
	  (function nonStandardBrowserEnv() {
	    return {
	      write: function write() {},
	      read: function read() { return null; },
	      remove: function remove() {}
	    };
	  })()
	);


/***/ }),
/* 36 */
/***/ (function(module, exports) {

	'use strict';
	
	/**
	 * Determines whether the specified URL is absolute
	 *
	 * @param {string} url The URL to test
	 * @returns {boolean} True if the specified URL is absolute, otherwise false
	 */
	module.exports = function isAbsoluteURL(url) {
	  // A URL is considered absolute if it begins with "<scheme>://" or "//" (protocol-relative URL).
	  // RFC 3986 defines scheme name as a sequence of characters beginning with a letter and followed
	  // by any combination of letters, digits, plus, period, or hyphen.
	  return /^([a-z][a-z\d\+\-\.]*:)?\/\//i.test(url);
	};


/***/ }),
/* 37 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	
	module.exports = (
	  utils.isStandardBrowserEnv() ?
	
	  // Standard browser envs have full support of the APIs needed to test
	  // whether the request URL is of the same origin as current location.
	  (function standardBrowserEnv() {
	    var msie = /(msie|trident)/i.test(navigator.userAgent);
	    var urlParsingNode = document.createElement('a');
	    var originURL;
	
	    /**
	    * Parse a URL to discover it's components
	    *
	    * @param {String} url The URL to be parsed
	    * @returns {Object}
	    */
	    function resolveURL(url) {
	      var href = url;
	
	      if (msie) {
	        // IE needs attribute set twice to normalize properties
	        urlParsingNode.setAttribute('href', href);
	        href = urlParsingNode.href;
	      }
	
	      urlParsingNode.setAttribute('href', href);
	
	      // urlParsingNode provides the UrlUtils interface - http://url.spec.whatwg.org/#urlutils
	      return {
	        href: urlParsingNode.href,
	        protocol: urlParsingNode.protocol ? urlParsingNode.protocol.replace(/:$/, '') : '',
	        host: urlParsingNode.host,
	        search: urlParsingNode.search ? urlParsingNode.search.replace(/^\?/, '') : '',
	        hash: urlParsingNode.hash ? urlParsingNode.hash.replace(/^#/, '') : '',
	        hostname: urlParsingNode.hostname,
	        port: urlParsingNode.port,
	        pathname: (urlParsingNode.pathname.charAt(0) === '/') ?
	                  urlParsingNode.pathname :
	                  '/' + urlParsingNode.pathname
	      };
	    }
	
	    originURL = resolveURL(window.location.href);
	
	    /**
	    * Determine if a URL shares the same origin as the current location
	    *
	    * @param {String} requestURL The URL to test
	    * @returns {boolean} True if URL shares the same origin, otherwise false
	    */
	    return function isURLSameOrigin(requestURL) {
	      var parsed = (utils.isString(requestURL)) ? resolveURL(requestURL) : requestURL;
	      return (parsed.protocol === originURL.protocol &&
	            parsed.host === originURL.host);
	    };
	  })() :
	
	  // Non standard browser envs (web workers, react-native) lack needed support.
	  (function nonStandardBrowserEnv() {
	    return function isURLSameOrigin() {
	      return true;
	    };
	  })()
	);


/***/ }),
/* 38 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	
	module.exports = function normalizeHeaderName(headers, normalizedName) {
	  utils.forEach(headers, function processHeader(value, name) {
	    if (name !== normalizedName && name.toUpperCase() === normalizedName.toUpperCase()) {
	      headers[normalizedName] = value;
	      delete headers[name];
	    }
	  });
	};


/***/ }),
/* 39 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	var utils = __webpack_require__(1);
	
	/**
	 * Parse headers into an object
	 *
	 * ```
	 * Date: Wed, 27 Aug 2014 08:58:49 GMT
	 * Content-Type: application/json
	 * Connection: keep-alive
	 * Transfer-Encoding: chunked
	 * ```
	 *
	 * @param {String} headers Headers needing to be parsed
	 * @returns {Object} Headers parsed into an object
	 */
	module.exports = function parseHeaders(headers) {
	  var parsed = {};
	  var key;
	  var val;
	  var i;
	
	  if (!headers) { return parsed; }
	
	  utils.forEach(headers.split('\n'), function parser(line) {
	    i = line.indexOf(':');
	    key = utils.trim(line.substr(0, i)).toLowerCase();
	    val = utils.trim(line.substr(i + 1));
	
	    if (key) {
	      parsed[key] = parsed[key] ? parsed[key] + ', ' + val : val;
	    }
	  });
	
	  return parsed;
	};


/***/ }),
/* 40 */
/***/ (function(module, exports) {

	'use strict';
	
	/**
	 * Syntactic sugar for invoking a function and expanding an array for arguments.
	 *
	 * Common use case would be to use `Function.prototype.apply`.
	 *
	 *  ```js
	 *  function f(x, y, z) {}
	 *  var args = [1, 2, 3];
	 *  f.apply(null, args);
	 *  ```
	 *
	 * With `spread` this example can be re-written.
	 *
	 *  ```js
	 *  spread(function(x, y, z) {})([1, 2, 3]);
	 *  ```
	 *
	 * @param {Function} callback
	 * @returns {Function}
	 */
	module.exports = function spread(callback) {
	  return function wrap(arr) {
	    return callback.apply(null, arr);
	  };
	};


/***/ }),
/* 41 */
/***/ (function(module, exports) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	var Color = {
	  WHITE: '#ffffff',
	  BLACK: '#000000',
	  TRANSPARENT: 'transparent',
	  DELTAGREY: {
	    50: '#e4e4e4',
	    100: '#bdbdbd',
	    200: '#a1a1a1',
	    300: '#7e7e7e',
	    400: '#6e6e6e',
	    500: '#5f5f5f',
	    600: '#505050',
	    700: '#404040',
	    800: '#313131',
	    900: '#222222',
	    A100: '#e4e4e4',
	    A200: '#bdbdbd',
	    A400: '#6e6e6e',
	    A700: '#404040'
	  },
	  DELTABLUE: {
	    50: '#e6f5ff',
	    100: '#9ad8ff',
	    200: '#62c2ff',
	    300: '#1aa7ff',
	    400: '#009afb',
	    500: '#0087dc',
	    600: '#0074bd',
	    700: '#00619f',
	    800: '#004f80',
	    900: '#003c62',
	    A100: '#e6f5ff',
	    A200: '#9ad8ff',
	    A400: '#009afb',
	    A700: '#00619f'
	  },
	  DELTAORANGE: {
	    50: '#FFFDFA',
	    100: '#FFDAAE',
	    200: '#FFC076',
	    300: '#FF9F2E',
	    400: '#FF9110',
	    500: '#F08200',
	    600: '#D17100',
	    700: '#B36100',
	    800: '#945000',
	    900: '#764000',
	    A100: '#FFFDFA',
	    A200: '#FFDAAE',
	    A400: '#FF9110',
	    A700: '#B36100'
	  },
	  DELTAGREEN: {
	    50: '#C5F8CD',
	    100: '#81EF91',
	    200: '#50E965',
	    300: '#1BD636',
	    400: '#17BB2F',
	    500: '#14A028',
	    600: '#118521',
	    700: '#0D6A1A',
	    800: '#0A4E14',
	    900: '#06330D',
	    A100: '#C5F8CD',
	    A200: '#81EF91',
	    A400: '#17BB2F',
	    A700: '#0D6A1A'
	  },
	  RED: {
	    50: 'ffebee',
	    100: '#ffcdd2',
	    200: '#ef9a9a',
	    300: '#e57373',
	    400: '#ef5350',
	    500: '#f44336',
	    600: '#e53935',
	    700: '#d32f2f',
	    800: '#c62828',
	    900: '#b71c1c',
	    A100: '#ff8a80',
	    A200: '#ff5252',
	    A400: '#ff1744',
	    A700: '#d50000'
	  },
	  PINK: {
	    50: '#fce4ec',
	    100: '#f8bbd0',
	    200: '#f48fb1',
	    300: '#f06292',
	    400: '#ec407a',
	    500: '#e91e63',
	    600: '#d81b60',
	    700: '#c2185b',
	    800: '#ad1457',
	    900: '#880e4f',
	    A100: '#ff80ab',
	    A200: '#ff4081',
	    A400: '#f50057',
	    A700: '#c51162'
	  },
	  PURPLE: {
	    50: '#f3e5f5',
	    100: '#e1bee7',
	    200: '#ce93d8',
	    300: '#ba68c8',
	    400: '#ab47bc',
	    500: '#9c27b0',
	    600: '#8e24aa',
	    700: '#7b1fa2',
	    800: '#6a1b9a',
	    900: '#4a148c',
	    A100: '#ea80fc',
	    A200: '#e040fb',
	    A400: '#d500f9',
	    A700: '#aa00ff'
	  },
	  DEEPPRUPLE: {
	    50: '#ede7f6',
	    100: '#d1c4e9',
	    200: '#b39ddb',
	    300: '#9575cd',
	    400: '#7e57c2',
	    500: '#673ab7',
	    600: '#5e35b1',
	    700: '#512DA8',
	    800: '#4527A0',
	    900: '#311B92',
	    A100: '#b388ff',
	    A200: '#7c4dff',
	    A400: '#651fff',
	    A700: '#6200ea'
	  },
	  INDIGO: {
	    50: '#e8eaf6',
	    100: '#c5cae9',
	    200: '#9fa8da',
	    300: '#7986cb',
	    400: '#5c6bc0',
	    500: '#3f51b5',
	    600: '#3949ab',
	    700: '#303F9F',
	    800: '#283593',
	    900: '#1A237E',
	    A100: '#8c9eff',
	    A200: '#536dfe',
	    A400: '#3d5afe',
	    A700: '#304ffe'
	  },
	  BLUE: {
	    50: '#e3f2fd',
	    100: '#bbdefb',
	    200: '#90caf9',
	    300: '#64b5f6',
	    400: '#42a5f5',
	    500: '#2196f3',
	    600: '#1e88e5',
	    700: '#1976d2',
	    800: '#1565c0',
	    900: '#0d47a1',
	    A100: '#82b1ff',
	    A200: '#448aff',
	    A400: '#2979ff',
	    A700: '#2962ff'
	  },
	  LIGHTBLUE: {
	    50: '#e1f5fe',
	    100: '#b3e5fc',
	    200: '#81d4fa',
	    300: '#4fc3f7',
	    400: '#29b6f6',
	    500: '#03a9f4',
	    600: '#039be5',
	    700: '#0288d1',
	    800: '#0277bd',
	    900: '#01579b',
	    A100: '#80d8ff',
	    A200: '#40c4ff',
	    A400: '#00b0ff',
	    A700: '#0091ea'
	  },
	  CYAN: {
	    50: '#e0f7fa',
	    100: '#b2ebf2',
	    200: '#80deea',
	    300: '#4dd0e1',
	    400: '#26c6da',
	    500: '#00bcd4',
	    600: '#00acc1',
	    700: '#0097a7',
	    800: '#00838f',
	    900: '#006064',
	    A100: '#84ffff',
	    A200: '#18ffff',
	    A400: '#00e5ff',
	    A700: '#00b8d4'
	  },
	  TEAL: {
	    50: '#e0f2f1',
	    100: '#b2dfdb',
	    200: '#80cbc4',
	    300: '#4db6ac',
	    400: '#26a69a',
	    500: '#009688',
	    600: '#00897b',
	    700: '#00796b',
	    800: '#00695c',
	    900: '#004d40',
	    A100: '#a7ffeb',
	    A200: '#64ffda',
	    A400: '#1de9b6',
	    A700: '#00bfa5'
	  },
	  GREEN: {
	    50: '#e8f5e9',
	    100: '#c8e6c9',
	    200: '#a5d6a7',
	    300: '#81c784',
	    400: '#66bb6a',
	    500: '#4caf50',
	    600: '#43a047',
	    700: '#388e3c',
	    800: '#2e7d32',
	    900: '#1b5e20',
	    A100: '#b9f6ca',
	    A200: '#69f0ae',
	    A400: '#00e676',
	    A700: '#00c853'
	  },
	  LIGHTGREEN: {
	    50: '#f1f8e9',
	    100: '#dcedc8',
	    200: '#c5e1a5',
	    300: '#aed581',
	    400: '#9ccc65',
	    500: '#8bc34a',
	    600: '#7cb342',
	    700: '#689f38',
	    800: '#558b2f',
	    900: '#33691e',
	    A100: '#ccff90',
	    A200: '#b2ff59',
	    A400: '#76ff03',
	    A700: '#64dd17'
	  },
	  LIME: {
	    50: '#f9fbe7',
	    100: '#f0f4c3',
	    200: '#e6ee9c',
	    300: '#dce775',
	    400: '#d4e157',
	    500: '#cddc39',
	    600: '#c0ca33',
	    700: '#afb42b',
	    800: '#9e9d24',
	    900: '#827717',
	    A100: '#f4ff81',
	    A200: '#eeff41',
	    A400: '#c6ff00',
	    A700: '#aeea00'
	  },
	  YELLOW: {
	    50: '#fffde7',
	    100: '#fff9c4',
	    200: '#fff59d',
	    300: '#fff176',
	    400: '#ffee58',
	    500: '#ffeb3b',
	    600: '#fdd835',
	    700: '#fbc02d',
	    800: '#f9a825',
	    900: '#f57f17',
	    A100: '#ffff8d',
	    A200: '#ffff00',
	    A400: '#ffea00',
	    A700: '#ffd600'
	  },
	  AMBER: {
	    50: '#fff8e1',
	    100: '#ffecb3',
	    200: '#ffe082',
	    300: '#ffd54f',
	    400: '#ffca28',
	    500: '#ffc107',
	    600: '#ffb300',
	    700: '#ffa000',
	    800: '#ff8f00',
	    900: '#ff6f00',
	    A100: '#ffe57f',
	    A200: '#ffd740',
	    A400: '#ffc400',
	    A700: '#ffab00'
	  },
	  ORANGE: {
	    50: '#fff3e0',
	    100: '#ffe0b2',
	    200: '#ffcc80',
	    300: '#ffb74d',
	    400: '#ffa726',
	    500: '#ff9800',
	    600: '#fb8c00',
	    700: '#f57c00',
	    800: '#ef6c00',
	    900: '#e65100',
	    A100: '#ffd180',
	    A200: '#ffab40',
	    A400: '#ff9100',
	    A700: '#ff6d00'
	  },
	  DEEPORANGE: {
	    50: '#fbe9e7',
	    100: '#ffccbc',
	    200: '#ffab91',
	    300: '#ff8a65',
	    400: '#ff7043',
	    500: '#ff5722',
	    600: '#f4511e',
	    700: '#e64a19',
	    800: '#d84315',
	    900: '#bf360c',
	    A100: '#ff9e80',
	    A200: '#ff6e40',
	    A400: '#ff3d00',
	    A700: '#dd2c00'
	  },
	  BROWN: {
	    50: '#efebe9',
	    100: '#d7ccc8',
	    200: '#bcaaa4',
	    300: '#a1887f',
	    400: '#8d6e63',
	    500: '#795548',
	    600: '#6d4c41',
	    700: '#5d4037',
	    800: '#4e342e',
	    900: '#3e2723',
	    A100: '#ece2df',
	    A200: '#cfb7af',
	    A400: '#8c6253',
	    A700: '#533a31'
	  },
	  BLUEGREY: {
	    50: '#eceff1',
	    100: '#cfd8dc',
	    200: '#b0bec5',
	    300: '#90a4ae',
	    400: '#78909c',
	    500: '#607d8b',
	    600: '#546e7a',
	    700: '#455a64',
	    800: '#37474f',
	    900: '#263238',
	    A100: '#f9fafb',
	    A200: '#ccd7dc',
	    A400: '#6e8d9b',
	    A700: '#475c67'
	  },
	  GREY: {
	    50: '#fafafa',
	    100: '#f5f5f5',
	    200: '#eeeeee',
	    300: '#e0e0e0',
	    400: '#bdbdbd',
	    500: '#9e9e9e',
	    600: '#757575',
	    700: '#616161',
	    800: '#424242',
	    900: '#212121',
	    A100: '#ffffff',
	    A200: '#fcfcfc',
	    A400: '#adadad',
	    A700: '#7f7f7f'
	  }
	};
	
	var WHITE = Color.WHITE,
	    BLACK = Color.BLACK,
	    TRANSPARENT = Color.TRANSPARENT,
	    DELTAGREY = Color.DELTAGREY,
	    DELTABLUE = Color.DELTABLUE,
	    DELTAORANGE = Color.DELTAORANGE,
	    DELTAGREEN = Color.DELTAGREEN,
	    RED = Color.RED,
	    PINK = Color.PINK,
	    PURPLE = Color.PURPLE,
	    DEEPPRUPLE = Color.DEEPPRUPLE,
	    INDIGO = Color.INDIGO,
	    BLUE = Color.BLUE,
	    LIGHTBLUE = Color.LIGHTBLUE,
	    CYAN = Color.CYAN,
	    TEAL = Color.TEAL,
	    GREEN = Color.GREEN,
	    LIGHTGREEN = Color.LIGHTGREEN,
	    LIME = Color.LIME,
	    YELLOW = Color.YELLOW,
	    AMBER = Color.AMBER,
	    ORANGE = Color.ORANGE,
	    DEEPORANGE = Color.DEEPORANGE,
	    BROWN = Color.BROWN,
	    BLUEGREY = Color.BLUEGREY,
	    GREY = Color.GREY;
	
	
	var primary = 500;
	var primaryColor = {
	  White: WHITE,
	  Black: BLACK,
	  Transparent: TRANSPARENT,
	  DeltaGrey: DELTAGREY[primary],
	  DeltaBlue: DELTABLUE[primary],
	  DeltaOrange: DELTAORANGE[primary],
	  DeltaGreen: DELTAGREEN[primary],
	  Red: RED[primary],
	  Pink: PINK[primary],
	  Purple: PURPLE[primary],
	  DeepPruple: DEEPPRUPLE[primary],
	  Indigo: INDIGO[primary],
	  Blue: BLUE[primary],
	  LightBlue: LIGHTBLUE[primary],
	  Cyan: CYAN[primary],
	  Teal: TEAL[primary],
	  Green: GREEN[primary],
	  LightGreen: LIGHTGREEN[primary],
	  Lime: LIME[primary],
	  Yellow: YELLOW[primary],
	  Amber: AMBER[primary],
	  Orange: ORANGE[primary],
	  DeepOrange: DEEPORANGE[primary],
	  Brown: BROWN[primary],
	  BlueGrey: BLUEGREY[primary],
	  Grey: GREY[primary]
	};
	
	var White = primaryColor.White,
	    Black = primaryColor.Black,
	    Transparent = primaryColor.Transparent,
	    DeltaGrey = primaryColor.DeltaGrey,
	    DeltaBlue = primaryColor.DeltaBlue,
	    DeltaOrange = primaryColor.DeltaOrange,
	    DeltaGreen = primaryColor.DeltaGreen,
	    Red = primaryColor.Red,
	    Pink = primaryColor.Pink,
	    Purple = primaryColor.Purple,
	    DeepPruple = primaryColor.DeepPruple,
	    Indigo = primaryColor.Indigo,
	    Blue = primaryColor.Blue,
	    LightBlue = primaryColor.LightBlue,
	    Cyan = primaryColor.Cyan,
	    Teal = primaryColor.Teal,
	    Green = primaryColor.Green,
	    LightGreen = primaryColor.LightGreen,
	    Lime = primaryColor.Lime,
	    Yellow = primaryColor.Yellow,
	    Amber = primaryColor.Amber,
	    Orange = primaryColor.Orange,
	    DeepOrange = primaryColor.DeepOrange,
	    Brown = primaryColor.Brown,
	    BlueGrey = primaryColor.BlueGrey,
	    Grey = primaryColor.Grey;
	exports.default = Object.assign({}, Color, primaryColor);
	exports.WHITE = WHITE;
	exports.BLACK = BLACK;
	exports.TRANSPARENT = TRANSPARENT;
	exports.DELTAGREY = DELTAGREY;
	exports.DELTABLUE = DELTABLUE;
	exports.DELTAORANGE = DELTAORANGE;
	exports.DELTAGREEN = DELTAGREEN;
	exports.RED = RED;
	exports.PINK = PINK;
	exports.PURPLE = PURPLE;
	exports.DEEPPRUPLE = DEEPPRUPLE;
	exports.INDIGO = INDIGO;
	exports.BLUE = BLUE;
	exports.LIGHTBLUE = LIGHTBLUE;
	exports.CYAN = CYAN;
	exports.TEAL = TEAL;
	exports.GREEN = GREEN;
	exports.LIGHTGREEN = LIGHTGREEN;
	exports.LIME = LIME;
	exports.YELLOW = YELLOW;
	exports.AMBER = AMBER;
	exports.ORANGE = ORANGE;
	exports.DEEPORANGE = DEEPORANGE;
	exports.BROWN = BROWN;
	exports.BLUEGREY = BLUEGREY;
	exports.GREY = GREY;
	exports.White = White;
	exports.Black = Black;
	exports.Transparent = Transparent;
	exports.DeltaGrey = DeltaGrey;
	exports.DeltaBlue = DeltaBlue;
	exports.DeltaOrange = DeltaOrange;
	exports.DeltaGreen = DeltaGreen;
	exports.Red = Red;
	exports.Pink = Pink;
	exports.Purple = Purple;
	exports.DeepPruple = DeepPruple;
	exports.Indigo = Indigo;
	exports.Blue = Blue;
	exports.LightBlue = LightBlue;
	exports.Cyan = Cyan;
	exports.Teal = Teal;
	exports.Green = Green;
	exports.LightGreen = LightGreen;
	exports.Lime = Lime;
	exports.Yellow = Yellow;
	exports.Amber = Amber;
	exports.Orange = Orange;
	exports.DeepOrange = DeepOrange;
	exports.Brown = Brown;
	exports.BlueGrey = BlueGrey;
	exports.Grey = Grey;

/***/ }),
/* 42 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.GroupLoad = undefined;
	
	var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
	
	var _react = __webpack_require__(6);
	
	var _react2 = _interopRequireDefault(_react);
	
	var _reactRedux = __webpack_require__(10);
	
	var _utils = __webpack_require__(3);
	
	var _ActionNames = __webpack_require__(2);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }
	
	function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }
	
	var GroupLoadView = function (_React$Component) {
	  _inherits(GroupLoadView, _React$Component);
	
	  function GroupLoadView() {
	    _classCallCheck(this, GroupLoadView);
	
	    return _possibleConstructorReturn(this, (GroupLoadView.__proto__ || Object.getPrototypeOf(GroupLoadView)).apply(this, arguments));
	  }
	
	  _createClass(GroupLoadView, [{
	    key: 'componentDidMount',
	    value: function componentDidMount() {
	      this.props.loadGroup();
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      return null;
	    }
	  }]);
	
	  return GroupLoadView;
	}(_react2.default.Component);
	
	var mapStateToProps = function mapStateToProps(state, ownProps) {
	  return {};
	};
	
	var mapDispatchToProps = function mapDispatchToProps(dispatch, ownProps) {
	  return {
	    loadGroup: function loadGroup() {
	      console.log("load group", ownProps);
	      var payload = ownProps.Data; //{entityName: ownProps.name, entityId: ownProps.id};
	      var meta = { serviceName: ownProps.service };
	      dispatch((0, _utils.createAction)(_ActionNames.ActionNames.GROUP_LOAD, payload, meta));
	    }
	  };
	};
	
	var GroupLoad = (0, _reactRedux.connect)(mapStateToProps, mapDispatchToProps)(GroupLoadView);
	
	exports.GroupLoad = GroupLoad;

/***/ }),
/* 43 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.LoginComponent = undefined;
	
	var _react = __webpack_require__(6);
	
	var _react2 = _interopRequireDefault(_react);
	
	var _LoginWeb = __webpack_require__(44);
	
	var _LoginWeb2 = _interopRequireDefault(_LoginWeb);
	
	var _md = __webpack_require__(65);
	
	var _md2 = _interopRequireDefault(_md);
	
	var _reactRedux = __webpack_require__(10);
	
	var _ActionNames = __webpack_require__(2);
	
	var _utils = __webpack_require__(3);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	var mapStateToProps = function mapStateToProps(state, ownProps) {
	  return {
	    realm: ownProps.realm,
	    renderLogin: ownProps.renderLogin,
	    signup: ownProps.signup
	  };
	};
	
	var mapDispatchToProps = function mapDispatchToProps(dispatch, ownProps) {
	  console.log("map dispatch of login compoent");
	  var realm = "";
	  if (ownProps.realm) {
	    realm = ownProps.realm;
	  }
	  return {
	    handleLogin: function handleLogin(email, password) {
	      var loginPayload = { "Username": email, "Password": (0, _md2.default)(password), "Realm": realm };
	      var loginMeta = { serviceName: ownProps.loginService };
	      dispatch((0, _utils.createAction)(_ActionNames.ActionNames.LOGIN, loginPayload, loginMeta));
	    },
	    handleOauthLogin: function handleOauthLogin(data) {
	      dispatch((0, _utils.createAction)(_ActionNames.ActionNames.LOGIN_SUCCESS, { userId: data.id, token: data.token, permissions: data.permissions }));
	    }
	  };
	};
	
	var LoginComponent = (0, _reactRedux.connect)(mapStateToProps, mapDispatchToProps)(_LoginWeb2.default);
	
	// Uncomment properties you need
	LoginComponent.propTypes = {
	  loginService: _react2.default.PropTypes.string.isRequired,
	  successpage: _react2.default.PropTypes.string,
	  realm: _react2.default.PropTypes.string,
	  signup: _react2.default.PropTypes.string
	};
	// LoginComponent.defaultProps = {};
	
	exports.LoginComponent = LoginComponent;

/***/ }),
/* 44 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	
	var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
	
	var _react = __webpack_require__(6);
	
	var _react2 = _interopRequireDefault(_react);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }
	
	function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }
	
	var LoginWeb = function (_React$Component) {
	  _inherits(LoginWeb, _React$Component);
	
	  function LoginWeb(props) {
	    _classCallCheck(this, LoginWeb);
	
	    var _this = _possibleConstructorReturn(this, (LoginWeb.__proto__ || Object.getPrototypeOf(LoginWeb)).call(this, props));
	
	    console.log("costructor of login web");
	    _this.state = {
	      "email": "",
	      "password": ""
	    };
	    _this.handleLogin = _this.handleLogin.bind(_this);
	    _this.handleChange = _this.handleChange.bind(_this);
	    /*
	        let getLocation = function(href) {
	            var l = document.createElement("a");
	            l.href = href;
	            return l;
	        };*/
	    var loginSite = "";
	    var realm = "";
	    if (props.realm) {
	      realm = "?Realm=" + props.realm;
	    }
	    /*this.oauthLogin = function(location) {
	      let instance = window.open(location+realm, '_blank','height=500,width=400,toolbar=no,resizable=yes,menubar=no,location=0')
	      var location = getLocation(location);
	      loginSite = location.protocol+'//'+location.hostname+(location.port ? ':'+location.port: '');
	    }
	     window.addEventListener("message", function(ev) {
	      if(ev.origin === loginSite && ev.data.message=="LoginSuccess") {
	        props.handleOauthLogin(ev.data)
	      }
	    });*/
	    return _this;
	  }
	
	  _createClass(LoginWeb, [{
	    key: 'handleChange',
	    value: function handleChange(e) {
	      var nextState = {};
	      nextState[e.target.name] = e.target.value;
	      this.setState(nextState);
	    }
	  }, {
	    key: 'handleLogin',
	    value: function handleLogin() {
	      console.log("this has been reached");
	      this.props.handleLogin(this.state.email, this.state.password);
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      console.log("rendering login web");
	      return this.props.renderLogin(this.state, this.handleChange, this.handleLogin, this.oauthLogin);
	    }
	  }]);
	
	  return LoginWeb;
	}(_react2.default.Component);
	
	LoginWeb.displayName = 'LoginComponent';
	
	// Uncomment properties you need
	LoginWeb.propTypes = {
	  handleOauthLogin: _react2.default.PropTypes.func.isRequired,
	  handleLogin: _react2.default.PropTypes.func.isRequired
	};
	
	exports.default = LoginWeb;

/***/ }),
/* 45 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.View = undefined;
	
	var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
	
	var _react = __webpack_require__(6);
	
	var _react2 = _interopRequireDefault(_react);
	
	var _reactRedux = __webpack_require__(10);
	
	var _ActionNames = __webpack_require__(2);
	
	var _utils = __webpack_require__(3);
	
	var _DataSource = __webpack_require__(5);
	
	var _ViewData = __webpack_require__(17);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }
	
	function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }
	
	var View = function (_React$Component) {
	  _inherits(View, _React$Component);
	
	  function View(props) {
	    _classCallCheck(this, View);
	
	    var _this = _possibleConstructorReturn(this, (View.__proto__ || Object.getPrototypeOf(View)).call(this, props));
	
	    _this.getItemGroup = _this.getItemGroup.bind(_this);
	    _this.getView = _this.getView.bind(_this);
	    _this.renderView = _this.renderView.bind(_this);
	    _this.getItem = _this.getItem.bind(_this);
	    _this.getHeader = _this.getHeader.bind(_this);
	    _this.getPagination = _this.getPagination.bind(_this);
	    _this.getFilter = _this.getFilter.bind(_this);
	    _this.addMethod = _this.addMethod.bind(_this);
	    _this.numItems = 0;
	    return _this;
	  }
	
	  _createClass(View, [{
	    key: 'addMethod',
	    value: function addMethod(name, method) {
	      return this.viewdata.addMethod(name, method);
	    }
	  }, {
	    key: 'getView',
	    value: function getView(header, groups, pagination, filter) {
	      if (this.props.getView) {
	        return this.props.getView(this, header, groups, pagination, filter);
	      }
	      return null;
	    }
	  }, {
	    key: 'getFilter',
	    value: function getFilter(filterTitle, filterForm, filterGo) {
	      if (this.props.getFilter) {
	        return this.props.getFilter(this, filterTitle, filterForm, filterGo);
	      }
	      return null;
	    }
	  }, {
	    key: 'getItemGroup',
	    value: function getItemGroup(x) {
	      if (this.props.getItemGroup) {
	        return this.props.getItemGroup(this, x);
	      }
	      return null;
	    }
	  }, {
	    key: 'getItem',
	    value: function getItem(x, i) {
	      if (this.props.getItem) {
	        return this.props.getItem(this, x, i);
	      }
	      return null;
	    }
	  }, {
	    key: 'getHeader',
	    value: function getHeader() {
	      if (this.props.getHeader) {
	        return this.props.getHeader(this);
	      }
	      return null;
	    }
	  }, {
	    key: 'getPagination',
	    value: function getPagination() {
	      if (this.props.paginate && this.props.getPagination) {
	        var pages = this.props.totalPages;
	        var page = this.props.currentPage;
	        return this.props.getPagination(this, pages, page);
	      }
	      return null;
	    }
	  }, {
	    key: 'renderView',
	    value: function renderView(viewdata, items, currentPage, totalPages) {
	      this.viewdata = viewdata;
	      var groups = [];
	      var groupsize = 1;
	      var group = [];
	
	      if (items) {
	        var keys = Object.keys(items);
	        this.numItems = keys.length;
	        for (var i in keys) {
	          var x = items[keys[i]];
	          if (x) {
	            var item = this.getItem(x, keys[i]);
	            group.push(item);
	            if (i % groupsize == 0) {
	              var itemGrp = this.getItemGroup(group);
	              groups.push(itemGrp);
	              group = [];
	            }
	          }
	        }
	      } else {
	        if (this.props.loader) {
	          groups.push(this.props.loader);
	        }
	      }
	      var header = this.getHeader();
	      var filterCtrl = this.getFilter(this.props.filterTitle, this.props.filterForm, this.props.filterGo, this.filter);
	      var pagination = this.getPagination();
	      return this.getView(header, groups, pagination, filterCtrl);
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      return _react2.default.createElement(_ViewData.ViewData, {
	        getView: this.renderView,
	        key: this.props.key,
	        reducer: this.props.reducer,
	        paginate: this.props.paginate,
	        pageSize: this.props.pageSize,
	        viewService: this.props.viewService,
	        urlParams: this.props.urlParams,
	        postArgs: this.props.postArgs,
	        defaultFilter: this.props.defaultFilter,
	        currentPage: this.props.currentPage,
	        style: this.props.style,
	        className: this.props.className,
	        incrementalLoad: this.props.incrementalLoad,
	        globalReducer: this.props.globalReducer });
	    }
	  }]);
	
	  return View;
	}(_react2.default.Component);
	
	exports.View = View;

/***/ }),
/* 46 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.Account = undefined;
	
	var _ActionNames = __webpack_require__(2);
	
	var _Globals = __webpack_require__(4);
	
	var initialSecState = {
	  status: "NotLogged",
	  token: "",
	  userId: "",
	  permissions: []
	};
	
	var Account = function Account(state, action) {
	  if (action.type) {
	    switch (action.type) {
	      case _ActionNames.ActionNames.LOGGING_IN:
	        return Object.assign({}, state, {
	          status: "LoggingIn"
	        });
	
	      case _ActionNames.ActionNames.LOGIN_SUCCESS:
	        if (state.authToken === action.payload.token) {
	          return state;
	        }
	        _Globals.Storage.auth = action.payload.token;
	        _Globals.Storage.permissions = action.payload.permissions;
	        _Globals.Storage.user = action.payload.userId;
	        return Object.assign({}, state, {
	          status: "LoggedIn",
	          authToken: action.payload.token,
	          userId: action.payload.userId,
	          permissions: action.payload.permissions
	        });
	
	      case _ActionNames.ActionNames.LOGIN_FAILURE:
	        _Globals.Storage.auth = "";
	        _Globals.Storage.permissions = [];
	        _Globals.Storage.user = "";
	        return initialSecState;
	
	      case _ActionNames.ActionNames.LOGOUT:
	        _Globals.Storage.auth = "";
	        _Globals.Storage.permissions = [];
	        _Globals.Storage.user = "";
	        return initialSecState;
	
	      default:
	        if (!state) {
	          if (_Globals.Storage.auth != null && _Globals.Storage.auth != "") {
	            return {
	              status: "LoggedIn",
	              authToken: _Globals.Storage.auth,
	              userId: _Globals.Storage.user,
	              permissions: _Globals.Storage.permissions
	            };
	          }
	          return initialSecState;
	        }
	        return state;
	    }
	  }
	};
	
	exports.Account = Account;

/***/ }),
/* 47 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.Reducers = undefined;
	
	var _redux = __webpack_require__(22);
	
	var _Security = __webpack_require__(46);
	
	var _View = __webpack_require__(19);
	
	var _Entity = __webpack_require__(18);
	
	/* Combine all available reducers to a single root reducer.
	 *
	 * CAUTION: When using the generators, this file is modified in some places.
	 *          This is done via AST traversal - Some of your formatting may be lost
	 *          in the process - no functionality should be broken though.
	 *          This modifications only run once when the generator is invoked - if
	 *          you edit them, they are not updated again.
	 */
	var Reducers = exports.Reducers = {
	  SecurityReducers: _Security.Account,
	  EntityReducer: _Entity.EntityReducer,
	  ViewReducer: _View.ViewReducer
	};

/***/ }),
/* 48 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.entitySaga = undefined;
	
	var _reduxSaga = __webpack_require__(7);
	
	var _effects = __webpack_require__(8);
	
	var _ActionNames = __webpack_require__(2);
	
	var _DataSource = __webpack_require__(5);
	
	var _utils = __webpack_require__(3);
	
	var _Globals = __webpack_require__(4);
	
	var _marked = [getEntityData, deleteEntityData, saveEntityData, putEntityData, updateEntityData, entitySaga].map(regeneratorRuntime.mark);
	
	function getEntityData(action) {
	  var resp;
	  return regeneratorRuntime.wrap(function getEntityData$(_context) {
	    while (1) {
	      switch (_context.prev = _context.next) {
	        case 0:
	          _context.prev = 0;
	          _context.next = 3;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GETTING, action.payload, { reducer: action.meta.reducer }));
	
	        case 3:
	          _context.next = 5;
	          return (0, _effects.call)(_DataSource.EntityData.GetEntity, action.payload.entityName, action.payload.entityId, action.payload.headers, action.payload.svc);
	
	        case 5:
	          resp = _context.sent;
	
	          resp.data.isOwner = resp.data.CreatedBy === _Globals.Storage.user;
	          _context.next = 9;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GET_SUCCESS, resp, action.meta));
	
	        case 9:
	          if (action.meta.successCallback) {
	            action.meta.successCallback({ resp: resp, payload: action.payload });
	          }
	          _context.next = 17;
	          break;
	
	        case 12:
	          _context.prev = 12;
	          _context.t0 = _context['catch'](0);
	          _context.next = 16;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GET_FAILED, _context.t0, action.meta));
	
	        case 16:
	          if (action.meta.failureCallback) {
	            action.meta.failureCallback(_context.t0);
	          } else {
	            if (window.handleError) {
	              window.handleError(_context.t0);
	            }
	          }
	
	        case 17:
	        case 'end':
	          return _context.stop();
	      }
	    }
	  }, _marked[0], this, [[0, 12]]);
	}
	
	function deleteEntityData(action) {
	  var resp;
	  return regeneratorRuntime.wrap(function deleteEntityData$(_context2) {
	    while (1) {
	      switch (_context2.prev = _context2.next) {
	        case 0:
	          _context2.prev = 0;
	          _context2.next = 3;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_DELETING, action.payload, { reducer: action.meta.reducer }));
	
	        case 3:
	          _context2.next = 5;
	          return (0, _effects.call)(_DataSource.EntityData.DeleteEntity, action.payload.entityName, action.payload.entityId, action.payload.headers, action.payload.svc);
	
	        case 5:
	          resp = _context2.sent;
	          _context2.next = 8;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_DELETE_SUCCESS, resp, action.meta));
	
	        case 8:
	          if (action.meta.successCallback) {
	            action.meta.successCallback({ resp: resp, payload: action.payload });
	          }
	          _context2.next = 16;
	          break;
	
	        case 11:
	          _context2.prev = 11;
	          _context2.t0 = _context2['catch'](0);
	          _context2.next = 15;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_DELETE_FAILURE, _context2.t0, action.meta));
	
	        case 15:
	          if (action.meta.failureCallback) {
	            action.meta.failureCallback(_context2.t0);
	          } else {
	            if (window.handleError) {
	              window.handleError(_context2.t0);
	            }
	          }
	
	        case 16:
	        case 'end':
	          return _context2.stop();
	      }
	    }
	  }, _marked[1], this, [[0, 11]]);
	}
	
	function saveEntityData(action) {
	  var resp;
	  return regeneratorRuntime.wrap(function saveEntityData$(_context3) {
	    while (1) {
	      switch (_context3.prev = _context3.next) {
	        case 0:
	          _context3.prev = 0;
	          _context3.next = 3;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_SAVING, action.payload, { reducer: action.meta.reducer }));
	
	        case 3:
	          _context3.next = 5;
	          return (0, _effects.call)(_DataSource.EntityData.SaveEntity, action.payload.entityName, action.payload.data, action.payload.headers, action.payload.svc);
	
	        case 5:
	          resp = _context3.sent;
	          _context3.next = 8;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_SAVE_SUCCESS, resp, { reducer: action.meta.reducer }));
	
	        case 8:
	          if (action.meta.successCallback) {
	            action.meta.successCallback({ resp: resp, payload: action.payload });
	          }
	          _context3.next = 16;
	          break;
	
	        case 11:
	          _context3.prev = 11;
	          _context3.t0 = _context3['catch'](0);
	          _context3.next = 15;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_SAVE_FAILURE, _context3.t0, action.meta));
	
	        case 15:
	          if (action.meta.failureCallback) {
	            action.meta.failureCallback(_context3.t0);
	          } else {
	            if (window.handleError) {
	              window.handleError(_context3.t0);
	            }
	          }
	
	        case 16:
	        case 'end':
	          return _context3.stop();
	      }
	    }
	  }, _marked[2], this, [[0, 11]]);
	}
	
	function putEntityData(action) {
	  var resp;
	  return regeneratorRuntime.wrap(function putEntityData$(_context4) {
	    while (1) {
	      switch (_context4.prev = _context4.next) {
	        case 0:
	          _context4.prev = 0;
	          _context4.next = 3;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_PUTTING, action.payload, { reducer: action.meta.reducer }));
	
	        case 3:
	          _context4.next = 5;
	          return (0, _effects.call)(_DataSource.EntityData.PutEntity, action.payload.entityName, action.payload.entityId, action.payload.data, action.payload.headers, action.payload.svc);
	
	        case 5:
	          resp = _context4.sent;
	          _context4.next = 8;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_PUT_SUCCESS, resp, action.meta));
	
	        case 8:
	          if (!action.meta.reload) {
	            _context4.next = 11;
	            break;
	          }
	
	          _context4.next = 11;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GET, action.payload, action.meta));
	
	        case 11:
	          if (action.meta.successCallback) {
	            action.meta.successCallback({ resp: resp, payload: action.payload });
	          }
	          _context4.next = 19;
	          break;
	
	        case 14:
	          _context4.prev = 14;
	          _context4.t0 = _context4['catch'](0);
	          _context4.next = 18;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_PUT_FAILURE, _context4.t0, action.meta));
	
	        case 18:
	          if (action.meta.failureCallback) {
	            action.meta.failureCallback(_context4.t0);
	          } else {
	            if (window.handleError) {
	              window.handleError(_context4.t0);
	            }
	          }
	
	        case 19:
	        case 'end':
	          return _context4.stop();
	      }
	    }
	  }, _marked[3], this, [[0, 14]]);
	}
	
	function updateEntityData(action) {
	  var resp;
	  return regeneratorRuntime.wrap(function updateEntityData$(_context5) {
	    while (1) {
	      switch (_context5.prev = _context5.next) {
	        case 0:
	          _context5.prev = 0;
	          _context5.next = 3;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_UPDATING, action.payload, { reducer: action.meta.reducer }));
	
	        case 3:
	          _context5.next = 5;
	          return (0, _effects.call)(_DataSource.EntityData.UpdateEntity, action.payload.entityName, action.payload.entityId, action.payload.data, action.payload.headers, action.payload.svc);
	
	        case 5:
	          resp = _context5.sent;
	          _context5.next = 8;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_UPDATE_SUCCESS, resp, action.meta));
	
	        case 8:
	          if (!action.meta.reload) {
	            _context5.next = 11;
	            break;
	          }
	
	          _context5.next = 11;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GET, action.payload, action.meta));
	
	        case 11:
	          if (action.meta.successCallback) {
	            action.meta.successCallback({ resp: resp, payload: action.payload });
	          }
	          _context5.next = 19;
	          break;
	
	        case 14:
	          _context5.prev = 14;
	          _context5.t0 = _context5['catch'](0);
	          _context5.next = 18;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_UPDATE_FAILURE, _context5.t0, action.meta));
	
	        case 18:
	          if (action.meta.failureCallback) {
	            action.meta.failureCallback(_context5.t0);
	          } else {
	            if (window.handleError) {
	              window.handleError(_context5.t0);
	            }
	          }
	
	        case 19:
	        case 'end':
	          return _context5.stop();
	      }
	    }
	  }, _marked[4], this, [[0, 14]]);
	}
	
	function entitySaga() {
	  return regeneratorRuntime.wrap(function entitySaga$(_context6) {
	    while (1) {
	      switch (_context6.prev = _context6.next) {
	        case 0:
	          _context6.next = 2;
	          return [(0, _reduxSaga.takeEvery)(_ActionNames.ActionNames.ENTITY_GET, getEntityData), (0, _reduxSaga.takeEvery)(_ActionNames.ActionNames.ENTITY_SAVE, saveEntityData), (0, _reduxSaga.takeEvery)(_ActionNames.ActionNames.ENTITY_UPDATE, updateEntityData), (0, _reduxSaga.takeEvery)(_ActionNames.ActionNames.ENTITY_PUT, putEntityData), (0, _reduxSaga.takeEvery)(_ActionNames.ActionNames.ENTITY_DELETE, deleteEntityData)];
	
	        case 2:
	        case 'end':
	          return _context6.stop();
	      }
	    }
	  }, _marked[5], this);
	}
	
	exports.entitySaga = entitySaga;

/***/ }),
/* 49 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.groupLoadSaga = undefined;
	
	var _reduxSaga = __webpack_require__(7);
	
	var _effects = __webpack_require__(8);
	
	var _ActionNames = __webpack_require__(2);
	
	var _DataSource = __webpack_require__(5);
	
	var _utils = __webpack_require__(3);
	
	var _Globals = __webpack_require__(4);
	
	var _marked = [loadGroup, groupLoadSaga].map(regeneratorRuntime.mark);
	
	function loadGroup(action) {
	  var request, req, resp, actions, i, _actions, _i;
	
	  return regeneratorRuntime.wrap(function loadGroup$(_context) {
	    while (1) {
	      switch (_context.prev = _context.next) {
	        case 0:
	          _context.prev = 0;
	          request = {};
	
	          console.log("load group saga", action.payload);
	          Object.keys(action.payload).forEach(function (key) {
	            var service = action.payload[key];
	            console.log("load group ", service, "key", key);
	            if (service.type == "entity") {
	              request[key] = { Params: { id: service.entityId }, Body: {} };
	            }
	            if (service.type == "view") {
	              request[key] = { Params: service.queryParams, Body: service.postArgs };
	            }
	          });
	          /*Object.keys(action.payload).forEach(function* (service, key) {
	            console.log("load group saga", service, "key", key)
	            if(service.type == "entity") {
	              yield put(createAction(ActionNames.ENTITY_GETTING, service.payload, {reducer: service.meta.reducer}));
	            }
	            if(service.type == "view") {
	              yield put(createAction(ActionNames.VIEW_FETCHING, service.payload,{reducer: service.meta.reducer}));
	            }
	          });*/
	          console.log("created request", request);
	          req = _DataSource.RequestBuilder.DefaultRequest(null, request);
	          _context.next = 8;
	          return (0, _effects.call)(_DataSource.DataSource.ExecuteService, action.meta.serviceName, req);
	
	        case 8:
	          resp = _context.sent;
	
	          console.log("resp", resp);
	          actions = new Array();
	
	          Object.keys(action.payload).forEach(function (key) {
	            var service = action.payload[key];
	            var servResponse = resp.data[key];
	            console.log("service ", service, servResponse);
	            var itemResp = { data: servResponse.Data, statuscode: servResponse.Status, info: servResponse.Info };
	            if (service.type == "entity") {
	              if (servResponse.Status == 200) {
	                itemResp.data.isOwner = itemResp.data.CreatedBy === _Globals.Storage.user;
	                actions.push((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GET_SUCCESS, itemResp, { reducer: service.meta.reducer }));
	              } else {
	                actions.push((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GET_FAILED, itemResp, { reducer: service.meta.reducer }));
	              }
	            }
	            if (service.type == "view") {
	              if (servResponse.Status == 200) {
	                actions.push((0, _utils.createAction)(_ActionNames.ActionNames.VIEW_FETCH_SUCCESS, itemResp.data, { info: itemResp.info, incrementalLoad: action.meta.incrementalLoad, reducer: service.meta.reducer }));
	              } else {
	                actions.push((0, _utils.createAction)(_ActionNames.ActionNames.VIEW_FETCH_FAILED, itemResp.data, { reducer: service.reducer }));
	              }
	            }
	          });
	          if (action.meta.successCallback) {
	            action.meta.successCallback({ resp: resp, payload: action.payload });
	          }
	          console.log("actions i", actions);
	          i = 0;
	
	        case 15:
	          if (!(i < actions.length)) {
	            _context.next = 22;
	            break;
	          }
	
	          console.log("actions i", i, actions[i]);
	          _context.next = 19;
	          return (0, _effects.put)(actions[i]);
	
	        case 19:
	          i++;
	          _context.next = 15;
	          break;
	
	        case 22:
	          _context.next = 38;
	          break;
	
	        case 24:
	          _context.prev = 24;
	          _context.t0 = _context['catch'](0);
	
	          if (!action.meta.services) {
	            _context.next = 37;
	            break;
	          }
	
	          _actions = new Array();
	
	          Object.keys(action.meta.services).forEach(function (key) {
	            var service = action.meta.services[key];
	            if (service.type == "entity") {
	              _actions.push((0, _utils.createAction)(_ActionNames.ActionNames.ENTITY_GET_FAILED, _context.t0, { reducer: service.reducer }));
	            }
	            if (service.type == "view") {
	              _actions.push((0, _utils.createAction)(_ActionNames.ActionNames.VIEW_FETCH_FAILED, _context.t0, { reducer: service.reducer }));
	            }
	          });
	          _i = 0;
	
	        case 30:
	          if (!(_i < _actions.length)) {
	            _context.next = 37;
	            break;
	          }
	
	          console.log("actions i", _i, _actions[_i]);
	          _context.next = 34;
	          return (0, _effects.put)(_actions[_i]);
	
	        case 34:
	          _i++;
	          _context.next = 30;
	          break;
	
	        case 37:
	          if (action.meta.failureCallback) {
	            action.meta.failureCallback(_context.t0);
	          } else {
	            if (window.handleError) {
	              window.handleError(_context.t0);
	            }
	          }
	
	        case 38:
	        case 'end':
	          return _context.stop();
	      }
	    }
	  }, _marked[0], this, [[0, 24]]);
	}
	
	function groupLoadSaga() {
	  return regeneratorRuntime.wrap(function groupLoadSaga$(_context2) {
	    while (1) {
	      switch (_context2.prev = _context2.next) {
	        case 0:
	          _context2.next = 2;
	          return [(0, _reduxSaga.takeEvery)(_ActionNames.ActionNames.GROUP_LOAD, loadGroup)];
	
	        case 2:
	        case 'end':
	          return _context2.stop();
	      }
	    }
	  }, _marked[1], this);
	}
	
	exports.groupLoadSaga = groupLoadSaga;

/***/ }),
/* 50 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.loginSaga = undefined;
	
	var _reduxSaga = __webpack_require__(7);
	
	var _effects = __webpack_require__(8);
	
	var _ActionNames = __webpack_require__(2);
	
	var _DataSource = __webpack_require__(5);
	
	var _utils = __webpack_require__(3);
	
	var _Globals = __webpack_require__(4);
	
	var _marked = [login, loginSaga].map(regeneratorRuntime.mark);
	
	function login(action) {
	  var req, resp, authToken, token, userId, permissions, loginaction;
	  return regeneratorRuntime.wrap(function login$(_context) {
	    while (1) {
	      switch (_context.prev = _context.next) {
	        case 0:
	          _context.prev = 0;
	          _context.next = 3;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.LOGGING_IN));
	
	        case 3:
	          req = _DataSource.RequestBuilder.DefaultRequest(null, action.payload);
	          _context.next = 6;
	          return (0, _effects.call)(_DataSource.DataSource.ExecuteService, action.meta.serviceName, req);
	
	        case 6:
	          resp = _context.sent;
	          authToken = _Globals.Application.Security.AuthToken.toLowerCase();
	          token = resp.info[authToken];
	          userId = resp.data.Id;
	          permissions = resp.data.Permissions;
	          loginaction = (0, _utils.createAction)(_ActionNames.ActionNames.LOGIN_SUCCESS, { userId: userId, token: token, permissions: permissions });
	          _context.next = 14;
	          return (0, _effects.put)(loginaction);
	
	        case 14:
	          console.log("dispatched login action &&&&");
	          _context.next = 21;
	          break;
	
	        case 17:
	          _context.prev = 17;
	          _context.t0 = _context['catch'](0);
	          _context.next = 21;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.LOGIN_FAILURE, _context.t0));
	
	        case 21:
	        case 'end':
	          return _context.stop();
	      }
	    }
	  }, _marked[0], this, [[0, 17]]);
	}
	
	function loginSaga() {
	  return regeneratorRuntime.wrap(function loginSaga$(_context2) {
	    while (1) {
	      switch (_context2.prev = _context2.next) {
	        case 0:
	          return _context2.delegateYield((0, _reduxSaga.takeLatest)(_ActionNames.ActionNames.LOGIN, login), 't0', 1);
	
	        case 1:
	        case 'end':
	          return _context2.stop();
	      }
	    }
	  }, _marked[1], this);
	}
	
	exports.loginSaga = loginSaga;

/***/ }),
/* 51 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.viewSaga = undefined;
	
	var _reduxSaga = __webpack_require__(7);
	
	var _effects = __webpack_require__(8);
	
	var _ActionNames = __webpack_require__(2);
	
	var _DataSource = __webpack_require__(5);
	
	var _utils = __webpack_require__(3);
	
	var _marked = [fetchViewData, viewSaga].map(regeneratorRuntime.mark);
	
	function fetchViewData(action) {
	  var req, resp;
	  return regeneratorRuntime.wrap(function fetchViewData$(_context) {
	    while (1) {
	      switch (_context.prev = _context.next) {
	        case 0:
	          _context.prev = 0;
	          _context.next = 3;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.VIEW_FETCHING, action.payload, { reducer: action.meta.reducer }));
	
	        case 3:
	          req = _DataSource.RequestBuilder.DefaultRequest(action.payload.queryParams, action.payload.postArgs, action.payload.headers);
	          _context.next = 6;
	          return (0, _effects.call)(_DataSource.DataSource.ExecuteService, action.meta.serviceName, req);
	
	        case 6:
	          resp = _context.sent;
	          _context.next = 9;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.VIEW_FETCH_SUCCESS, resp.data, { info: resp.info, incrementalLoad: action.meta.incrementalLoad, reducer: action.meta.reducer }));
	
	        case 9:
	          _context.next = 16;
	          break;
	
	        case 11:
	          _context.prev = 11;
	          _context.t0 = _context['catch'](0);
	          _context.next = 15;
	          return (0, _effects.put)((0, _utils.createAction)(_ActionNames.ActionNames.VIEW_FETCH_FAILED, _context.t0, action.meta));
	
	        case 15:
	          if (action.meta.failureCallback) {
	            action.meta.failureCallback(_context.t0);
	          } else {
	            if (window.handleError) {
	              window.handleError(_context.t0);
	            }
	          }
	
	        case 16:
	        case 'end':
	          return _context.stop();
	      }
	    }
	  }, _marked[0], this, [[0, 11]]);
	}
	
	function viewSaga() {
	  return regeneratorRuntime.wrap(function viewSaga$(_context2) {
	    while (1) {
	      switch (_context2.prev = _context2.next) {
	        case 0:
	          return _context2.delegateYield((0, _reduxSaga.takeEvery)(_ActionNames.ActionNames.VIEW_FETCH, fetchViewData), 't0', 1);
	
	        case 1:
	        case 'end':
	          return _context2.stop();
	      }
	    }
	  }, _marked[1], this);
	}
	
	exports.viewSaga = viewSaga;

/***/ }),
/* 52 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports.Sagas = undefined;
	exports.runSagas = runSagas;
	
	var _View = __webpack_require__(51);
	
	var _Security = __webpack_require__(50);
	
	var _Entity = __webpack_require__(48);
	
	var _GroupLoad = __webpack_require__(49);
	
	var Sagas = exports.Sagas = {
	  ViewSaga: _View.viewSaga,
	  LoginSaga: _Security.loginSaga,
	  GroupLoadSaga: _GroupLoad.groupLoadSaga,
	  EntitySaga: _Entity.entitySaga
	};
	
	function runSagas(sagaMiddleware, sagas) {
	  sagas.map(function (x, i) {
	    sagaMiddleware.run(x);
	  });
	}

/***/ }),
/* 53 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	    value: true
	});
	exports.EntityDataService = undefined;
	
	var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
	
	var _Globals = __webpack_require__(4);
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	var EntityDataService = exports.EntityDataService = function () {
	    function EntityDataService(DataSource, RequestBuilder) {
	        _classCallCheck(this, EntityDataService);
	
	        this.DataSource = DataSource;
	        this.RequestBuilder = RequestBuilder;
	        this.GetEntity = this.GetEntity.bind(this);
	        this.SaveEntity = this.SaveEntity.bind(this);
	        this.DeleteEntity = this.DeleteEntity.bind(this);
	        this.PutEntity = this.PutEntity.bind(this);
	        this.UpdateEntity = this.UpdateEntity.bind(this);
	    }
	
	    _createClass(EntityDataService, [{
	        key: 'GetEntity',
	        value: function GetEntity(entityName, id, headers, svc) {
	            if (svc) {
	                var req = this.RequestBuilder.URLParamsRequest({ ":id": id }, null, headers);
	                return this.DataSource.ExecuteService(svc, req);
	            } else {
	                var service = {};
	                service.method = "GET";
	                service.url = _Globals.Application.EntityPrefix + entityName.toLowerCase() + "/" + id;
	                var req = this.RequestBuilder.DefaultRequest(null, null, headers);
	                return this.DataSource.ExecuteServiceObject(service, req);
	            }
	        }
	    }, {
	        key: 'SaveEntity',
	        value: function SaveEntity(entityName, data, headers, svc) {
	            var req = this.RequestBuilder.DefaultRequest(null, data, headers);
	            if (svc) {
	                return this.DataSource.ExecuteService(svc, req);
	            } else {
	                var service = {};
	                service.method = "POST";
	                service.url = _Globals.Application.EntityPrefix + entityName.toLowerCase();
	                return this.DataSource.ExecuteServiceObject(service, req);
	            }
	        }
	    }, {
	        key: 'DeleteEntity',
	        value: function DeleteEntity(entityName, id, headers, svc) {
	            if (svc) {
	                var req = this.RequestBuilder.URLParamsRequest({ ":id": id }, null, headers);
	                return this.DataSource.ExecuteService(svc, req);
	            } else {
	                var service = {};
	                service.method = "DELETE";
	                service.url = _Globals.Application.EntityPrefix + entityName.toLowerCase() + "/" + id;
	                var req = this.RequestBuilder.DefaultRequest(null, null, headers);
	                return this.DataSource.ExecuteServiceObject(service, req);
	            }
	        }
	    }, {
	        key: 'PutEntity',
	        value: function PutEntity(entityName, id, data, headers, svc) {
	            if (svc) {
	                var req = this.RequestBuilder.URLParamsRequest({ ":id": id }, null, headers);
	                return this.DataSource.ExecuteService(svc, req);
	            } else {
	                var service = {};
	                service.method = "PUT";
	                service.url = _Globals.Application.EntityPrefix + entityName.toLowerCase() + "/" + id;
	                var req = this.RequestBuilder.DefaultRequest(null, data, headers);
	                return this.DataSource.ExecuteServiceObject(service, req);
	            }
	        }
	    }, {
	        key: 'UpdateEntity',
	        value: function UpdateEntity(entityName, id, fieldmap, headers, svc) {
	            if (svc) {
	                var req = this.RequestBuilder.URLParamsRequest({ ":id": id }, null, headers);
	                return this.DataSource.ExecuteService(svc, req);
	            } else {
	                var service = {};
	                service.method = "PUT";
	                service.url = _Globals.Application.EntityPrefix + entityName.toLowerCase() + "/" + id;
	                var req = this.RequestBuilder.DefaultRequest(null, fieldmap, headers);
	                return this.DataSource.ExecuteServiceObject(service, req);
	            }
	        }
	    }]);

	    return EntityDataService;
	}();

/***/ }),
/* 54 */
/***/ (function(module, exports) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	
	var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();
	
	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
	
	var RequestBuilderService = exports.RequestBuilderService = function () {
	  function RequestBuilderService() {
	    _classCallCheck(this, RequestBuilderService);
	
	    this.ParameterSeparatorRequest = this.ParameterSeparatorRequest.bind(this);
	    this.DefaultRequest = this.DefaultRequest.bind(this);
	    this.URLParamsRequest = this.URLParamsRequest.bind(this);
	  }
	
	  _createClass(RequestBuilderService, [{
	    key: 'ParameterSeparatorRequest',
	    value: function ParameterSeparatorRequest(params, data, urlparams, headers) {
	      var parameterSeparator = {};
	      if (data == null) {
	        data = {};
	      }
	      parameterSeparator.params = params;
	      parameterSeparator.data = data;
	      parameterSeparator.urlparams = urlparams;
	      parameterSeparator.headers = headers;
	      parameterSeparator.GetRequest = function (protocol) {
	        if (protocol == 'http') {
	          var http = {};
	          http.data = parameterSeparator.data;
	          http.headers = parameterSeparator.headers;
	          if (parameterSeparator.params == null) {
	            http.params = null;
	            http.urlparams = null;
	            return http;
	          }
	          var httpUrlParams = {};
	          var httpParams = {};
	          var count = 0;
	          if (parameterSeparator.urlparams != null) {
	            for (var param in parameterSeparator.urlparams) {
	              if (param in parameterSeparator.params) {
	                httpUrlParams[param] = parameterSeparator.params[param];
	                count = count + 1;
	              }
	            }
	          }
	          if (count > 0) {
	            var remaincount = 0;
	            for (var param in parameterSeparator.params) {
	              if (param in httpUrlParams) {
	                continue;
	              } else {
	                httpParams[param] = parameterSeparator.params[param];
	                remaincount = remaincount + 1;
	              }
	            }
	            if (remaincount > 0) {
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
	  }, {
	    key: 'DefaultRequest',
	    value: function DefaultRequest(params, data, headers) {
	      var defaultRequest = {};
	      if (data == null) {
	        data = {};
	      }
	      defaultRequest.params = params;
	      defaultRequest.data = data;
	      defaultRequest.headers = headers;
	      defaultRequest.GetRequest = function (protocol) {
	        var request = {};
	        request.data = defaultRequest.data;
	        request.params = defaultRequest.params;
	        request.urlparams = null;
	        request.headers = defaultRequest.headers;
	        return request;
	      };
	      return defaultRequest;
	    }
	  }, {
	    key: 'URLParamsRequest',
	    value: function URLParamsRequest(urlparams, data, headers) {
	      var urlparamsRequest = {};
	      if (data == null) {
	        data = {};
	      }
	      urlparamsRequest.data = data;
	      urlparamsRequest.urlparams = urlparams;
	      urlparamsRequest.headers = headers;
	      urlparamsRequest.GetRequest = function (protocol) {
	        if (protocol == 'http') {
	          var http = {};
	          http.data = urlparamsRequest.data;
	          http.params = null;
	          http.urlparams = urlparamsRequest.urlparams;
	          http.headers = urlparamsRequest.headers;
	          return http;
	        } else {
	          var socket = {};
	          socket.data = urlparamsRequest.data;
	          socket.params = Object.assign({}, urlparamsRequest.urlparams, urlparamsRequest.headers);
	          return socket;
	        }
	      };
	      return urlparamsRequest;
	    }
	  }]);

	  return RequestBuilderService;
	}();

/***/ }),
/* 55 */
/***/ (function(module, exports) {

	"use strict";
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	
	var punjabiKeymap = ["੧", '"', "੩", "੪", "੫", "੭", "‘", "੯", "੦", "੮", "+", ",", "-", ".", "/", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "ਞ", "ਙ", "ਖ਼", "=", "ੴ", "?", "੨", "ੳ", "ਭ", "ਛ", "ਧ", "ਓ", "ਢ", "ਘ", "ੵ", "ੀ", "ਝ", "ਖ", "ਲ਼", "ੰ", "ਂ", "ੌ", "ਫ", "ਥ", "੍", "ਸ਼", "ਠ", "ੂ", "ੜ", "ਆ", "ਯ", "ੈ", "ਗ਼", "[", "\\", "]", "^", "_", "ਫ਼", "ਅ", "ਬ", "ਚ", "ਦ", "ੲ", "ਡ", "ਗ", "ਹ", "ਿ", "ਜ", "ਕ", "ਲ", "ਮ", "ਨ", "ੋ", "ਪ", "ਤ", "ਰ", "ਸ", "ਟ", "ੁ", "ਵ", "ਾ", "ਣ", "ੇ", "ਜ਼", "(", "|", ")", "ੱ", ""];
	
	exports.default = punjabiKeymap;

/***/ }),
/* 56 */
/***/ (function(module, exports) {

	'use strict'
	
	exports.byteLength = byteLength
	exports.toByteArray = toByteArray
	exports.fromByteArray = fromByteArray
	
	var lookup = []
	var revLookup = []
	var Arr = typeof Uint8Array !== 'undefined' ? Uint8Array : Array
	
	var code = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'
	for (var i = 0, len = code.length; i < len; ++i) {
	  lookup[i] = code[i]
	  revLookup[code.charCodeAt(i)] = i
	}
	
	revLookup['-'.charCodeAt(0)] = 62
	revLookup['_'.charCodeAt(0)] = 63
	
	function placeHoldersCount (b64) {
	  var len = b64.length
	  if (len % 4 > 0) {
	    throw new Error('Invalid string. Length must be a multiple of 4')
	  }
	
	  // the number of equal signs (place holders)
	  // if there are two placeholders, than the two characters before it
	  // represent one byte
	  // if there is only one, then the three characters before it represent 2 bytes
	  // this is just a cheap hack to not do indexOf twice
	  return b64[len - 2] === '=' ? 2 : b64[len - 1] === '=' ? 1 : 0
	}
	
	function byteLength (b64) {
	  // base64 is 4/3 + up to two characters of the original data
	  return b64.length * 3 / 4 - placeHoldersCount(b64)
	}
	
	function toByteArray (b64) {
	  var i, j, l, tmp, placeHolders, arr
	  var len = b64.length
	  placeHolders = placeHoldersCount(b64)
	
	  arr = new Arr(len * 3 / 4 - placeHolders)
	
	  // if there are placeholders, only get up to the last complete 4 chars
	  l = placeHolders > 0 ? len - 4 : len
	
	  var L = 0
	
	  for (i = 0, j = 0; i < l; i += 4, j += 3) {
	    tmp = (revLookup[b64.charCodeAt(i)] << 18) | (revLookup[b64.charCodeAt(i + 1)] << 12) | (revLookup[b64.charCodeAt(i + 2)] << 6) | revLookup[b64.charCodeAt(i + 3)]
	    arr[L++] = (tmp >> 16) & 0xFF
	    arr[L++] = (tmp >> 8) & 0xFF
	    arr[L++] = tmp & 0xFF
	  }
	
	  if (placeHolders === 2) {
	    tmp = (revLookup[b64.charCodeAt(i)] << 2) | (revLookup[b64.charCodeAt(i + 1)] >> 4)
	    arr[L++] = tmp & 0xFF
	  } else if (placeHolders === 1) {
	    tmp = (revLookup[b64.charCodeAt(i)] << 10) | (revLookup[b64.charCodeAt(i + 1)] << 4) | (revLookup[b64.charCodeAt(i + 2)] >> 2)
	    arr[L++] = (tmp >> 8) & 0xFF
	    arr[L++] = tmp & 0xFF
	  }
	
	  return arr
	}
	
	function tripletToBase64 (num) {
	  return lookup[num >> 18 & 0x3F] + lookup[num >> 12 & 0x3F] + lookup[num >> 6 & 0x3F] + lookup[num & 0x3F]
	}
	
	function encodeChunk (uint8, start, end) {
	  var tmp
	  var output = []
	  for (var i = start; i < end; i += 3) {
	    tmp = (uint8[i] << 16) + (uint8[i + 1] << 8) + (uint8[i + 2])
	    output.push(tripletToBase64(tmp))
	  }
	  return output.join('')
	}
	
	function fromByteArray (uint8) {
	  var tmp
	  var len = uint8.length
	  var extraBytes = len % 3 // if we have 1 byte left, pad 2 bytes
	  var output = ''
	  var parts = []
	  var maxChunkLength = 16383 // must be multiple of 3
	
	  // go through the array every three bytes, we'll deal with trailing stuff later
	  for (var i = 0, len2 = len - extraBytes; i < len2; i += maxChunkLength) {
	    parts.push(encodeChunk(uint8, i, (i + maxChunkLength) > len2 ? len2 : (i + maxChunkLength)))
	  }
	
	  // pad the end with zeros, but make sure to not forget the extra bytes
	  if (extraBytes === 1) {
	    tmp = uint8[len - 1]
	    output += lookup[tmp >> 2]
	    output += lookup[(tmp << 4) & 0x3F]
	    output += '=='
	  } else if (extraBytes === 2) {
	    tmp = (uint8[len - 2] << 8) + (uint8[len - 1])
	    output += lookup[tmp >> 10]
	    output += lookup[(tmp >> 4) & 0x3F]
	    output += lookup[(tmp << 2) & 0x3F]
	    output += '='
	  }
	
	  parts.push(output)
	
	  return parts.join('')
	}


/***/ }),
/* 57 */
/***/ (function(module, exports, __webpack_require__) {

	/* WEBPACK VAR INJECTION */(function(global) {/*!
	 * The buffer module from node.js, for the browser.
	 *
	 * @author   Feross Aboukhadijeh <feross@feross.org> <http://feross.org>
	 * @license  MIT
	 */
	/* eslint-disable no-proto */
	
	'use strict'
	
	var base64 = __webpack_require__(56)
	var ieee754 = __webpack_require__(58)
	var isArray = __webpack_require__(59)
	
	exports.Buffer = Buffer
	exports.SlowBuffer = SlowBuffer
	exports.INSPECT_MAX_BYTES = 50
	
	/**
	 * If `Buffer.TYPED_ARRAY_SUPPORT`:
	 *   === true    Use Uint8Array implementation (fastest)
	 *   === false   Use Object implementation (most compatible, even IE6)
	 *
	 * Browsers that support typed arrays are IE 10+, Firefox 4+, Chrome 7+, Safari 5.1+,
	 * Opera 11.6+, iOS 4.2+.
	 *
	 * Due to various browser bugs, sometimes the Object implementation will be used even
	 * when the browser supports typed arrays.
	 *
	 * Note:
	 *
	 *   - Firefox 4-29 lacks support for adding new properties to `Uint8Array` instances,
	 *     See: https://bugzilla.mozilla.org/show_bug.cgi?id=695438.
	 *
	 *   - Chrome 9-10 is missing the `TypedArray.prototype.subarray` function.
	 *
	 *   - IE10 has a broken `TypedArray.prototype.subarray` function which returns arrays of
	 *     incorrect length in some situations.
	
	 * We detect these buggy browsers and set `Buffer.TYPED_ARRAY_SUPPORT` to `false` so they
	 * get the Object implementation, which is slower but behaves correctly.
	 */
	Buffer.TYPED_ARRAY_SUPPORT = global.TYPED_ARRAY_SUPPORT !== undefined
	  ? global.TYPED_ARRAY_SUPPORT
	  : typedArraySupport()
	
	/*
	 * Export kMaxLength after typed array support is determined.
	 */
	exports.kMaxLength = kMaxLength()
	
	function typedArraySupport () {
	  try {
	    var arr = new Uint8Array(1)
	    arr.__proto__ = {__proto__: Uint8Array.prototype, foo: function () { return 42 }}
	    return arr.foo() === 42 && // typed array instances can be augmented
	        typeof arr.subarray === 'function' && // chrome 9-10 lack `subarray`
	        arr.subarray(1, 1).byteLength === 0 // ie10 has broken `subarray`
	  } catch (e) {
	    return false
	  }
	}
	
	function kMaxLength () {
	  return Buffer.TYPED_ARRAY_SUPPORT
	    ? 0x7fffffff
	    : 0x3fffffff
	}
	
	function createBuffer (that, length) {
	  if (kMaxLength() < length) {
	    throw new RangeError('Invalid typed array length')
	  }
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    // Return an augmented `Uint8Array` instance, for best performance
	    that = new Uint8Array(length)
	    that.__proto__ = Buffer.prototype
	  } else {
	    // Fallback: Return an object instance of the Buffer class
	    if (that === null) {
	      that = new Buffer(length)
	    }
	    that.length = length
	  }
	
	  return that
	}
	
	/**
	 * The Buffer constructor returns instances of `Uint8Array` that have their
	 * prototype changed to `Buffer.prototype`. Furthermore, `Buffer` is a subclass of
	 * `Uint8Array`, so the returned instances will have all the node `Buffer` methods
	 * and the `Uint8Array` methods. Square bracket notation works as expected -- it
	 * returns a single octet.
	 *
	 * The `Uint8Array` prototype remains unmodified.
	 */
	
	function Buffer (arg, encodingOrOffset, length) {
	  if (!Buffer.TYPED_ARRAY_SUPPORT && !(this instanceof Buffer)) {
	    return new Buffer(arg, encodingOrOffset, length)
	  }
	
	  // Common case.
	  if (typeof arg === 'number') {
	    if (typeof encodingOrOffset === 'string') {
	      throw new Error(
	        'If encoding is specified then the first argument must be a string'
	      )
	    }
	    return allocUnsafe(this, arg)
	  }
	  return from(this, arg, encodingOrOffset, length)
	}
	
	Buffer.poolSize = 8192 // not used by this implementation
	
	// TODO: Legacy, not needed anymore. Remove in next major version.
	Buffer._augment = function (arr) {
	  arr.__proto__ = Buffer.prototype
	  return arr
	}
	
	function from (that, value, encodingOrOffset, length) {
	  if (typeof value === 'number') {
	    throw new TypeError('"value" argument must not be a number')
	  }
	
	  if (typeof ArrayBuffer !== 'undefined' && value instanceof ArrayBuffer) {
	    return fromArrayBuffer(that, value, encodingOrOffset, length)
	  }
	
	  if (typeof value === 'string') {
	    return fromString(that, value, encodingOrOffset)
	  }
	
	  return fromObject(that, value)
	}
	
	/**
	 * Functionally equivalent to Buffer(arg, encoding) but throws a TypeError
	 * if value is a number.
	 * Buffer.from(str[, encoding])
	 * Buffer.from(array)
	 * Buffer.from(buffer)
	 * Buffer.from(arrayBuffer[, byteOffset[, length]])
	 **/
	Buffer.from = function (value, encodingOrOffset, length) {
	  return from(null, value, encodingOrOffset, length)
	}
	
	if (Buffer.TYPED_ARRAY_SUPPORT) {
	  Buffer.prototype.__proto__ = Uint8Array.prototype
	  Buffer.__proto__ = Uint8Array
	  if (typeof Symbol !== 'undefined' && Symbol.species &&
	      Buffer[Symbol.species] === Buffer) {
	    // Fix subarray() in ES2016. See: https://github.com/feross/buffer/pull/97
	    Object.defineProperty(Buffer, Symbol.species, {
	      value: null,
	      configurable: true
	    })
	  }
	}
	
	function assertSize (size) {
	  if (typeof size !== 'number') {
	    throw new TypeError('"size" argument must be a number')
	  } else if (size < 0) {
	    throw new RangeError('"size" argument must not be negative')
	  }
	}
	
	function alloc (that, size, fill, encoding) {
	  assertSize(size)
	  if (size <= 0) {
	    return createBuffer(that, size)
	  }
	  if (fill !== undefined) {
	    // Only pay attention to encoding if it's a string. This
	    // prevents accidentally sending in a number that would
	    // be interpretted as a start offset.
	    return typeof encoding === 'string'
	      ? createBuffer(that, size).fill(fill, encoding)
	      : createBuffer(that, size).fill(fill)
	  }
	  return createBuffer(that, size)
	}
	
	/**
	 * Creates a new filled Buffer instance.
	 * alloc(size[, fill[, encoding]])
	 **/
	Buffer.alloc = function (size, fill, encoding) {
	  return alloc(null, size, fill, encoding)
	}
	
	function allocUnsafe (that, size) {
	  assertSize(size)
	  that = createBuffer(that, size < 0 ? 0 : checked(size) | 0)
	  if (!Buffer.TYPED_ARRAY_SUPPORT) {
	    for (var i = 0; i < size; ++i) {
	      that[i] = 0
	    }
	  }
	  return that
	}
	
	/**
	 * Equivalent to Buffer(num), by default creates a non-zero-filled Buffer instance.
	 * */
	Buffer.allocUnsafe = function (size) {
	  return allocUnsafe(null, size)
	}
	/**
	 * Equivalent to SlowBuffer(num), by default creates a non-zero-filled Buffer instance.
	 */
	Buffer.allocUnsafeSlow = function (size) {
	  return allocUnsafe(null, size)
	}
	
	function fromString (that, string, encoding) {
	  if (typeof encoding !== 'string' || encoding === '') {
	    encoding = 'utf8'
	  }
	
	  if (!Buffer.isEncoding(encoding)) {
	    throw new TypeError('"encoding" must be a valid string encoding')
	  }
	
	  var length = byteLength(string, encoding) | 0
	  that = createBuffer(that, length)
	
	  var actual = that.write(string, encoding)
	
	  if (actual !== length) {
	    // Writing a hex string, for example, that contains invalid characters will
	    // cause everything after the first invalid character to be ignored. (e.g.
	    // 'abxxcd' will be treated as 'ab')
	    that = that.slice(0, actual)
	  }
	
	  return that
	}
	
	function fromArrayLike (that, array) {
	  var length = array.length < 0 ? 0 : checked(array.length) | 0
	  that = createBuffer(that, length)
	  for (var i = 0; i < length; i += 1) {
	    that[i] = array[i] & 255
	  }
	  return that
	}
	
	function fromArrayBuffer (that, array, byteOffset, length) {
	  array.byteLength // this throws if `array` is not a valid ArrayBuffer
	
	  if (byteOffset < 0 || array.byteLength < byteOffset) {
	    throw new RangeError('\'offset\' is out of bounds')
	  }
	
	  if (array.byteLength < byteOffset + (length || 0)) {
	    throw new RangeError('\'length\' is out of bounds')
	  }
	
	  if (byteOffset === undefined && length === undefined) {
	    array = new Uint8Array(array)
	  } else if (length === undefined) {
	    array = new Uint8Array(array, byteOffset)
	  } else {
	    array = new Uint8Array(array, byteOffset, length)
	  }
	
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    // Return an augmented `Uint8Array` instance, for best performance
	    that = array
	    that.__proto__ = Buffer.prototype
	  } else {
	    // Fallback: Return an object instance of the Buffer class
	    that = fromArrayLike(that, array)
	  }
	  return that
	}
	
	function fromObject (that, obj) {
	  if (Buffer.isBuffer(obj)) {
	    var len = checked(obj.length) | 0
	    that = createBuffer(that, len)
	
	    if (that.length === 0) {
	      return that
	    }
	
	    obj.copy(that, 0, 0, len)
	    return that
	  }
	
	  if (obj) {
	    if ((typeof ArrayBuffer !== 'undefined' &&
	        obj.buffer instanceof ArrayBuffer) || 'length' in obj) {
	      if (typeof obj.length !== 'number' || isnan(obj.length)) {
	        return createBuffer(that, 0)
	      }
	      return fromArrayLike(that, obj)
	    }
	
	    if (obj.type === 'Buffer' && isArray(obj.data)) {
	      return fromArrayLike(that, obj.data)
	    }
	  }
	
	  throw new TypeError('First argument must be a string, Buffer, ArrayBuffer, Array, or array-like object.')
	}
	
	function checked (length) {
	  // Note: cannot use `length < kMaxLength()` here because that fails when
	  // length is NaN (which is otherwise coerced to zero.)
	  if (length >= kMaxLength()) {
	    throw new RangeError('Attempt to allocate Buffer larger than maximum ' +
	                         'size: 0x' + kMaxLength().toString(16) + ' bytes')
	  }
	  return length | 0
	}
	
	function SlowBuffer (length) {
	  if (+length != length) { // eslint-disable-line eqeqeq
	    length = 0
	  }
	  return Buffer.alloc(+length)
	}
	
	Buffer.isBuffer = function isBuffer (b) {
	  return !!(b != null && b._isBuffer)
	}
	
	Buffer.compare = function compare (a, b) {
	  if (!Buffer.isBuffer(a) || !Buffer.isBuffer(b)) {
	    throw new TypeError('Arguments must be Buffers')
	  }
	
	  if (a === b) return 0
	
	  var x = a.length
	  var y = b.length
	
	  for (var i = 0, len = Math.min(x, y); i < len; ++i) {
	    if (a[i] !== b[i]) {
	      x = a[i]
	      y = b[i]
	      break
	    }
	  }
	
	  if (x < y) return -1
	  if (y < x) return 1
	  return 0
	}
	
	Buffer.isEncoding = function isEncoding (encoding) {
	  switch (String(encoding).toLowerCase()) {
	    case 'hex':
	    case 'utf8':
	    case 'utf-8':
	    case 'ascii':
	    case 'latin1':
	    case 'binary':
	    case 'base64':
	    case 'ucs2':
	    case 'ucs-2':
	    case 'utf16le':
	    case 'utf-16le':
	      return true
	    default:
	      return false
	  }
	}
	
	Buffer.concat = function concat (list, length) {
	  if (!isArray(list)) {
	    throw new TypeError('"list" argument must be an Array of Buffers')
	  }
	
	  if (list.length === 0) {
	    return Buffer.alloc(0)
	  }
	
	  var i
	  if (length === undefined) {
	    length = 0
	    for (i = 0; i < list.length; ++i) {
	      length += list[i].length
	    }
	  }
	
	  var buffer = Buffer.allocUnsafe(length)
	  var pos = 0
	  for (i = 0; i < list.length; ++i) {
	    var buf = list[i]
	    if (!Buffer.isBuffer(buf)) {
	      throw new TypeError('"list" argument must be an Array of Buffers')
	    }
	    buf.copy(buffer, pos)
	    pos += buf.length
	  }
	  return buffer
	}
	
	function byteLength (string, encoding) {
	  if (Buffer.isBuffer(string)) {
	    return string.length
	  }
	  if (typeof ArrayBuffer !== 'undefined' && typeof ArrayBuffer.isView === 'function' &&
	      (ArrayBuffer.isView(string) || string instanceof ArrayBuffer)) {
	    return string.byteLength
	  }
	  if (typeof string !== 'string') {
	    string = '' + string
	  }
	
	  var len = string.length
	  if (len === 0) return 0
	
	  // Use a for loop to avoid recursion
	  var loweredCase = false
	  for (;;) {
	    switch (encoding) {
	      case 'ascii':
	      case 'latin1':
	      case 'binary':
	        return len
	      case 'utf8':
	      case 'utf-8':
	      case undefined:
	        return utf8ToBytes(string).length
	      case 'ucs2':
	      case 'ucs-2':
	      case 'utf16le':
	      case 'utf-16le':
	        return len * 2
	      case 'hex':
	        return len >>> 1
	      case 'base64':
	        return base64ToBytes(string).length
	      default:
	        if (loweredCase) return utf8ToBytes(string).length // assume utf8
	        encoding = ('' + encoding).toLowerCase()
	        loweredCase = true
	    }
	  }
	}
	Buffer.byteLength = byteLength
	
	function slowToString (encoding, start, end) {
	  var loweredCase = false
	
	  // No need to verify that "this.length <= MAX_UINT32" since it's a read-only
	  // property of a typed array.
	
	  // This behaves neither like String nor Uint8Array in that we set start/end
	  // to their upper/lower bounds if the value passed is out of range.
	  // undefined is handled specially as per ECMA-262 6th Edition,
	  // Section 13.3.3.7 Runtime Semantics: KeyedBindingInitialization.
	  if (start === undefined || start < 0) {
	    start = 0
	  }
	  // Return early if start > this.length. Done here to prevent potential uint32
	  // coercion fail below.
	  if (start > this.length) {
	    return ''
	  }
	
	  if (end === undefined || end > this.length) {
	    end = this.length
	  }
	
	  if (end <= 0) {
	    return ''
	  }
	
	  // Force coersion to uint32. This will also coerce falsey/NaN values to 0.
	  end >>>= 0
	  start >>>= 0
	
	  if (end <= start) {
	    return ''
	  }
	
	  if (!encoding) encoding = 'utf8'
	
	  while (true) {
	    switch (encoding) {
	      case 'hex':
	        return hexSlice(this, start, end)
	
	      case 'utf8':
	      case 'utf-8':
	        return utf8Slice(this, start, end)
	
	      case 'ascii':
	        return asciiSlice(this, start, end)
	
	      case 'latin1':
	      case 'binary':
	        return latin1Slice(this, start, end)
	
	      case 'base64':
	        return base64Slice(this, start, end)
	
	      case 'ucs2':
	      case 'ucs-2':
	      case 'utf16le':
	      case 'utf-16le':
	        return utf16leSlice(this, start, end)
	
	      default:
	        if (loweredCase) throw new TypeError('Unknown encoding: ' + encoding)
	        encoding = (encoding + '').toLowerCase()
	        loweredCase = true
	    }
	  }
	}
	
	// The property is used by `Buffer.isBuffer` and `is-buffer` (in Safari 5-7) to detect
	// Buffer instances.
	Buffer.prototype._isBuffer = true
	
	function swap (b, n, m) {
	  var i = b[n]
	  b[n] = b[m]
	  b[m] = i
	}
	
	Buffer.prototype.swap16 = function swap16 () {
	  var len = this.length
	  if (len % 2 !== 0) {
	    throw new RangeError('Buffer size must be a multiple of 16-bits')
	  }
	  for (var i = 0; i < len; i += 2) {
	    swap(this, i, i + 1)
	  }
	  return this
	}
	
	Buffer.prototype.swap32 = function swap32 () {
	  var len = this.length
	  if (len % 4 !== 0) {
	    throw new RangeError('Buffer size must be a multiple of 32-bits')
	  }
	  for (var i = 0; i < len; i += 4) {
	    swap(this, i, i + 3)
	    swap(this, i + 1, i + 2)
	  }
	  return this
	}
	
	Buffer.prototype.swap64 = function swap64 () {
	  var len = this.length
	  if (len % 8 !== 0) {
	    throw new RangeError('Buffer size must be a multiple of 64-bits')
	  }
	  for (var i = 0; i < len; i += 8) {
	    swap(this, i, i + 7)
	    swap(this, i + 1, i + 6)
	    swap(this, i + 2, i + 5)
	    swap(this, i + 3, i + 4)
	  }
	  return this
	}
	
	Buffer.prototype.toString = function toString () {
	  var length = this.length | 0
	  if (length === 0) return ''
	  if (arguments.length === 0) return utf8Slice(this, 0, length)
	  return slowToString.apply(this, arguments)
	}
	
	Buffer.prototype.equals = function equals (b) {
	  if (!Buffer.isBuffer(b)) throw new TypeError('Argument must be a Buffer')
	  if (this === b) return true
	  return Buffer.compare(this, b) === 0
	}
	
	Buffer.prototype.inspect = function inspect () {
	  var str = ''
	  var max = exports.INSPECT_MAX_BYTES
	  if (this.length > 0) {
	    str = this.toString('hex', 0, max).match(/.{2}/g).join(' ')
	    if (this.length > max) str += ' ... '
	  }
	  return '<Buffer ' + str + '>'
	}
	
	Buffer.prototype.compare = function compare (target, start, end, thisStart, thisEnd) {
	  if (!Buffer.isBuffer(target)) {
	    throw new TypeError('Argument must be a Buffer')
	  }
	
	  if (start === undefined) {
	    start = 0
	  }
	  if (end === undefined) {
	    end = target ? target.length : 0
	  }
	  if (thisStart === undefined) {
	    thisStart = 0
	  }
	  if (thisEnd === undefined) {
	    thisEnd = this.length
	  }
	
	  if (start < 0 || end > target.length || thisStart < 0 || thisEnd > this.length) {
	    throw new RangeError('out of range index')
	  }
	
	  if (thisStart >= thisEnd && start >= end) {
	    return 0
	  }
	  if (thisStart >= thisEnd) {
	    return -1
	  }
	  if (start >= end) {
	    return 1
	  }
	
	  start >>>= 0
	  end >>>= 0
	  thisStart >>>= 0
	  thisEnd >>>= 0
	
	  if (this === target) return 0
	
	  var x = thisEnd - thisStart
	  var y = end - start
	  var len = Math.min(x, y)
	
	  var thisCopy = this.slice(thisStart, thisEnd)
	  var targetCopy = target.slice(start, end)
	
	  for (var i = 0; i < len; ++i) {
	    if (thisCopy[i] !== targetCopy[i]) {
	      x = thisCopy[i]
	      y = targetCopy[i]
	      break
	    }
	  }
	
	  if (x < y) return -1
	  if (y < x) return 1
	  return 0
	}
	
	// Finds either the first index of `val` in `buffer` at offset >= `byteOffset`,
	// OR the last index of `val` in `buffer` at offset <= `byteOffset`.
	//
	// Arguments:
	// - buffer - a Buffer to search
	// - val - a string, Buffer, or number
	// - byteOffset - an index into `buffer`; will be clamped to an int32
	// - encoding - an optional encoding, relevant is val is a string
	// - dir - true for indexOf, false for lastIndexOf
	function bidirectionalIndexOf (buffer, val, byteOffset, encoding, dir) {
	  // Empty buffer means no match
	  if (buffer.length === 0) return -1
	
	  // Normalize byteOffset
	  if (typeof byteOffset === 'string') {
	    encoding = byteOffset
	    byteOffset = 0
	  } else if (byteOffset > 0x7fffffff) {
	    byteOffset = 0x7fffffff
	  } else if (byteOffset < -0x80000000) {
	    byteOffset = -0x80000000
	  }
	  byteOffset = +byteOffset  // Coerce to Number.
	  if (isNaN(byteOffset)) {
	    // byteOffset: it it's undefined, null, NaN, "foo", etc, search whole buffer
	    byteOffset = dir ? 0 : (buffer.length - 1)
	  }
	
	  // Normalize byteOffset: negative offsets start from the end of the buffer
	  if (byteOffset < 0) byteOffset = buffer.length + byteOffset
	  if (byteOffset >= buffer.length) {
	    if (dir) return -1
	    else byteOffset = buffer.length - 1
	  } else if (byteOffset < 0) {
	    if (dir) byteOffset = 0
	    else return -1
	  }
	
	  // Normalize val
	  if (typeof val === 'string') {
	    val = Buffer.from(val, encoding)
	  }
	
	  // Finally, search either indexOf (if dir is true) or lastIndexOf
	  if (Buffer.isBuffer(val)) {
	    // Special case: looking for empty string/buffer always fails
	    if (val.length === 0) {
	      return -1
	    }
	    return arrayIndexOf(buffer, val, byteOffset, encoding, dir)
	  } else if (typeof val === 'number') {
	    val = val & 0xFF // Search for a byte value [0-255]
	    if (Buffer.TYPED_ARRAY_SUPPORT &&
	        typeof Uint8Array.prototype.indexOf === 'function') {
	      if (dir) {
	        return Uint8Array.prototype.indexOf.call(buffer, val, byteOffset)
	      } else {
	        return Uint8Array.prototype.lastIndexOf.call(buffer, val, byteOffset)
	      }
	    }
	    return arrayIndexOf(buffer, [ val ], byteOffset, encoding, dir)
	  }
	
	  throw new TypeError('val must be string, number or Buffer')
	}
	
	function arrayIndexOf (arr, val, byteOffset, encoding, dir) {
	  var indexSize = 1
	  var arrLength = arr.length
	  var valLength = val.length
	
	  if (encoding !== undefined) {
	    encoding = String(encoding).toLowerCase()
	    if (encoding === 'ucs2' || encoding === 'ucs-2' ||
	        encoding === 'utf16le' || encoding === 'utf-16le') {
	      if (arr.length < 2 || val.length < 2) {
	        return -1
	      }
	      indexSize = 2
	      arrLength /= 2
	      valLength /= 2
	      byteOffset /= 2
	    }
	  }
	
	  function read (buf, i) {
	    if (indexSize === 1) {
	      return buf[i]
	    } else {
	      return buf.readUInt16BE(i * indexSize)
	    }
	  }
	
	  var i
	  if (dir) {
	    var foundIndex = -1
	    for (i = byteOffset; i < arrLength; i++) {
	      if (read(arr, i) === read(val, foundIndex === -1 ? 0 : i - foundIndex)) {
	        if (foundIndex === -1) foundIndex = i
	        if (i - foundIndex + 1 === valLength) return foundIndex * indexSize
	      } else {
	        if (foundIndex !== -1) i -= i - foundIndex
	        foundIndex = -1
	      }
	    }
	  } else {
	    if (byteOffset + valLength > arrLength) byteOffset = arrLength - valLength
	    for (i = byteOffset; i >= 0; i--) {
	      var found = true
	      for (var j = 0; j < valLength; j++) {
	        if (read(arr, i + j) !== read(val, j)) {
	          found = false
	          break
	        }
	      }
	      if (found) return i
	    }
	  }
	
	  return -1
	}
	
	Buffer.prototype.includes = function includes (val, byteOffset, encoding) {
	  return this.indexOf(val, byteOffset, encoding) !== -1
	}
	
	Buffer.prototype.indexOf = function indexOf (val, byteOffset, encoding) {
	  return bidirectionalIndexOf(this, val, byteOffset, encoding, true)
	}
	
	Buffer.prototype.lastIndexOf = function lastIndexOf (val, byteOffset, encoding) {
	  return bidirectionalIndexOf(this, val, byteOffset, encoding, false)
	}
	
	function hexWrite (buf, string, offset, length) {
	  offset = Number(offset) || 0
	  var remaining = buf.length - offset
	  if (!length) {
	    length = remaining
	  } else {
	    length = Number(length)
	    if (length > remaining) {
	      length = remaining
	    }
	  }
	
	  // must be an even number of digits
	  var strLen = string.length
	  if (strLen % 2 !== 0) throw new TypeError('Invalid hex string')
	
	  if (length > strLen / 2) {
	    length = strLen / 2
	  }
	  for (var i = 0; i < length; ++i) {
	    var parsed = parseInt(string.substr(i * 2, 2), 16)
	    if (isNaN(parsed)) return i
	    buf[offset + i] = parsed
	  }
	  return i
	}
	
	function utf8Write (buf, string, offset, length) {
	  return blitBuffer(utf8ToBytes(string, buf.length - offset), buf, offset, length)
	}
	
	function asciiWrite (buf, string, offset, length) {
	  return blitBuffer(asciiToBytes(string), buf, offset, length)
	}
	
	function latin1Write (buf, string, offset, length) {
	  return asciiWrite(buf, string, offset, length)
	}
	
	function base64Write (buf, string, offset, length) {
	  return blitBuffer(base64ToBytes(string), buf, offset, length)
	}
	
	function ucs2Write (buf, string, offset, length) {
	  return blitBuffer(utf16leToBytes(string, buf.length - offset), buf, offset, length)
	}
	
	Buffer.prototype.write = function write (string, offset, length, encoding) {
	  // Buffer#write(string)
	  if (offset === undefined) {
	    encoding = 'utf8'
	    length = this.length
	    offset = 0
	  // Buffer#write(string, encoding)
	  } else if (length === undefined && typeof offset === 'string') {
	    encoding = offset
	    length = this.length
	    offset = 0
	  // Buffer#write(string, offset[, length][, encoding])
	  } else if (isFinite(offset)) {
	    offset = offset | 0
	    if (isFinite(length)) {
	      length = length | 0
	      if (encoding === undefined) encoding = 'utf8'
	    } else {
	      encoding = length
	      length = undefined
	    }
	  // legacy write(string, encoding, offset, length) - remove in v0.13
	  } else {
	    throw new Error(
	      'Buffer.write(string, encoding, offset[, length]) is no longer supported'
	    )
	  }
	
	  var remaining = this.length - offset
	  if (length === undefined || length > remaining) length = remaining
	
	  if ((string.length > 0 && (length < 0 || offset < 0)) || offset > this.length) {
	    throw new RangeError('Attempt to write outside buffer bounds')
	  }
	
	  if (!encoding) encoding = 'utf8'
	
	  var loweredCase = false
	  for (;;) {
	    switch (encoding) {
	      case 'hex':
	        return hexWrite(this, string, offset, length)
	
	      case 'utf8':
	      case 'utf-8':
	        return utf8Write(this, string, offset, length)
	
	      case 'ascii':
	        return asciiWrite(this, string, offset, length)
	
	      case 'latin1':
	      case 'binary':
	        return latin1Write(this, string, offset, length)
	
	      case 'base64':
	        // Warning: maxLength not taken into account in base64Write
	        return base64Write(this, string, offset, length)
	
	      case 'ucs2':
	      case 'ucs-2':
	      case 'utf16le':
	      case 'utf-16le':
	        return ucs2Write(this, string, offset, length)
	
	      default:
	        if (loweredCase) throw new TypeError('Unknown encoding: ' + encoding)
	        encoding = ('' + encoding).toLowerCase()
	        loweredCase = true
	    }
	  }
	}
	
	Buffer.prototype.toJSON = function toJSON () {
	  return {
	    type: 'Buffer',
	    data: Array.prototype.slice.call(this._arr || this, 0)
	  }
	}
	
	function base64Slice (buf, start, end) {
	  if (start === 0 && end === buf.length) {
	    return base64.fromByteArray(buf)
	  } else {
	    return base64.fromByteArray(buf.slice(start, end))
	  }
	}
	
	function utf8Slice (buf, start, end) {
	  end = Math.min(buf.length, end)
	  var res = []
	
	  var i = start
	  while (i < end) {
	    var firstByte = buf[i]
	    var codePoint = null
	    var bytesPerSequence = (firstByte > 0xEF) ? 4
	      : (firstByte > 0xDF) ? 3
	      : (firstByte > 0xBF) ? 2
	      : 1
	
	    if (i + bytesPerSequence <= end) {
	      var secondByte, thirdByte, fourthByte, tempCodePoint
	
	      switch (bytesPerSequence) {
	        case 1:
	          if (firstByte < 0x80) {
	            codePoint = firstByte
	          }
	          break
	        case 2:
	          secondByte = buf[i + 1]
	          if ((secondByte & 0xC0) === 0x80) {
	            tempCodePoint = (firstByte & 0x1F) << 0x6 | (secondByte & 0x3F)
	            if (tempCodePoint > 0x7F) {
	              codePoint = tempCodePoint
	            }
	          }
	          break
	        case 3:
	          secondByte = buf[i + 1]
	          thirdByte = buf[i + 2]
	          if ((secondByte & 0xC0) === 0x80 && (thirdByte & 0xC0) === 0x80) {
	            tempCodePoint = (firstByte & 0xF) << 0xC | (secondByte & 0x3F) << 0x6 | (thirdByte & 0x3F)
	            if (tempCodePoint > 0x7FF && (tempCodePoint < 0xD800 || tempCodePoint > 0xDFFF)) {
	              codePoint = tempCodePoint
	            }
	          }
	          break
	        case 4:
	          secondByte = buf[i + 1]
	          thirdByte = buf[i + 2]
	          fourthByte = buf[i + 3]
	          if ((secondByte & 0xC0) === 0x80 && (thirdByte & 0xC0) === 0x80 && (fourthByte & 0xC0) === 0x80) {
	            tempCodePoint = (firstByte & 0xF) << 0x12 | (secondByte & 0x3F) << 0xC | (thirdByte & 0x3F) << 0x6 | (fourthByte & 0x3F)
	            if (tempCodePoint > 0xFFFF && tempCodePoint < 0x110000) {
	              codePoint = tempCodePoint
	            }
	          }
	      }
	    }
	
	    if (codePoint === null) {
	      // we did not generate a valid codePoint so insert a
	      // replacement char (U+FFFD) and advance only 1 byte
	      codePoint = 0xFFFD
	      bytesPerSequence = 1
	    } else if (codePoint > 0xFFFF) {
	      // encode to utf16 (surrogate pair dance)
	      codePoint -= 0x10000
	      res.push(codePoint >>> 10 & 0x3FF | 0xD800)
	      codePoint = 0xDC00 | codePoint & 0x3FF
	    }
	
	    res.push(codePoint)
	    i += bytesPerSequence
	  }
	
	  return decodeCodePointsArray(res)
	}
	
	// Based on http://stackoverflow.com/a/22747272/680742, the browser with
	// the lowest limit is Chrome, with 0x10000 args.
	// We go 1 magnitude less, for safety
	var MAX_ARGUMENTS_LENGTH = 0x1000
	
	function decodeCodePointsArray (codePoints) {
	  var len = codePoints.length
	  if (len <= MAX_ARGUMENTS_LENGTH) {
	    return String.fromCharCode.apply(String, codePoints) // avoid extra slice()
	  }
	
	  // Decode in chunks to avoid "call stack size exceeded".
	  var res = ''
	  var i = 0
	  while (i < len) {
	    res += String.fromCharCode.apply(
	      String,
	      codePoints.slice(i, i += MAX_ARGUMENTS_LENGTH)
	    )
	  }
	  return res
	}
	
	function asciiSlice (buf, start, end) {
	  var ret = ''
	  end = Math.min(buf.length, end)
	
	  for (var i = start; i < end; ++i) {
	    ret += String.fromCharCode(buf[i] & 0x7F)
	  }
	  return ret
	}
	
	function latin1Slice (buf, start, end) {
	  var ret = ''
	  end = Math.min(buf.length, end)
	
	  for (var i = start; i < end; ++i) {
	    ret += String.fromCharCode(buf[i])
	  }
	  return ret
	}
	
	function hexSlice (buf, start, end) {
	  var len = buf.length
	
	  if (!start || start < 0) start = 0
	  if (!end || end < 0 || end > len) end = len
	
	  var out = ''
	  for (var i = start; i < end; ++i) {
	    out += toHex(buf[i])
	  }
	  return out
	}
	
	function utf16leSlice (buf, start, end) {
	  var bytes = buf.slice(start, end)
	  var res = ''
	  for (var i = 0; i < bytes.length; i += 2) {
	    res += String.fromCharCode(bytes[i] + bytes[i + 1] * 256)
	  }
	  return res
	}
	
	Buffer.prototype.slice = function slice (start, end) {
	  var len = this.length
	  start = ~~start
	  end = end === undefined ? len : ~~end
	
	  if (start < 0) {
	    start += len
	    if (start < 0) start = 0
	  } else if (start > len) {
	    start = len
	  }
	
	  if (end < 0) {
	    end += len
	    if (end < 0) end = 0
	  } else if (end > len) {
	    end = len
	  }
	
	  if (end < start) end = start
	
	  var newBuf
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    newBuf = this.subarray(start, end)
	    newBuf.__proto__ = Buffer.prototype
	  } else {
	    var sliceLen = end - start
	    newBuf = new Buffer(sliceLen, undefined)
	    for (var i = 0; i < sliceLen; ++i) {
	      newBuf[i] = this[i + start]
	    }
	  }
	
	  return newBuf
	}
	
	/*
	 * Need to make sure that buffer isn't trying to write out of bounds.
	 */
	function checkOffset (offset, ext, length) {
	  if ((offset % 1) !== 0 || offset < 0) throw new RangeError('offset is not uint')
	  if (offset + ext > length) throw new RangeError('Trying to access beyond buffer length')
	}
	
	Buffer.prototype.readUIntLE = function readUIntLE (offset, byteLength, noAssert) {
	  offset = offset | 0
	  byteLength = byteLength | 0
	  if (!noAssert) checkOffset(offset, byteLength, this.length)
	
	  var val = this[offset]
	  var mul = 1
	  var i = 0
	  while (++i < byteLength && (mul *= 0x100)) {
	    val += this[offset + i] * mul
	  }
	
	  return val
	}
	
	Buffer.prototype.readUIntBE = function readUIntBE (offset, byteLength, noAssert) {
	  offset = offset | 0
	  byteLength = byteLength | 0
	  if (!noAssert) {
	    checkOffset(offset, byteLength, this.length)
	  }
	
	  var val = this[offset + --byteLength]
	  var mul = 1
	  while (byteLength > 0 && (mul *= 0x100)) {
	    val += this[offset + --byteLength] * mul
	  }
	
	  return val
	}
	
	Buffer.prototype.readUInt8 = function readUInt8 (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 1, this.length)
	  return this[offset]
	}
	
	Buffer.prototype.readUInt16LE = function readUInt16LE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 2, this.length)
	  return this[offset] | (this[offset + 1] << 8)
	}
	
	Buffer.prototype.readUInt16BE = function readUInt16BE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 2, this.length)
	  return (this[offset] << 8) | this[offset + 1]
	}
	
	Buffer.prototype.readUInt32LE = function readUInt32LE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 4, this.length)
	
	  return ((this[offset]) |
	      (this[offset + 1] << 8) |
	      (this[offset + 2] << 16)) +
	      (this[offset + 3] * 0x1000000)
	}
	
	Buffer.prototype.readUInt32BE = function readUInt32BE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 4, this.length)
	
	  return (this[offset] * 0x1000000) +
	    ((this[offset + 1] << 16) |
	    (this[offset + 2] << 8) |
	    this[offset + 3])
	}
	
	Buffer.prototype.readIntLE = function readIntLE (offset, byteLength, noAssert) {
	  offset = offset | 0
	  byteLength = byteLength | 0
	  if (!noAssert) checkOffset(offset, byteLength, this.length)
	
	  var val = this[offset]
	  var mul = 1
	  var i = 0
	  while (++i < byteLength && (mul *= 0x100)) {
	    val += this[offset + i] * mul
	  }
	  mul *= 0x80
	
	  if (val >= mul) val -= Math.pow(2, 8 * byteLength)
	
	  return val
	}
	
	Buffer.prototype.readIntBE = function readIntBE (offset, byteLength, noAssert) {
	  offset = offset | 0
	  byteLength = byteLength | 0
	  if (!noAssert) checkOffset(offset, byteLength, this.length)
	
	  var i = byteLength
	  var mul = 1
	  var val = this[offset + --i]
	  while (i > 0 && (mul *= 0x100)) {
	    val += this[offset + --i] * mul
	  }
	  mul *= 0x80
	
	  if (val >= mul) val -= Math.pow(2, 8 * byteLength)
	
	  return val
	}
	
	Buffer.prototype.readInt8 = function readInt8 (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 1, this.length)
	  if (!(this[offset] & 0x80)) return (this[offset])
	  return ((0xff - this[offset] + 1) * -1)
	}
	
	Buffer.prototype.readInt16LE = function readInt16LE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 2, this.length)
	  var val = this[offset] | (this[offset + 1] << 8)
	  return (val & 0x8000) ? val | 0xFFFF0000 : val
	}
	
	Buffer.prototype.readInt16BE = function readInt16BE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 2, this.length)
	  var val = this[offset + 1] | (this[offset] << 8)
	  return (val & 0x8000) ? val | 0xFFFF0000 : val
	}
	
	Buffer.prototype.readInt32LE = function readInt32LE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 4, this.length)
	
	  return (this[offset]) |
	    (this[offset + 1] << 8) |
	    (this[offset + 2] << 16) |
	    (this[offset + 3] << 24)
	}
	
	Buffer.prototype.readInt32BE = function readInt32BE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 4, this.length)
	
	  return (this[offset] << 24) |
	    (this[offset + 1] << 16) |
	    (this[offset + 2] << 8) |
	    (this[offset + 3])
	}
	
	Buffer.prototype.readFloatLE = function readFloatLE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 4, this.length)
	  return ieee754.read(this, offset, true, 23, 4)
	}
	
	Buffer.prototype.readFloatBE = function readFloatBE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 4, this.length)
	  return ieee754.read(this, offset, false, 23, 4)
	}
	
	Buffer.prototype.readDoubleLE = function readDoubleLE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 8, this.length)
	  return ieee754.read(this, offset, true, 52, 8)
	}
	
	Buffer.prototype.readDoubleBE = function readDoubleBE (offset, noAssert) {
	  if (!noAssert) checkOffset(offset, 8, this.length)
	  return ieee754.read(this, offset, false, 52, 8)
	}
	
	function checkInt (buf, value, offset, ext, max, min) {
	  if (!Buffer.isBuffer(buf)) throw new TypeError('"buffer" argument must be a Buffer instance')
	  if (value > max || value < min) throw new RangeError('"value" argument is out of bounds')
	  if (offset + ext > buf.length) throw new RangeError('Index out of range')
	}
	
	Buffer.prototype.writeUIntLE = function writeUIntLE (value, offset, byteLength, noAssert) {
	  value = +value
	  offset = offset | 0
	  byteLength = byteLength | 0
	  if (!noAssert) {
	    var maxBytes = Math.pow(2, 8 * byteLength) - 1
	    checkInt(this, value, offset, byteLength, maxBytes, 0)
	  }
	
	  var mul = 1
	  var i = 0
	  this[offset] = value & 0xFF
	  while (++i < byteLength && (mul *= 0x100)) {
	    this[offset + i] = (value / mul) & 0xFF
	  }
	
	  return offset + byteLength
	}
	
	Buffer.prototype.writeUIntBE = function writeUIntBE (value, offset, byteLength, noAssert) {
	  value = +value
	  offset = offset | 0
	  byteLength = byteLength | 0
	  if (!noAssert) {
	    var maxBytes = Math.pow(2, 8 * byteLength) - 1
	    checkInt(this, value, offset, byteLength, maxBytes, 0)
	  }
	
	  var i = byteLength - 1
	  var mul = 1
	  this[offset + i] = value & 0xFF
	  while (--i >= 0 && (mul *= 0x100)) {
	    this[offset + i] = (value / mul) & 0xFF
	  }
	
	  return offset + byteLength
	}
	
	Buffer.prototype.writeUInt8 = function writeUInt8 (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 1, 0xff, 0)
	  if (!Buffer.TYPED_ARRAY_SUPPORT) value = Math.floor(value)
	  this[offset] = (value & 0xff)
	  return offset + 1
	}
	
	function objectWriteUInt16 (buf, value, offset, littleEndian) {
	  if (value < 0) value = 0xffff + value + 1
	  for (var i = 0, j = Math.min(buf.length - offset, 2); i < j; ++i) {
	    buf[offset + i] = (value & (0xff << (8 * (littleEndian ? i : 1 - i)))) >>>
	      (littleEndian ? i : 1 - i) * 8
	  }
	}
	
	Buffer.prototype.writeUInt16LE = function writeUInt16LE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 2, 0xffff, 0)
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset] = (value & 0xff)
	    this[offset + 1] = (value >>> 8)
	  } else {
	    objectWriteUInt16(this, value, offset, true)
	  }
	  return offset + 2
	}
	
	Buffer.prototype.writeUInt16BE = function writeUInt16BE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 2, 0xffff, 0)
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset] = (value >>> 8)
	    this[offset + 1] = (value & 0xff)
	  } else {
	    objectWriteUInt16(this, value, offset, false)
	  }
	  return offset + 2
	}
	
	function objectWriteUInt32 (buf, value, offset, littleEndian) {
	  if (value < 0) value = 0xffffffff + value + 1
	  for (var i = 0, j = Math.min(buf.length - offset, 4); i < j; ++i) {
	    buf[offset + i] = (value >>> (littleEndian ? i : 3 - i) * 8) & 0xff
	  }
	}
	
	Buffer.prototype.writeUInt32LE = function writeUInt32LE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 4, 0xffffffff, 0)
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset + 3] = (value >>> 24)
	    this[offset + 2] = (value >>> 16)
	    this[offset + 1] = (value >>> 8)
	    this[offset] = (value & 0xff)
	  } else {
	    objectWriteUInt32(this, value, offset, true)
	  }
	  return offset + 4
	}
	
	Buffer.prototype.writeUInt32BE = function writeUInt32BE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 4, 0xffffffff, 0)
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset] = (value >>> 24)
	    this[offset + 1] = (value >>> 16)
	    this[offset + 2] = (value >>> 8)
	    this[offset + 3] = (value & 0xff)
	  } else {
	    objectWriteUInt32(this, value, offset, false)
	  }
	  return offset + 4
	}
	
	Buffer.prototype.writeIntLE = function writeIntLE (value, offset, byteLength, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) {
	    var limit = Math.pow(2, 8 * byteLength - 1)
	
	    checkInt(this, value, offset, byteLength, limit - 1, -limit)
	  }
	
	  var i = 0
	  var mul = 1
	  var sub = 0
	  this[offset] = value & 0xFF
	  while (++i < byteLength && (mul *= 0x100)) {
	    if (value < 0 && sub === 0 && this[offset + i - 1] !== 0) {
	      sub = 1
	    }
	    this[offset + i] = ((value / mul) >> 0) - sub & 0xFF
	  }
	
	  return offset + byteLength
	}
	
	Buffer.prototype.writeIntBE = function writeIntBE (value, offset, byteLength, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) {
	    var limit = Math.pow(2, 8 * byteLength - 1)
	
	    checkInt(this, value, offset, byteLength, limit - 1, -limit)
	  }
	
	  var i = byteLength - 1
	  var mul = 1
	  var sub = 0
	  this[offset + i] = value & 0xFF
	  while (--i >= 0 && (mul *= 0x100)) {
	    if (value < 0 && sub === 0 && this[offset + i + 1] !== 0) {
	      sub = 1
	    }
	    this[offset + i] = ((value / mul) >> 0) - sub & 0xFF
	  }
	
	  return offset + byteLength
	}
	
	Buffer.prototype.writeInt8 = function writeInt8 (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 1, 0x7f, -0x80)
	  if (!Buffer.TYPED_ARRAY_SUPPORT) value = Math.floor(value)
	  if (value < 0) value = 0xff + value + 1
	  this[offset] = (value & 0xff)
	  return offset + 1
	}
	
	Buffer.prototype.writeInt16LE = function writeInt16LE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 2, 0x7fff, -0x8000)
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset] = (value & 0xff)
	    this[offset + 1] = (value >>> 8)
	  } else {
	    objectWriteUInt16(this, value, offset, true)
	  }
	  return offset + 2
	}
	
	Buffer.prototype.writeInt16BE = function writeInt16BE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 2, 0x7fff, -0x8000)
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset] = (value >>> 8)
	    this[offset + 1] = (value & 0xff)
	  } else {
	    objectWriteUInt16(this, value, offset, false)
	  }
	  return offset + 2
	}
	
	Buffer.prototype.writeInt32LE = function writeInt32LE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 4, 0x7fffffff, -0x80000000)
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset] = (value & 0xff)
	    this[offset + 1] = (value >>> 8)
	    this[offset + 2] = (value >>> 16)
	    this[offset + 3] = (value >>> 24)
	  } else {
	    objectWriteUInt32(this, value, offset, true)
	  }
	  return offset + 4
	}
	
	Buffer.prototype.writeInt32BE = function writeInt32BE (value, offset, noAssert) {
	  value = +value
	  offset = offset | 0
	  if (!noAssert) checkInt(this, value, offset, 4, 0x7fffffff, -0x80000000)
	  if (value < 0) value = 0xffffffff + value + 1
	  if (Buffer.TYPED_ARRAY_SUPPORT) {
	    this[offset] = (value >>> 24)
	    this[offset + 1] = (value >>> 16)
	    this[offset + 2] = (value >>> 8)
	    this[offset + 3] = (value & 0xff)
	  } else {
	    objectWriteUInt32(this, value, offset, false)
	  }
	  return offset + 4
	}
	
	function checkIEEE754 (buf, value, offset, ext, max, min) {
	  if (offset + ext > buf.length) throw new RangeError('Index out of range')
	  if (offset < 0) throw new RangeError('Index out of range')
	}
	
	function writeFloat (buf, value, offset, littleEndian, noAssert) {
	  if (!noAssert) {
	    checkIEEE754(buf, value, offset, 4, 3.4028234663852886e+38, -3.4028234663852886e+38)
	  }
	  ieee754.write(buf, value, offset, littleEndian, 23, 4)
	  return offset + 4
	}
	
	Buffer.prototype.writeFloatLE = function writeFloatLE (value, offset, noAssert) {
	  return writeFloat(this, value, offset, true, noAssert)
	}
	
	Buffer.prototype.writeFloatBE = function writeFloatBE (value, offset, noAssert) {
	  return writeFloat(this, value, offset, false, noAssert)
	}
	
	function writeDouble (buf, value, offset, littleEndian, noAssert) {
	  if (!noAssert) {
	    checkIEEE754(buf, value, offset, 8, 1.7976931348623157E+308, -1.7976931348623157E+308)
	  }
	  ieee754.write(buf, value, offset, littleEndian, 52, 8)
	  return offset + 8
	}
	
	Buffer.prototype.writeDoubleLE = function writeDoubleLE (value, offset, noAssert) {
	  return writeDouble(this, value, offset, true, noAssert)
	}
	
	Buffer.prototype.writeDoubleBE = function writeDoubleBE (value, offset, noAssert) {
	  return writeDouble(this, value, offset, false, noAssert)
	}
	
	// copy(targetBuffer, targetStart=0, sourceStart=0, sourceEnd=buffer.length)
	Buffer.prototype.copy = function copy (target, targetStart, start, end) {
	  if (!start) start = 0
	  if (!end && end !== 0) end = this.length
	  if (targetStart >= target.length) targetStart = target.length
	  if (!targetStart) targetStart = 0
	  if (end > 0 && end < start) end = start
	
	  // Copy 0 bytes; we're done
	  if (end === start) return 0
	  if (target.length === 0 || this.length === 0) return 0
	
	  // Fatal error conditions
	  if (targetStart < 0) {
	    throw new RangeError('targetStart out of bounds')
	  }
	  if (start < 0 || start >= this.length) throw new RangeError('sourceStart out of bounds')
	  if (end < 0) throw new RangeError('sourceEnd out of bounds')
	
	  // Are we oob?
	  if (end > this.length) end = this.length
	  if (target.length - targetStart < end - start) {
	    end = target.length - targetStart + start
	  }
	
	  var len = end - start
	  var i
	
	  if (this === target && start < targetStart && targetStart < end) {
	    // descending copy from end
	    for (i = len - 1; i >= 0; --i) {
	      target[i + targetStart] = this[i + start]
	    }
	  } else if (len < 1000 || !Buffer.TYPED_ARRAY_SUPPORT) {
	    // ascending copy from start
	    for (i = 0; i < len; ++i) {
	      target[i + targetStart] = this[i + start]
	    }
	  } else {
	    Uint8Array.prototype.set.call(
	      target,
	      this.subarray(start, start + len),
	      targetStart
	    )
	  }
	
	  return len
	}
	
	// Usage:
	//    buffer.fill(number[, offset[, end]])
	//    buffer.fill(buffer[, offset[, end]])
	//    buffer.fill(string[, offset[, end]][, encoding])
	Buffer.prototype.fill = function fill (val, start, end, encoding) {
	  // Handle string cases:
	  if (typeof val === 'string') {
	    if (typeof start === 'string') {
	      encoding = start
	      start = 0
	      end = this.length
	    } else if (typeof end === 'string') {
	      encoding = end
	      end = this.length
	    }
	    if (val.length === 1) {
	      var code = val.charCodeAt(0)
	      if (code < 256) {
	        val = code
	      }
	    }
	    if (encoding !== undefined && typeof encoding !== 'string') {
	      throw new TypeError('encoding must be a string')
	    }
	    if (typeof encoding === 'string' && !Buffer.isEncoding(encoding)) {
	      throw new TypeError('Unknown encoding: ' + encoding)
	    }
	  } else if (typeof val === 'number') {
	    val = val & 255
	  }
	
	  // Invalid ranges are not set to a default, so can range check early.
	  if (start < 0 || this.length < start || this.length < end) {
	    throw new RangeError('Out of range index')
	  }
	
	  if (end <= start) {
	    return this
	  }
	
	  start = start >>> 0
	  end = end === undefined ? this.length : end >>> 0
	
	  if (!val) val = 0
	
	  var i
	  if (typeof val === 'number') {
	    for (i = start; i < end; ++i) {
	      this[i] = val
	    }
	  } else {
	    var bytes = Buffer.isBuffer(val)
	      ? val
	      : utf8ToBytes(new Buffer(val, encoding).toString())
	    var len = bytes.length
	    for (i = 0; i < end - start; ++i) {
	      this[i + start] = bytes[i % len]
	    }
	  }
	
	  return this
	}
	
	// HELPER FUNCTIONS
	// ================
	
	var INVALID_BASE64_RE = /[^+\/0-9A-Za-z-_]/g
	
	function base64clean (str) {
	  // Node strips out invalid characters like \n and \t from the string, base64-js does not
	  str = stringtrim(str).replace(INVALID_BASE64_RE, '')
	  // Node converts strings with length < 2 to ''
	  if (str.length < 2) return ''
	  // Node allows for non-padded base64 strings (missing trailing ===), base64-js does not
	  while (str.length % 4 !== 0) {
	    str = str + '='
	  }
	  return str
	}
	
	function stringtrim (str) {
	  if (str.trim) return str.trim()
	  return str.replace(/^\s+|\s+$/g, '')
	}
	
	function toHex (n) {
	  if (n < 16) return '0' + n.toString(16)
	  return n.toString(16)
	}
	
	function utf8ToBytes (string, units) {
	  units = units || Infinity
	  var codePoint
	  var length = string.length
	  var leadSurrogate = null
	  var bytes = []
	
	  for (var i = 0; i < length; ++i) {
	    codePoint = string.charCodeAt(i)
	
	    // is surrogate component
	    if (codePoint > 0xD7FF && codePoint < 0xE000) {
	      // last char was a lead
	      if (!leadSurrogate) {
	        // no lead yet
	        if (codePoint > 0xDBFF) {
	          // unexpected trail
	          if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
	          continue
	        } else if (i + 1 === length) {
	          // unpaired lead
	          if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
	          continue
	        }
	
	        // valid lead
	        leadSurrogate = codePoint
	
	        continue
	      }
	
	      // 2 leads in a row
	      if (codePoint < 0xDC00) {
	        if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
	        leadSurrogate = codePoint
	        continue
	      }
	
	      // valid surrogate pair
	      codePoint = (leadSurrogate - 0xD800 << 10 | codePoint - 0xDC00) + 0x10000
	    } else if (leadSurrogate) {
	      // valid bmp char, but last char was a lead
	      if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
	    }
	
	    leadSurrogate = null
	
	    // encode utf8
	    if (codePoint < 0x80) {
	      if ((units -= 1) < 0) break
	      bytes.push(codePoint)
	    } else if (codePoint < 0x800) {
	      if ((units -= 2) < 0) break
	      bytes.push(
	        codePoint >> 0x6 | 0xC0,
	        codePoint & 0x3F | 0x80
	      )
	    } else if (codePoint < 0x10000) {
	      if ((units -= 3) < 0) break
	      bytes.push(
	        codePoint >> 0xC | 0xE0,
	        codePoint >> 0x6 & 0x3F | 0x80,
	        codePoint & 0x3F | 0x80
	      )
	    } else if (codePoint < 0x110000) {
	      if ((units -= 4) < 0) break
	      bytes.push(
	        codePoint >> 0x12 | 0xF0,
	        codePoint >> 0xC & 0x3F | 0x80,
	        codePoint >> 0x6 & 0x3F | 0x80,
	        codePoint & 0x3F | 0x80
	      )
	    } else {
	      throw new Error('Invalid code point')
	    }
	  }
	
	  return bytes
	}
	
	function asciiToBytes (str) {
	  var byteArray = []
	  for (var i = 0; i < str.length; ++i) {
	    // Node's code seems to be doing this and not & 0x7F..
	    byteArray.push(str.charCodeAt(i) & 0xFF)
	  }
	  return byteArray
	}
	
	function utf16leToBytes (str, units) {
	  var c, hi, lo
	  var byteArray = []
	  for (var i = 0; i < str.length; ++i) {
	    if ((units -= 2) < 0) break
	
	    c = str.charCodeAt(i)
	    hi = c >> 8
	    lo = c % 256
	    byteArray.push(lo)
	    byteArray.push(hi)
	  }
	
	  return byteArray
	}
	
	function base64ToBytes (str) {
	  return base64.toByteArray(base64clean(str))
	}
	
	function blitBuffer (src, dst, offset, length) {
	  for (var i = 0; i < length; ++i) {
	    if ((i + offset >= dst.length) || (i >= src.length)) break
	    dst[i + offset] = src[i]
	  }
	  return i
	}
	
	function isnan (val) {
	  return val !== val // eslint-disable-line no-self-compare
	}
	
	/* WEBPACK VAR INJECTION */}.call(exports, (function() { return this; }())))

/***/ }),
/* 58 */
/***/ (function(module, exports) {

	exports.read = function (buffer, offset, isLE, mLen, nBytes) {
	  var e, m
	  var eLen = nBytes * 8 - mLen - 1
	  var eMax = (1 << eLen) - 1
	  var eBias = eMax >> 1
	  var nBits = -7
	  var i = isLE ? (nBytes - 1) : 0
	  var d = isLE ? -1 : 1
	  var s = buffer[offset + i]
	
	  i += d
	
	  e = s & ((1 << (-nBits)) - 1)
	  s >>= (-nBits)
	  nBits += eLen
	  for (; nBits > 0; e = e * 256 + buffer[offset + i], i += d, nBits -= 8) {}
	
	  m = e & ((1 << (-nBits)) - 1)
	  e >>= (-nBits)
	  nBits += mLen
	  for (; nBits > 0; m = m * 256 + buffer[offset + i], i += d, nBits -= 8) {}
	
	  if (e === 0) {
	    e = 1 - eBias
	  } else if (e === eMax) {
	    return m ? NaN : ((s ? -1 : 1) * Infinity)
	  } else {
	    m = m + Math.pow(2, mLen)
	    e = e - eBias
	  }
	  return (s ? -1 : 1) * m * Math.pow(2, e - mLen)
	}
	
	exports.write = function (buffer, value, offset, isLE, mLen, nBytes) {
	  var e, m, c
	  var eLen = nBytes * 8 - mLen - 1
	  var eMax = (1 << eLen) - 1
	  var eBias = eMax >> 1
	  var rt = (mLen === 23 ? Math.pow(2, -24) - Math.pow(2, -77) : 0)
	  var i = isLE ? 0 : (nBytes - 1)
	  var d = isLE ? 1 : -1
	  var s = value < 0 || (value === 0 && 1 / value < 0) ? 1 : 0
	
	  value = Math.abs(value)
	
	  if (isNaN(value) || value === Infinity) {
	    m = isNaN(value) ? 1 : 0
	    e = eMax
	  } else {
	    e = Math.floor(Math.log(value) / Math.LN2)
	    if (value * (c = Math.pow(2, -e)) < 1) {
	      e--
	      c *= 2
	    }
	    if (e + eBias >= 1) {
	      value += rt / c
	    } else {
	      value += rt * Math.pow(2, 1 - eBias)
	    }
	    if (value * c >= 2) {
	      e++
	      c /= 2
	    }
	
	    if (e + eBias >= eMax) {
	      m = 0
	      e = eMax
	    } else if (e + eBias >= 1) {
	      m = (value * c - 1) * Math.pow(2, mLen)
	      e = e + eBias
	    } else {
	      m = value * Math.pow(2, eBias - 1) * Math.pow(2, mLen)
	      e = 0
	    }
	  }
	
	  for (; mLen >= 8; buffer[offset + i] = m & 0xff, i += d, m /= 256, mLen -= 8) {}
	
	  e = (e << mLen) | m
	  eLen += mLen
	  for (; eLen > 0; buffer[offset + i] = e & 0xff, i += d, e /= 256, eLen -= 8) {}
	
	  buffer[offset + i - d] |= s * 128
	}


/***/ }),
/* 59 */
/***/ (function(module, exports) {

	var toString = {}.toString;
	
	module.exports = Array.isArray || function (arr) {
	  return toString.call(arr) == '[object Array]';
	};


/***/ }),
/* 60 */
/***/ (function(module, exports) {

	// shim for using process in browser
	var process = module.exports = {};
	
	// cached from whatever global is present so that test runners that stub it
	// don't break things.  But we need to wrap it in a try catch in case it is
	// wrapped in strict mode code which doesn't define any globals.  It's inside a
	// function because try/catches deoptimize in certain engines.
	
	var cachedSetTimeout;
	var cachedClearTimeout;
	
	function defaultSetTimout() {
	    throw new Error('setTimeout has not been defined');
	}
	function defaultClearTimeout () {
	    throw new Error('clearTimeout has not been defined');
	}
	(function () {
	    try {
	        if (typeof setTimeout === 'function') {
	            cachedSetTimeout = setTimeout;
	        } else {
	            cachedSetTimeout = defaultSetTimout;
	        }
	    } catch (e) {
	        cachedSetTimeout = defaultSetTimout;
	    }
	    try {
	        if (typeof clearTimeout === 'function') {
	            cachedClearTimeout = clearTimeout;
	        } else {
	            cachedClearTimeout = defaultClearTimeout;
	        }
	    } catch (e) {
	        cachedClearTimeout = defaultClearTimeout;
	    }
	} ())
	function runTimeout(fun) {
	    if (cachedSetTimeout === setTimeout) {
	        //normal enviroments in sane situations
	        return setTimeout(fun, 0);
	    }
	    // if setTimeout wasn't available but was latter defined
	    if ((cachedSetTimeout === defaultSetTimout || !cachedSetTimeout) && setTimeout) {
	        cachedSetTimeout = setTimeout;
	        return setTimeout(fun, 0);
	    }
	    try {
	        // when when somebody has screwed with setTimeout but no I.E. maddness
	        return cachedSetTimeout(fun, 0);
	    } catch(e){
	        try {
	            // When we are in I.E. but the script has been evaled so I.E. doesn't trust the global object when called normally
	            return cachedSetTimeout.call(null, fun, 0);
	        } catch(e){
	            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error
	            return cachedSetTimeout.call(this, fun, 0);
	        }
	    }
	
	
	}
	function runClearTimeout(marker) {
	    if (cachedClearTimeout === clearTimeout) {
	        //normal enviroments in sane situations
	        return clearTimeout(marker);
	    }
	    // if clearTimeout wasn't available but was latter defined
	    if ((cachedClearTimeout === defaultClearTimeout || !cachedClearTimeout) && clearTimeout) {
	        cachedClearTimeout = clearTimeout;
	        return clearTimeout(marker);
	    }
	    try {
	        // when when somebody has screwed with setTimeout but no I.E. maddness
	        return cachedClearTimeout(marker);
	    } catch (e){
	        try {
	            // When we are in I.E. but the script has been evaled so I.E. doesn't  trust the global object when called normally
	            return cachedClearTimeout.call(null, marker);
	        } catch (e){
	            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error.
	            // Some versions of I.E. have different rules for clearTimeout vs setTimeout
	            return cachedClearTimeout.call(this, marker);
	        }
	    }
	
	
	
	}
	var queue = [];
	var draining = false;
	var currentQueue;
	var queueIndex = -1;
	
	function cleanUpNextTick() {
	    if (!draining || !currentQueue) {
	        return;
	    }
	    draining = false;
	    if (currentQueue.length) {
	        queue = currentQueue.concat(queue);
	    } else {
	        queueIndex = -1;
	    }
	    if (queue.length) {
	        drainQueue();
	    }
	}
	
	function drainQueue() {
	    if (draining) {
	        return;
	    }
	    var timeout = runTimeout(cleanUpNextTick);
	    draining = true;
	
	    var len = queue.length;
	    while(len) {
	        currentQueue = queue;
	        queue = [];
	        while (++queueIndex < len) {
	            if (currentQueue) {
	                currentQueue[queueIndex].run();
	            }
	        }
	        queueIndex = -1;
	        len = queue.length;
	    }
	    currentQueue = null;
	    draining = false;
	    runClearTimeout(timeout);
	}
	
	process.nextTick = function (fun) {
	    var args = new Array(arguments.length - 1);
	    if (arguments.length > 1) {
	        for (var i = 1; i < arguments.length; i++) {
	            args[i - 1] = arguments[i];
	        }
	    }
	    queue.push(new Item(fun, args));
	    if (queue.length === 1 && !draining) {
	        runTimeout(drainQueue);
	    }
	};
	
	// v8 likes predictible objects
	function Item(fun, array) {
	    this.fun = fun;
	    this.array = array;
	}
	Item.prototype.run = function () {
	    this.fun.apply(null, this.array);
	};
	process.title = 'browser';
	process.browser = true;
	process.env = {};
	process.argv = [];
	process.version = ''; // empty string to avoid regexp issues
	process.versions = {};
	
	function noop() {}
	
	process.on = noop;
	process.addListener = noop;
	process.once = noop;
	process.off = noop;
	process.removeListener = noop;
	process.removeAllListeners = noop;
	process.emit = noop;
	process.prependListener = noop;
	process.prependOnceListener = noop;
	
	process.listeners = function (name) { return [] }
	
	process.binding = function (name) {
	    throw new Error('process.binding is not supported');
	};
	
	process.cwd = function () { return '/' };
	process.chdir = function (dir) {
	    throw new Error('process.chdir is not supported');
	};
	process.umask = function() { return 0; };


/***/ }),
/* 61 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	exports.__esModule = true;
	exports.UNDEFINED_INPUT_ERROR = exports.INVALID_BUFFER = exports.isEnd = exports.END = undefined;
	
	var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };
	
	exports.emitter = emitter;
	exports.channel = channel;
	exports.eventChannel = eventChannel;
	exports.stdChannel = stdChannel;
	
	var _utils = __webpack_require__(9);
	
	var _buffers = __webpack_require__(20);
	
	var _scheduler = __webpack_require__(63);
	
	var CHANNEL_END_TYPE = '@@redux-saga/CHANNEL_END';
	var END = exports.END = { type: CHANNEL_END_TYPE };
	var isEnd = exports.isEnd = function isEnd(a) {
	  return a && a.type === CHANNEL_END_TYPE;
	};
	
	function emitter() {
	  var subscribers = [];
	
	  function subscribe(sub) {
	    subscribers.push(sub);
	    return function () {
	      return (0, _utils.remove)(subscribers, sub);
	    };
	  }
	
	  function emit(item) {
	    var arr = subscribers.slice();
	    for (var i = 0, len = arr.length; i < len; i++) {
	      arr[i](item);
	    }
	  }
	
	  return {
	    subscribe: subscribe,
	    emit: emit
	  };
	}
	
	var INVALID_BUFFER = exports.INVALID_BUFFER = 'invalid buffer passed to channel factory function';
	var UNDEFINED_INPUT_ERROR = exports.UNDEFINED_INPUT_ERROR = 'Saga was provided with an undefined action';
	
	if (false) {
	  exports.UNDEFINED_INPUT_ERROR = UNDEFINED_INPUT_ERROR += '\nHints:\n    - check that your Action Creator returns a non-undefined value\n    - if the Saga was started using runSaga, check that your subscribe source provides the action to its listeners\n  ';
	}
	
	function channel() {
	  var buffer = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : _buffers.buffers.fixed();
	
	  var closed = false;
	  var takers = [];
	
	  (0, _utils.check)(buffer, _utils.is.buffer, INVALID_BUFFER);
	
	  function checkForbiddenStates() {
	    if (closed && takers.length) {
	      throw (0, _utils.internalErr)('Cannot have a closed channel with pending takers');
	    }
	    if (takers.length && !buffer.isEmpty()) {
	      throw (0, _utils.internalErr)('Cannot have pending takers with non empty buffer');
	    }
	  }
	
	  function put(input) {
	    checkForbiddenStates();
	    (0, _utils.check)(input, _utils.is.notUndef, UNDEFINED_INPUT_ERROR);
	    if (closed) {
	      return;
	    }
	    if (!takers.length) {
	      return buffer.put(input);
	    }
	    for (var i = 0; i < takers.length; i++) {
	      var cb = takers[i];
	      if (!cb[_utils.MATCH] || cb[_utils.MATCH](input)) {
	        takers.splice(i, 1);
	        return cb(input);
	      }
	    }
	  }
	
	  function take(cb) {
	    checkForbiddenStates();
	    (0, _utils.check)(cb, _utils.is.func, 'channel.take\'s callback must be a function');
	
	    if (closed && buffer.isEmpty()) {
	      cb(END);
	    } else if (!buffer.isEmpty()) {
	      cb(buffer.take());
	    } else {
	      takers.push(cb);
	      cb.cancel = function () {
	        return (0, _utils.remove)(takers, cb);
	      };
	    }
	  }
	
	  function flush(cb) {
	    checkForbiddenStates(); // TODO: check if some new state should be forbidden now
	    (0, _utils.check)(cb, _utils.is.func, 'channel.flush\' callback must be a function');
	    if (closed && buffer.isEmpty()) {
	      cb(END);
	      return;
	    }
	    cb(buffer.flush());
	  }
	
	  function close() {
	    checkForbiddenStates();
	    if (!closed) {
	      closed = true;
	      if (takers.length) {
	        var arr = takers;
	        takers = [];
	        for (var i = 0, len = arr.length; i < len; i++) {
	          arr[i](END);
	        }
	      }
	    }
	  }
	
	  return { take: take, put: put, flush: flush, close: close,
	    get __takers__() {
	      return takers;
	    },
	    get __closed__() {
	      return closed;
	    }
	  };
	}
	
	function eventChannel(subscribe) {
	  var buffer = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : _buffers.buffers.none();
	  var matcher = arguments[2];
	
	  /**
	    should be if(typeof matcher !== undefined) instead?
	    see PR #273 for a background discussion
	  **/
	  if (arguments.length > 2) {
	    (0, _utils.check)(matcher, _utils.is.func, 'Invalid match function passed to eventChannel');
	  }
	
	  var chan = channel(buffer);
	  var close = function close() {
	    if (!chan.__closed__) {
	      if (unsubscribe) {
	        unsubscribe();
	      }
	      chan.close();
	    }
	  };
	  var unsubscribe = subscribe(function (input) {
	    if (isEnd(input)) {
	      close();
	      return;
	    }
	    if (matcher && !matcher(input)) {
	      return;
	    }
	    chan.put(input);
	  });
	  if (chan.__closed__) {
	    unsubscribe();
	  }
	
	  if (!_utils.is.func(unsubscribe)) {
	    throw new Error('in eventChannel: subscribe should return a function to unsubscribe');
	  }
	
	  return {
	    take: chan.take,
	    flush: chan.flush,
	    close: close
	  };
	}
	
	function stdChannel(subscribe) {
	  var chan = eventChannel(function (cb) {
	    return subscribe(function (input) {
	      if (input[_utils.SAGA_ACTION]) {
	        cb(input);
	        return;
	      }
	      (0, _scheduler.asap)(function () {
	        return cb(input);
	      });
	    });
	  });
	
	  return _extends({}, chan, {
	    take: function take(cb, matcher) {
	      if (arguments.length > 1) {
	        (0, _utils.check)(matcher, _utils.is.func, 'channel.take\'s matcher argument must be a function');
	        cb[_utils.MATCH] = matcher;
	      }
	      chan.take(cb);
	    }
	  });
	}

/***/ }),
/* 62 */
/***/ (function(module, exports, __webpack_require__) {

	'use strict';
	
	exports.__esModule = true;
	exports.throttle = exports.takeLatest = exports.takeEvery = undefined;
	exports.takeEveryHelper = takeEveryHelper;
	exports.takeLatestHelper = takeLatestHelper;
	exports.throttleHelper = throttleHelper;
	
	var _channel = __webpack_require__(61);
	
	var _utils = __webpack_require__(9);
	
	var _io = __webpack_require__(21);
	
	var _buffers = __webpack_require__(20);
	
	var done = { done: true, value: undefined };
	var qEnd = {};
	
	function fsmIterator(fsm, q0) {
	  var name = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : 'iterator';
	
	  var updateState = void 0,
	      qNext = q0;
	
	  function next(arg, error) {
	    if (qNext === qEnd) {
	      return done;
	    }
	
	    if (error) {
	      qNext = qEnd;
	      throw error;
	    } else {
	      updateState && updateState(arg);
	
	      var _fsm$qNext = fsm[qNext](),
	          q = _fsm$qNext[0],
	          output = _fsm$qNext[1],
	          _updateState = _fsm$qNext[2];
	
	      qNext = q;
	      updateState = _updateState;
	      return qNext === qEnd ? done : output;
	    }
	  }
	
	  return (0, _utils.makeIterator)(next, function (error) {
	    return next(null, error);
	  }, name, true);
	}
	
	function safeName(patternOrChannel) {
	  if (_utils.is.channel(patternOrChannel)) {
	    return 'channel';
	  } else if (Array.isArray(patternOrChannel)) {
	    return String(patternOrChannel.map(function (entry) {
	      return String(entry);
	    }));
	  } else {
	    return String(patternOrChannel);
	  }
	}
	
	function takeEveryHelper(patternOrChannel, worker) {
	  for (var _len = arguments.length, args = Array(_len > 2 ? _len - 2 : 0), _key = 2; _key < _len; _key++) {
	    args[_key - 2] = arguments[_key];
	  }
	
	  var yTake = { done: false, value: (0, _io.take)(patternOrChannel) };
	  var yFork = function yFork(ac) {
	    return { done: false, value: _io.fork.apply(undefined, [worker].concat(args, [ac])) };
	  };
	
	  var action = void 0,
	      setAction = function setAction(ac) {
	    return action = ac;
	  };
	
	  return fsmIterator({
	    q1: function q1() {
	      return ['q2', yTake, setAction];
	    },
	    q2: function q2() {
	      return action === _channel.END ? [qEnd] : ['q1', yFork(action)];
	    }
	  }, 'q1', 'takeEvery(' + safeName(patternOrChannel) + ', ' + worker.name + ')');
	}
	
	function takeLatestHelper(patternOrChannel, worker) {
	  for (var _len2 = arguments.length, args = Array(_len2 > 2 ? _len2 - 2 : 0), _key2 = 2; _key2 < _len2; _key2++) {
	    args[_key2 - 2] = arguments[_key2];
	  }
	
	  var yTake = { done: false, value: (0, _io.take)(patternOrChannel) };
	  var yFork = function yFork(ac) {
	    return { done: false, value: _io.fork.apply(undefined, [worker].concat(args, [ac])) };
	  };
	  var yCancel = function yCancel(task) {
	    return { done: false, value: (0, _io.cancel)(task) };
	  };
	
	  var task = void 0,
	      action = void 0;
	  var setTask = function setTask(t) {
	    return task = t;
	  };
	  var setAction = function setAction(ac) {
	    return action = ac;
	  };
	
	  return fsmIterator({
	    q1: function q1() {
	      return ['q2', yTake, setAction];
	    },
	    q2: function q2() {
	      return action === _channel.END ? [qEnd] : task ? ['q3', yCancel(task)] : ['q1', yFork(action), setTask];
	    },
	    q3: function q3() {
	      return ['q1', yFork(action), setTask];
	    }
	  }, 'q1', 'takeLatest(' + safeName(patternOrChannel) + ', ' + worker.name + ')');
	}
	
	function throttleHelper(delayLength, pattern, worker) {
	  for (var _len3 = arguments.length, args = Array(_len3 > 3 ? _len3 - 3 : 0), _key3 = 3; _key3 < _len3; _key3++) {
	    args[_key3 - 3] = arguments[_key3];
	  }
	
	  var action = void 0,
	      channel = void 0;
	
	  var yActionChannel = { done: false, value: (0, _io.actionChannel)(pattern, _buffers.buffers.sliding(1)) };
	  var yTake = function yTake() {
	    return { done: false, value: (0, _io.take)(channel) };
	  };
	  var yFork = function yFork(ac) {
	    return { done: false, value: _io.fork.apply(undefined, [worker].concat(args, [ac])) };
	  };
	  var yDelay = { done: false, value: (0, _io.call)(_utils.delay, delayLength) };
	
	  var setAction = function setAction(ac) {
	    return action = ac;
	  };
	  var setChannel = function setChannel(ch) {
	    return channel = ch;
	  };
	
	  return fsmIterator({
	    q1: function q1() {
	      return ['q2', yActionChannel, setChannel];
	    },
	    q2: function q2() {
	      return ['q3', yTake(), setAction];
	    },
	    q3: function q3() {
	      return action === _channel.END ? [qEnd] : ['q4', yFork(action)];
	    },
	    q4: function q4() {
	      return ['q2', yDelay];
	    }
	  }, 'q1', 'throttle(' + safeName(pattern) + ', ' + worker.name + ')');
	}
	
	var deprecationWarning = function deprecationWarning(helperName) {
	  return 'import { ' + helperName + ' } from \'redux-saga\' has been deprecated in favor of import { ' + helperName + ' } from \'redux-saga/effects\'.\nThe latter will not work with yield*, as helper effects are wrapped automatically for you in fork effect.\nTherefore yield ' + helperName + ' will return task descriptor to your saga and execute next lines of code.';
	};
	var takeEvery = exports.takeEvery = (0, _utils.deprecate)(takeEveryHelper, deprecationWarning('takeEvery'));
	var takeLatest = exports.takeLatest = (0, _utils.deprecate)(takeLatestHelper, deprecationWarning('takeLatest'));
	var throttle = exports.throttle = (0, _utils.deprecate)(throttleHelper, deprecationWarning('throttle'));

/***/ }),
/* 63 */
/***/ (function(module, exports) {

	"use strict";
	
	exports.__esModule = true;
	exports.asap = asap;
	exports.suspend = suspend;
	exports.flush = flush;
	
	var queue = [];
	/**
	  Variable to hold a counting semaphore
	  - Incrementing adds a lock and puts the scheduler in a `suspended` state (if it's not
	    already suspended)
	  - Decrementing releases a lock. Zero locks puts the scheduler in a `released` state. This
	    triggers flushing the queued tasks.
	**/
	var semaphore = 0;
	
	/**
	  Executes a task 'atomically'. Tasks scheduled during this execution will be queued
	  and flushed after this task has finished (assuming the scheduler endup in a released
	  state).
	**/
	function exec(task) {
	  try {
	    suspend();
	    task();
	  } finally {
	    release();
	  }
	}
	
	/**
	  Executes or queues a task depending on the state of the scheduler (`suspended` or `released`)
	**/
	function asap(task) {
	  queue.push(task);
	
	  if (!semaphore) {
	    suspend();
	    flush();
	  }
	}
	
	/**
	  Puts the scheduler in a `suspended` state. Scheduled tasks will be queued until the
	  scheduler is released.
	**/
	function suspend() {
	  semaphore++;
	}
	
	/**
	  Puts the scheduler in a `released` state.
	**/
	function release() {
	  semaphore--;
	}
	
	/**
	  Releases the current lock. Executes all queued tasks if the scheduler is in the released state.
	**/
	function flush() {
	  release();
	
	  var task = void 0;
	  while (!semaphore && (task = queue.shift()) !== undefined) {
	    exec(task);
	  }
	}

/***/ }),
/* 64 */
/***/ (function(module, exports) {

	module.exports = __WEBPACK_EXTERNAL_MODULE_64__;

/***/ }),
/* 65 */
/***/ (function(module, exports) {

	module.exports = __WEBPACK_EXTERNAL_MODULE_65__;

/***/ })
/******/ ])
});
;
//# sourceMappingURL=index.js.map