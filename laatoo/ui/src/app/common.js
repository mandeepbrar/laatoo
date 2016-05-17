document.logout = function() {
    localStorage.auth = "";
    localStorage.user = "";
    window.location.href = document.Application.HomePage;
};

document.isLoggedIn = function() {
	return localStorage.auth != "";
};

document.buildPermissions = function(){
    var permissions = [];
    if (localStorage.permissions != null) {
        permissions = localStorage.permissions.split(",");
    }
	var isAllowed = function(action) {
	    return permissions.indexOf(action) > -1;
	};
	var actions = document.Actions;
    for (var i in actions) {
        var action = actions[i];
        if (action.permission) {
            if (action.permission == "none") {
                action.allowed = true;
            } else {
                action.allowed = isAllowed(action.permission);
            }
        } else {
            action.allowed = false;
        }
    }
};

/*
document.startApplication = function($http, token, modules) {
    angular.element(document).ready(function() {
        angular.bootstrap(document, modules);
    });
};*/

document.promptAuthentication = function(appinitialized) {
	if(window.location.href != document.Application.Security.AuthPage) {
		window.location.href = document.Application.Security.AuthPage
	}
}

document.initializeApp = function(app, sp, hp, dp, tp, provide) {
    if (document.Application.Security.AuthRequired) {
        var token = localStorage.auth;
        if (token == null || token.length == 0) {
            //document.promptAuthentication();
			//return;
        }
    	//document.buildPermissions();
    } 
	
    app.stateProvider = sp;
    dp.useBackdrop('static');
    dp.setSize('sm');
    tp.preferredLanguage('en-US');
    provide.factory('myHttpInterceptor', function($q) {
        var token = localStorage.auth;
        if (token && token != null && token.length > 0) {
            hp.defaults.headers.common[document.Application.Security.AuthToken] = token;
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
                    //document.promptAuthentication();
                    return;
                }
                // do something on error
                return $q.reject(rejection);
            }
        };
    });
    hp.defaults.useXDomain = document.Application.Security.UseXDomain;
    hp.interceptors.push('myHttpInterceptor');
};


document.runApp = function(app, templateCache, dataservice, reqbuilderservice, state) {
	var req = reqbuilderservice.DefaultRequest(null, null)
	var partialsDownloaded = function(response) {
		try {
			var files = response.data.Files;
			for(var partialname in files) {
				var partial = files[partialname];
				templateCache.put(partial.Info.webroute, partial.Content);	
			}					
		}catch(ex){console.log(ex);}		

		try {
			var actions = document.Actions;
			if(actions) {
				for( var actionName in actions) {
					var value = actions[actionName];
					console.log("setting value true");
					value.allowed = true;
					var url = value.url;		
					var templatepath = value.templatepath;		
					var actiontype = value.actiontype;		
					if(actiontype === 'openpartialpage') {
						var viewmode = value.viewmode;		
						var obj = {};
						if(value.views) {
							obj = value.views;
						} 
						else {
							var view = value.view;
							var templatepath = value.templatepath;		
							obj[view] = { templateUrl: templatepath };
						}		
						var cfunc = function($scope, $stateParams) {
									for(var i = 0; i< len($stateParams); i++) {
										if(!$scope.params) {
											$scope.params = {};
										}
										var key = $stateParams[i];
										var val = $stateParams[key];
						     			$scope.params[key] = val;							
									}
					         	};
					    app.stateProvider.state(actionName, 
							{ url: url, views : obj, params: value.params, controller: cfunc });	
					}
				}		  
				state.go('Home');
			}
		}
		catch(exc) {
			console.log(exc);
		}

	
	};
	var partialsFailure = function(response) {
		console.log("Couldnt load partials");
	};
	var runcomplete = dataservice.ExecuteService(document.Application.PartialsSvc, req, partialsDownloaded, partialsFailure);
	document.RunComplete = runcomplete;
};

