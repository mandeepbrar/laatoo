(function() {
  'use strict';

  angular
    .module('view')
    .directive('view', view);

  /** @ngInject */
  function view() {
    var directive = {
      restrict: 'E',
      templateUrl: function($element, $attrs) {
			if($attrs.viewtype === 'ul') {
				return 'app/components/links/ul.view.html';
			}
			if($attrs.viewtype === 'table') {
				return 'app/components/view/table.view.html';
			}
			return 'app/components/view/view.view.html';
		},
      replace: true,
	  scope: { },
      transclude: true,
      controller: ViewCtrl,
      controllerAs: 'view',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function ViewCtrl($scope, $element, $attrs, ViewService, $http) {
      var name = $attrs.name;
	  $scope.params = {};
	  $scope.modelname = 'viewrows';
      if($attrs.class) {
      	$scope.class = "class="+$attrs.class;
      }
  	  if($attrs.args) {
		$scope.params = angular.fromJson($attrs.args);
	  }
  	  if($attrs.modelname) {
		$scope.modelname = $attrs.modelname;
	  }
      if($attrs.editable) {
      	$scope.editable = ($attrs.editable == 'true');
		$scope.submitText = "Save";
		var actionUrl = "";
	  	if($attrs.action) {
			actionUrl = $attrs.action;
	  	}
		$scope.onSubmit = function() {
			console.log("submit view");
			$http.put(actionUrl, $scope[$scope.modelname]).then(
		       function(response) {
					console.log(response);
		       },
		       function(errorResponse) {
		         	console.log("error communicating with server");
		       }
		   );
		};
      }
	  $scope.params['viewname'] = name;
	  $scope.refreshView = function() {
		  $scope[$scope.modelname] = ViewService.query($scope.params);	  		
	  };
	  $scope.refreshView();
    }
  }

})();