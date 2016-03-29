(function() {
  'use strict';

  var mainapp = angular
    .module('main')
    .run(runBlock);

  /** @ngInject */
  function runBlock($log, $templateCache) {
	if(window.pageConf.partials) {
		for(var i=0; i<window.pageConf.partials.length;i++ ) {
			var partial = window.pageConf.partials[i];
			$templateCache.put(partial.Path, partial.Template);	
		}		
	}

	try {
		var actions = window.pageConf.actions;
		for( var actionName in actions) {
			var value = actions[actionName];
			var view = value.view;
			var url = value.url;		
			var templatepath = value.templatepath;		
			var actiontype = value.actiontype;		
			if(actiontype === 'openpartialpage') {
				var viewmode = value.viewmode;
				var obj = {};
				obj[view] = 	{ templateUrl: templatepath };
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
	
			    mainapp.stateProvider.state(actionName, 
					{ url: url, views : obj, params: value.params, controller: cfunc });
				
			}
		}		  
		mainapp.urlRouteProvider.otherwise('/');
	}
	catch(exc) {
		console.log(exc);
	}
	
			
    $log.debug('runBlock end');
  }

})();
