(function() {

    'use strict';
    var mainApp = angular.module('main', ['ngAnimate', 'ngCookies', 'ngTouch', 'ngSanitize', 'ngResource', 'ui.router', 'ui.bootstrap', 'login', 'view', 'actions', 'entity', 'media']);

    bootstrapApplication();
	
	window.logout = function() {
		if(localStorage.auth != null && localStorage.auth.length > 0 ) {				
			localStorage.auth = "";				
		}			
	};

    function bootstrapApplication() {
		if(window.pageConf.AuthRequired) {
			var token = localStorage.auth;
			if(token && token!=null && token.length > 0) {				
				startApplication();
			}			
			else {
				window.location.href = window.pageConf.AuthPage;				
			}
		}
		else {
			startApplication();
		}
	}
    function startApplication() {
        var initInjector = angular.injector(["ng"]);
        var $http = initInjector.get("$http");
		var docUrl = document.location.href;
		var loc = docUrl.indexOf("#");
		if(loc >0) {
			docUrl = docUrl.substring(0, loc);
		}
        var confUrl = docUrl + "/conf";
        return $http.get(confUrl).then(
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
              console.log("error communicating with server");
            }
        );
    }
}());

