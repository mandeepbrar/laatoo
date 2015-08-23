(function() {
  'use strict';

  angular
    .module('media')
    .directive('media', media);

  /** @ngInject */
  function media() {
    var directive = {
      restrict: 'E',
      templateUrl: 'app/components/media/media.view.html',
      replace: true,
	  scope: {},
      transclude: true,
      controller: MediaCtrl,
      controllerAs: 'media',
      bindToController: true
    };

    return directive;

    /** @ngInject */
    function MediaCtrl($scope, $element, $attrs, $http ) {
	  if($attrs.type) {
		console.log("setting type");
		$scope.type = $attrs.type;
	  }		
      if($attrs.class) {
        $scope.class = "class="+$attrs.class;
      }
	  if($attrs.height) {
		$scope.height = $attrs.height;	
	  }
	  if($attrs.width) {
		$scope.width = $attrs.width;	
	  }
      var source;
      if($attrs.source) {
        $scope.source = $attrs.source;
      } 
	  $attrs.$observe('source', function(passedval) {
		passedval = passedval.replace("watch?v=", "v/");
		console.log("setting source");
		console.log(passedval);
		  $scope.source = passedval;
		
      });		
	  $attrs.$observe('type', function(passedval) {
		console.log("setting type " + passedval);
		  $scope.type = passedval;
		
      });		
	}	
  }

})();