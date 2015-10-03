(function() {
  'use strict';

  var app = angular.module('login');
  app.controller('LoginController', LoginController);
  app.controller('OAuthLoginController', OAuthLoginController);

  /** @ngInject */ 
  function LoginController ($scope, $state, $location, AuthenticationService, dialogs, $http) {
      $scope.login = function() {
		var id = $scope.username;
		var pass = $scope.password;
        var result = AuthenticationService.login({"Id": id,"Password":pass}, function(data, headers) {
			localStorage.auth = headers(pageConf.AuthToken);
			localStorage.permissions = data.Permissions;
			window.location.href = pageConf.SuccessRedirect;
		},
		function(err) {
			if(err.status == 401) {
				localStorage.auth = "";				
				localStorage.permissions = null;
			}
			dialogs.error('Error','Login unsuccessful.');
			console.log(err);
		});
      };
	  $scope.oauthLogin = function(url) {
		/*var dlg = dialogs.create('app/components/login/oauthlogin.html','OAuthLoginController',{url: url},{key: false,back: 'static'});
    		dlg.result.then(function(result) {
			console.log(result);
            $scope.name = name;
    		},function(){
        		//$scope.name = 'You decided not to enter in your name, that makes me sad.';
    		});*/
		var newwindow=window.open(url,'name', 'modal=true,height=600,width=600');
		if (window.focus) {newwindow.focus()}
	  };
	  window.oauthLogin = function(type, state, code) {
		var url = $scope.google;
		if(type === 'facebook') {
			url = $scope.facebook;
		}
		var data = {state: state, code:code};
        var result = $http.post(url, JSON.stringify(data)).then( function(response) {
			localStorage.auth = response.headers(pageConf.AuthToken);
			localStorage.permissions = response.data.Permissions;
			window.location.href = pageConf.SuccessRedirect;
		},
		function(err) {
			if(err.status == 401) {
				localStorage.auth = "";				
				localStorage.permissions = null;
			}
			dialogs.error('Error','Login unsuccessful.');
			console.log(err);
		});
	  }
  }

  /** @ngInject */ 
  function OAuthLoginController ($scope, data) {
		$scope.url = data.url;
  }

})();
