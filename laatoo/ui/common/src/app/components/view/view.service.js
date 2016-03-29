(function() {
  'use strict';

  var mod = angular.module('view');

  mod.factory('ViewService', ViewService);

  /** @ngInject */
  function ViewService($resource){
	var data = $resource(pageConf.ViewsServer);
    return data;	
  }

	

})();
