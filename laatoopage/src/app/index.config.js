(function() {
  'use strict';

  var mainapp = angular
    .module('main')
    .config(config);

  /** @ngInject */
  function config($stateProvider, $urlRouterProvider, $httpProvider) {
	mainapp.stateProvider = $stateProvider;
	mainapp.urlRouteProvider = $urlRouterProvider;
	var token = localStorage.auth;
	if(token && token!=null && token.length > 0) {				
		$httpProvider.defaults.headers.common[pageConf.AuthToken] = token;
	}
  }

})();
