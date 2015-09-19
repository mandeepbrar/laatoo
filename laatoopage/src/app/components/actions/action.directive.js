(function() {
  'use strict';

  var actionsapp = angular
    .module('actions');
  actionsapp.directive('action', action);

  /** @ngInject */
  function action() {
    var directive = {
      restrict: 'E',
      templateUrl: function($element, $attrs) {
			if($attrs.widget === 'button') {
				return 'app/components/actions/button.view.html';
			}
			if($attrs.widget === 'link') {
				return 'app/components/actions/link.view.html';
			}
			return 'app/components/actions/button.view.html';
		},
      replace: true,
	  transclude: true,
	  scope: { },
      controller: ActionCtrl,
      controllerAs: 'action',
      bindToController: true
    };
    return directive;

    /** @ngInject */
    function ActionCtrl($scope, $element, $attrs, $http, $templateCache, $state, dialogs ) {
      var name = $attrs.name;
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
  	  $scope.action = window.pageConf.actions[name];
	  $scope.actionFunc = function() {};
	  var params = {};
	  var actiontype = $scope.action.actiontype;
	  if(actiontype === 'restcall') {
		$scope.route = '#';
		var successstate = $attrs.successstate;
		$scope.actionFunc = function() { 
			var url = $scope.action.url;
			for(var prop in params) {				
				var key = prop;
				var re = new RegExp("{"+prop+"}", 'g');				
				url = url.replace(re, params[prop]);
			}

		    var dlg = dialogs.confirm('Confirm', 'Are you sure you want to proceed?');
		    dlg.result.then(function(btn){
				 $http({method: $scope.action.restmethod, url: url }).then(
			        function(response) {
						if(successstate) {
							$state.go(successstate);							
						} else {
							$state.go($state.current, {}, {reload: true});						
						}
			        },
			        function(errorResponse) {
			            if(errorResponse.status == 0) {
						  dialogs.error('Error','Could not connect to the website. Please check your internet connection or the website is offline.');
			               return;
			            }
						dialogs.error('Error','Action could not be completed. ' + errorResponse.statusText);
			        }
			    );
			});														
		};
	  }
	  else {
		  $scope.route = $state.href(name);		
	  }
	  $attrs.$observe('params', function(passedVal){
		  	params = JSON.parse(passedVal);
	  		if(actiontype != 'restcall') {
			  $scope.route = $state.href(name, params);
			}
	  });
    }
  }

})();