function logout() {
    localStorage.auth = "";
    localStorage.user = "";
    window.location.href = window.pageConf.HomePage;
};

function bootstrapApplication(modules) {
    var initInjector = angular.injector(["ng"]);
    var $httpProvider = initInjector.get("$http");
    if (window.pageConf.AuthRequired) {
        var token = localStorage.auth;
        if (token && token != null && token.length > 0) {
            startApplication($httpProvider, token, modules);
        } else {
            window.location.href = window.pageConf.AuthPage;
        }
    } else {
        startApplication($httpProvider, '', modules);
    }
};



function buildPermissions() {
    var permissions = [];
    if (localStorage.permissions != null) {
        permissions = localStorage.permissions.split(",");
    }
	window.isAllowed = function(action) {
	    return permissions.indexOf(action) > -1;
	};
    for (var i in window.pageConf.actions) {
        var action = window.pageConf.actions[i];
        if (action.permission) {
            if (action.permission == "none") {
                action.allowed = true;
            } else {
                action.allowed = window.isAllowed(action.permission);
            }
        } else {
            action.allowed = false;
        }
    }
};


function startApplication($http, token, modules) {
	console.log("start app");
    var docUrl = document.location.href;
    var loc = docUrl.indexOf("#");
    if (loc > 0) {
        docUrl = docUrl.substring(0, loc);
    }
    var confUrl = docUrl + "/conf";
    var authheaders = {};
    authheaders[window.pageConf.AuthToken] = token;
    return $http.get(confUrl, {
        headers: authheaders
    }).then(
        function(response) {
            if (response.data.partials) {
                window.pageConf.partials = response.data.partials;
            } else {
                window.pageConf.partials = [];
            }
            angular.element(document).ready(function() {
                angular.bootstrap(document, modules);
            });
        },
        function(errorResponse) {
            if (errorResponse.status == 401) {
                window.location.href = window.pageConf.AuthPage;
                localStorage.auth = "";
                localStorage.user = "";
            }
            console.log(errorResponse);
        }
    );
};

function initializeApp(app, sp, hp, dp, tp, provide) {
    app.stateProvider = sp;
    //app.urlRouteProvider = $urlRouterProvider;
    window.buildPermissions();
    dp.useBackdrop('static');
    dp.setSize('sm');
    tp.preferredLanguage('en-US');
    provide.factory('myHttpInterceptor', function($q) {
        var token = localStorage.auth;
        if (token && token != null && token.length > 0) {
            hp.defaults.headers.common[pageConf.AuthToken] = token;
        }
        return {
            'response': function(response) {
                // do something on success
                return response || $q.when(response);
            },
            'responseError': function(rejection) {
                console.log("response error");
                console.log(rejection);
                if (rejection.status == 401) {
                    localStorage.auth = "";
                    localStorage.user = "";
                    window.location.href = window.pageConf.AuthPage;
                    return;
                }
                // do something on error
                return $q.reject(rejection);
            }
        };
    });
    hp.defaults.useXDomain = true;
    hp.interceptors.push('myHttpInterceptor');
};