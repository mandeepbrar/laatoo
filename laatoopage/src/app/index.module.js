(function() {

    'use strict';
    var mainApp = angular.module('main', ['ngAnimate', 'ngCookies', 'ngTouch', 'ngSanitize', 'ngResource', 'ui.router', 'ui.bootstrap', 'dialogs.main', 'pascalprecht.translate', 'dialogs.default-translations', 'login', 'view', 'actions', 'entity', 'media', 'smart-table', 'uigrid']);

    bootstrapApplication();
	
	window.logout = function() {
		if(localStorage.auth != null && localStorage.auth.length > 0 ) {				
			localStorage.auth = "";				
		}			
	};	

    function bootstrapApplication() {
        var initInjector = angular.injector(["ng"]);
        var $httpProvider = initInjector.get("$http");
		if(window.pageConf.AuthRequired) {
			var token = localStorage.auth;
			if(token && token!=null && token.length > 0) {				
				startApplication($httpProvider, token);
			}			
			else {
				window.location.href = window.pageConf.AuthPage;				
			}
		}
		else {
			startApplication($httpProvider, '');
		}
	}
	
    function startApplication($http, token) {
		var docUrl = document.location.href;
		var loc = docUrl.indexOf("#");
		if(loc >0) {
			docUrl = docUrl.substring(0, loc);
		}
        var confUrl = docUrl + "/conf";
		var authheaders = {};
		authheaders[window.pageConf.AuthToken] = token;
        return $http.get(confUrl, {headers: authheaders}).then(
            function(response) {
				if(response.data.partials) {
					window.pageConf.partials = response.data.partials;					
				} else {
					window.pageConf.partials = [];					
				}
		        angular.element(document).ready(function() {
		            angular.bootstrap(document, ["main"]);
		        });
            },
            function(errorResponse) {
			  if(errorResponse.status == 401) {
				window.location.href = window.pageConf.AuthPage;								
			  }
			  console.log(errorResponse);
            }
        );
    }
}());

