(function() {
  'use strict';

  var mod = angular
    .module('designerui', ['ngAnimate', 'ngCookies', 'ngTouch', 'ngSanitize', 'ngResource', 'ui.router', 'ui.bootstrap', 'login']);


  mod.run(function($http) {
    //$http.defaults.headers.common.AuthToken = 'Basic YmVlcDpib29w'
  });

})();
