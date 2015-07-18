(function() {
  'use strict';

  var app = angular.module('designerui');
  app.controller("LoginController", function ($scope, $state, AuthenticationService) {
      $scope.login = function() {
        var result = AuthenticationService.authenticateLocal("someuser", "somepass");
        if(result) {
            $state.go("home");
        }
      };
  });


})();
