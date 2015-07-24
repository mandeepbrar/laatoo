(function() {
  'use strict';

  var mod = angular
    .module('loginui', ['ngAnimate', 'ngCookies', 'ngTouch', 'ngSanitize', 'ngStorage', 'ngResource', 'ui.router', 'ui.bootstrap', 'login']);


  mod.run(function($http) {
    //$http.defaults.headers.common.AuthToken = 'Basic YmVlcDpib29w'
  });

})();
