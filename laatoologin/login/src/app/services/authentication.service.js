(function() {
  'use strict';

  var mod = angular.module('loginui');

  mod.factory('AuthenticationService', AuthenticationService);

  /** @ngInject */
  function AuthenticationService($resource){
	var data = $resource(LocalAuthServer, {}, {
	      login:{
	          method:'POST'
	       }
      });
      return data;	
  }

	

})();
