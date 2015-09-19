(function() {
  'use strict';

  var mainapp = angular
    .module('main')
    .config(config);

  /** @ngInject */
  function config($stateProvider, $urlRouterProvider, $httpProvider, dialogsProvider, $translateProvider, $provide) {
	mainapp.stateProvider = $stateProvider;
	mainapp.urlRouteProvider = $urlRouterProvider;
	var token = localStorage.auth;
	var permissions = [];
	if(localStorage.permissions != null) {
		permissions = localStorage.permissions.split(",");		
	}
	window.isAllowed = function(action) {
		return permissions.indexOf(action)>-1;
	}
	for(var i in window.pageConf.actions) {
		var action = window.pageConf.actions[i];
		if(action.permission) {
			if(action.permission =="none") {
				action.allowed = true;
			} else {
				action.allowed = window.isAllowed(action.permission);							
			}
		} else {
			action.allowed = false;
		}
	}
	if(token && token!=null && token.length > 0) {				
		$httpProvider.defaults.headers.common[pageConf.AuthToken] = token;
	}
	
	dialogsProvider.useBackdrop('static');
	dialogsProvider.setSize('sm');
	$translateProvider.preferredLanguage('en-US');	
	$provide.factory('myHttpInterceptor', function($q) {
	  return {
	    'response': function(response) {
	      // do something on success
	      return response || $q.when(response);
	    },
	
	   'responseError': function(rejection) {
            if(rejection.status == 401) {
				window.location.href = window.pageConf.AuthPage;								
                return;
            }
	      // do something on error
	      return $q.reject(rejection);
	    }
	  };
	});
	
	$httpProvider.interceptors.push('myHttpInterceptor');
	
	
	
  }

})();
