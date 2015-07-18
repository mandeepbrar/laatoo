(function() {
  'use strict';

  var mod = angular.module('designerui');

  mod.service('AuthenticationService', AuthenticationService);

  function AuthenticationService($rootScope ){
    this.authenticateLocal = function(user, pass) {
      $rootScope.currentUser = 'mandeep';
	  return true;
	}
  }

})();
