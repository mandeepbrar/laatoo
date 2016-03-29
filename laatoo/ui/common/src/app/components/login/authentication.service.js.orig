(function() {
  'use strict';

  var mod = angular.module('login');

  mod.factory('AuthenticationService', AuthenticationService);

  /** @ngInject */
  function AuthenticationService($resource){
	var data = $resource(pageConf.LocalAuthServer, {}, {
	      login:{
	          method:'POST'
	       }
      });
      return data;	
  }

	

})();
