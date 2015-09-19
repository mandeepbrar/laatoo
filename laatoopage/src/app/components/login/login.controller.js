(function() {
  'use strict';

  var app = angular.module('login');
  app.controller('LoginController', LoginController);

  /** @ngInject */ 
  function LoginController ($scope, $state, $location, AuthenticationService, dialogs) {
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
  }


})();
