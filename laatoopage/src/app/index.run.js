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
			var viewmode = value.viewmode;
			var obj = {};
			obj[view] = 	{ templateUrl: templatepath };
		    mainapp.stateProvider.state(actionName, 
				{ url: url, views : obj });
		}		  
		mainapp.urlRouteProvider.otherwise('/');
	}
	catch(exc) {
		console.log(exc);
	}
	
			
    $log.debug('runBlock end');
  }

})();
