(function() {
  'use strict';

  var app = angular.module('loginui');
  app.controller('LoginController', LoginController);

  /** @ngInject */ 
  function LoginController ($scope, $state, $location, $localStorage, AuthenticationService) {
      $scope.login = function() {
		var id = $scope.username;
		var pass = $scope.password;
        var result = AuthenticationService.login({"Id": id,"Password":pass}, function(data, headers) {
			localStorage.auth = headers(AuthToken);
			window.location.href = SuccessRedirect;
		},
		function(err) {
			console.log("err" + err);
		} );
        if(result) {
            $state.go('home');
        }
      };
  }


})();
