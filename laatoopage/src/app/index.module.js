(function() {

    'use strict';
    var mainApp = angular.module('main', ['ngAnimate', 'ngCookies', 'ngTouch', 'ngSanitize', 'ngResource', 'ui.router', 'ui.bootstrap', 'login', 'view', 'actions', 'entity', 'media']);

    bootstrapApplication();

    function bootstrapApplication() {
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
				window.pageConf.partials = response.data.partials;
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

