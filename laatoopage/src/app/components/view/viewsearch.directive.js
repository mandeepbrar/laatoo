(function() {
  'use strict';

  angular
    .module('view')
    .directive('viewsearch', viewsearch);

  /** @ngInject */
  function viewsearch() {
    var directive = {
      restrict: 'E',
      template:'<div ng-transclude></div>',
      replace: true,
      transclude: true,
      controller: ViewSearchCtrl,
      controllerAs: 'viewsearch',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function ViewSearchCtrl($scope, $element, $attrs) {
	  if($attrs.popupsearch) {
		console.log("popup search");	
	  }
	  $scope.params = {};
	  $scope.applyFilter = function() {
		var parent = $scope.$parent;
		parent.params.args = $scope.params;
		parent.refreshView();
	  };
	  $scope.resetFilter = function() {
		parent = $scope.$parent;
		parent.params.args = $scope.params = {};
		parent.refreshView();
	  };
    }
  }

})();