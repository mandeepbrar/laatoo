(function() {

    'use strict';
    var mainApp = angular.module('main', ['ngAnimate', 'ngCookies', 'ngTouch', 'ngSanitize', 'ngResource', 'ui.router', 'ui.bootstrap', 'login']);

    fetchData().then(bootstrapApplication);

    function fetchData() {
        var initInjector = angular.injector(["ng"]);
        var $http = initInjector.get("$http");
        var confUrl = document.location.href + "/conf";
        return $http.get(confUrl).then(
            function(response) {
                window.pageConf = response.data;
            },
            function(errorResponse) {
              console.log("error communicating with server");
            }
        );
    }

    function bootstrapApplication() {
        angular.element(document).ready(function() {
            angular.bootstrap(document, ["main"]);
        });
    }
}());

