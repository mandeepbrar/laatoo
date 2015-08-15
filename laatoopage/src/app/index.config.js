(function() {
  'use strict';

  var mainapp = angular
    .module('main')
    .config(config);

  /** @ngInject */
  function config($stateProvider, $urlRouterProvider, $httpProvider) {
	mainapp.stateProvider = $stateProvider;
	mainapp.urlRouteProvider = $urlRouterProvider;
  }

})();
