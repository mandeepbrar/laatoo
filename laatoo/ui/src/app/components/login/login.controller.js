(function() {
    'use strict';

    var app = angular.module('login');
    app.controller('LoginController', LoginController);
    app.controller('OAuthLoginController', OAuthLoginController);

    /** @ngInject */
    function LoginController($scope, $state, $location, DataService, RequestBuilderService, dialogs, $http) {
        $scope.login = function() {
            var id = $scope.username;
            var pass = $scope.password;
            localStorage.auth = null;
			var loginSuccess = function(response) {
				console.log(response);
				
				//headers method is changing them to lower case
                var token = response.info[document.Application.Security.AuthToken.toLowerCase()];
				
				localStorage.auth = token;
                localStorage.permissions = response.data.Permissions;
                localStorage.user = response.data.Id;
                document.buildPermissions();
				window.location.href = document.Application.HomePage;
			};
			var loginFailure = function(response) {
				if (response.code == 2) {
					localStorage.auth = "";
					localStorage.user = null;
					localStorage.permissions = null;
				}
				dialogs.error('Error', 'Login unsuccessful.');
				console.log(response);
			};
			var req = RequestBuilderService.DefaultRequest(null, {"Id": id, "Password": pass});

			DataService.ExecuteService('login', req, loginSuccess, loginFailure);
        };
        $scope.register = function() {
            var name = $scope.NewUserName;
            var id = $scope.NewUserEmail;
            var pass = $scope.NewUserPass;
            var confirmpass = $scope.NewUserConfirmPass;
            if (confirmpass != pass) {
                return;
            }
            var usertoberegistered = {
                "Id": id,
                "Password": pass,
                "Name": name,
                "Email": id
            };
			var registrationSuccess = function(resposne) {
                    window.location.href = document.Application.RegSuccessRedirect;
                };
			var registrationFailure = function(err) {
                    dialogs.error('Error', 'Registration unsuccessful.');
                    console.log(err);
                };
			var req = RequestBuilderService.DefaultRequest(null, usertoberegistered);

			DataService.ExecuteService('register', req, registrationSuccess, registrationFailure);
        };
        $scope.oauthLogin = function(url) {
            /*var dlg = dialogs.create('app/components/login/oauthlogin.html','OAuthLoginController',{url: url},{key: false,back: 'static'});
                dlg.result.then(function(result) {
            $scope.name = name;
                },function(){
                        //$scope.name = 'You decided not to enter in your name, that makes me sad.';
                });*/
            var newwindow = window.open(url, 'name', 'modal=true,height=600,width=600');
            if (window.focus) {
                newwindow.focus()
            }
            var receiveMessage = function(event) {
                if (event.source != newwindow) {
                    return;
                }
                document.oauthLogin(event.data.type, event.data.state, event.data.code);
            };
            window.addEventListener("message", receiveMessage, false);
        };
        document.oauthLogin = function(type, state, code) {
            var url = $scope.google;
            if (type === 'facebook') {
                url = $scope.facebook;
            }
            var data = {
                state: state,
                code: code
            };
            var result = $http.post(url, JSON.stringify(data)).then(function(response) {
                    localStorage.auth = response.headers(pageConf.AuthToken);
                    localStorage.permissions = response.data.Permissions;
                    window.location.href = pageConf.HomePage;
                },
                function(err) {
                    if (err.status == 401) {
                        localStorage.auth = "";
                        localStorage.permissions = null;
                    }
                    dialogs.error('Error', 'Login unsuccessful.');
                    console.log(err);
                });
        }
    }

    /** @ngInject */
    function OAuthLoginController($scope, data) {
        $scope.url = data.url;
    }

})();